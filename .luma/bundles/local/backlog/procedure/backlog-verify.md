---
type: procedure
title: Confirm an outcome holds
description: Record evidence that an outcome's desired state is true — use when work looks finished, before closing a work item as delivered, or when asked whether something is actually done. Do NOT use to assert that work was performed; verification is about the world being in the desired state, not about effort spent.
---

# Confirm an outcome holds

```
luma-backlog outcome verify <ref> -e "<what the confirmation rests on>"
```

**Verification accumulates.** Several actors confirming the same outcome is the
normal case, and a human entry raises the derived trust tier with no special
handling. Verifying again is not an error.

## What counts as evidence

**Evidence is what somebody else could check.** A command that was run and what
it printed, a file and what it now contains, a behaviour that was tried and what
happened. Write it so a reader who does not trust you could repeat it.

**"Tests pass" is usually not evidence.** Tests written by whoever wrote the code
are the weakest confirmation available — they assert what the author already
believed. Say which test, what it exercises, and why it would fail if the outcome
did not hold.

**Read `verify_by` first and do what it says.** The outcome states how it is to
be checked. Checking something else and recording that is how an outcome comes to
be marked passing on the strength of a different question being answered.

**If `verify_by` cannot be followed, say so and do not verify.** An outcome that
cannot be checked as written is a defect in the outcome, and the honest move is
to fix the outcome or record why it is unverifiable — not to substitute a check
nobody agreed to.

## The judgment

**An outcome is a condition, so the answer is true or false.** "Mostly", "for the
common case" and "except when" are not verifications; they are the discovery that
the outcome was written as work rather than as a state.

**Verify what is true now, not what will be true.** An outcome confirmed on the
strength of a change that has not landed is a record that will quietly become
wrong.
