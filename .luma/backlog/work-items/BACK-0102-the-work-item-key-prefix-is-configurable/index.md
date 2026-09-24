---
type: work-item
key: BACK-0102
title: The work item key prefix is configurable
workflow_status: closed
rank: 070.0210.000
kind: change
stage: provisional
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:16:12Z'}
description: 'I want to be able to change the default WORK key in config. What does Jira allow for their keys — I probably want to follow similar rules. Decided while capturing: Jira Cloud style.'
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:53:46Z'}
closed: {on: 2026-09-20, as: completed, by: 'agent:claude-fable-5/luma-backlog'}
former_keys: ["WORK-0102"]
---

# The work item key prefix is configurable

## The problem

**The prefix is a constant.** `internal/corpus/key.go` — `const KeyPrefix = "WORK"`,
written into every record at creation. A project adopting this tool gets `WORK`
whether it wants it or not, and the constant's own comment already anticipates
this record: *"a repository that later chooses its own prefix changes what gets
written and not what already exists."*

## What is being delivered

**A configuration setting for the key prefix**, validated by Jira Cloud rules —
decided at capture: the prefix starts with an uppercase letter, then uppercase
letters or digits, two to ten characters (`^[A-Z][A-Z0-9]{1,9}$`).

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### What Jira allows, since that was the question

- **Jira Cloud (fixed):** 2–10 characters, an uppercase letter first, then only
  uppercase letters and digits. No hyphens, no underscores.
- **Jira Data Center (admin-configurable):** default `([A-Z][A-Z+]+)` — letters
  only; admins may widen to letters, digits and underscore, first character
  always a letter, length capped separately at 10.

**Cloud style was chosen**, and it is also the safe choice here: underscore is
excluded, and `-` stays unambiguous as the prefix/number separator.

### Most of the enforcement lands before this starts

[[work-items/BACK-0082-one-key-written-many-ways-must-resolve-to-one-record]]
is teaching the engine to *recognize* any Jira Cloud prefix when parsing and
resolving keys. Once that lands, this record reduces to a config key feeding
`FormatKey` and allocation — validation of the configured value against the
same rule, and nothing else new.

### The open question the capture left open

**Whether an existing corpus can change its prefix later, or the setting only
applies at init.** "Change the default" reads either way. Changing it later is
a migration — every stored key is the old prefix, and a key is meant never to
change — which puts it in
[[work-items/BACK-0022-migrate-a-corpus-when-the-vocabulary-changes]] and
[[work-items/BACK-0037-old-records-get-migrated-as-the-system-improves]]
territory. Left ambiguous on purpose; preparation decides.

## Out of scope

**How configuration is written at all** —
[[work-items/BACK-0049-whether-configuration-is-edited-by-command-or-by-hand]]
decides command versus hand-editing; this record adds a setting to whichever
answer wins.

## Constraints

- **A configured prefix changes what gets written, never what exists** — the
  constant's comment, kept as the rule. No silent renaming of a corpus.
- **The prefix must validate against the Jira Cloud rule** the engine enforces
  after WORK-0082, or a configured value would create keys the tool itself
  cannot parse.

## References

- `internal/corpus/key.go` — `KeyPrefix`, `FormatKey`, and the comment that
  predicted this record.
- [[work-items/BACK-0082-one-key-written-many-ways-must-resolve-to-one-record]]
  — the recognition half, in flight when this was captured.
