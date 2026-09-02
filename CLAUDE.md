# luma-backlog

A git-native backlog worked by people and agents at the same time. Records are markdown in `.backlog/`, conforming to the Luma Knowledge Format.

**Phase: implementation.** The specification is settled enough to build against. `.backlog/` is live and this project tracks itself in it — read it before starting work.

**Bootstrap order:** lead with a skill, backfill the command, then rewrite the skill to call it. The skill is not thrown away — it ends up holding *when and why*, while the command holds *how*. **Every command must work standalone**, with no skill involved: there is no privileged path (`principles.md`). What promotes a step from prose to command is **not friction alone** — the maintainer works alone, so multi-actor failures produce no friction here. Three drivers: friction where it appears, **divergence** where we provoke it (same instruction, two agents, diff the records), and **invariants prose cannot hold** — measured compliance with prose-only rules runs far below what a guarantee requires.

## Where things live

- `docs/development.md` — setup, tests, conventions. **Start here to work on the code.**
- `docs/principles.md` — the values decisions are argued against.
- `docs/spec.md` — the design. Normative.
- `docs/lifecycle.md` — the loop. Non-normative, expected to move to the workflow project.
- `docs/open-questions.md` — what is unsettled, and **why settled things were settled.** Read before reopening anything.

## Using the tool on this project

`.backlog/` is kept by the binary. Build it with `go build -o ./luma-backlog ./cmd/luma-backlog`.

**Set `LUMA_BACKLOG_ACTOR` before writing anything**, or every record you create is attributed to the machine's human owner:

```
export LUMA_BACKLOG_ACTOR="agent:<model>/luma-backlog"
```

Provenance is the point of that field. A record that says a person confirmed something they never saw is worse than one with no attribution at all.

## Working rules

- **Never name a competing project in committed output.** This is about rivals, not references — tools we depend on, borrow technique from, or interoperate with are named normally.
- **Never abbreviate terminology to initials.** Spell every phrase out.
- **American spelling.** *organization*, not *organisation*. Applies to prose, code comments, and field values alike.
- **Discuss before writing.** On anything exploratory — naming, modeling, structure — reach agreement first. Do not draft the answer and present it as the discussion.
- **Do not anchor on the maintainer's earlier research.** It contains ideas already abandoned. Reason from the problem instead.
- **Optimize for readability.** Short, unless clarity genuinely costs more.
- **Propose format changes freely, until this project reaches 1.0.** We are the knowledge format's first consumer and it is pre-1.0. Hitting a limit is evidence about the format, not a reason to contort around it — record the ask in `docs/format-requests.md` and, where sensible, ship ahead of it.
- **Record a path not taken as deferred with a re-open trigger, never as rejected.** *Rejected* reads as permanent and stops the option ever being raised again.

## Conventions

- Argue from the principles, not from what other tools do.
- Keep section numbers and cross-references consistent after any insert.
- **Do not treat an example as a decision.** Sample names and values in layout or configuration blocks show shape only. If a decision is needed, check `open-questions.md` or ask — reasoning from an illustration invents settled facts nobody agreed to.
- Commit messages say **why**, not what.
