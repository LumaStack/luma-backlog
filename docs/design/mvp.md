# Design MVP

See the [vision.md](vision.md) first to understand where we will end up so we can have an expandable open ended MVP.

## First features

First features that make this tool usable.

    CAPTURE
    Capture things we may need to do
        ↓
    TRIAGE
    Understand and classify it
        ↓
    RANK
    Determine its relative importance
        ↓
    ADVANCE
    Move through workflow status
        ↓
    EXECUTE
    Create tasks and mark outcomes as complete
        ↓
    RESOLVE
    Reach and record an outcome

The first loop we should strive towards is:

    Capture work → prepare it → prioritize it → collaborate on it → execute it → prove it → monitor it → report on it → learn from what happened.

### Capture (MVP)

The important idea is that captured work shouldn't need to be well-formed yet. Make capture extremely cheap.

Minimum capabilities:

- Create work item
- Short title
- Freeform description
- Creator
- Created timestamp

Ideally this includes (nice to have):

- Quick capture

### Triage (MVP)

Triage turns raw input into something the organization understands.

| Triage choice | What happens |
|---------------|--------------|
| Keep          | Remains a work candidate, moves past captured |
| Defer         | Remains captured, not worth focus right now |
| Reject        | No longer visible as captured, does not destroy |
| De-duplicate  | Combine, supercede, or link work items |

Minimum capabilities:

- Review captured work in a list and by opening (show) it
- Edit work items
- Ability to Reject or Defer work items
- Change workflow status
- CRUD outcomes
- CRUD tasks

Nice to have but probably skipped for MVP:

- De-duplicate

### Rank (MVP)

Don't confuse priority with rank.

Eventually we will want both priority and rank, let's start with rank.

    Priority: Critical / High / Normal / Low
    Rank: #1 → #2 → #3 → #4

Minimum capabilities:

- Rank by moving work items up and down in a list
- Basic filtering by Workflow status

### Advance

Minimum capabilities:

- Display work items in 2-3 columns grouped by workflow status
- Ability to move the workflow status columns forward and backwards (left and right)
- Ability to highlight a work item and move it left or right, using `M` key (by default)
- A moved work item should remain highlighted after moving it, press `M` again or `Enter` to set it back down
- Moving your cursor to a different workflow column should return to the work item that you had previously selected OR directly below the last work item that was moved out so you don't lose your place where you were at - we want to avoid making moving multiple items annoying because the column jumps to the top each time you enter it after moving.
- I think this means we'll need to save where you were last in each column in some kind of interface cache, and respect it as much as we can - within reason.  And if some external action moves items out of a column, we should approximate where we were.  Keep it simple for MVP and maybe this will become more perfect over time. (nice to have for MVP)

### Execute

- CRUD tasks
- CRUD outcomes
- Mark outcomes as successful or failed
- Mark outcomes as verified and by whom

### Resolve

Minimum capabilities:

- Mark work item complete
- Basic validation 
- Ensure no duplicate work items exist upon completion
  
## The command surface

Everything the first release ships. **This is the scope; `spec.md` §9 is what
each verb means.** A command absent here is not absent from the design — it is
absent from the first cut.

**Noun then verb**, and verb only where no noun applies
([[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]]).
Marked **new** where nothing exists yet, **changed** where a shipped command's
shape moves.

### On every command

| Flag | Does |
|---|---|
| `-h`, `--help` | Help. Examples first. |
| `--version` | Version. Never `-v`, which is ambiguous. |
| `--json` | The machine contract (`spec.md` §9.3). Formatted. |
| `--plain` | One record per line, for `grep` and `awk`. **new** |
| `--no-color` | Also honors `NO_COLOR`, `TERM=dumb`, and a non-terminal. **new** |
| `-q`, `--quiet` | Suppress non-essential output. **new** |

**No command prompts.** Missing input is a usage error naming what was needed.
There is no `--prompt` and no `--no-input` in the first release; the design for
adding them is on ADR-0006.

### Work items

| Command | Does | Arguments | Flags |
|---|---|---|---|
| `work-item new` | Create | title, positionally or `--title` | `-k/--kind`, `--title` **new** |
| `work-item show` | Read one | reference | |
| `work-item list` | Read many | | `-s/--status`, `-k/--kind`, `--open` |
| `work-item set` | Change fields | reference, `field=value` | `--unset`, `--if-unchanged` |
| `work-item edit` | Open in an editor | reference | **new** |
| `work-item rank` | Reorder | reference | `--before`, `--after`, `--top`, `--bottom` **new** |
| `work-item close` | End it | reference, `completed\|rejected\|canceled\|superseded` | `--reason` (prose), `--force` **changed** |

