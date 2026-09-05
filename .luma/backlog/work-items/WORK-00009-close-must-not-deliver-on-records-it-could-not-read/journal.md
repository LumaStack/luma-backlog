# Journal — Close must not deliver on records it could not read

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-04

found while fixing: Completion.Skipped took every skip from the walk, not just this work item's outcomes, so a corrupt file anywhere would have blocked every close
refused for want of an answer rather than on an opinion — the tool cannot count what it cannot read, and delivered would otherwise come out clean on evidence nobody has seen
no force flag: canceled, superseded and abandoned claim nothing about evidence, so none needs a count and any of them is the way out when a file is beyond repair
Complete() itself now returns false when anything was skipped, so a future caller of the arithmetic cannot be fooled the way close was
