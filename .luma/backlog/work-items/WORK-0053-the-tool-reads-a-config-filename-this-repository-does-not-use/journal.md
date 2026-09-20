# Journal — The tool reads a config filename this repository does not use

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-20

the undecided name decided itself: docs/spec.md §8.1 already said luma-backlog.yaml — one file per tool, named for the tool. the code was the deviation, and the capture's guess about how (the internal/backlog → internal/corpus rename dragging the filename with it) fits
the fix was one constant plus deleting the twin. the rename was free because no other repository uses the tool yet (WORK-0070) — the window WORK-0037 warned about, used before it closes
proof ran inverted from the capture's: instead of setting a value and watching it be ignored, work_item_key: BACK went into luma-backlog.yaml and the next record was keyed BACK-0103
the second defect in the capture — an unrecognized config file ignored in silence — is NOT fixed by the rename and left this record so it cannot be lost in the close: captured as BACK-0104
