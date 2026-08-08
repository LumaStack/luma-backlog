# Journal — first usable build

> The deliverable's memory. Newest entry first; everything below the top block is historical. Append, never curate. Shape: `SPEC.md` §5.5.

---

## ▶ 2026-08-08 — the backlog exists; what "earns a command" means

**Where things stand.** The specification is settled across units, record shapes, layout, concurrency, the policy boundary, scaffolding, containment, and testing. No Go code exists. `.backlog/` was created by hand today, which is itself the first act of dogfooding.

**Next, in order.**

1. Author the outcomes for this deliverable — they are the acceptance criteria for the first build.
2. Write the record-authoring skill. It doubles as the specification of what `new` must do, so it is not throwaway.
3. Scaffold the Go module only once something is annoying enough to justify it.

**Decided: what promotes a rule from skill to command.**

The obvious answer was friction — automate what annoys us. That is insufficient, and the reason is structural rather than a matter of degree: **the maintainer works alone, and the failures this tool exists to prevent are multi-actor.** Two agents diverging on record shape produces no friction for a single user. It would be felt only once it had become a corpus.

So three drivers, not one:

- **Friction**, where it appears. Necessary, not sufficient.
- **Divergence, provoked rather than awaited.** Give the same instruction to two agents with fresh context and diff the records. Any difference is exactly the failure the tool exists to eliminate, made observable long before it would surface naturally.
- **Invariants prose cannot hold.** Research into a mature system that measured this found compliance with prose-only rules running **11–22%**. Anything the skill must *guarantee* rather than suggest is therefore already a command — and no amount of friction would have revealed that, because the failures are silent.

**What was ruled out:** waiting for friction alone. Not because it is wrong, but because it is blind in precisely the environment this project is being built in. Re-open trigger: none needed — friction remains one of the three.

**Decided: skills and commands are a lifecycle, not a handover.**

Lead with a skill, backfill the command, then rewrite the skill to call it. The skill is never discarded — it settles into holding *when and why*, while the command holds *how*.

That is the same split the design already applies to itself (`SPEC.md` §5.0): the command is mechanism, the skill is the opinion about using it. It also inherits the same rule — **every command must work standalone**, because there is no privileged path (`PRINCIPLES.md`). A command that only works when driven by a skill would be an internal API wearing a public name.

**Open questions carried in.** Whether a repository-level journal accumulates anything (`OPEN-QUESTIONS.md` §2); whether tasks are worth storing at all (§18); default body sections (§17). All three are expected to answer themselves through this deliverable rather than through discussion.
