---
type: decision
title: An outcome carries the doer's assertion and the checker's verdict separately
decided: 2026-09-05
stage: provisional
reopen_trigger: the two axes are found to always agree in real use, which would mean the split is recording a distinction nobody makes
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T23:40:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T00:24:45Z'}
---

# ADR-0007: An outcome carries the doer's assertion and the checker's verdict separately

## Summary

An outcome holds **two independent records**: what the doer claims, and what a
checker found. They may disagree, and the disagreement is the point. Closing a
work item as *completed* gates on the **verdict only**.

## Problem

**`verify` could only record success.** `internal/cli/verify.go` appends to
`verified`; an outcome with entries passes and one without has not been
checked. There was no way to say *I looked and it was false*.

`spec.md` §4.4 made that deliberate:

> There is no separate pass or fail field, because there is nothing to store
> that the verification record does not already say.

That is wrong, and the design already knew it elsewhere. `close.go` refuses to
deliver over an unreadable outcome with exactly this reasoning: *"An unreadable
outcome might be failing, and nothing here can tell."* Absence of a verdict and
a negative verdict are different facts, and §4.4 conflates them everywhere the
refusal does not.

Underneath sat a second gap. **Doers and checkers are separate parties** —
often a person checking an agent's work. The doer's claim had nowhere to live
at all, so the design distrusted assertions of completion by having no field
for them, which loses the claim rather than examining it.

## Decision

### Two axes, both events

```yaml
asserted:
  - {by: 'agent:opus-5/luma-backlog', at: …, as: failed}
  - {by: 'agent:opus-5/luma-backlog', at: …, as: succeeded}
verified:
  - {by: 'human:maintainer', at: …, as: proven}
```

- **`asserted`** — the doer, `as: succeeded` or `as: failed`. **A list, appended
  never replaced.** Absent means never attempted; the current claim is the last
  entry; the attempt count is the length.

**Both axes append.** Overwriting `asserted` would be the one place in this
design that destroys history — everywhere else the rule is create, never
overwrite: `verified` accumulates (§4.7), promotion copies (§4.8.1), succession
writes a new record (§4.6), and closing appends so a reopen cannot erase it
([[backlog/work-items/WORK-0020-reopen-a-work-item-that-was-closed]]).

It also costs nothing. `verified` is already an append-only list and the code
that appends to it exists; a single field would have been a second shape for no
gain, and would have made *failed once, then succeeded* indistinguishable from
*succeeded first time*.

**A second attempt is therefore visible**, which a single field could not
represent. What is still not represented is an attempt **in flight** — during
one, the last entry reads `failed`, which is stale rather than false. That
belongs to waves (§2.3) or to a claim (§6.5), both outside the first release.
- **`verified`** — the checker, `as: proven`, `as: disproven`, or
  `as: inconclusive`. Already a list.

**One key for all three events.** Every field here is a past participle and `as`
is what English puts after it — *asserted as succeeded, verified as proven,
closed as completed*. The values remain distinct per axis, so a value alone
still identifies which axis it belongs to; that was never the key's job.

