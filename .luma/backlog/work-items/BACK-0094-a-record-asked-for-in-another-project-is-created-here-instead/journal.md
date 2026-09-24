# Journal — A record asked for in another project is created here instead

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

the instruction named the destination — 'create a new idea in luma-foreman' — and the agent created it here, then wrote in the record that it belongs to luma-foreman. so the information was present in the sentence and in the record and still did not reach the action
backlog-capture never asks which project a record belongs to. it chooses between new and append, quick and considered, and searches for duplicates — always in the corpus it is standing in. the question is not on the page at any step
the tool's confirmation would not have caught it either: 'created backlog/work-items/...' is a relative path with no project in it, identical in every repository. naming the corpus in the output is the cheap half and is independent of the procedure change
two cases that are not the same, and only one is absolute: a destination was NAMED, where creating it elsewhere is disregarding an instruction; and no destination was named but the subject is plainly another project, where asking once is right. conflating them would make capture ask every time, which WORK-0084 says is how capture stops happening
and it cannot be fully fixed yet — neither luma-foreman nor luma-catalog has a backlog corpus, so refusing to create here would leave nowhere to create at all. the immediate fix is narrower: stop, say it belongs elsewhere, ask. the rest waits on WORK-0070
TO EXPLORE, added by the maintainer: what should happen when the receiving project does not have luma-backlog set up yet
and it is the ORDINARY case, not the edge one — both siblings checked today have .luma and no .luma/backlog, so a fix that only works on an already-initialised destination would have helped zero times
four shapes, differing in who decides: refuse and say why; offer to initialise there; hold it here marked for elsewhere; or write it somewhere that is not a backlog at all
the third is the only one that keeps capture cheap AND the only one that quietly builds a pile of records in the wrong place — it is what happened four times today by accident, so the question is whether doing it deliberately and visibly is better or merely tidier
and holding a record for elsewhere has a prerequisite: there is no field for 'this belongs to another project'. WORK-0025 is a work item that cannot point at a work item; this is a work item that cannot point at a PROJECT
MAINTAINER'S FIRST GUESS: error or warn, and ask two questions — install luma-backlog in the receiving project, or create it here for now and migrate it later
that is shapes one and three from the four above, offered as a choice rather than picked — refusing outright is not on the list, and neither is deciding silently
and the second option names the thing the agent's version of it was missing: MIGRATE IT OVER LATER. holding a record here is only acceptable if the move is a real operation somebody will run, which is WORK-0090's copy-then-supersede, which needs WORK-0037's migration mechanism. otherwise 'for now' is where records go to stay
