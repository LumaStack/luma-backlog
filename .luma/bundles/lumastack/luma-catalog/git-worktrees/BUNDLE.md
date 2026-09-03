---
type: bundle
title: lumastack/luma-catalog/git-worktrees
version: 0.8.1
published: 2026-09-02
stage: draft
consumers: [project]
description: Isolated worktrees for concurrent agents in one repository — where they live, what has to be provisioned, and how to tear them down without leaving wreckage.
---

# Git worktrees

Two agents editing one checkout stop working within minutes. A worktree gives
each its own directory and branch over one shared object database, and Git
enforces the part that matters: **a branch can be checked out in exactly one
worktree**, so a collision fails loudly instead of silently interleaving edits.

What Git does *not* do is make the new checkout runnable. **Only tracked files
come with a worktree**, and everything that makes a repository work — `.env`,
installed dependencies, generated config, local databases — is untracked. A
fresh worktree has none of it.

That gap is what this bundle closes.

## What is here

- [[worktree-isolation]] — where worktrees live, how they are named, what is
  shared, and what has to be provisioned. Read first.
- [[create-worktree]] — create one, provision it, verify it before starting.
- [[remove-worktree]] — tear it down completely, at merge.
- [[recover-worktree]] — reclaim one left by a crashed session or a failed
  setup: a lock nobody holds, a half-created checkout.
- [[repair-worktrees]] — the metadata states worktrees get stuck in, and which
  fix applies to which.

## The rules that eliminate the edge cases

**Default to a worktree without being asked.** Any task that edits or commits
starts in one. By the time a collision is visible it has already happened — an
edit swept into the wrong commit, a rename breaking another session's working
directory, an unexplained divergence nobody can trace.

**The shared checkout is not a workspace.** It holds `main` and worktrees are
created from it; nobody edits there, including the person supervising. It is the
one directory with no collision detector, and an edit made in it is invisible to
every worktree — so a merge reverts it with nothing in the diff to show why.

**Create them through the path that runs the whole lifecycle**, never a bare
`git worktree add`. That path provisions, locks while live, sweeps up, and
resumes; a bare `add` gives you a directory and none of it, and the missing
provisioning is silent.

**The slug is the identity** of the directory, the branch, the port offset and
every namespaced resource. One name, so nothing can drift out of sync.

**Ports derive from a hash of the slug, not from position** in `git worktree
list`. Position changes when any other worktree is removed, silently reassigning
the port of a process already running.

**Provision from `.worktreeinclude`**, in `.gitignore` syntax at the repository
root — and copy only files that match **and are already gitignored**. That
second condition means a tracked file can never be duplicated, whatever the
patterns say.

**Never invent a missing credential.** Stop and say so. A worktree that comes up
with a placeholder produces failures that look like bugs in the code, hours
later and somewhere else.

**One task, one branch, one pull request — merged serially.** Parallel merges
produce conflicts nobody caused: each agent was correct against the `main` it
started from, and only the second one through finds out otherwise.

**Scope every `git add`.** Never `-A` or `-u` in a shared checkout — it sweeps
whatever another agent wrote that second into your commit. Costs nothing alone,
so it is unconditional rather than something to switch on.

## Where "it just works" is not achievable

**Submodules.** They are not inherited, multi-worktree support is still
incomplete, and a project using them should expect rough edges. Said plainly
rather than papered over — a procedure claiming to handle them would be lying.

## Consumers

`project` only. Worktrees are a property of one repository's checkout.

## Corrections to published practice

Each is a failure that exists in tooling or guides in current use:

**Position-derived ports.** Several guides compute a port from a worktree's
index in `git worktree list`. That index shifts when any *other* worktree is
removed, silently reassigning the port of a running process.

**Decimal digits from a hash.** A widely copied setup script extracts digits
with `tr -d -c '0-9'` and does arithmetic on them. A result beginning with `0`
is read as octal, so any `8` or `9` is a fatal error — intermittent, and
determined by the branch name. Use hex with an explicit `0x`.

**Truncating a provisioned file.** The same script copies `.env.local` and then
writes the derived port with `cat >`, destroying the file it just copied. The
failure looks exactly like the copy never happened. **Append.**

## Version

`0.6.4` — **the manifest declares `lifecycle: draft`.** The field was absent, and
absent reads as `unknown` — *nobody has said*. Something was known: this is
developed by its maintainers for their own use, and its shape can reverse
without notice.

**Publication did not promote it.** Being reachable by somebody who did not
write it makes the question live rather than answering it, and the answer here
is *still a draft* — which is a legitimate thing to publish, and says more than
silence did.

Patch: a fact written down. Nothing an adopter is obliged to do has changed, and
`unknown` promised nothing that `draft` withdraws.

`0.6.3` — **references to the knowledge format name sections instead of numbering them.** The format removed section numbers, so every `§n` here pointed at a position that no longer exists — and a stale number resolves to the wrong section rather than to nothing, which is why none of them were reported. Decorative citations are dropped; the rest name what they meant.

