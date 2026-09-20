# Journal — The work item key prefix is configurable

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-20

captured, prepared, worked and closed in one session. setting named work_item_key (maintainer's spelling, underscores matching the file's other keys), optional, WORK when absent, this repository set to BACK
the validator and the key reader share one regexp fragment, config.KeyPrefixRule — two copies of the Jira Cloud rule would drift, and a prefix the validator accepted but the reader could not parse would create records the tool cannot find
numbering continues across a prefix change — BACK-0103 followed WORK-0102 — because the corpus keeps one sequence (highestKey's own comment: a spoken number has to mean one record). restarting at 1 was never seriously on the table
the capture's open question (init-only or changeable later) answered itself: changeable any time, because WORK-0082 made mixed prefixes readable first. order mattered — recognition before configurability meant no migration was forced
migration is deliberately NOT implied by changing the setting — the maintainer's own reason, external systems hold old keys we do not control. deferred to BACK-0103 with that reason on the record
scope grew twice mid-work, both maintainer-driven: no placeholder config (lkf_version, type_namespace, columns deleted from code too — parsed by nothing, Qualify had zero callers), then workflow_status out of the FILES until thought through — it stays in code because transition and rank read it; the file copy duplicated the defaults byte for byte and asserted the vocabulary was settled, which it is not (WORK-0083)
spec §8.1 already named luma-backlog.yaml; §8.2's example gained work_item_key
