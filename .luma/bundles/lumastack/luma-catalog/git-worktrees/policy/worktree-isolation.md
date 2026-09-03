---
type: policy
title: Worktree isolation
description: Where worktrees live, how they are named, and what is shared versus isolated — so concurrent agents in one repository can never collide.
matches: eager
---

# Worktree isolation

Several agents working one repository at once need real isolation, not
discipline. A worktree gives each its own directory and branch over one shared
object database — but only the *tracked* files come with it. **Everything that
makes a checkout runnable is untracked, and none of it arrives on its own.**

These rules exist so no situation requires judgement. Where one looks excessive,
it is closing an edge case that has bitten somebody.

## Default to a worktree, without being asked

**Any task that edits or commits starts in its own worktree.** Not when someone
announces parallel work — always, because by the time a collision is visible it
has already happened: an edit swept into the wrong commit, a rename in one
session breaking another's working directory, a mystery *ahead 4* with no
explanation.

**Opt out only for** read-only or answer-only work, where there is nothing to
collide, or a genuinely throwaway edit when you are *certain* no other agent is
in the repository.

The overhead is small and the failures are not. A clobbered edit costs more to
diagnose than every worktree you will create that week — and it costs it later,
in a session that has forgotten the cause.

## Create them through the path that runs the whole lifecycle

**Not a bare `git worktree add`.**

Whatever creates worktrees for you — a harness command, a project script — is
not a convenience wrapper around `git worktree add`. It runs the steps that make
the checkout usable: provisioning from `.worktreeinclude`, locking while a
session is live, sweeping up on exit, resuming an interrupted one.

A bare `add` gives you a directory and none of that. **The missing provisioning
is silent**: the worktree comes up looking correct and fails on the first
command that needs an environment file, by which point the cause is several
steps behind you.

If nothing provides this, write the script once. The point is one path that
always runs every step, not which tool runs it.

**Worktrees live in a gitignored directory inside the repository**, wherever
that path-of-record puts them.

Keeping them outside looks safer and forfeits the lifecycle, which is a bad
trade. The risk it appears to remove — tools walking the tree and finding *n*
copies of the codebase — **belongs to those tools**. Anything scanning a
repository should ask git rather than walking the filesystem, or it will find
`node_modules` and build output too. Fixing one scanner is cheaper than
constraining every project's layout, and it fixes the other cases at the same
time.

## One branch, one worktree, one agent

**Git enforces the first part**: a branch can be checked out in exactly one
worktree, and a second attempt fails. That is a feature — it is the collision
detector — but only if branch names cannot accidentally coincide.

```
agent/<task-slug>          branch
<repo>.worktrees/<task-slug>/   directory
```

The slug appears in both, so a directory maps to a branch by inspection and
neither can drift. Namespacing under `agent/` keeps them out of the way of
branches people create by hand.

**An agent must be able to tell where it is without asking:**

```sh
git rev-parse --show-toplevel      # which worktree
git branch --show-current          # which branch
git rev-parse --git-common-dir     # the shared object database
```

## The shared checkout is not a workspace

**It exists to create worktrees from and to hold `main`. Nobody edits in it** —
not the agents, and not the person watching them.

**Because there is no collision detector on it.** Git's one-branch-one-worktree
rule is what makes concurrent work safe, and it does not apply here: the shared
checkout sits on `main`, which nothing else has checked out, so a second writer
in that directory collides with nothing and is reported by nothing.

**And a worktree is branched from the committed state.** Edits sitting
uncommitted in the shared checkout are invisible inside every worktree. The
sequence that follows is not exotic:

1. Somebody edits a file in the shared checkout.
2. An agent, in a worktree, edits the same file — reading the committed version.
3. The worktree's branch merges, reverting the first edit **with nothing in the
   diff to show it happened.**

Nobody sees a conflict, because to git there was never a conflict. This also
breaks the `--ff-only` rule above: a shared checkout with local edits cannot be
kept cleanly on `main`.

