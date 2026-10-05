---
type: work-item
key: BACK-0114
title: Harvest what the infra repository's project system proved in use
workflow_status: captured
rank: 010.0930.000
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-10-04T17:21:15Z'}
description: Is there anything worth copying from the infra repository's committed project-folder system and its project-new / project-close skills? Capture everything the comparison surfaced as considerations, with how the analysis was arrived at, so there is one place to decide what this project adopts.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-10-04T17:25:02Z'}
---

# Harvest what the infra repository's project system proved in use

> **"infra" is the codeword this record uses for a private repository of this
> maintainer's** — a second project that adopted this tool, and that ran a
> hand-written project-folder system of its own for months before that. Its real
> name is deliberately absent here, and so are its hosts, addresses, service
> names and credential locations. Ask the maintainer if you need to find it.

## The problem

**Is there anything worth copying?** infra keeps a committed project-folder
system that has been worked daily for months, plus two skills — `project-new`
and `project-close` — that drive it. Its backlog has since migrated into this
tool, so the two systems now sit next to each other and can be compared on
evidence rather than on taste.

Everything the comparison surfaced belongs in one place, with **how the analysis
was arrived at**, so there is a single record to decide against instead of a
conversation nobody can re-read.

## What is being delivered

**Understanding, and whatever work items it justifies.** A recommendation on
each consideration below: take it, leave it, or fold it into a record that
already exists. Concluding that most of them should not be copied is a complete
result.

## Out of scope

**Nothing is adopted by this record.** It collects and recommends; adopting any
of it is a separate decision, and several of the candidates belong to records
that already exist (see *How this relates to what exists*).

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### How the comparison was done

1. Read infra's project-system README, its retired backlog tombstone, and all
   five of its file templates.
2. Inventoried which files its 37 active and 48 archived projects actually grew
   — README 48, tasks 43, journal 38, decisions 11, plan 8, exploration 7 — to
   separate designed practice from surviving practice.
3. Read its deepest project end to end: a 1078-line journal, plus that project's
   README, task list and decision log. Read a second project's decision log for
   a different style.
4. Read all 347 lines of its agent instructions file.
5. Read both skills (`project-new`, `project-close`) against this project's
   `backlog-capture` and `backlog-transition` procedures.
6. Re-read `spec.md` §2.2, §4.1, §4.8, §5.2, §5.3, §5.5, §7.2, `lifecycle.md`,
   `open-questions.md` §2 and `format-requests.md`, and sampled this corpus
   (BACK-0003, and the bodies of BACK-0032, BACK-0033, BACK-0034, BACK-0038,
   BACK-0071) so nothing already captured is presented as new.
7. Checked staleness with `git log -1` per active project, to keep the
   comparison fair in both directions.

### The framing correction

**infra's backlog is not doing anything better — it is a tombstone.** All 178
entries migrated into this tool, and its own documentation marks the
project-folder system deprecated. What is still live there is the **project
folder**, which is the layer that sits *under* a work item. So the real
comparison is one of its project directories against one of our
`work-items/<key>/` directories, and that is where the transferable material is.

### Tier 1 — structural, and the reason this is worth looking at

**1. A mutable current-state head above an append-only log.** Its longest
journal opens with a block headed *resume here*, labelled *"Kept current;
everything below this section is the chronological log. Last updated
&lt;date&gt;"*, carrying stable sub-headings — how to reach the thing, live state
as measured on a date, what to do next in order, open questions, what was pushed
out of scope, gotchas, a map of the documentation — and **below it** the
chronological log, oldest-first.

`spec.md` §5.5 mandates the opposite on both axes: pure newest-first prepend,
*"append, never curate"*, *"headings are named after what they settle, not drawn
from a fixed template"*. And `format-requests.md` records that our resume
pointer was harvested **from these journals**. It was — but the file shows the
practice moved one step further than what we took.

**The reason is mechanical: standing knowledge cannot survive in a prepend-only
file.** Something learned in entry 3 is buried by entry 20 unless it is
re-copied into every new entry, and nobody does that. infra's answer to drift
is not to forbid the block but to **stamp it** — *last updated*, *measured on* —
so a reader can price its age.

**2. Gotchas as a standing, numbered, accumulating artifact.** Ten of them,
under *"⚠ Gotchas that will waste your time if you forget them"*, each shaped:
symptom → the wrong guess → **the tell** → cause → fix. One example, with the
product genericized: *an API returns success for enum values it silently
ignores, which left a replication feature completely dead behind a green deploy
— so always read settings back and assert.*

