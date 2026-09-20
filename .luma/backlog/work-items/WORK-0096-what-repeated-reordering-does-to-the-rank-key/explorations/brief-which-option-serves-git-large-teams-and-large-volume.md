---
type: exploration
type_version: "0.0.1"
title: 'Brief: which option serves git, large teams and large volume'
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-18T03:21:21Z'}
---

# Brief: which option serves git, large teams and large volume

> **Superseded by `brief-vet-the-three-designs`.** This covered two options and
> was shaped around one question. It is kept as the record of what was actually
> put to the reader whose answer is filed as
> `second-opinion-on-git-teams-and-volume` --- a brief and the answer to it
> should be readable together.


> **A self-contained brief for a fresh reader.** Everything needed is here.
> **Nothing else in this repository should be read** --- the rest of it contains
> a week of our own opinions, and the point of asking somebody else is to get an
> answer that is not ours.
>
> **We have a preference. We are deliberately not telling you what it is.** If
> you want it afterwards, it is in this work item's journal --- **after** you
> have formed a view, not before.
>
> **If the question is wrong, say so.** If both options are bad, if there is a
> third we have not considered, or if one of our measurements does not mean what
> we think it means --- that is more useful than an answer to the question as
> asked.

## The question

**Two designs are on the table for keeping work items in order. Which one serves
git, large teams, and a large volume of work better --- and why?**

Not which is more elegant. Which one would you rather be running with fifty
people and two million work items, and what breaks first.

## The situation

A board of work items, grouped into columns --- roughly: captured, being
prepared, ready, in progress, done. **Within a column the items have an order
somebody chose.**

**Each work item is its own markdown file in a git repository**, with key-value
fields at the top. There is **no server and no database**. People and agents
reorder the board independently, on different machines and in different
branches, and discover each other's changes when git merges.

**Two columns grow without limit**: the intake column, because new items never
stop arriving, and the done column, because work keeps finishing.

**Some items never move, and they sit in the middle of a column.** Something
captured years ago that nobody will pick up and nobody will delete. Work is
reordered around it indefinitely, and **it cannot be renumbered, because
renumbering it means rewriting files that did not change** --- and a large diff
of ordering data cannot be reviewed by a person.

**Scale to design for:** a high-volume project doing ten thousand items a day
for a hundred years is 3.65 × 10⁸ items. Assume operations at both ends of a
column and in any gap inside it.

---

## Option 1 --- an address in each item

**An item's position is a list of numbers**, like a street number, then an
apartment, then a room. To fit an item between two others, go one level deeper
and count there.

| step | address | what is stored in the file |
| --- | --- | --- |
| three items created at the back | `[1]` `[2]` `[3]` | `n1-w3` `n2-w3` `n3-w3` |
| insert between the first two --- no whole number fits, so open a level below | `[1, 0]` | `n1m-w3` |
| insert again --- **count downward** | `[1, -1]` | `n1l8-w3` |
| and again | `[1, -2]` | `n1l7-w3` |

**Inserting repeatedly at one spot becomes counting, not gap-halving.** That is
the whole idea.

The front of a column goes negative --- `[-1]`, `[-1000000]` --- so there is no
floor.

**Each address ends with a short tag naming who wrote it**, so two people
computing the same address store different values.

**The letters exist only so that text sorting matches number sorting**: sorted
as text, `10` comes before `9`, so a letter states how many digits follow.
Negative numbers use letters below `m` with their digits inverted.

### Measured

| | |
| --- | --- |
| a million insertions in front of an item that never moves | **12 characters** |
| a million items added to the front of a column | 11 characters |
| 20,000 mixed operations including moving items | order held, no duplicates, worst value **9 characters** |
| sorting | matches `LC_ALL=C sort` exactly, over 84,030 keys |
| 300 trials x 120 random adversarial insertions | order held, worst value 17 characters |

### Known costs and unknowns

- **Nobody has shipped this.** The nearest published algorithm was measured and
  costs **2 characters per insertion** on the never-moves workload --- 200
  million characters at 10⁸ --- because it appends to a path where this counts.
  **Adopting this means owning an ordering algorithm.**
- **The stored value is unreadable to a person.**
- **One pattern is still linear:** repeatedly inserting between your own two most
  recent insertions. Nothing on this board does that.
- **The implementation is a sketch.** Its counter has a ceiling that crashes
  rather than degrading. The idea is tested; the code is not.
- Behaviour beyond 10⁶ at one spot is extrapolated from a clean logarithmic
  curve.

---

## Option 2 --- one ordered file per column

**The column's order is a list of item names in a file. The order is the line
order. Moving an item means moving a line.**

```
upgrade-the-parser
fix-the-login-bug
write-the-docs
```

**Only items somebody deliberately placed are listed.** Everything else is
ordered by a timestamp the item already carries. So **creating an item writes
the item and nothing else**, and **advancing an item to another column writes
the item and nothing else** --- both land at the back of their column by
timestamp. Only a deliberate reordering touches the file.

**That is what keeps the unbounded columns out of it**: intake's untriaged tail
and the done column are never deliberately ordered, so they have no file at all.

### Evidenced

- **Every plain-text board tool works this way.** Two independent literature
  searches found that no git-native project allocates positional values; they
  all use line order and let git merge it.
- **There is no growth curve to measure**, because there is no value.

### Known costs and unknowns

- **An item is no longer self-contained.** You need the column's file as well as
  the item to know where it sits.
- **The file is a contention point.** Every deliberate reordering in a column
  touches one file.
- **Two mechanisms instead of one**: a file for placed items, timestamps for the
  rest.
- **Unknown:** how often real concurrent reordering of one column actually
  collides.
- **Unknown:** whether git's line merging can ever reorder lines *wrongly* under
  an unlucky three-way merge, rather than duplicating or dropping one.

---

## What git actually does --- measured, not assumed

**Real merges, run for both options.**

| two people, concurrently | ordered file | addresses |
| --- | --- | --- |
| move different, distant items | **clean**, both moves applied correctly | **clean**, both applied |
| move **adjacent** items | **conflict** | **clean**, both applied |
| move the **same** item | **conflict** | **conflict** |

**Addresses conflict when two people write the same item's file.** The ordered
file conflicts for that too, and additionally when two people move neighbouring
items --- which is compatible work.

**The one case where addresses stay quiet:** two people inserting *different*
items at the *same* spot. Both items exist; the only ambiguity is which of the
two comes first, which neither person expressed an opinion about.

---

## What we want from you

**Answer on three axes and say which dominates:**

1. **git** --- merge behaviour, diff reviewability, what a conflict costs a
   person to resolve, and anything either design does to git that we have not
   measured.
2. **large teams** --- fifty people and agents working in parallel branches and
   worktrees. Contention, false conflicts, lost intent, and whether *quiet but
   traceable* is acceptable or whether only *loud* is.
3. **large volume** --- 10⁸ items and operations. What degrades, when, and what
   the repair costs.

**And tell us what we have got wrong.** Three confident claims in this work item
have already been disproved by running something rather than reasoning about it,
and two of those were about merge behaviour.

**If you want to measure something, measure it.** Every number above came from
running code, and none of it should be taken on trust.

## How to write it

**Two rules about the output, because they are ours and you could not otherwise
know them.**

- **American spelling.**
- **Do not name a competing product.** Tools whose technique is being borrowed
  or that we interoperate with are named normally; a rival is not. If a
  comparison needs one, describe it rather than name it.
