---
type: outcome
title: A prefix outside the Jira Cloud rule is refused, naming the rule
desired_state: ""
verify_by: "Test: work_item_key values 'back', 'X', '2AB', 'ABCDEFGHIJK' and 'WORK_ITEMS' each fail config.Parse with an error citing the rule ^[A-Z][A-Z0-9]{1,9}$ — an uppercase letter, then uppercase letters or digits, two to ten characters."
work_item: '[[work-items/WORK-0102-the-work-item-key-prefix-is-configurable]]'
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:39:38Z'}
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:39:52Z'}
verified:
  - as: proven
    at: "2026-09-20T17:52:09Z"
    by: agent:claude-fable-5/luma-backlog
evidence:
  - at: "2026-09-20T17:52:09Z"
    by: agent:claude-fable-5/luma-backlog
    what: 'TestParseRefusesAPrefixOutsideTheJiraCloudRule: back, X, 2AB, ABCDEFGHIJK, WORK_ITEMS, BA-CK and explicit empty each fail config.Parse with an error citing ^[A-Z][A-Z0-9]{1,9}$. Repeated live: work_item_key: back makes every command refuse with ''is not a legal key prefix: an uppercase letter, then uppercase letters or digits, two to ten characters''. The validator and the key reader share one rule fragment, config.KeyPrefixRule, so they cannot drift.'
---

# A prefix outside the Jira Cloud rule is refused, naming the rule

Why this matters, and anything needed to read the check correctly.
