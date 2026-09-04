---
type: type_definition
defines: work-item
fields:
  kind:            {field_presence: optional, field_type: enum, values: [defect, request, idea, change], desc: "What sort of work item this is. A kind says what has to happen before the record can be judged; see the body. Absent means nobody has classified it, which is not the same as `change`."}
  workflow_status: {field_presence: recommended, field_type: enum, values: [captured, unprepared, preparing, prepared, todo, in_progress, closed], desc: "Where the work is. Absent means the first configured value — captured. Configurable per repository; the tool attaches no meaning to the values. See docs/workflow-status.md."}
  blocked:         {field_presence: optional, desc: "Present means blocked. A list of { on, why}, or a single entry written bare. Undeclared shape — the format has no composite field type yet." }
  paused:          {field_presence: optional, desc: "Present means deliberately paused. { on, why}. Undeclared shape, as above." }
---

# Work item

The unit of delivery, and the thing that sits on a backlog. Judged on its outcomes and never on its tasks.

**Body sections:** *The problem* · *What is being delivered* · *Out of scope* · *Constraints*. Leave one out rather than writing nothing under it.

## Three kinds, and the test for a fourth

**A kind says what has to happen before the record can be judged.** That is the
test, and it is why the list is short.

| | |
| --- | --- |
| **`defect`** | Something does not work, and nobody planned for it. A desired state was already supposed to hold and does not, which is a different shape from declaring a new one. **Judgeable now:** is it worth fixing? |
| **`request`** | Somebody asked for something nobody had thought of. **Judgeable now:** is it a real problem, is there a solution, does it align, is it worth doing — and it may be answered *no*, which ordinary work never is. You simply do not do ordinary work; you decline a request. |
| **`idea`** | Neither of the above, and **not judgeable yet**. A thought worth not losing, whose capture is not finished: somebody has to develop it before there is anything to evaluate, and what it becomes is one of the rows above or ordinary work. |
| **`change`** | **None of the above.** Nothing broke, nobody asked, and it is formed enough to judge. Most of what a team builds. |
| *absent* | **Nobody has classified it.** Not a fourth state — a missing answer. |

**`idea` is not a restatement of `workflow_status: captured`.** The rung says
nobody has *decided*; the kind says the record is not a complete statement of
anything yet. A bug at `captured` is fully described and merely unjudged. An idea
at `captured` is neither.

**And `idea` is not a peer of the other two — it sits upstream of them.** A bug
or a request arrives judgeable, one gate away from `unprepared`. An idea has to
be developed first, and what it becomes is a bug, a request, or ordinary work
with no kind at all. **So the kind changes**, which no other kind does, and that
is the tell: `idea` describes how finished the capture is, while `bug` and
`request` describe what the work is.

**That development happens above the first gate**, which is another reason the
captured zone is a zone and not a moment — and why the distance to `unprepared`
is not the same for every record.

**`change` is defined by exclusion, and that is why it reads weakly.** The other
three each carry a fact: something broke, somebody asked, it is not formed yet.
`change` carries none — it is the remainder, *work that is none of the other
three*. All work changes something, so the word is true; it is just not doing the
work the others do.

**It was chosen as least-bad rather than argued for**, and a future reader should
know that rather than assume it was reasoned to. The search failed for a
structural reason: a negatively-defined category has no positive noun, so every
candidate named something narrower than the category.

Two failed on opposite halves of it. **`improvement`** cannot cover creation —
building a first version improves nothing. **`original`** covers creation and
reads wrong for a rename. `opportunity` and `elective` were serviceable and
carried sales and medical flavors respectively. `own`, `native` and
`self-originated` describe origin rather than the work, and every stance word —
`chosen`, `planned`, `committed`, `intended` — fails because the stance applies
to all four once the work is accepted.

**The live cost:** `spec.md` uses *change* 73 times as ordinary English, and each
one is now slightly ambiguous. **Re-open when a better word turns up**; nothing
depends on this one beyond the value itself.

**Why four is the whole set, checked the way `spec.md` §2.1 checks units.** Three
name what has to happen before a record can be judged — fix it, answer them,
develop it — and the fourth is everything already judgeable that arrived by none
of those routes. Between them they cover every state a record can be in on the
way to the first gate, which is a completeness argument rather than a preference
for short lists.

**`story` is not a kind.** It is a narrative template for *describing* work, not
a statement about where the work came from or what has to happen next. A team
that says *story* is relabeling the unit (`spec.md` §2.1), and ADR-0001 declined
the word as the unit's own name precisely because it carries that template.

**Internal against external is not the line; *is anybody owed an answer* is.** A
request from another department and one from a customer are both requests, and
`request` is the right word for both — it is external to the team even when it is
internal to the organization. Work somebody on the team decided to do is **not**
a request, and recording it as one would make an author role-play a requester.
That is the strain that set aside `request` and `ask` as names for the unit
itself (`open-questions.md` §16).

**Absence now means what it says.** Before `change` existed, a blank field meant
either ordinary work or an unclassified record and nothing told them apart. It
means unclassified, and ordinary work says `change`.

**What would earn a fourth.** Something whose next step is none of *fix it*,
*answer them*, or *develop it*. Candidates that look like kinds and are not:
a **question** is a request for information; an **incident** is a bug plus
urgency, and urgency is a different axis; **debt** and **chores** are ordinary
work with a mood attached.

**`issue` is not a kind.** It names where a record came from, not what it is — a
record arriving from an external tracker may be any of the three, or none. That
belongs on a provenance axis, and putting it beside `bug` would import a source
system's vocabulary as though it were a distinction we need.

**Internal and external requests are the same kind.** Who asked is provenance,
and the tool treats them alike. An organization will not: an external request
may need a reply, a public status, a promise about time. **That difference is
ADR-0001's recorded trigger** for splitting requests from work items — an intake
population distinct from the people working the backlog, needing its own
lifecycle of answered, declined and duplicate. Until that population exists, one
kind and a provenance field is the smaller answer.

## What this deliberately does not declare

**Core fields** — `title`, `description`, `created`, `modified`, `stage` — are inherited and must not be restated here.

**Outcomes, waves, and tasks are not listed.** Membership lives on the member: a task names its work item, never the reverse. A field here would be a second copy of the same fact, and the two would disagree.

**`blocked` and `paused` carry no `field_type`** because the format has none that fits a small map. They are described rather than typed, which is legal — a Type Definition publishes intent rather than enforcing it, and the format does the same with `sources`.

> **Written from one record.** This describes `work items/first-usable-build/`, the only work item that exists. It is a description, not a prediction, and it should grow as real records need more — not in advance of them.