**`rank` replaces `move`**, and `set` refuses the rank field — the caller never
computes an ordering key
([[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]).

**`close` changes shape.** The disposition becomes a positional and `--reason`
becomes the prose; `delivered` becomes `completed` and `abandoned` is dropped.
Only `completed` is gated.

**Advancing is `set`, not a verb of its own.** Workflow status is a field, `set`
is the field verb, and `--if-unchanged` handles the stale-read race — which is
the real hazard, not skipping a rung.

**`--open`** — everything not closed. WORK-0016 records why: it is the first
question anybody asks and cannot be expressed today.

### Outcomes

| Command | Does | Arguments | Flags |
|---|---|---|---|
| `outcome new` | Create | title | `-w/--work-item` |
| `outcome show` / `list` | Read | reference | |
| `outcome set` | Change fields | reference, `field=value` | `--unset`, `--if-unchanged` |
| `outcome assert` | The doer's claim | reference, `succeeded\|failed` | **new** |
| `outcome verify` | The checker's finding | reference, `proven\|disproven\|inconclusive` | `-e/--evidence` **changed** |
| `outcome archive` | Retire it | reference | **new** |

**`verify` changes shape** — the verdict becomes a required positional. Recording
proof has to be said out loud
([[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]]).

**`archive` is why `close` can refuse and still be usable** — its message says
*"retire the ones that no longer apply,"* which needs a command to name.

### Tasks

| Command | Does | Arguments | Flags |
|---|---|---|---|
| `task new` | Create | title | `-w/--work-item`, `--advances` |
| `task show` / `list` | Read | reference | |
| `task set` | Change fields | reference, `field=value` | `--unset`, `--if-unchanged` |

**Taking is not in the first release**
([[records/decisions/ADR-0008-taking-a-task-expires-and-owning-a-work-item-does-not]]),
so `take`, `release` and `steal` do not ship, and exit code `6` stays reserved.

### Decisions

| Command | Does | Arguments | Flags |
|---|---|---|---|
| `decision new` | Create, numbered | title | `-w/--work-item`, `--project` |
| `decision show` / `list` | Read | reference | |
| `decision set` | Change fields | reference, `field=value` | `--unset` |

### Anything, and no noun

| Command | Does | Arguments | Flags |
|---|---|---|---|
| `journal` | Append a line, or show it | reference, text | `-w/--work-item` |
| `init` | Create `.luma/` and a configuration | | |
| `board` | Open the board. Also the no-argument behavior **when a terminal is attached**; concise help otherwise | | |
| `contract` | Emit the whole interface | | `--json` |

**`contract` earns its place in the first release** even though no stage needs
it: `spec.md` §9.7 exists so an actor arriving in an unfamiliar repository
bootstraps from the binary rather than from documentation somebody forgot to
update. That matters most while the interface is moving.

### Deliberately absent

`serve` — the browser interface (§11.7) is a follow-up.
`check` — conditions are reported where records are read rather than asked for
separately.
`config` — configuration is a file people edit; a command for it is not owed yet.
`log` — history is git's, until something needs it portable.
`wave` — waves are not in the first release, so every verb on them is moot.
`promote` — decision promotion (§4.8.1) has no first-release trigger.

### Still open

**Whether `--dry-run` ships**, and on what. `spec.md` §9.6's multi-record
operations are the case for it; none of them is a first-release command, so it
may have nothing to attach to yet.

## The board

`docs/design/sketches.md` draws these screens. It is concept art — where it and
this section disagree, this one is the scope.

### Must

- **A board view of two or three columns.** Three is fine.
- **Navigation between columns, and up and down within a column.**
- **Navigation to columns that are offscreen.** Seven statuses do not fit at
  once; `captured` and `closed` are the bookends.
- **Computed completion on the card** — how many of a work item's live outcomes
  are proven, out of how many there are. `show` and `list` carry it in their
  structured output (`spec.md` §9.3); `CompletionOf` already computes it and
  `close` already layers its gating on top, so this is plumbing an existing
  function into the read path. This is the board's reason to exist
  (`spec.md` §11.1): it is derived by counting evidence, appears nowhere on
  disk, and without it a column of identifiers and titles shows exactly what a
  directory listing shows. `CompletionOf` already computes it.
- **A detail view of one work item** — its outcomes, their state, its tasks.
- **Creating a work item, an outcome, and a task.**
- **Moving a work item between columns, and ranking it within one**, with the
  arrow keys.

### Should

- **One visible column beside the detail view**, as an alternative arrangement
  to the columns.
- **A loading screen** while work items are read from disk, and while whatever
  worktrees and branches turn out to require is resolved.
- **Asserting an outcome.**
- **Verifying an outcome.**
- **Opening a work item in an editor.**

### Nice to have

- **Filters.**
- **Menus.**
- **Saying which person is present**, for attribution. Until then this is
  uncommitted per-machine configuration.
- **Basic support for people working at the same time.**

### What each capability calls

Every mutation resolves to the same request a command produces
([[records/decisions/ADR-0004-every-interface-is-an-adapter-over-one-application-layer]]).
A line here either names a command or says it is view state; what it may not do
is neither.

| Capability | Calls |
|---|---|
| Board view of two or three columns | `work-item list --json` |
| Navigating between and within columns | **view state** |
| Navigating to offscreen columns | **view state** |
| Computed completion on the card | derived from the listing — see *Still open* under the command surface |
| Detail view of one work item | `work-item show`, plus `outcome list` and `task list` scoped to it |
| Creating a work item | `work-item new` |
| Creating an outcome | `outcome new --work-item <ref>` |
| Creating a task | `task new --work-item <ref>` |
| Moving between columns | `work-item set <ref> workflow_status=<value> --if-unchanged <hash>` |
| Ranking within a column | `work-item rank <ref> --before <ref>` |
| One visible column beside the detail | **view state** — an arrangement, not a capability |
| Loading screen | **view state** |
| Asserting an outcome | `outcome assert <ref> <succeeded\|failed>` |
| Verifying an outcome | `outcome verify <ref> <verdict> --evidence …` |
| Opening in an editor | `work-item edit <ref>` |

**Moving a card always passes `--if-unchanged`.** The hazard is not skipping a
rung — nothing forbids that — it is the stale read: the board renders
`preparing`, computes that the next column is `prepared`, and between those
moments an agent moves the record to `todo`. Without the hash the card silently
moves backwards.

**"Move to the top of this column" is not `--top`.** `--top` is global; the
board wants `--before <the first card in this column>`, which it knows because
it rendered the column. Two different operations wearing the same word, so the
board's label and the command's flag deliberately differ.

**Every view-state line is one the board must never acquire a command for**
(`spec.md` §11.6). A cursor position is not a fact about the work.

### Three things moved out of the tier they were first given

**The detail view is a must, not a should.** Creating an outcome means adding it
*to* a work item, which means standing on one and seeing what it already has —
and completion on the card is a pointer that demands somewhere to resolve it.
*Four of six proven* invites exactly one question, and a board that cannot
answer it has put a number on screen to no purpose.

**Moving and ranking are one interaction, so they are one scope item.**
`spec.md` §11.2 describes a single modal move mode — *"enter it, reposition with
the arrow keys across columns and within one, confirm or cancel."* Shipping the
mode gives both; splitting them across tiers means building half a mode and
paying for it twice. And *advance* is one of this document's six stages, so a
board that cannot move a card leaves a stage with no board at all.

**Asserting and verifying stay a should.** Once they are there the board runs
the whole loop, which is the right eventual shape and more than a first release
needs. The split that remains is deliberate: **the board arranges the work and
the command line records the evidence.** Creating and arranging is the expensive
typing; asserting is one short command. Both ship as commands regardless — this
is only about which surface reaches them first.

### What this leaves out, and why it is not an oversight

**A wave view** (`spec.md` §11.2). Waves are not in the first release, so the
view has nothing to show.

**A health view.** The conditions in `spec.md` §5.2 are worth surfacing and are
not worth a screen yet.

**Formation as visual sharpness** (§11.2). A genuinely good idea, and expensive:
it must cost no horizontal space, need no legend, and survive without color.

**Coalesced updates** (§11.3). Redrawing once when a burst of writes settles
needs file watching and debouncing. Related to *people working at the same time*
above, and deferred with it.

**Hiding drafts.** §11.2 says the backlog view hides them by default — *"capture
is generous, the working surface is not."* This board does the opposite: the
pile is simply the leftmost column, which is what makes `captured` a bookend.
The intent behind §11.2 survives elsewhere, since nothing forces a person to
look at that column.

### Still open

**How horizontal movement works.** Whether moving the cursor past the last
visible column shifts the viewport by one, or whether paging the column set is
its own key. The sketches imply the first, which reads better and couples the
cursor to the viewport.
