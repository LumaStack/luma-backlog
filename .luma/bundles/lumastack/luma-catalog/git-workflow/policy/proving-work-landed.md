---
type: policy
title: Proving work has landed
description: The commands that answer "is this landed" and "is anything stranded", why they run against the remote ref, and why a report is not an answer.
matches:
  - topic: checking whether work has landed
  - topic: finishing a piece of work
---

# Proving work has landed

**A commit is not a landed change.** It satisfies every sense of *written down*
— the file exists, git has it, nothing is lost — and it is still **one
`checkout` from invisible** to anybody working from the integration branch.

**The gap between those two is where work disappears**, and it disappears
quietly: nothing reports a branch that was never merged.

## Ask git; do not remember

**Whoever did the work is the worst witness to whether it landed.** They
committed it, so they recall it as done — a sincere report of an intention
rather than an observation of a state.

**So the answer is a command's output, not a recollection.** Four of them,
returning in milliseconds:

```sh
git fetch -q
git status --short                              # uncommitted work
git log --oneline origin/<integration>..HEAD    # committed, not landed
git branch --no-merged origin/<integration>     # stranded elsewhere
git worktree list                               # another checkout holding work
```

**Empty output is the answer you want from the first three.** `git worktree
list` shows the one you are in and nothing else holding unfinished work — see
`git-worktrees` for what to do when it does.

## Fetch first, and compare against the remote ref

**A local integration branch is stale the moment somebody else merges.** Against
a stale `main`, `git log main..HEAD` counts commits that are already landed and
`git branch --no-merged main` names branches that were merged days ago.

**Both are false positives, and a false positive is worse than no check.** It
teaches the reader that the check cries wolf, and the next real finding gets
waved through with the noise.

**`git fetch -q` costs a network round trip and removes the whole class.**

## Show the output; a summary of it is a report again

**Running the check privately and saying *verified* is the same failure one
level down** — the reader is back to trusting a recollection, only now it is a
recollection of having looked.

**Put the output where it can be read.** Three lines of `git status --short`
saying nothing is cheaper for everybody than a sentence claiming the same.

## A check with nothing gated on it is a suggestion

**Gate in prose by default**: *do not do X until the output above is shown.* It
costs a sentence and applies everywhere, and it is exactly as strong as the
willingness to follow it — **so put it where the violation happens**, not beside
the rule it enforces. The check for unlanded work belongs at the moment the next
branch is cut, not in the section about finishing.

**Gate in a hook where the failure earns it.** A hook is real machinery:
installed, kept working, and met by somebody who has to understand it. **Spend
it where a failure is both silent and expensive** — silent because nothing else
reports it, expensive because it surfaces late and wrong.

**Most failures are neither**, and a hook for them costs more than they do.

## This is why ancestry has to work

**Every command above is an ancestry question**, which is what
[[merge-commits]] is protecting. Under squash or rebase merges, `git branch
--no-merged` names branches that did land and `git log origin/main..HEAD`
counts commits that are already in — **the proof stops being a proof**, and the
only remaining answer is `git cherry` and patch-id comparison by hand.

**A merge-commit history is what makes these four commands cheap enough to
run every time.**
