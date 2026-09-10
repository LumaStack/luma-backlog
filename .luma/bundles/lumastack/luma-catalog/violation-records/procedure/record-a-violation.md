---
type: procedure
title: Record a violation
description: File a violation record in under a minute — the directory, the six fields, and the one question that decides what it means. Use when an agent did something that was not wanted, whether or not a written rule covers it.
---

# Record a violation

**This should take under a minute.** A violation earns a line and a count, not
an investigation — and a procedure that costs more than that produces an empty
register, which is the only way this fails.

## 1. Make the directory

```sh
mkdir -p .luma/records/violations
mkdir .luma/records/violations/$(date -u +%Y-%m-%d-%H%M%S)-<short-name>
```

`<short-name>` is two to five words, kebab-case, naming **what was breached**
rather than what happened — `wikilink-across-bundles`, not `agent-made-mistake`.

**The second `mkdir` has no `-p`, deliberately.** If that directory exists,
somebody has already filed this breach this second and the create must fail
rather than succeed silently — the policy explains why a collision here is the
feature and not a defect. It is rare at this resolution and free to keep. Add
nothing to the name to get past it; go and read the record that is already
there.

## 2. Answer the one question that matters

**Was the rule there, and did it reach the actor?**

| answer | `delivery` |
| --- | --- |
| the rule existed and was in context | **`delivered`** |
| the rule existed and never arrived | **`undelivered`** |
| there was no rule — this was wanted and never written down | **`unwritten`** |

**Check before answering `delivered`.** *The project has adopted that bundle* is
not the same as *the document was in front of the actor* — a policy with a
`matches` that never fired is `undelivered`, and that is a different bug with a
different owner. Look at what was actually loaded.

**`unwritten` is a full answer, not a gap.** Most first-time filers reach for the
nearest almost-relevant rule rather than admitting there was none. **Do not** —
a wrong `policy` value is worse than an absent one, because it gets counted.

## 3. Fill the record

Copy [the template](../templates/violation.md) to `violation.md` in the
directory you just made. Six fields and two short sections.

**`expectation` is written as a rule, in one line**, even when `delivery` is
`unwritten` — especially then, because that line is the first draft of the policy
this violation is asking for. *Wikilinks never point outside their own bundle*,
not *the agent linked wrongly*.

**`violating_actor` names the model**, not just `agent:`. The register is read per model,
and a version that behaves differently is a thing worth seeing.

**`policy` carries the version that was in force**, not just the document —
`local/backlog procedure/backlog-move 0.36.0`. A citation without one points at
a moving target, and stops being checkable the first time somebody rewords the
rule. Take it from the bundle's own `BUNDLE.md`.

**`violating_commit` is the commit the breach is embodied in** — a bare SHA.
**A violation has two commits**, and this is not the other one: the commit that
carries this record is later, and naming it here would say nothing.

The act itself is usually invisible — nobody can see a check that was not run —
so the commit is where it becomes reviewable, which is what makes it worth
citing. Leave it out where no commit contains the breach.

**Where the breach is visible in work you are about to commit, commit that
first, separately.** Step 5 says to commit a violation with the work that
produced it; where the work *is* the breach, following that would make this
field point at its own record.

**`noticed_by` is honest about who caught it.** If you are the actor recording
your own, say so — a register where nothing external ever appears is itself the
finding.

## 4. Put evidence beside it, not inside it

Anything longer than a few lines — the turn, the diff, the output — goes in the
directory as its own file. The record stays scannable, which is what makes forty
of them readable.

## 5. Stamp `created_using` and commit

```sh
luma-foreman bundle show violation-records
```

**Only you can answer it.** The catalog cannot know what you adopted, so the
template ships a placeholder rather than a value that would be wrong for
somebody by construction. Nothing rewrites it afterwards, including a migration.

**Commit it with the work that produced it** where you can. A violation committed
alongside the change it happened during is far easier to reconstruct later than
one filed in isolation.

## 6. Stop

**Do not investigate, grade, or assign it.** There is no severity and nothing to
resolve. If it seems to warrant a postmortem, that is the signal it was an
incident — see the `incident-records` bundle, and file both.

**Do not batch them either.** Four violations in a session is four records, not
one summary — the count is the product, and a summary destroys it.
