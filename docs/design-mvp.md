# Design MVP

See the [design-vision.md] first to understand where we will end up so we can have an expandable open ended MVP.

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
  
## The board

`docs/design-sketches.md` draws these screens. It is concept art — where it and
this section disagree, this one is the scope.

### Must

- **A board view of two or three columns.** Three is fine.
- **Navigation between columns, and up and down within a column.**
- **Navigation to columns that are offscreen.** Seven statuses do not fit at
  once; `captured` and `closed` are the bookends.
- **Creating a work item.**
- **Creating an outcome.**
- **Creating a task.**

### Should

- **One visible column beside a detail view.**
- **Ranking, and moving between columns, with the arrow keys.**
- **A loading screen** while work items are read from disk, and while whatever
  worktrees and branches turn out to require is resolved.
- **Asserting an outcome.**
- **Verifying an outcome.**
- **Opening and editing a work item.**

### Nice to have

- **Filters.**
- **Menus.**
- **Saying which person is present**, for attribution. Until then this is
  uncommitted per-machine configuration.
- **Basic support for people working at the same time.**

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

### Two readings this section assumes

**Creating an outcome does not require the detail view.** Creating is a must and
the detail view is a should, so the create flow opens from a selected card
rather than from inside an opened work item. If that turns out to be wrong, the
detail view moves up a tier rather than creation moving down.

**The first board builds plans; the command line closes them.** Creating work
items, outcomes and tasks is a must, while asserting and verifying are a should.
That is a coherent split — creating is the expensive typing, completing is one
short command — and it is deliberate rather than an omission.

### Still open

**Whether a card shows computed completion.** `spec.md` §11.1 says this is the
board's reason to exist — *"opening a record in an editor tells you what it
says; the board tells you what is true"* — and completion is derived by counting
outcomes with evidence, appearing nowhere on disk. A board of identifiers and
titles shows nothing a directory listing does not. The arithmetic already exists
(`CompletionOf`), so the cost is display rather than computation.

**How horizontal movement works.** Whether moving the cursor past the last
visible column shifts the viewport by one, or whether paging the column set is
its own key. The sketches imply the first, which reads better and couples the
cursor to the viewport.
