# Journal — Close promises a force flag it does not have

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-09

found 2026-09-09 while wiring force-tracking on transition: close's refusal message names --force, backlog-move devotes four paragraphs to how it behaves, and the flag was never built
so the refusal it carries is currently absolute — a completed close over an unmet outcome cannot be done at all, where the design says it can be done and recorded. that is a stricter tool than anybody chose
same shape as WORK-0075: a guarantee table asserting behaviour the binary does not have. this one is a message promising an escape hatch that does not exist, which is worse than silence — it tells the user to type something that fails
