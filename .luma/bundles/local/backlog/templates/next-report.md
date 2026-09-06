# What's next template

**Where the work stands, then an opinion.** Copy the shape, not this file.
Listings inside it follow [[listing]].

```
## THE BIG PICTURE

### Last touched

<prose: which record, roughly when, and what it was>

### In Progress (0)

`luma-backlog work-item list --status in_progress`

### To Do (0)

`luma-backlog work-item list --status todo`

### Preparation considerations

○ WORK-0031 · Reshape the command surface
○ WORK-0039 · What closed work items cost as the corpus grows

<one sentence, at the bottom, on how close any of it is>

### Risks

<what you noticed that nobody asked about --- or omit the section>

---

## NEXT OPTIONS

1. **<the thing>** — <one clause of why>
2. **<the thing>** — <one clause of why>
3. **<the thing>** — <one clause of why>

---

## RECOMMENDATION

**<the one>** — <why it beat the others>. <what would change it>
```

## The rules the shape depends on

**No numbers on headings.** They are sections, not steps, and a reader who
arrives at *3.* wonders what they missed.

**Every heading is `###`.** One level, so the four sections above the rule read
as four sections rather than a hierarchy --- there is no containment between
them and a size difference would imply one.

**Last touched is prose, not a listing.** That is what sets it apart, and it is
enough: what happened is usually more than a record and a timestamp, and it is
first, which is the only prominence it needs.

**In Progress and To Do carry no prose.** Heading, rows, command. Nothing
explaining what the section is: a reader who cannot tell what *In Progress*
means will not be helped by a sentence saying it.

**Shown even when empty**, with their commands. Two empty listings *are* the
finding, and deleting them hides it.

**Preparation considerations** --- a label, like the headings above it, rather
than a sentence with a hedge in it. *Consideration* carries *not chosen* and
does not promise these are close, which *candidate* slightly does. Rows only.
**Up to three.**

**It is the last section above the rule**, because it is the one a reader with
work in progress skips. Everything they came for is above it.

**One sentence at the bottom if it is worth saying**, never in the middle: how
close any of this is to producing work. *"Neither has outcomes, so both are
several conversations away."* Below the rows, so somebody who only wanted the
names has already got them.

**Never offered as work.** It has not crossed the second gate, and offering it
mistakes a pile for a queue.

## Three divisions, named

**The big picture, next options, a recommendation.** They answer different
questions and a reader usually wants one of them, so each is named and a reader
can stop at the one they came for.

**The big picture** is where the work stands and what is worth worrying about.
Somebody checking in reads this and nothing else. **Risks belong here** --- a
concern is part of the picture, not part of the choice.

**Next options is a menu.** Options, ranked, and the reader chooses. **Not "where
to start"** --- nobody arriving at an existing backlog is starting, they are
continuing, and a heading that says otherwise is written for a project that
does not exist yet.

**Recommendation is one thing.** If the menu were enough there would be no need
to pick, and handing somebody three ranked options without saying which is
answering a question with a question.

**Divisions are uppercase; sections inside them are not.** `##` and `###` render
identically in a terminal, so heading level alone divides nothing there. Case is
what survives --- and it is what the tool's own help already uses, so this is the
existing convention rather than a second one.

**A `---` between divisions.** Three signals, each doing different work: an
uppercase heading says *new division*, a normal one says *section*, and the rule
is the hard break. The rules were dropped once on the argument that named
headings divide; they do not, in the place this is read.

**Not between the sections inside a division.** Five rules under the big picture
would be noise, and the case difference already separates them.

**Risks is omitted when there are none.** A manufactured concern costs more than
the section is worth and teaches a reader to skip it.

**Next options is one to five alternatives, ranked, a clause each.** Not a
plan --- a menu somebody can act on without reading twice. **The recommendation
is not among them.**

**Options, not candidates.** *Options* is the word somebody uses for a menu they
are choosing from; *candidate* implies the things are being assessed for
suitability, which is what the preparation list does and this one does not.

**Neither list asserts. Only the recommendation does.** Considerations are what
might be worth working out, options are what might be worth doing, and both are
handed over rather than urged.

**The recommendation is not repeated in the menu.** The menu holds what you
would do *instead*; the recommendation holds what you would do. Listing the pick
twice spends a line saying nothing --- a reader who has read one has read the
other.

**It does not have to come from the menu at all.** Sometimes the right answer is
none of the options, and forcing it into the list first, so it can be pulled out
again, is ceremony.

**Say what would change it** when something obvious would, usually a decision
nobody has made.

**Recommend nothing when nothing is right**, and say why. An empty
recommendation with a reason is worth more than a filled one without.
