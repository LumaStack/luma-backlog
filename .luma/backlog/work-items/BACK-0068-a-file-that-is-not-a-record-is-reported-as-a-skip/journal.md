# Journal — A file that is not a record is reported as a skip

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-08

the immediate breakage is fixed — evidence/ is excluded in corpus.isRecordPath and in the test that walks this project's own records; TestTheProjectsOwnRecordsParse was FAILING, not merely noisy, so committing the transcript had turned this from warning fatigue into a red suite
left open deliberately: evidence/ is a local convention and not a luma-layout tier, so where attachments belong is still the real question and the exclusion is a name this project chose today rather than one the layout blesses