We have *"exact commands, values, and gotchas — verbatim"* as **one row** in a
table of what a journal entry *may* carry. That is a licence, not a surface.
Theirs is a list that grows and stays at the top.

**3. Same-symptom discrimination, and a feedback path back to the project's own
rules.** One gotcha records that two different failures share a single symptom,
names the discriminator (*connecting by address works while connecting by name
fails*), and then says the expensive wrong guess is the one **the project's own
agent instructions tell you to make first**.

Two things we have no slot for: **differential diagnosis** as a knowledge shape,
and a channel for *our own standing instruction misfires in this case*. We have
`format-requests.md` for asks of the format and violation records for an agent
doing something unwanted — nothing for *the rule was followed and the rule was
wrong here*.

**4. Does the fix travel?** A fix is annotated *"⚠ This fix does not travel"*,
with the reason (it landed in untracked local client configuration on one
machine, so anything else hits the same wall) and the durable alternative named
as a deliberately-undecided backlog candidate.

Our evidence model records **that an outcome holds**. It never records **where
the fix lives, or whether it survives a rebuild** — which for a tool whose whole
premise is committed-and-portable is a natural field on evidence, and the
difference between *verified* and *verified in one place and nowhere else*.

**5. An operational access block.** A kept-current table of how to reach the
thing: consoles, where the credential lives (a pointer to a password manager,
with *deliberately not in the repo* stated), the automation identity, the one
command that applies configuration, plus a verbatim snippet for a fiddly
authentication step and two warnings about calling conventions. Our
`references` (§4.1.2) is explicitly *pointers this tool does not follow* — what
to read, not how to operate. **Weaker transfer for a Go command line than for
infrastructure**, but the shape generalizes to *the commands and conventions you
would otherwise reconstruct*.

### Tier 2 — working rules, candidates for this project's own instructions

**6. Answering versus acting, as a closed two-bucket rule.** The test is one
sentence — *is the user telling me to do something, or asking me about
something?* — with the both-in-one-message case worked through, and the default:
*unsure → treat it as a question; asking costs a sentence, an unwanted change
costs cleanup and repo drift.* We have *discuss before writing*, scoped only to
exploratory work.

**7. A closed list of permitted exceptions, plus no self-authorizing outside
it.** Their strongest default names **six** good reasons to depart from it
(discovery, prototyping, urgency, no good path, disproportionate effort,
bootstrapping) and then: *if your reason is not one of these six, do not
self-authorize it — surface it and ask. It may be a genuine new case (worth
adding to this list), or a sign the default should hold.* That final clause
makes the list **grow by evidence**. Our strong defaults have no exception list
and no escalation instruction, so an agent either obeys mechanically or departs
silently — and we learn nothing either way.

**8. A standing end-of-response obligation for work left out of sync.** Theirs:
after editing anything deployable, end the response with a one-line *pending
deploy* summary, *to prevent the edited-locally-forgot-to-apply failure mode.*
Our analogue is documented in our own corpus: BACK-0003's journal records that
*several decisions moved after the code was written.* Specification-moved,
implementation-lags is the same failure with nothing catching it.

**9. An accountability log for repeated advice.** One project carries a
*recommendation log* — recorded at the maintainer's request — tallying three
successive hardware recommendations that did not work, with the honest summary
that the sequence had been a costly slog. The agent tracking its own batting
average, inside the work item. **Not a violation record**: nothing was unwanted,
the advice was simply wrong repeatedly, and that only shows in aggregate.

### Tier 3 — technique

**10. The tombstone.** The retired backlog file was replaced in place by a
record of its own retirement: the entry count, that none were discarded, the
split between individually-reviewed and mechanically-imported, both commit
hashes, the exact command to read the original, the **load-bearing** import
marker with *do not strip it*, the handful of records flagged as already marked
done in the source, one redaction and **how it had survived** (the pre-push
secret guard did not scan that directory), and why it was retired rather than
run in parallel — *two backlogs is worse than either*, evidenced by the very
first capture into the new system needing a hand cross-check. The now-obsolete
template is kept *only until the rest of the system retires.*

**11. A one-line sizing rule at the task boundary.** *If a task needs a full
write-up — problem, value, success criteria — it is really a project, not a
task.* Directly usable on the question BACK-0026 is holding.

### The close-out skill — four things ours does not do

