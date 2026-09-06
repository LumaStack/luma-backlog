# Interface sketches

**Concept art, not a specification.** These are a starting point for arguing
about the interface — drawn to be thrown away, and expected to be short-lived.
Nothing here is decided.

## How to read this

**Do not treat a sketch as a decision.** `CLAUDE.md` states this as a rule for
the whole project, and this document is where it matters most: sample names,
values, labels and layouts show **shape only**. If you need a settled answer,
it is in `.luma/records/decisions/` or `docs/spec.md` — and if it is in neither,
ask rather than reasoning from a drawing.

**A decision in force outranks anything here**, without exception. These screens
were drawn before the decisions that now govern the interface and have not been
revised against them.

**Some of this is more durable than the rest.** The reserved hotkeys, the
box-drawing characters and the color guidance are conventions that may well
outlive the mockups; as each is settled it moves into `docs/spec.md` and stops
being provisional. **The screens themselves are not expected to survive contact
with a working board** — they exist so that disagreement has something concrete
to point at.

The simplification is deliberate. ASCII is a rough medium here, and a real
implementation should use whatever renders more readable, functional and
intuitive than what is drawn below.

We want the interface to be navigated by Arrow keys but also clickable when it is pragmatic to do so:

    Search: ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒  <- a box I can type in, text should be transposed over the box as I type

    Kind: ▒▒▒▒▒▒▒▒▒▒▒▒▒▼▒  <- a dropdown I can choose from, text should be visible and transposed over the light background box

For example instead of drawing a box like:

    |----------|
    |          |
    |----------|

We should probably instead make it pretty with:

    Top-Left Corner ╭ U+256D
    Top-Right Corner ╮ U+256E
    Bottom-Left Corner ╰ U+256F
    Bottom-Right Corner ╯ U+2570
    Horizontal Border ─ U+2500
    Vertical Border │ U+2502

In order to draw a pane above existing interface (vertically layered pane),
consider giving it a different visual treatment, such as surrounding the pain
with an extra border to give it an ASCII shadow effect, like:

    ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
    ▒              ▒
    ▒              ▒
    ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒

Dropdown options will likely be centered in the middle of the screen in one of
these panes rather than trying to inline them.  If the box has more text than 
can fit on the screen it should scroll within the pane.

Where color makes sense, let's use it.  Sparringly.  For example a decent place to start might be (this is an idea not a requirement):
- Dark grey - Highlight code
- Muted blue - Highlight hotkeys or section titles
- Muted light blue - Highlight links
- Muted yellow - Highlight paths
- Muted light green - Highlight done or checkmarks

## Reserved hotkeys

The design should make it easy to change default key bindings.  For MVP it's not important but it should come soon so we can try different bindings out.  And eventually this may become configurable for end users as well.

I think as much as possible we should reserve `HJKL` keys for potential navigation.

- `Arrow keys`: Navigation
- `Enter`: Select
- `Esc` or `q`: Back or exit
- `/`: Search
- `?`: Menu
- `Tab`: Switch views or modes
- `N`: New item
- `E`: Edit tiem
- `M`: Move item, typically between columns

## Capture view (example)

In most cases I believe two columns are best, but in some cases we might want 3 
or more.  For MVP let's stick with 2 column max.

    |--- Filters ---------------------------------------------------------------------------------------------------------|
    | Search: ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒                                      Kind: ▒▒▒▒▒▒▒▒▒▒▒▒▒▼▒ |
    |---------------------------------------------------------------------------------------------------------------------|
    
    |--- Captured (5) ---------------------------------------|--- Unprepared (1) -----------------------------------------|
    | WORK-0001  Customer keeps asking for CSV export        | WORK-0002    Bootstrap the styleguide                      |
    | WORK-0003  Investigate database spikes                 |                                                            |
    | WORK-0004  Maybe replace authentication provider       |                                                            |
    | WORK-0005  Sarah mentioned onboarding problem          |                                                            |
    | WORK-0006  Upgrade core library to v1.8                |                                                            |
    |---------------------------------------------------------------------------------------------------------------------|

    [q/Esc] Back | [Tab] Switch | [/] Search | [N] New | [E] Edit | [M] Move | [Arrows] Navigate [?] Menu

## Navigating workflow columns

When navigating to the right it would show the next two columns over.

    |--- Filters ---------------------------------------------------------------------------------------------------------|
    | Search: ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒                                      Kind: ▒▒▒▒▒▒▒▒▒▒▒▒▒▼▒ |
    |---------------------------------------------------------------------------------------------------------------------|
    
    |--- Unprepared (5) -------------------------------------|--- Preparing (1) ------------------------------------------|
    | WORK-0001  Customer keeps asking for CSV export        | WORK-0002    Bootstrap the styleguide                      |
    | WORK-0003  Investigate database spikes                 |                                                            |
    |---------------------------------------------------------------------------------------------------------------------|

    [q/Esc] Back | | [Tab] Switch | [/] Search | [N] New | [E] Edit | [M] Move | [Arrows] Navigate | [Enter] Select

