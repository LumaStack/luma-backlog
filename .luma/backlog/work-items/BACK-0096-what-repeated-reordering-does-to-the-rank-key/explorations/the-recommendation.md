---
type: exploration
type_version: "0.0.1"
title: The recommendation
work_item: '[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-18T02:58:17Z'}
---

# The recommendation

> ## Superseded on 2026-09-17, by measurement
>
> **The recommendation below was to adopt position-strings. It is wrong, and the
> measurement asked for before deciding is what caught it.**
>
> **Position-strings grows logarithmically when appending and linearly when
> inserting in front of a card that never moves** --- which is the workload this
> whole work item exists for. Measured against the real library
> (`npm install position-strings`), not a reimplementation:
>
> | workload | result |
> | --- | --- |
> | append at the back, 10,000,000 times | **11 characters** --- logarithmic, exactly as published |
> | insert in front of a card that never moves | **2.00 characters per insertion, flat, forever** |
> | insert into one interior gap | 2.00 characters per insertion |
> | same, but the fixed card is one I created | 2.00 characters per insertion --- ownership is irrelevant |
>
> At 20,000 insertions the value is **40,011 characters**. At 10⁸ it would be
> **200 million characters** --- and an attempt to run 10⁷ exhausted 3.5 GB of
> heap and killed the process.
>
> **Why: the library's own documentation says the optimisation is for a
> "left-to-right sequence."** Its `createBetween` has a case for *"left child of
> right"* which the source comments **"this always appends a waypoint"** ---
> so each leftward insertion lengthens the path by a whole waypoint rather than
> incrementing a counter. Position-strings was built for collaborative text
> editing, where typing forward dominates. **Our hardest case is the one it does
> not optimise.**
>
> **I generalised the published claim, and so did the literature search that
> found it.** Neither of us checked. That is the third confident
> generalisation in this work item to be wrong, and the first that would have
> been shipped.
>
> ### What measurement says instead
>
> | scheme | in front of a card that never moves, 10⁶ times | sorts with `sort`? |
> | --- | --- | --- |
> | **the stepping key derived here** | **12 characters** | **yes** --- verified over 84,030 keys |
> | Stern--Brocot mediants | 15 characters | **no** --- comparing fractions needs cross-multiplication |
> | position-strings | **~2,000,000 characters** | yes |
> | our current decimal scheme | refuses after 204 | yes |
>
> **The blind derivation beats the published state of the art on this workload**,
> and the reason is legible: position-strings *appends a path segment* when you
> insert to the left of a foreign card, while the derived scheme *descends once
> and then counts*. Counting is what makes it logarithmic.
>
> **That is a real contribution rather than a lucky guess** --- requirement 7 was
> written into the brief as the thing that decides this, so the agent designed
> for a workload the published work was not built for.
>
> ### What this changes
>
> **Do not adopt position-strings.** It is excellent work aimed at a different
> access pattern.
>
> **The stepping key is now the leading candidate on evidence** --- with the
> caveat that its implementation is a sketch whose counter ceiling crashes
> rather than degrading, found by ten minutes of adversarial testing. **The idea
> is validated; the code is not.**
>
> **A modified position-strings would also work** --- make the leftward case
> count instead of appending --- but that is no longer *adopt published work with
> a proof*. It is publishing a variant of our own, with the risk that carries.
>
> **Everything below is retained as written**, because the reasoning that led to
> the wrong answer is worth as much as the answer.

---


**Adopt a published algorithm called position-strings. Stop designing our own.**

Each card keeps one text value. Sorting those values with `sort` gives the
order. Nothing ever needs renumbering, at any volume, in any position.

---

## The problem, in plain terms

Every card has to know where it sits in its column, written in its own file.
There is no database to ask, and two people may reorder the board at the same
time on different machines and then merge.

**The obvious approach is numbering**: 10, 20, 30. Put something between 10 and
20 and you use 15. Again, 12. Again, 11. Now you are out of whole numbers, so
you go to decimals: 10.5, 10.25, 10.125.

**Each insertion halves the space left, so each one needs more digits than the
last.** That is what we do today. Measured, it stops working after about 200
insertions in one place.

## Why numbers were the wrong idea

**Halving is the problem --- not decimals, and not how long the numbers are.**

If you could *count* instead of halve, you would never run out: a million steps
is a seven-digit number. The question is how to arrange things so that you are
always counting.

