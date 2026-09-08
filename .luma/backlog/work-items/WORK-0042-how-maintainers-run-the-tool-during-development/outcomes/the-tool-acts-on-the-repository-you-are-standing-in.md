---
type: outcome
title: The tool acts on the repository you are standing in
desired_state: The tool resolves its corpus from the caller's working directory, whichever repository that is, and refuses outside one.
verify_by:
  - Run from a subdirectory of this repository; expect this project's records.
  - Run from a different project's checkout; expect that project's records, not this one's.
  - Run from a directory in no repository; expect exit 2 and a message naming the path.
  - Confirm the wrapper uses `go build -C` and not `go run -C` --- the latter follows the program into execution and pins every call to this repository.
work_item: '[[work-items/WORK-0042-how-maintainers-run-the-tool-during-development]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:36:24Z'}
verified:
  - at: "2026-09-06T16:37:20Z"
    by: agent:claude-opus-5/luma-backlog
    as: proven
evidence:
  - at: "2026-09-06T16:37:20Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'From luma-backlog/docs: this project`s work items. From luma-foreman: that project`s corpus (ideas/, plans/, sweeps/) and its two frontmatter-less files reported as skipped. From /tmp: exit 2, "no git repository here or above /tmp". go run -C was tested and rejected --- from /tmp it listed luma-backlog`s records rather than refusing.'
---

# The tool acts on the repository you are standing in

Two behaviors depend on the working directory: root discovery walks up from
it, and `new` derives the work item from where it was run. A wrapper that
sets the directory for its own convenience breaks both, and breaks them
quietly --- records filed under the wrong parent, answers about the wrong
repository, no error either time.
