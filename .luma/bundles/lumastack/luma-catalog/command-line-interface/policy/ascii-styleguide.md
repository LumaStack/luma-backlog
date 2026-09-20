---
type: policy
type_version: "0.0.1"
title: ASCII style guide
description: The marks a command line uses to show state — six glyphs, what each one means, and the rules that keep them readable when colour, fonts or width are not available.
matches: eager
---

# ASCII style guide

**A state is shown as a mark, not a word.** One mark in a fixed first column
scans in a single pass; a word per row has to be read, and a column of read
words is a paragraph nobody asked for.

**Two sets, and a project picks one.** They carry the same six meanings in the
same order, so a reader who knows one can read the other. Pick per project, not
per command — mixing them inside one tool is the only way to get this wrong.

## Basic

Bracketed ASCII. Nothing to install, nothing to render, correct in any
terminal, any font, any pipe, and in a plain-text file forty years from now.

```
┌─────┬──────────────────────────────────┐
│ [ ] │ not started, empty               │
│ [~] │ under way, in progress, working  │
│ [x] │ delivered, success, proven       │
│ [>] │ superseded, replaced             │
│ [-] │ cancelled, missing, not found    │
│ [!] │ failed, error, disproven         │
└─────┴──────────────────────────────────┘
```

**`[ ]`** — an empty box. The universal unchecked affordance, and the one every
reader already knows from a paper list and from markdown.

**`[~]`** — a tilde, which means *approximately* everywhere else it appears.
Something is in the box but it is not settled.

**`[x]`** — the box marked. Markdown task lists made this *done* for a very
large number of people, and that convention wins over any other reading.

**`[>]`** — pointing on. The work went somewhere; look there.

**`[-]`** — struck through. Deliberately nothing, and neutral about it — a dash
is the mark for *not applicable*, not for *went wrong*.

**`[!]`** — the only mark that raises its voice, and the only state that
deserves to.

**The cost of this set is its width.** Three columns per row against one, which
matters in a narrow terminal and matters more in a tree where the marks indent.

**And `[x]` is a trap across sets.** Here it means *done*; in the glyph set the
crossed mark `✘` means *failed*. A reader who learned one and meets the other
will read a success as a failure. **This is the reason to pick one and not
mix.**

## Glyph

Unicode marks. Narrower, quieter, and they hold their meaning at a glance
without being parsed.

```
┌─────┬──────────────────────────────────┐
│ ○   │ not started, empty               │
│ ◐   │ under way, in progress, working  │
│ ✔   │ delivered, success, proven       │
│ ↪   │ superseded, replaced             │
│ ⊘   │ cancelled, missing, not found    │
│ ✘   │ failed, error, disproven         │
└─────┴──────────────────────────────────┘
```

**`○`** — an empty circle. Nothing has happened to it.

**`◐`** — the same circle, half filled. **The fill is the progress**, which is
why these two share a shape: they are one thing at two points.

**`✔`** — a different mark entirely, on purpose. A filled circle would join the
unfinished family and be mistaken for an empty one down a column; **done has to
separate from not-done before anything is read.**

**`↪`** — a hooked arrow: *picked up over there*. It is also the prompt — name
what superseded it on the same line, because a supersession with no successor
named is the one shape of it that is genuinely lost.

**`⊘`** — a circle with a line through it. Switched off, or never there.
**Cancelled and missing read the same to somebody looking**: nothing will come
of it here, and neither is anybody's failure.

**`✘`** — the counterpart to `✔`, same weight and size. *It came out* and *it
did not come out*, so an ending that went badly reads as an ending rather than
an interruption.

**The cost of this set is that it is not ASCII**, despite the name of this file.
A terminal without the glyphs shows tofu. Each is chosen from an old, widely
covered block over a prettier alternative for that reason — `◐` U+25D0 rather
than a quarter-filled variant, `↪` U+21AA rather than a curving arrow from a
later block.

## Why these six, either way

**`✔` and `✘`, or `[x]` and `[!]`, are the two verdicts.** *It came out* and
*it did not*.

**`↪` and `⊘`, or `[>]` and `[-]`, are the endings that are neither.**
Superseded work did not fail; it moved, and something else covers it now.
Cancelled work did not fail either; somebody chose. **Marking either as a
failure reports a loss where there was a decision.**

## The rules that keep it readable

**Pick one set per project and do not mix them.** They share meanings and order
so a reader can move between projects; they do not survive being interleaved
inside one, because `[x]` means *done* and `✘` means *failed*.

**In the glyph set, use the heavy marks — `✔` U+2714 and `✘` U+2718, not `✓`
U+2713 and `✗` U+2717.** The light pair is visibly thinner than the circles
beside it, so a finished row reads as fainter than an unfinished one. That is
backwards: finished work is what a scan skips past, and it should be the easiest
thing to skip.

**One fixed order, endings last** — `○ ◐ ✔ ↪ ⊘ ✘`, or
`[ ] [~] [x] [>] [-] [!]`. Best outcome first among the endings and the bad one
last of all — it is worth arriving at deliberately
rather than meeting halfway down a list. **A view whose order changes between
readings cannot be scanned**, which is the whole reason to fix it.

**Within a state, keep the order the data arrived in.** Do not re-sort. Any
ordering the records carry is one somebody chose, and re-sorting discards it.

**The mark is the meaning; colour is a shortcut.** Where colour is available,
in either set — not started dim, under way yellow, delivered green, superseded
blue, cancelled dim, failed red — it is deliberately redundant with the mark.
Output stays correct in a pipe, in a log, and for anybody who cannot distinguish
the hues. **Colour makes a scan faster; it never
makes one possible.**

## Do not invent a state the record does not claim

**A mark is a claim about what happened**, so use the one the data supports and
no stronger. Most systems cannot distinguish every state here: *cancelled* and
*abandoned* are usually one field, *failed* and *unproven* usually the same
absence of a pass.

**Absence is `○`, not `✘`.** Nobody having looked is not a bad result — it is no
result. **The distinction worth building toward** is between *not checked* and
*checked and undecidable*: the second is a defect in what was asked, it never
improves by waiting, and collapsing the two hides it. Test frameworks keep
*inconclusive* separate from both *failure* and *not run* for exactly that
reason.
