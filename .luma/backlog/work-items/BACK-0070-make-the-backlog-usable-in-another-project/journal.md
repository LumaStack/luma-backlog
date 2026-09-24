# Journal — Make the backlog usable in another project

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-23

### Bundle and binary sync — settled as prose and waiting, and it changes this delivery's shape

**The precondition list here says "the bundle leaves `local/`". That is still true and it is no longer the whole ask** — the discussion concluded the bundle should be **shrunk before it is published**, not published as it stands. Reasoning is in `docs/open-questions.md` §25; this is what it changes for this record.

**The arrangement makes lockstep unreachable, and that is the load-bearing fact.** One binary per machine, one vendored bundle per project, on different update paths. Skew is the steady state rather than an edge case, so the goal is not preventing drift but shrinking what can drift.

**What the bundle should carry after the shrink: only what the binary cannot emit.** Anything the tool prints or enforces becomes a pointer rather than a copy — the bootstrap order (skill holds when and why, command holds how) scaled up to the artifact. `showing-records` already carries a note saying this about its own state marks, so the first instance is identified.

**Three directions deferred, each with a trigger**, so none of them gets re-argued from scratch:

- **A declared tool-version floor checked by the binary.** Right home if ever needed, because the binary is the only thing that sees both the vendored bundle and its own version, per project. *Reopen when* a published procedure reaches for a command an installed binary lacks and the refusal is not enough to recover from.
- **Lockstep distribution — binary embeds and emits the bundle.** Collapses the opt-out that `principles.md` protects (command usable without the bundle) and puts a second distributor beside the one that already vendors. *Reopen if* re-adoption proves unreliable in practice.
- **Pinning the binary per project.** The actual fix, disproportionate while one person owns both ends. *Reopen when* somebody outside this repository runs a different version.

**A mistake worth recording: the first answer given was wrong.** The initial recommendation was a version floor on the bundle plus a warning from the binary. That was abandoned once the failure modes were separated by whether they are loud — procedures that invoke commands fail loudly and the refusals merged in PR #153 (`parseRefusal`, near-miss suggestions, `See every command with`) already recover from it better than any version warning could. Only the quiet failures needed anything, and those are removed by deleting duplication rather than detected. **The version-floor instinct was machinery for a problem the error path already solved.**

**Consequence for the other open question here: bundle version and tool version should not be the same number.** After the shrink they are not describing the same thing, and the bundle carries no floor on the binary.

**Nothing above should be built yet.** No second project has ever held a stale bundle, so the promotion test — friction, divergence, or an invariant prose cannot hold — is unmet on all three counts. Settled by publishing and watching what actually breaks.

## ▶ 2026-09-20

BACK-0070-make-the-backlog-usable-in-another-project captured → unprepared: selected

## ▶ 2026-09-10

BACK-0070-make-the-backlog-usable-in-another-project unprepared → captured: sent back to the pile so there is one place to choose from; this had crossed the first gate and that decision is being re-made rather than lost

## ▶ 2026-09-08

created knowingly in the wrong shape — this is a delivery and work items are all we have, which is what WORK-0069 exists to fix; recorded on the record itself so a reader does not mistake it for a piece of work
the order is forced by WORK-0037: the window for migration being a one-repo problem closes on first use elsewhere, not on a date — so the breaking changes finish before anybody else starts, or every later shape change becomes a distributed migration on data we cannot see
crossed the first gate — this will become work, and what crossing commits to is working out what it is rather than doing it; the pieces are already recorded, what is unworked is the delivery itself
the second item is done — --open ships, so a stranger can ask what is open in their first hour rather than piping --json through a second language