## Create work item

When pressing [N] for new work item.

Maybe Ctrl+N so we can give creating work items a universal hotkey.  It would be ideal if this new work item
pane showed up over the existing interface by display the previous interface underneath - padding the new pane on top, 
right, left by 4 characters worth of blank space so it appears on top. (nice to have)
     
    |--- Create Work Item ------------------------------------------------------------------------------------------[Esc]-|
    |                                                                                                                     | 
    | Under:   Captured ▼                                                                                                 |
    | Kind:    Change ▼                                                                                                   |
    | Title:   ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒                                                           |
    |                                                                                                                     |
    | ## Description:                                                                                                     |
    | █                                                                                                                   |
    |---------------------------------------------------------------------------------------------------------------------|

    [q/Esc] Back | [Arrows] Navigate | [Enter] Select

## Search view (using Detail view)

Here is an example of what pressing Tab and switching views or modes might do.  Or entrying stuff in the Search bar does.

    |--- Filters ---------------------------------------------------------------------------------------------------------|
    | Search: ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒                                      Kind: ▒▒▒▒▒▒▒▒▒▒▒▒▒▼▒ |
    |---------------------------------------------------------------------------------------------------------------------|
    
    |--- Captured (5) -----------------|--- Details ----------------------------------------------------------------------|
    | WORK-0001  Customer keeps asking | ○ Captured WORK-0004  Maybe replace authentication provider                      |   
    | WORK-0003  Investigate database  | -------------------------------------------------------------------------------- |
    | WORK-0004  Maybe replace authent | Details                                                                          |
    | WORK-0005  Sarah mentioned onboa | Created:  2026-01-01 23:39                                                       |
    | WORK-0006  Upgrade core library  | Updated:  2026-01-01 23:39                                                       |
    |                                  | Kind:     Inquire                                                                |
    |                                  | Priority: Medium                                                                 |
    |                                  | Owner:    @jane                                                                  |
    |                                  |                                                                                  |
    |                                  | Description                                                                      |
    |                                  | Spike the authentication provider to see if there is a better solution and if we |
    |                                  | can get it done in less than a month.                                            |
    |                                  |                                                                                  |
    |                                  | Preparation                                                                      |
    |                                  | ✔ Gather the authentication provider documentation                               |
    |                                  | ○ Refine outomes to be measurable and provable                                   | 
    |                                  |                                                                                  |
    |                                  | Outcomes                                                                         |
    |                                  | ○ Authenticaion provider evaluated                                               |
    |                                  | ○ Timeline provided                                                              |
    |                                  | ○ Cost analysis                                                                  | 
    |                                  |                                                                                  |
    |                                  | References                                                                       |
    |                                  |   https://example.com/authentication/provider                                    |
    |                                  |                                                                                  |
    |---------------------------------------------------------------------------------------------------------------------|

    [q/Esc] Back | [Tab] Switch | [/] Search | [N] New | [E] Edit | [M] Move | [Arrows] Navigate [?] Menu

If decisions, tasks, explorations, or other sub-item types are present, we should display those as well.
We may also need a way to add more somehow using an `[+]` or `[Add]` button maybe? (nice to have for MVP)

## Empty results

Here is an example of it looks like when work items are missing or the filters didn't match.

    |--- Filters ---------------------------------------------------------------------------------------------------------|
    | Search: ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒                                      Kind: ▒▒▒▒▒▒▒▒▒▒▒▒▒▼▒ |
    |---------------------------------------------------------------------------------------------------------------------|
    
    |--- Work Items (0) ---------------|--- Details ----------------------------------------------------------------------|
    | No matching work items.          | No work items to display.                                                        |   
    |                                  |                                                                                  |
    |  • Status: Done                  | Try adjusting search or clearing filters.                                        |
    |  • Kind:   Defect                |                                                                                  |
    |                                  |                                                                                  |
    |                                  |                                                                                  |
    |---------------------------------------------------------------------------------------------------------------------|

    [q/Esc] Back | [Tab] Switch | [/] Search | [N] New | [E] Edit | [M] Move | [Arrows] Navigate [?] Menu

## Menu example

    ┌─ Main Menu ────────────┐
    │   File Manager         │  
    │ ▶ Settings             │  <- Active Selection
    │   Exit                 │  
    └────────────────────────┘