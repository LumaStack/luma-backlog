# Journal — review and audit the implementation

> The work item's memory. Newest entry first; everything below the top block is historical. Append, never curate.

---

## ▶ 2026-09-04

classified as an inquiry — the kind that did not exist when this record was left blank, and the record that argued it into existence

## ▶ 2026-08-10 — the debt is now the whole build

**All eleven tasks are done and none of the code has been read.** When this was parked, six tasks were complete; the first build is now finished and running against this repository's own backlog.

**What that adds to the scope:** the completion arithmetic and the tool's only refusal, the containment story, the record parser that every other package depends on, and the JSON shapes now pinned as contract by golden files. Golden files are worth naming separately — **they lock in whatever the output happened to be at the moment it was first captured**, and nobody has checked that those shapes are the ones we want.

**One more piece of evidence for the case, from today.** Dogfooding found within minutes that every record an agent wrote was attributed to the machine's human owner. Not caught by any test, because no test asks *who does this say did it*. That is precisely the class of defect a reading finds and a suite does not.

**Still open: when.** Unchanged.

---

## 2026-08-09 — parked, with the debt stated plainly

**Where things stand.** Six of eleven implementation tasks are done and none of the code has been read by anyone but the agent that wrote it. The maintainer said so directly: *approving without reviewing for now, and a review and audit will be necessary at some point.*

**Recorded because the alternative is that it disappears.** Unreviewed code stops looking unreviewed within days — it accumulates history, gets built on, and becomes indistinguishable from code somebody vouched for. Writing it down now is the only moment it is obvious.

**The sharpest version of the risk:** tests were written by the same author as the code, which is the weakest form of confirmation available. And that is not theoretical here — **three bugs in `new` were found by running the tool, not by the tests, which passed while the output was wrong.** They asserted on fields rather than on the file a person reads. Nothing suggests that pattern has stopped.

**Still open: when.** Before another person uses it, before 1.0, before it writes to a repository that matters, or when reading it all stops being cheap. Those are different scopes, and choosing between them now would be a guess. The trigger is the decision, not the checklist.
