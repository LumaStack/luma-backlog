# Journal — Session save and session end, and moving a record to the project it belongs in

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

captured in luma-backlog and it belongs to luma-catalog — session-manager is a bundle there. that is not an accident of where the conversation was; it is the second half of the record demonstrating the first half
the rename would overturn a stated design rather than tidy it: session-continuity splits the three procedures by WHO READS the note — you minutes later, a named successor, a stranger who may be you having forgotten — and close deliberately writes no next steps because a stale plan is unfalsifiable
the strongest argument FOR collapsing is that the reader is not knowable at the time. you save, and afterwards it turns out to have been a checkpoint or a handoff — which is the maintainer's point that it does not matter which
one thing must not collapse: something has to be the LAST procedure, because the design rests on notes being drained and deleted and session-continuity calls that its weakest structural risk. a save that sometimes deletes is worse than two verbs
and end beats close on a spent-word test — a work item closes, close is a command with dispositions and refusals, so a session closing collides. end is unspent, shorter, and pairs correctly: save repeatedly, end once
on moving records between projects: the hard part is not the move. spec 4.8.1 says promotion copies and never moves because a path is an identity, and across a project boundary the tool cannot rewrite the other side's links at all. so the shape is probably copy-and-leave-a-pointer, which is what promotion already does for decisions
DECIDED shape for moving a record between projects: COPY it to the destination, then CLOSE the origin as superseded, with the superseded link pointing at the record in the other project
which is good because it reuses machinery that exists rather than inventing any — superseded is already a close disposition, backlog-move already says 'something else covers it, link to what', and nothing is deleted so the origin keeps its journal and its history
THE UNSOLVED PART is the link itself. a wikilink resolves inside one corpus, so [[work-items/WORK-0090-...]] means nothing across a project boundary. a cross-project reference has no form yet and that is the actual work here, not the copying
and it interacts with keys: WORK-0090 here and WORK-0090 there are different records with the same handle, so a cross-project reference has to carry the project as well as the key — which is the same shortfall as WORK-0013's collision problem seen from outside one corpus
