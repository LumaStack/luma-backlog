---
type: type_definition
type_version: "0.0.1"
defines: exploration
version: "0.0.1"
fields:
  work_item: {field_presence: required, field_type: wikilink, desc: "The work item this investigation belongs to."}
---

# Exploration

Ideas, research, spikes, and investigations — including the ones that went
nowhere. It lives in `explorations/` inside the work item it belongs to.

**Its own type, because the whole risk is leakage.** An idea recorded while
thinking must never be mistaken for something the team committed to — a work
item is judged on its outcomes and on nothing else, and keeping exploration
visibly apart makes that true on inspection as well as structurally.

**Nothing leaves exploration except by an explicit act.** Turning an
investigation into work means someone deliberately creating an outcome or a
task from it — promotion copies, never relocates, so the reasoning stays where
it was written at the exact moment it becomes worth having. Both endings are
non-destructive: an exploration either produces work or it does not, and
neither is a deletion.
