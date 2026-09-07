---
type: decision
title: A symbol that must mean one thing is assigned in one place
decided: 2026-09-06
stage: provisional
reopen_trigger: "a registry grows large enough that finding a symbol in it is harder than finding the collision it prevents, or a surface arrives whose symbols are genuinely scoped rather than global"
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:50:31Z'}
---

# ADR-0009: A symbol that must mean one thing is assigned in one place

## Summary

**Flag shorthands are assigned in one registry. Board hotkeys are assigned in
one registry.** Neither is spelled inline at the place it is used.

## Problem

A single letter carries a whole meaning, and the meaning is invisible where it
is written. `-w` at a call site is just a string. Whether it already means
something else somewhere in the tree cannot be seen from there, and the cost of
being wrong is not a build failure --- it is a letter that means one thing under
`work-item` and another under `wave`, discovered by a person who typed the
wrong one.

The same is true of a hotkey. `spec.md` §11 gives the board views, editing and
navigation; each will want keys, and they are drawn from a smaller alphabet
than flag names are.

**The failure mode is silence.** Nothing errors. Two commands simply disagree
about what a letter means, and both work.

## Decision

**One registry per namespace, and the symbol is never spelled inline.**

```go
// Flag shorthands are assigned here, once, so a letter means the same flag
// everywhere in the tree. Adding a command? Take a shorthand from this list
// or add one here --- never spell a letter inline.
const (
	shortWorkItem = "w"
	shortKind     = "k"
	// ...
)
```

The registry is the place the author already has to visit, so it is where the
guidance belongs. A comment in `list.go` does not reach whoever writes
`wave.go` next month; **a comment guards the file it is written in, and a
registry guards the namespace.**

**Two registries, not one.** Flags and hotkeys are separate alphabets used by
separate surfaces; `w` meaning `--work-item` has no bearing on what `w` does on
the board.

**A test is the backstop, not the mechanism.** Whoever bypasses the registry is
exactly who the registry cannot help, so a check that a name maps to one
shorthand tree-wide is worth having. It is second in line.

## What this is not

**Not a claim that flags are uniform across nouns.** They are not, and they
should not be --- `--kind` belongs only to `work-item new`, `--project` only to
`decision new`. A test asserting identical flag sets would be false almost
immediately and would get weakened until it asserted nothing. The invariant is
**one name, one shorthand**, which survives flags diverging because divergence
adds symbols rather than reusing them.

**Not a claim that a name means the same thing everywhere.** `--work-item` is
already a filter on `list`, an assignment on `new`, and a selector on
`journal`. That is fine. The letter must be stable; the sentence around it need
not be.

## Deferred

**Which letters get which flags.** Assigning a shorthand scheme against a
partial flag set means redoing it. **Re-open when the last command in
[[work-items/WORK-0031-reshape-the-command-surface]] lands** and the full set
can be seen at once.

*Input for that decision, recorded as it turns up:*

- **`-r` stays with `--reason`.** It was briefly contested: `-r` is the listing
  convention for reversing (`sort -r`, `ls -r`), and a `--reverse` flag was
  proposed. That flag does not exist --- direction is part of the sort key
  instead, `--sort=-updated` --- so nothing competes for the letter.
- **The episode is kept because the shape of it will recur.** `-r` belonged to
  `--reason` because `close` was written before anything wanted to reverse a
  listing, not because `--reason` had the better claim. Had the letters been
  settled then, a later command with a stronger claim would have found the
  conventional one taken. That is the argument for assigning them once, at the
  end, rather than as each command arrives.

**Hotkeys entirely.** The board does not exist. Re-open with it.

## References

- [[work-items/WORK-0031-reshape-the-command-surface]] --- where the registry
  first gets built.
- [[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]]
  --- CLIG's standard flag names are an input to the assignment, not a
  substitute for the registry.