**People are the likely offender, and it is not carelessness.** Agents are told
to work in a worktree. The person reviewing alongside them has a perfectly good
checkout already open, and editing it is the obvious thing to do.

### Check the shared checkout before every commit

Not only when something looks wrong. **The failure is silent, and by merge time
the evidence is gone.**

```sh
MAIN="$(git rev-parse --git-common-dir)/.."
git -C "$MAIN" status --short --untracked-files=no
```

Anything listed that the committing session did not put there is misplaced work.

**`--untracked-files=no` is what makes this usable.** The shared checkout always
has untracked entries — the worktrees directory itself, editor state, build
output — and a check that cries wolf on every run is one people stop reading.
Modified tracked files are the case that reverts silently, so that is the signal.
**Untracked files in the shared checkout are worth a separate look**, less
urgently: they are not in any diff, so they cannot be reverted by a merge, but
they also will not travel.

### Move it, then say so

**Never ask for the work to be redone, and never quietly work around it.**
Redoing it loses whatever was better about the first attempt; working around it
makes the wrong tree the normal one.

```sh
MAIN=$(git rev-parse --git-common-dir)/..
git -C "$MAIN" diff > /tmp/misplaced.patch     # tracked edits
git apply /tmp/misplaced.patch                 # from inside the worktree
```

**Verify the content arrived before resetting anything.** Untracked files are not
in that diff and have to be moved separately. Then clear the shared checkout, and
**say plainly what happened and where to edit next time** — a repeat is the same
silent revert, and the whole value of detecting it is that somebody finds out.

## What is shared, and what is not

| | |
| --- | --- |
| **Shared** — one copy, all worktrees | object database, refs, remotes, `config`, hooks, stash |
| **Per-worktree** — separate, and empty on creation | working files, index, `HEAD`, **everything untracked** |

That second row is the whole problem. `.env`, installed dependencies, build
output, local databases and generated config are all untracked, so a new
worktree has **none of them** and will not run until something puts them there.

## Provision from `.worktreeinclude`

A project declares what a worktree needs in a **`.worktreeinclude`** at its
root, in `.gitignore` syntax:

```gitignore
.env
.env.*
!.env.example
config/local.yml
certs/*.pem
```

**Only files that match a pattern *and* are already gitignored are copied.**
That second condition is the safety rule and it is doing real work: a tracked
file can never be duplicated into a worktree, so the pattern list cannot cause
two copies of something git is managing.

**Create worktrees from a checkout that is on `main`:**

```sh
git checkout main && git pull --ff-only
```

`--ff-only` refuses rather than quietly creating a merge you did not ask for, so
a source checkout that has drifted announces itself instead of being repaired
behind your back.
 The provisioning list is
read from the *source* checkout you invoke the command in, so creating from a
stale branch that predates `.worktreeinclude` carries **nothing** — silently. The
new worktree comes up with no environment at all and the first failure looks
like a broken tool. Keep the shared checkout on `main` and branch task worktrees
off it.

Everything else is absent by design. **Dependencies are reinstalled, not
copied** — `node_modules`, `.venv`, `dist`, `target` and build caches are
frequently invalid in a new path, and some embed the absolute directory they
were built in. Slower once, correct always.

Copy with **mode preserved** and parent directories created. A `0600` key that
arrives as `0644` is a security regression introduced by the provisioning step
itself, which is the worst place to have one.

## Secrets are copied deliberately, and never widened

Copying `.env` into a sibling directory does not meaningfully increase exposure
— it is already unencrypted on the same disk, owned by the same user. The risks
are different:

- **A worktree outliving its task**, holding credentials nobody remembers.
  Removing a worktree the moment its branch merges is the mitigation, and it is
  why teardown is a step rather than a habit.
- **A secret reaching the repository** because a worktree sat inside it. Ruled
  out by the location rule above.

**Never generate or invent credentials for a worktree.** If a required file is
missing from the main checkout, stop and say so. A worktree that silently comes
up with a placeholder produces failures that look like bugs in the code.