## What position-strings does instead

**Think of house addresses.** Rather than squeezing a number between 10 and 11,
you write *"10, apartment 3"*. That gives you a fresh set of whole numbers
inside 10 to count through. If that ever fills up: *"10, apartment 3, room 5"*.

**You add a level instead of subdividing a gap.** Adding a level is cheap.
Subdividing is what runs out.

**Two details make it work in our situation.**

Each level is labelled with **who created it**. So when two people add a card at
the same spot at the same time, they produce two *different* addresses rather
than the same one --- both valid, and each traceable to whoever wrote it.

And an address can go **below** the first card as well as above the last ---
*"apartment −3"* --- which is why the front of a column never runs out. That was
the actual defect that started this: our front had a floor at zero and nothing
to step down into.

## Why I trust this answer

**It is not our idea, and that is the point.**

- **Three independent reasoners and two independent literature searches**, none
  able to see the others, all arrived at the same underlying result: you cannot
  have a short sortable value, no renumbering, *and* unlimited insertion in one
  spot. Something has to give. The literature has known this since 1981 and has
  proved the limits.
- **Two of our three reasoners independently reinvented position-strings** ---
  working from the problem alone, with searching explicitly forbidden. They
  produced the same structure: path levels, a per-writer label, counting rather
  than halving. **We arrived at the field's answer twice without looking at
  it.**
- **It is published work with a proof.** Matthew Weidner's Fugue, and *The Art
  of the Fugue* with Martin Kleppmann, prove that concurrent insertions cannot
  interleave --- a property nobody had thought to ask for and which matters for
  bulk moves.
- **There is a written specification, a reference implementation and
  benchmarks.** Real-world traces produce values of 10 to 100 characters.

**The alternative was one of our agents' clever sketches**, and an
adversarial test found a case where it crashes rather than degrading --- ten
minutes of looking. Adopting published work with a proof is a different quality
of decision.

## What it costs

**Three things, and they should be accepted knowingly.**

**The value stops being readable.** Today a card says `010.0020.000`. It would
say something like `r95sU223.B7`. Sortable, unique, traceable --- and meaningless
to look at. **This is the real cost**, because these files are read by people in
text editors.

**Concurrent insertions at one spot merge quietly.** Two people adding a card at
the same place produce different, valid, traceable addresses --- and git says
nothing, because they edited different files. The order between those two cards
is consistent but chosen by nobody.

**It is worth being precise about why that is acceptable.** The failure that hurt
us was two cards ending up with the *identical* value and the order silently
lost. **That cannot happen here** --- the per-writer label makes identical values
impossible, and the condition is findable afterwards. What remains is two cards
in an order nobody picked, with nothing lost.

**Existing cards need renumbering once.** One large, mechanical change. We
already have the command that does it.

## What we still do not know

**One thing, and it is precisely located.** The published benchmarks cover
collaborative text editing. **Nobody has characterised what happens after 10⁸
insertions into a single gap** --- which is the volume this board is being
designed for. Growth is expected to stay logarithmic and has not been measured
there.

**We can measure it ourselves.** The algorithm is written down; the measurement
is a short program. **This should be done before committing**, because
extrapolating a curve is exactly the kind of assumption that has already been
wrong twice in this work item.

## The alternative, if one thing is non-negotiable

**If git must stop and force somebody to choose whenever two people reorder the
same spot, then no per-card scheme can ever satisfy that** --- two cards are two
files, and git has nothing to conflict over.

In that case the answer is different: **keep the order in one small text file
per column**, listing card names one per line. Moving a card means moving a
line. Two people editing the same lines get a real conflict. Every plain-text
board tool in existence works this way, which is genuine field evidence.

Its cost is that **a card no longer knows where it sits** --- you need the list
file as well as the card --- and that two people *moving* the same card can
quietly duplicate a line.

**That is the whole decision**: quiet-but-traceable with self-contained cards,
against loud-but-the-order-lives-elsewhere.

## What happens next

1. **Answer one question.** Must git stop and ask, or is *different, traceable
   and findable afterwards* enough? This picks the family and nothing else can
   be settled first.
2. **Measure the unknown** --- 10⁸ insertions into one gap, against the
   published algorithm.
3. **Implement it** behind the one function that allocates positions today, so
   the change is contained.
4. **Renumber the existing cards** once, with the command that already exists.
5. **Write the decision down**, naming what was given up.