Patch: wording only. No rule, field or procedure changed.

`0.6.2` — **`entry_point` is now `entrypoint`.** One word, so the same word names the same thing at every level it appears.

Patch: one key renamed. Same value, same meaning, same `optional` presence, and `luma-foreman` reads both spellings while the rename lands.

`0.6.1` — **bundle IDs in this catalog gained their namespace.** A bundle here
is `lumastack/luma-catalog/<name>` rather than `luma/<name>`, because the
namespace now derives from where the catalog lives instead of being declared.
Every reference in this bundle's prose is updated.

**A fork can no longer publish under this catalog's name.** It lives somewhere
else, so it is named something else, and its bundles sit beside these in a
project rather than colliding with them.

*Type names are unaffected.* `type: luma/catalog` and its siblings name the
format, not this catalog, and resolve separately.

Patch: nothing but the identifiers a reference points at.

`0.6.0` — **`applies_to` is now `matches`.** The old name obliged an author to
write a false sentence: `applies_to: everything` claims a rule governs
everything, and none does — what a rule governs is stated in its body, where no
frontmatter value reaches. The field says what makes a Document *surface*, which
is smaller and true, and it reads as a sentence in every form it takes: matches
`git commit`, matches always, matches nothing.

**The default reverses with it.** A Document that says nothing is now available
on request rather than loaded into every session. Nothing here is affected —
every rule in this bundle already states what surfaces it — but a rule that
genuinely should always be present now says `matches: always` rather than
staying silent and being treated as though it had.

Minor. Nothing a reader is obliged to do has changed; the field it is declared
in has been renamed, and `applies_to` is still read while the rename finishes.

`0.5.0` — **vocabulary.** `moment` becomes `event` — a moment is a point in
time and `applies_to` takes nouns. `compliance` is dropped wherever it was
saying nothing: a policy binds unless it says otherwise, so only a strong
default declares `recommended`, and a procedure's steps bind by being steps.
Type Definitions use `field_presence: required` for what was
`obligation: mandatory`, matching the format.

Minor. Nothing a reader is obliged to do has changed; what declares it has.

`0.4.0` — **`preload` is replaced by `compliance` and `applies_to`.** An author
now says how strongly a rule binds and when it governs; *when it is delivered* is
computed from those and never declared. Every rule here could state when it
applies, so **nothing in this bundle is loaded unconditionally any more** — it
arrives when the work matches and costs nothing before then.

Minor: a consumer reading `preload` finds nothing, and the loading behaviour of
every document changes.

`0.3.0` — **the manifest is `BUNDLE.md`.** Reserved markdown files are now
ALL CAPS across the estate, because nobody types all caps by accident: a file
becomes load-bearing only when somebody deliberately made it so, and writing
`bundle.md` now fails in the safe direction — ignored rather than silently wired
into machinery. Minor rather than patch, and pre-1.0 that is the tier for a
breaking change: anything naming the old path by hand stops resolving.

`0.2.0` — **the shared checkout is declared off limits for editing**, with the
check that catches it and the repair that does not lose the work.

The bundle had covered agents colliding with each other, which git reports, and
missed the case it cannot: **a second writer in the shared checkout collides with
nothing**, because that checkout sits on `main` and nothing else has it checked
out. The edit is invisible inside every worktree, and the merge reverts it
silently.

**Found by a person doing it** — reviewing alongside an agent, with a checkout
already open, editing the obvious file in the obvious place. The near-miss was a
rewrite that would have vanished at merge with nothing in the diff to explain it.

The repair is *move the work and say so*, deliberately: asking for it to be
redone loses whatever was better about the first attempt, and quietly working
around it makes the wrong tree the normal one.

`0.1.1` — a heading no longer says how many things are beneath it. Wording only.

Patch: no normative sentence moved and a reader who correctly understood
`0.1.0` behaves identically. See `writing-style` in `lumastack/luma-catalog/project-documentation`
for the rule and the failure it prevents.

`0.1.0`. Assembled from current practice, corrected where that practice is
fragile, and **run by no fleet of agents on a real project yet.** The submodule
section in particular is a placeholder rather than a solution.

## Sources

- [Git Worktree Isolation Patterns for Parallel AI Agent Development](https://zylos.ai/research/2026-02-22-git-worktree-parallel-ai-development/) — failure modes and per-worktree provisioning
- [git-worktree documentation](https://git-scm.com/docs/git-worktree) — `prune`, `repair`, `lock`, `move`
- [Git Worktrees for AI Coding Agents](https://nimbalyst.com/blog/git-worktrees-for-ai-coding-agents-complete-guide/)
- [Git Worktree: Pros, Cons, and the Gotchas Worth Knowing](https://joshtune.com/posts/git-worktree-pros-cons/) — branch exclusivity, stale metadata
- [Fix: Git Worktree Not Working](https://fixdevs.com/blog/git-worktree-not-working/) — submodules, locked worktrees