Their `project-close` is weaker than our closing procedure almost everywhere:
no gating, no dispositions, no forcing discipline, no journal-first, no
show-then-ask-before-filing. Four exceptions.

**a. Promotion is the headline, and the candidate list is wider than ours.**
Their framing: *"The easy-to-skip, highest-value step is promotion. If you just
move the folder to archived, the knowledge dies with it — buried in a journal
nobody reads."* Our close names **two** registers that outlive the work item — a
decision and a violation — and stops. Theirs names four homes. The two we are
missing:

- **The project's own documentation.** A closed work item that changed how
  `spec.md` or `workflow-status.md` should read has no step that asks.
  `lifecycle.md` §2 has it as *Propagate* — *"update documentation and
  references, mark things stale"* — and `spec.md` §5.5 requires *"on close,
  where knowledge was promoted, so the archived record still points at the
  durable version."* **The design has it; the procedure does not carry it.**
  That is an omission rather than an open question.
- **Leftovers as new work.** *"Re-home leftovers. Follow-ups not worth a doc →
  the backlog. Do not leave actionable items stranded in an archived folder."*
  Our close resolves tasks to terminal statuses and warns on unsuccessful ones,
  but nothing asks **what of this should become a new work item**. Their deepest
  project exercised exactly this, pushing two items out at close under a heading
  for what was deliberately out of scope.

**b. Record where each promotion landed.** *"As you promote each piece, note in
the final journal entry where it landed, so the archived project is a map to its
own residue."* One checkable instruction that turns a closed record into an
index of its own promotions, and the thing that would make §5.5's requirement
verifiable rather than aspirational.

**c. Read the whole record before closing.** Their first step, with the reason:
*"You need its full state to know what is durable and what is disposable."* Ours
opens with *write the journal entry first* — right priority, wrong starting
point, because what is promotable cannot be judged without re-reading the
explorations and decisions, which are precisely the files nobody opens at close.

**d. The downstream checkpoint as a procedure step.** Their final step: a closed
project is a natural known-good anchor, so **always weigh** cutting a release —
with the ordering constraint (only after the close-out is committed and pushed,
because the release tool needs a clean tree), the discretion boundary (patch at
the agent's discretion, confirm anything larger), the agent-specific path (a dry
run first to get the structural reference, then write the notes from it), and an
explicit note that this step **leaves local-edit mode and becomes
outward-facing**. We have `publish-release` as a skill with no link from close.

> **Worth noticing what that is.** `spec.md` §5.4 lists `work-item.closed` as a
> hook boundary whose typical use is *"promote decisions, mark things stale,
> archive, update references"* — near-verbatim. infra implements the whole
> thing as **prose steps in a skill, with no machinery at all.** That is working
> evidence for the cheaper alternative `open-questions.md` §22 keeps open
> against hooks, and it has been running in a real repository for months.

**Minor:** both skills declare blast radius in the body — *this is a local-edit
action, never a deploy.* Ours do not say whether a procedure writes outside
`.luma/`.

### The new-project skill — a negative result, and why

`project-new` yields nothing, and the reason is the interesting part. Its
central move is a fork — *is this ready to be worked now?* — branching to either
a parked backlog entry or a scaffolded project folder, on the grounds that *an
empty active project is just clutter*. Promotion across that fork then carries
fields over and **deletes the original entry** so it lives in exactly one place.

**Both mechanisms exist only because that backlog was not lossless.** §2.2.1
removes the fork by construction: capture is cheap, archiving is lossless, ideas
stay off the main board by default, and one record walks the ladder so there is
never a second copy to delete. `backlog-capture` is right not to ask, and this
is a case where the comparison confirms a design rather than challenging it.

The rest is covered or deliberately different: naming (our keys and derived
slugs), seeding first tasks at creation (ours sits behind the formation bar in
`backlog-refine`, on purpose), and the lean-no-stubs rule (we always create
`journal.md`, argued in §7.2).

**One small thing worth taking:** their descriptions end with an anti-skip
clause — *so do not skip this even if the user just says "archive X" or "mark X
done."* Our descriptions use negative routing (*do not use to…*), which is
better for disambiguation but does not defend against a terse request bypassing
the procedure entirely.

### The tension this record carries, stated plainly

**Finding 1 contradicts a settled position in our own specification.** §5.5
mandates *append, never curate* and *headings named after what they settle, not
drawn from a fixed template.* The journal we harvested the resume pointer from
now has a mutable, fixed-section head.

