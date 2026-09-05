# Journal — Do not invent a workflow status

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-04

starting the end-to-end pass: this work item is run through commands only, to verify the-backlog-is-kept-by-the-tool
the line for this pass: commands own structure and state, prose is authored in place — CLAUDE.md's bootstrap order says the skill writes the body and the command makes the shape, so writing body prose is not a failure of the outcome; hand-editing frontmatter would be
friction: new prints a path relative to .luma/ (backlog/work-items/x/index.md), which cannot be pasted into cat from the repository root
the fixture never held a non-worked, stage-less record, so the bug was uncovered by construction — init creates no PROJECT.md and new cannot make a luma/project
friction: close takes a work item and an enum reason, so a task is closed with set workflow_status=closed instead — two different gestures for finishing something, and I reached for the wrong one first
fixed in three places, not one: Status() returns empty for a non-worked type with no stage, the JSON tag gained omitempty, and the text show drops the row rather than printing a bare label
the pre-existing goldens did not move — only the two new ones — which is the evidence the change is contained to the case it was aimed at
delivered: close checked the outcomes and let it through on the arithmetic, not on my say-so
friction: the work item sat at idea for the whole pass and closed straight from there — nothing asks you to move it to in_progress, so the rung vocabulary went unexercised without any complaint
friction confirmed: journal still appends one line at a time, so this wrap-up is a stack of lines rather than an entry with a shape — the 2026-08-10 finding stands
stale finding, now false: tool-written entries DO carry the --- separator and the date heading; only the summary after the date is still missing
unrelated and pre-existing: go test ./internal/record fails on main — corpus_check walks every .md under .luma/ and the adopted bundles ship INDEX.md files and templates that deliberately have no frontmatter
