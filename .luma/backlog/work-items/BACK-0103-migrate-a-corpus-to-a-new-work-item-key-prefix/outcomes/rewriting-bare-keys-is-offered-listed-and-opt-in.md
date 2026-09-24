---
type: outcome
title: Rewriting bare keys is offered, listed, and opt-in
description: "Bare keys are left alone by default because rewriting one can be wrong: it may name a work item in another project using the same prefix. But for a single-project repository it is right nearly every time, so refusing to offer it makes the operator do by hand what the tool could do reliably. The resolution is the same shape as --renumber: the run reports, a flag acts, and nobody is asked to decide without seeing what would change. Grouped by file rather than listed line by line, because the decision is per corpus and nobody makes it by scrolling 550 entries. Journals are included, confirmed rather than assumed: 45% of what the flag touches is lines like 'Journaled on BACK-0039', and a key in a journal is an address rather than a quotation — the record still exists and answers to a new name, git holds what was literally written, and a reference that resolves is worth more than one preserved in amber."
desired_state: "A run lists the bare old keys it did not rewrite, grouped by file with counts, and names the flag that would rewrite them. With --include-bare-keys, it rewrites them and reports what it changed instead."
verify_by: "Over a corpus with bare old keys in prose: a default run leaves every one byte-identical and lists them grouped by file, ending with the command that would rewrite them. A run with --include-bare-keys rewrites exactly those occurrences and nothing else, and a key belonging to no record in this corpus is reported rather than rewritten."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T05:52:59Z'}
verified:
  - as: proven
    at: "2026-09-24T07:51:50Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-24T07:51:50Z"
    by: agent:claude-opus-5/luma-backlog
    what: Both halves run against this repository. The default run left all bare keys byte-identical and listed them grouped by file with the flag named in the output. The opt-in run rewrote 138 files, listing every one with its change count; .git was fingerprinted before and after and was byte-identical; all 102 redirects survived. A key belonging to no record here is reported rather than rewritten — TestIncludeBareKeysRewritesProse asserts WORK-0999 survives. TestIncludeBareKeysWorksAfterTheMigration covers the later-run case, which was a silent no-op until former_keys supplied the mapping.
---

# Rewriting bare keys is offered, listed, and opt-in

Why this matters, and anything needed to read the check correctly.