**An outcome is a proposition, not a process** (§2.4 — *"a statement of a
condition that must hold"*), so a checker does not succeed or fail. They
determine whether a statement is true or false. `succeeded` and `failed` cannot
express a check that **ran correctly and found the condition false** — a
successful verification with a negative finding — and would collapse it into
the same word as *the check broke*.

**`proven` and `disproven` are both findings**: somebody looked and established
something. That is what frees the third value to mean the one case where
nothing was established.

**`inconclusive`** covers that case without claiming whose problem it is — a
`verify_by` that is prose rather than a check, an impossible measurement, or a
check that errored are all *no verdict yet*, and which one it was belongs in
`--evidence` rather than in the enum. It is the only verdict that reports an
absence.

`failed` was considered for this slot and rejected: the doer's axis already
uses it. Since both events carry the same `as` key, one record would hold
`as: failed` and `as: failed` on adjacent lines meaning different things —
identical in key *and* value, distinguishable only by which field they sat
under.

**The two axes use deliberately different words.** A doer says `succeeded` or
`failed`; a checker says `proven`, `disproven` or `inconclusive`. The doer is
reporting an attempt, the checker is establishing a fact, and no word appears on
both axes — so an assertion can never be misread as a verdict, including by
somebody scanning the raw file with only the value in front of them.

### The commands

```
outcome assert <ref> <succeeded|failed>
outcome verify <ref> <proven|disproven|inconclusive> [--evidence …]
```

**The verdict is required, not defaulted.** Recording proof must be said out
loud. The mistakes are asymmetric: a wrongly recorded failure is noise somebody
corrects, while a wrongly recorded success is the exact claim this design
exists to distrust — and `close … completed` trusts it without asking. A tool
whose thesis is that unbacked assertions are untrustworthy should not have
*asserts proof* as its zero-argument behavior.

### Closing gates on verification only

**Never on the assertion.** Gating on the doer's claim would gate on the thing
the design distrusts, and would let a doer close their own work.

`work-item close <ref> completed` is refused unless **at least one live outcome
exists and every one is proven.** Unreadable, unverified, disproven and
inconclusive are all the same answer: not proven. This collapses the three
refusal paths in `close.go` into one.

**Zero outcomes stays an explicit clause**, not a consequence — "every outcome
is proven" is vacuously true of none. `completed` is itself a declaration, and
claiming delivery of something never defined is a record contradicting itself
in one word, which is what keeps this refusal inside `spec.md` §5.0 rather than
carving an exception to it.

**The refusal message still names the files.** One rule in the logic, grouped
specifics in the output — *fix it* is only actionable if the caller learns
which outcome is unreadable and which are unverified.

### `--force` never touches the outcomes

The tempting implementation marks outcomes verified so the arithmetic comes out
clean. That destroys the record. Instead the outcomes are left exactly as they
are, completion still computes *two of five*, and the **work item** carries the
forced close. A reader afterwards sees a completed work item whose own
arithmetic disagrees with it — which is the truth, and is what should be
visible.

`--force` without `--reason` should say so, the way `verify` already notes a
missing `--evidence`. Forcing is when the paper trail is worth most and when
nobody feels like typing.

### Four dispositions

`completed` · `rejected` · `canceled` · `superseded`

```yaml
closed:
  - {by: 'human:maintainer', at: …, as: completed}
```

**A list, appended never replaced**, matching `asserted` and `verified`. `as`
rather than `reason`, which carries the prose. **`completed` rather than
`delivered`** — the tool cannot observe a handover and can compute a count, and
a record should never carry a word claiming more than the tool can defend.

| | accepted? | work started? |
| --- | --- | --- |
| `rejected` | no — never a consideration | no |
| `canceled` | yes, then changed our minds | no |
| `superseded` | either — replaced by another record | either |
| `completed` | yes | yes, and proven |

**`abandoned` is dropped.** The rule: *the enum carries what the record cannot
derive.* Whether work started is derivable — from whether it reached
`in_progress`, whether tasks exist, whether any outcome was asserted — so
storing it duplicates something computable. *Never a consideration* is a
statement of intent and cannot be derived, so `rejected` earns its place.
`superseded` survives on a technicality that matters: supersession is carried
by the successor (§4.6, §3.2), so the closed record never learns it was
replaced, and the successor may not exist yet at close time.

**Reject is a disposition, not a rung.** Triage's *reject* closes the record;
the ladder is untouched. Triage's *defer* writes nothing at all — it is the
board moving on, which is view state and not a mutation.

### Self-verification is observed, never blocked

`outcome.self-verified` joins the conditions in §5.2 when the asserter and the
only verifier are the same actor.

**Nothing is refused.** The first release ships no gate, which keeps §5.4's
*"a repository that declares nothing gets nothing"* literally true and needs no
amendment to §5.0. Enforcement — including the stricter posture for agents than
for people — is configuration, deferred.

**Correction appends, never erases.** A later entry with a different verdict
supersedes; the history stays whole, and *an agent said proven and a person
disagreed* remains visible. Same rule as promotion (§4.8.1) and succession
(§4.6).

## Why

**It makes the founding principle structural rather than aspirational.**
`principles.md` calls an unbacked assertion of completion *"the claim this
design is most interested in distrusting."* Today the design distrusts it by
having nowhere to put it. Recording the claim beside the evidence makes the gap
countable — *six asserted, two proven* is exactly what §11.1 says a board
exists to show and a file cannot.

**The same shape appears at both levels**, which is evidence the model is
coherent:

| Level | The assertion | The evidence |
| --- | --- | --- |
| Outcome | `asserted: [{as: succeeded}]` | `verified: [{as: proven}]` |
| Work item | closed as completed, forced | the completion arithmetic |

Claim and proof are recorded separately and allowed to disagree, at both.

**On the name.** `claim` was unavailable — `spec.md` §4.4 already rejected it
for an outcome field because it collides with claiming a task (§4.5), and the
word is load-bearing across §9.2, §5.2, §6.5 and `open-questions.md` §8, where
its storage is still unsettled. `attempted` carries no trust signal.
**`asserted` is the specification's own word for the untrusted thing**, and an
assertion is a claim made with confidence — which is the agent behavior being
distrusted.

## Alternatives

| Candidate | Set aside because |
| --- | --- |
| **One axis, as §4.4 has it** | Cannot distinguish *nobody checked* from *checked and failed*. The same test correctly removes `abandoned` from the dispositions and correctly adds verdicts here; §4.4 applied it and got this case wrong. |
| **`verify --disproven` / `--inconclusive`, success as default** | Backwards compatible and ergonomic. Set aside on the asymmetry above. *Reopened if requiring the verdict proves to cost more than the accidental-proof it prevents.* |
| **A separate verb for a negative verdict** | Reads better in isolation; splits one evidence trail across two commands and grows a verb per verdict. `check` is unavailable — §9.2 assigns it to evaluating conditions. |
| **Blocking agent self-verification** | What the maintainer first wanted, and it would be the first gate the tool ships — §5.0 and §5.4 both need amending for it. Deferred to configuration. *Reopened when somebody wants the strict posture enforced rather than reported.* |
| **A single `asserted` field, replaced on each attempt** | Considered and reversed. It cannot represent a second attempt, it loses whether an outcome took one try or four, and it would be the only field in the model that overwrites. The list costs the same code. |
| **Keeping `abandoned`** | Derivable. *Reopened if the distinction between stopped-before-starting and stopped-midway turns out to be one people make and the record cannot reconstruct.* |

## Revisit When

- The two axes always agree in practice — the split would then be recording a
  distinction nobody makes.
- Waves arrive, which would bind each assertion to the attempt it belonged to —
  the remaining half of
  [[backlog/work-items/WORK-0019-a-ledger-of-attempts-against-an-outcome]].
- Somebody needs self-verification enforced rather than observed.

## References

- `docs/spec.md` §2.4, §4.4, §4.7 — the outcome, and the sentence this amends.
- `docs/spec.md` §5.0, §5.2, §5.3 — refusals, conditions, and closing.
- `docs/principles.md` — completion is evidenced, not asserted.
- `docs/format-requests.md` — a verdict on a verification event is likely the
  second ask of the knowledge format, after evidence itself.
