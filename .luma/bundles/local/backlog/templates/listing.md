# Listing template

**A list of records, anywhere.** Copy the shape, not this file. The marks and
the ordering are [[showing-records]]; this is what the block looks like.

```
**<Column or filter> (<count>)**

○ WORK-0022 · Migrate a corpus when the vocabulary changes
◐ WORK-0031 · Reshape the command surface
✔ WORK-0018 · Extract the application layer

`luma-backlog work-item list --status <status>`
```

**The heading is the label and a count in parentheses** --- *Closed (13)*, not
*Closed — 13 work items*. The parenthetical scans as a label and does not repeat
the noun the rows already are.

**Add a tally after the count where the states differ** --- *Tasks (23) — 9
delivered, 14 not started* --- so the split is read once rather than counted.

**Records with no key drop it**, since outcomes and tasks have none:

```
○ Every command is noun then verb
✔ Build the noun-verb command tree
```

**The command goes directly below the list, always** --- empty results included.
Below, never above: the answer comes first and provenance follows it.

**An empty listing keeps its heading and its command.**

```
**In Progress (0)**

`luma-backlog work-item list --status in_progress`
```

Nothing else. *"Nothing is in progress"* under the command that looked is a
fact; without the command it is indistinguishable from a wrong question.
