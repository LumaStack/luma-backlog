---
type: procedure
title: Create a worktree
description: Create an isolated worktree for a task, provision what it needs to run, and verify it before starting work. Use before beginning any task that runs alongside another agent.
---

# Create a worktree

Every step here closes a failure that otherwise appears later as something
unrelated. Run them in order; do not skip verification.

## 1. Establish the slug

```sh
SLUG=<short-kebab-task-name>
REPO=$(git rev-parse --show-toplevel)
TREE="${REPO}.worktrees/${SLUG}"
BRANCH="agent/${SLUG}"
```

The slug is the identity of the whole thing — directory, branch, port offset,
container names. Keep it short, kebab-case, and derived from the task.

**Run this from the main checkout.** Creating a worktree from inside another
worktree works, but the relative paths below will be wrong.

## 2. Fail fast on collisions

```sh
git show-ref --verify --quiet "refs/heads/${BRANCH}" && echo "branch exists"
test -e "$TREE" && echo "directory exists"
git worktree list --porcelain | grep -q "^worktree ${TREE}$" && echo "worktree exists"
```

Any of these means another agent is already on this task, or a previous one did
not clean up. **Stop and resolve it** — do not append `-2` to the slug. A second
worktree for the same task is how two agents end up doing the same work and
merging conflicting versions of it.

## 3. Create it

```sh
git worktree add -b "$BRANCH" "$TREE"
```

This creates the branch from the current `HEAD` and checks it out in one step.
Branch from an explicit base when the default is not what you want:

```sh
git worktree add -b "$BRANCH" "$TREE" origin/main
```

## 4. Provision what the checkout needs to run

Read the project's `.worktreeinclude` — `.gitignore` syntax, at the repository
root — and copy every file that matches **and is already gitignored**:

```sh
[ -f "$REPO/.worktreeinclude" ] || echo "no .worktreeinclude — nothing will be provisioned"

(cd "$REPO" && git ls-files --others --ignored --exclude-standard) |
while IFS= read -r f; do
  [ -n "$f" ] || continue
  (cd "$REPO" && git check-ignore -q --no-index      --exclude-from=.worktreeinclude "$f") || continue
  mkdir -p "$TREE/$(dirname "$f")"
  cp -p "$REPO/$f" "$TREE/$f"
done
```

Three details, each closing a real failure:

- **`--exclude-standard` first** restricts candidates to files git already
  ignores, so a tracked file can never be duplicated no matter what the
  patterns say.
- **`cp -p`** preserves mode. Without it a `0600` key arrives world-readable —
  a security regression created by the provisioning step itself.
- **`mkdir -p`** so nested entries like `config/local.yml` work rather than
  failing on a missing directory.

**If a file the project needs is missing from the main checkout, stop and say
which.** Do not fall back to `.env.example` and do not invent values. A worktree
that comes up with placeholder credentials fails hours later, somewhere else,
looking like a bug in the code.

**Never copy dependencies or build output**, whatever the patterns say. They are
regenerated in the next step.

## 5. Install dependencies

```sh
cd "$TREE" && <the project's install command>
```

Prefer the offline or frozen-lockfile form. This is the slow step, and it is the
price of a checkout that is actually correct.

## 6. Claim a port and namespace anything shared

```sh
OFFSET=$(( 0x$(printf '%s' "$SLUG" | shasum | cut -c1-4) % 6900 ))
PORT=$(( 3100 + OFFSET ))
```

**Hex with an explicit `0x`, not extracted decimal digits.** The common recipe —
strip non-digits from a hash and do arithmetic — breaks when the result starts
with `0`, because the shell reads it as octal and any `8` or `9` is then a fatal
error. It works until a branch name happens to produce one.

Store it where the worktree can find it:

```sh
git config --global extensions.worktreeConfig true    # once, per repository
git -C "$TREE" config --worktree luma.port "$PORT"
```

If the application reads an env file instead, **append**:

```sh
printf '\nDEV_PORT=%s\n' "$PORT" >> "$TREE/.env.local"
```

**Never `>`.** Truncating a file provisioned in step 4 destroys it, and the
failure presents as though the copy never happened — a mistake present in more
than one published setup script.

Apply the slug as a prefix to every other shared namespace the project touches:
container names, compose project, database name, cache directory.

## 7. Initialise submodules, if there are any

```sh
git submodule update --init --recursive
```

They are not inherited. Multi-worktree submodule support is incomplete, so
verify rather than assume this worked.

## 8. Verify before starting work

```sh
cd "$TREE"
git rev-parse --show-toplevel        # this worktree, not the main one
git branch --show-current            # agent/<slug>
<the project's test command>
```

**A clean test baseline before the first edit is the highest-value step here.**
Without it, the first failure is ambiguous — a pre-existing break and a break
you just caused look identical, and distinguishing them costs more than this
step ever will.
