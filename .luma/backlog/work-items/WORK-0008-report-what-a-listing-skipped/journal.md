# Journal — Report what a listing skipped

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-04

starting: the skip already happens and is already right — only the reason is discarded, so this is plumbing rather than validation
reported from list and close only; Resolve and journal's work-item derivation take the skips and drop them, with a comment naming this work item rather than staying silent about the silence
skips are collected whatever the filter — narrowing a listing must not narrow the warning and hide the very record somebody is hunting for
verified by breaking it: silencing the call in list.go fails TestListNamesWhatItSkipped, and TestCloseReports still passes because close carries its own call — which is the evidence the two paths are independent
