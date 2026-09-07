---
type: work-item
key: WORK-0031
title: Reshape the command surface
workflow_status: prepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T14:00:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:51:08Z'}
---

# Reshape the command surface

## The problem

The specification and the binary disagree, and the specification won
([[backlog/work-items/WORK-0017-specify-the-minimum-viable-product]]). **The
code is what changes.**

Split out of
[[backlog/work-items/WORK-0018-extract-the-application-layer]] deliberately.
That work was behaviour-preserving and proved it by leaving every golden file
untouched. **Everything here moves the golden files**, so mixing the two would
have meant an unchanged suite could no longer prove the extraction was clean.

## What is being delivered

**Noun then verb** (`spec.md` §9.1, ADR-0006). `work-item close`,
`outcome verify`, `task new` — with verb-only where no noun applies: `init`,
`board`, `contract`, `config`, `check`, `log`, `serve`.

**The changed command shapes**, each settled and none built:

| Command | Change |
| --- | --- |
| `work-item rank` | Replaces `move`; `set` refuses the rank field (ADR-0005) |
| `work-item close <ref> <as>` | Disposition becomes positional; `--reason` becomes prose; `delivered` → `completed`; `abandoned` dropped, `rejected` added (ADR-0007) |
| `outcome assert <ref> <as>` | New — the doer's claim |
| `outcome verify <ref> <as>` | Verdict becomes a required positional |
| `outcome archive` | New — what `close`'s refusal already tells people to do |
| `work-item new` | Accepts a positional title **and** `--title`, erroring if both |

**Reference resolution**
([[backlog/work-items/WORK-0023-refer-to-a-record-by-the-path-a-person-would-type]]).
Key-scoped paths — `WORK-0017/outcomes/<slug>` — accepted and emitted; bare
names accepted only while unambiguous. It belongs in `internal/app`, because
every surface needs the same forms.

**Completion in the read path.** `CompletionOf` has one caller and it is the
one that refuses; `show` and `list` carry it as counts (`spec.md` §9.3).

**Six exit codes, not seven.** `6` stays reserved until taking ships (ADR-0006).

## Split out, and why

**[[work-items/WORK-0043-work-the-backlog-from-a-terminal]]** took five tasks:
showing a work item's whole state, filtering by column, sorting and grouping,
paging, and colouring the marks.

**None of them was settled anywhere.** This work item reshapes commands whose
shape a decision already fixed; those five were found by writing the procedures
and discovering the reading commands were too thin to write against. They are
`spec.md` §11.2's views arriving in the terminal, which is a different piece of
work with a different argument.

**[[work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]]** took
`rank repair`. It is ADR-0005's third mechanism, and what it repairs is drift
caused by editing the status vocabulary --- which is that work item's subject
rather than this one's.

**This work item began with nine tasks and reached twenty-three**, eleven of
them discovered by doing it. That is the reshape working as intended and also
the reason to split: a work item that absorbs everything found while doing it
never closes.

## Out of scope

**The board.** This is the surface it will be built against, not the board
itself.

**Prompting.** No `--prompt`, no `--no-input`, no configuration key — the design
is recorded on ADR-0006 for when it is wanted.

**Taking.** `take` / `release` / `steal` do not ship (ADR-0008).

## Constraints

- **The golden files move, and that is the point.** Every diff in them is a
  breaking change being made deliberately (`spec.md` §9a.5), so each should be
  reviewed rather than regenerated.
- **Every departure is already argued.** Nothing here is a fresh design
  decision; where the shape looks wrong, the record that settled it is the
  place to argue.
- **Adapters still may not import the engine.** The containment test in
  `internal/policy` holds throughout.
- **`--json` shapes are contract** (`spec.md` §9.3, §9.9). Additions are free;
  changes are breaking and there are several here.

## References

- `docs/design/mvp.md` — the command surface, in one table.
- `[[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]]`
- `[[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]`
- `[[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]]`
