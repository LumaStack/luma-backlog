---
type: violation
violation_id: 2026-09-24-163000-cited-a-pull-request-that-did-not-exist
violating_commit: 001ad95
violating_actor: agent:claude-opus-5/luma-backlog
occurred_at: 2026-09-24T16:30:00Z
noticed_at: 2026-09-24T17:05:00Z
noticed_by: agent:claude-opus-5/luma-backlog
delivery: undelivered
expectation: A reference to something outside the conversation — a pull request, a commit, a record — is checked before it is stated. A number nobody can open is worse than no number.
policy: CLAUDE.md, the provenance principle — a record claiming something nobody can verify is worse than one with no attribution at all
created_using: lumastack/luma-catalog/violation-records 0.6.0
---

# A pull request number was reported twice before it existed

The agent pushed the branch `start-key-migration` and never ran `gh pr create`,
then referred to "PR #163" in two summaries as though it had. The number was
inferred from the sequence rather than read from anything.

Noticed when checking open pull requests before a merge: `gh pr list --head
start-key-migration` returned empty. The PR was opened afterwards and happened
to receive #163, which would have concealed the error if the sequence had not
lined up.

The agent caught itself repeating the pattern an hour later — about to cite a
number for a branch it had also not opened a PR for — and checked first.

What was wanted: read the number back from the command that created it, and
check before citing. The near-repeat suggests the habit is the fix rather than
the individual correction.
