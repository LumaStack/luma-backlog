---
type: policy
title: Adopting a rule the corpus does not meet
description: What to do when a new rule would fail records that already exist — backfill them or grandfather them, never neither, and how to choose.
matches: topic:adopting a rule, tightening a policy, or changing what a record must contain
---

# Adopting a rule the corpus does not meet

**Almost every rule worth adopting is one the existing corpus fails.** That is
not an argument against the rule. It is the moment the rule is decided, because
what happens to the records already written is decided at the same time or not
at all.

## There are two options, and doing neither is the default

**Backfill** — bring the existing records up to the rule.
**Grandfather** — declare the rule applies from here, and say so on the record.

**Picking neither is what actually happens**, and it is the failure this policy
exists to prevent. A rule adopted over a corpus that fails it is decorative:
every reader sees records violating it, concludes it is aspirational, and the
next rule is read the same way. **One unenforced rule devalues the others**,
which is why this is not a matter of tidiness.

**Observed:** the two-readers comprehension test was proposed for the first
gate on 2026-09-07, while six of the seven newest work items had an empty
`## The problem`. The rule and its counterexamples were written the same week.

## Lean backfill

**Backfill unless there is a reason not to.** A corpus that satisfies its own
rules can be trusted by a reader who has not read the rules, which is most
readers most of the time.

**Preference order, and the difference is reproducibility:**

1. **Mechanical and reproducible** — a migration that can be re-run and
   converges. `spec.md` §9.6 holds the guarantees — history never partial, the
   working tree possibly so, re-running converges.
2. **A model, given written instructions.** Legitimate, and the instructions are
   the artifact — **keep them with the change.** A model run nobody can repeat
   is a one-time edit wearing a migration's clothes, and the instructions are
   the only thing that makes it inspectable afterwards.

**Never automatic, and `--dry-run` first.** A corpus rewrite that happens because
a binary was upgraded is the least recoverable thing available, and both rules
hold however the backfill is performed.

## Grandfathering is legitimate and must be explicit

**Say it out loud, in the same change that adopts the rule.** Which records are
exempt, and why. An unrecorded grandfather clause is indistinguishable from a
rule nobody enforces — that is the whole reason it has to be written rather
than understood.

**A grandfather clause with no end is a permanent second class of record**, and
a reader has to be able to tell which class they are holding. If the exemption
is meant to expire, say what ends it.

## Open

**When grandfathering beats backfilling is not settled**, and neither is who
decides. Cost is the obvious axis and probably not the only one — a rule about
what a record *means* may be unbackfillable at any price, because nobody can
reconstruct what an author intended.
