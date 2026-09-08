# Journal — Review the backlog procedures

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-08

backlog-journal gained a 'say where it went' section at bundle 0.24.0 — the maintainer noticed the agent doing it ad hoc and wanted it consistent; the argument that made it a rule rather than a courtesy is that journalling is invisible AND unasked, so not reporting it leaves a reader with no model of what is being written in their name
corrected within minutes at 0.24.1 — the first version asked for one line at the end and the maintainer wanted it per point, per blurb; a footer names the records and hides the mapping, and correcting placement is the whole reason for reporting, which nobody can do for a mapping they were never shown
the say-where-it-went rule contradicted itself — 'what was journalled, where it went' followed by 'where, not what' — and the agent followed the second, producing a footer naming a record and saying nothing about it; sixth instance today of a rule failing in practice, and the first where the PROSE was wrong rather than the compliance
settled by the maintainer: inline carries the substance by position and needs nothing more, a footer has no position so it must reference the point — which is why 'journaled on WORK-0031' fails as a footer and would have been fine inline
two rules added to backlog-move at 0.25.0, both from live failure: finish a work item before starting another unless asked otherwise, and never pick the next one — asking a question is not confirmation and neither is answering one, which is exactly how it went wrong today
corrected at 0.25.1 — 'never pick' forbade the opinion, and the maintainer wants it: recommend freely, then stop. the harm is the SURPRISE, not the recommendation, and an orchestrator is surprised by unannounced work exactly as a person is

## ▶ 2026-09-07

seven procedures, three policies and three templates were written in one session with no work item — so no outcomes, and nothing says what a good procedure is; that is why reviewed had no meaning until this record invented one
four files have a single commit and were never read again: backlog-capture at 381 lines, backlog-refine, showing-records, record-view — the longest file in the bundle is among them
backlog-move and backlog-refine are the ones that matter most and got the least: move carries both selection gates, which is where the ladder model lives, and refine carries what an outcome is, which is what this project gets wrong most often
only backlog-show and backlog-next have ever been run against the real corpus, and both changed every single time — a procedure nobody has followed is a draft whatever it looks like, and that is probably the whole definition of reviewed
commit count is a weak proxy: backlog-capture was iterated heavily inside one commit through direct feedback, while backlog-show earned several of its twenty-four on a single heading — use it to find candidates, not to judge them
the table in this record names backlog-next and next-report, both renamed at local/backlog 0.17.0 — backlog-rundown and rundown; the line counts and commit counts still hold, only the filenames moved
