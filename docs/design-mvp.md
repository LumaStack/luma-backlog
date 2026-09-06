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
  