## Anything holding a port or a name needs a per-worktree value

Two agents running a dev server on the same port is the most common collision,
and it surfaces as an unrelated error.

**Derive it from the slug, not from position in the worktree list.** Position
shifts the moment any other worktree is removed, silently reassigning the port
of a process already running.

```sh
OFFSET=$(( 0x$(printf '%s' "$SLUG" | shasum | cut -c1-4) % 6900 ))
PORT=$(( 3100 + OFFSET ))
```

**Hex, with an explicit `0x`.** Extracting decimal digits from a hash and doing
arithmetic on them is a common recipe and it is broken: digits beginning with
`0` are read as octal, so any hash yielding `08…` or `09…` fails outright. It
works until a branch name happens to produce one.

The same prefixing applies to every shared namespace — container names, compose
project, database name, cache directory. Anything two agents could both claim
takes the slug.

**Where the derived value lives.** Git has per-worktree configuration, gated
behind an extension enabled once per repository:

```sh
git config --global extensions.worktreeConfig true   # once, per repo
git config --worktree luma.port "$PORT"
```

That is the right home for a value that must differ per worktree. Where the
application reads an env file instead, **append — never rewrite.** Truncating a
file that was just provisioned destroys it, and the failure looks like the copy
never happened.

## One task, one branch, one pull request — merged serially

**One worktree per task, one branch per worktree, one pull request per branch.**
Never two agents in one directory; git prevents the branch case and nothing
prevents the directory case but this rule.

**Merge them one at a time.** When several agents finish together, parallel
merges produce conflicts that none of them caused and none can resolve alone —
each was correct against the `main` it started from, and only the second one
through discovers otherwise. Serialising makes every merge a fast-forward or a
clean three-way, and it costs a few minutes of queueing.

**Pull before you push.**

```sh
git pull --rebase origin main    # your own unpushed commits, onto current main
```

Without this every push is a race, and the loser resolves a rejection under time
pressure with a worktree full of state.

*Rebasing your own unpushed commits is not the same thing as rebase-merging a
pull request.* The first replays work nobody has seen; the second rewrites
commits others may already reference, and breaks merged-detection. Do the first,
never the second.

## Scope every `git add`, even working alone

**Never `git add -A` or `git add -u` in a shared checkout.** Name the paths:

```sh
git add src/thing.py docs/thing.md
```

`-A` sweeps whatever another agent happens to have written that second into your
commit. The result is a commit containing someone else's half-finished work,
attributed to your task, and discovered much later by whoever is confused by it.

This costs nothing when working alone and it is the single cheapest guard
against the failure, so it is unconditional rather than something to remember
when parallel work starts.

## Know what runs from a worktree and what does not

A worktree is for **editing, committing and pushing**. Operations that act on
the outside world — deploying, publishing, anything touching shared
infrastructure — should run from the canonical checkout unless you have
deliberately made them worktree-safe.

The failure this prevents is subtle: an operation that reads its configuration
from the current directory picks up a task worktree's provisioned copy, which
was derived for isolation rather than for production. It will usually work,
which is what makes it dangerous.

Decide this per operation and write it down, because the answer is not
guessable from the command.

## Cleanup is part of the task, not a periodic chore

**Remove the worktree when the branch merges.** Not weekly, not when disk fills
— at merge, as the last step of the work.

Two things surprise people, and both are deliberate:

**Removing a worktree never deletes its branch.** Git will not discard history
as a side effect of cleaning a directory.

**Deleting the directory by hand leaves metadata behind** in the shared git
directory, and that entry keeps the branch locked to a worktree that no longer
exists. `git worktree remove` handles both; `rm -rf` requires
`git worktree prune` afterwards.

## Submodules do not come along

A new worktree does not initialise submodules, and multi-worktree submodule
support is still incomplete. A project using them must initialise per worktree
and should expect rough edges — this is the one case where "it just works" is
not achievable and the honest answer is to say so.