It also sits against §2.4 by way of BACK-0038, which weighs *put something in
the file* and leans away from it because anything written down can drift from
the outcomes it summarizes. infra's answer — **price the drift with a stamp
rather than forbid the block** — is a third option BACK-0038 does not consider.

So this is not only *is there anything to copy*. It carries evidence against a
rule this project has already decided, which is an argument for it being an
inquiry rather than a pile of ideas.

### Where this project is already ahead, so the comparison is fair

- **`blocked` as a flag rather than a status** (§4.1.1). infra puts a blocked
  reason inline on tasks, which is fine, but at project level a paused project
  lives in the archive alongside finished ones — conflating paused with done,
  and destroying exactly the information §4.1.1 protects.
- **Outcomes as records with evidence and verifiers; completion computed rather
  than declared**, and self-verification detected. Theirs has a success-criteria
  bullet list and honest prose — no arithmetic, and no refusal.
- **The conditions set** — `record.stale`, `journal.stale`, `drifted`,
  `not-converging`, `churning`. Their own repository is the argument for it: 37
  projects in the active directory, most untouched for two to three months, and
  one sitting in a soak window that was supposed to close eight weeks ago, with
  nothing surfacing any of it.

### Honest shortfall grading — close, but ours is stronger

One project's README grades its own success criteria **in place**: *partially
met*, the resilience drill deliberately moved to the backlog, *the project
closes without proving its central claim end to end — recorded rather than
quietly dropped.* Good practice, and BACK-0032 and BACK-0071 already cover the
mechanism better than prose does. **The one thing theirs adds is cosmetic but
real:** the shortfall is written on the record's own face, not only in a field
or a journal entry.

### How this relates to what exists

**Seven overlaps, no duplicates, and no record-level conflicts.**

- **BACK-0050** *Review the backlog procedures* — the nearest neighbour, and the
  seam is direction of travel. BACK-0050 reads our own procedures critically
  from the inside; this brings outside evidence against one of them. The
  close-procedure amendment is a finding BACK-0050 should **receive**, not own.
- **BACK-0107** *The journal earns a type definition* — decides what a journal
  is, what is noise, and how an entry is formatted. This supplies field evidence
  about which shape survived 1078 lines.
- **BACK-0038** *Where a work item stands is hard to see in the file* — asks the
  question and leans against putting state in the file. This carries a working
  counter-example and a third option.
- **BACK-0106** *A work item's decisions live with it and promote at close* —
  covers decision promotion. This covers the other promotion targets:
  documentation, and leftovers as new work items.
- **BACK-0033** *Closing a work item records what was learned* — the journal
  entry at close. This is where the knowledge goes **after** the journal, and
  the instruction that makes it checkable.
- **BACK-0051** *A retro skill, name undecided* — their close chaining into a
  downstream outward-facing act is a worked example for the naming-and-scope
  question BACK-0051 holds.
- **BACK-0084** *Improve the work item capturing skill* — narrow: `project-new`
  was examined and yields nothing, for a structural reason. A negative result
  worth recording rather than re-deriving.

Lighter touches: **BACK-0071** and **BACK-0032** (shortfall honesty, above),
**BACK-0064** (leftovers at close), **BACK-0100** (the tombstone as a worked
migration output), **BACK-0062** (journal quality).

### Two things this capture ran into

- **There is no search over record contents.** Finding these overlaps meant
  reading the listing for titles and then opening everything plausible, because
  several of them match on body rather than title. That is a gap rather than a
  limit.
- **The source material could not be quoted freely.** infra is private, and its
  journals carry hosts, addresses, key fingerprints, service names and
  credential locations. Every example above is genericized deliberately; anyone
  re-reading the originals should expect more detail than this record can hold,
  and should not copy that detail here.

## References

- `docs/spec.md` §5.5 — the journal, the resume pointer, and *append, never curate*.
- `docs/spec.md` §5.4 and `docs/open-questions.md` §22 — hooks, and the cheaper alternative.
- `docs/spec.md` §4.1.2 — `references`, and what this tool will not resolve.
- `docs/spec.md` §4.1.1 — what belongs in a status.
- `docs/spec.md` §2.2.1 — cheap capture, lossless archiving, ideas off the board.
- `docs/lifecycle.md` §2 — *Propagate*, the phase the closing procedure does not carry.
- `docs/format-requests.md` — where the resume pointer was harvested from.
- `.luma/bundles/lumastack/luma-catalog/backlog/procedure/backlog-transition.md` — the closing procedure this would amend.
