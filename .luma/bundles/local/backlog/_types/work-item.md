---
type: type_definition
defines: work-item
fields:
  key:             {field_presence: recommended, field_type: text, desc: "The handle somebody quotes — WORK-0002. Allocated at creation from one project-wide sequence and written into the record, never derived, so a later change of prefix cannot rename what already exists. The path stays the identity for linking; the key is the identity for finding."}
  kind:            {field_presence: optional, field_type: enum, values: [defect, request, idea, inquiry, change], desc: "What sort of work item this is. A kind says what has to happen before the record can be judged; see the body. Absent means nobody has classified it, which is not the same as `change`."}
  workflow_status: {field_presence: recommended, field_type: enum, values: [captured, unprepared, preparing, prepared, todo, in_progress, closed], desc: "Where the work is. Absent means the first configured value — captured. Configurable per repository; the tool attaches no meaning to the values. See docs/workflow-status.md."}
  blocked:         {field_presence: optional, desc: "Present means blocked. A list of { on, why}, or a single entry written bare. Undeclared shape — the format has no composite field type yet." }
  paused:          {field_presence: optional, desc: "Present means deliberately paused. { on, why}. Undeclared shape, as above." }
---

# Work item

The unit of delivery, and the thing that sits on a backlog. Judged on its outcomes and never on its tasks.

**Body sections:** *The problem* · *What is being delivered* · *Out of scope* · *Constraints*. Leave one out rather than writing nothing under it.

## Five kinds, and what separates them

**What separates a kind is what it produces.** That is the test, and it is what
makes the set complete rather than merely short. The second column is what the
kind is *for*: scanning the backlog, it tells you at a glance how much vetting
is owed and of what sort.

| | produces | what it needs vetted |
| --- | --- | --- |
| `defect` | a fix | is it real, and does it reproduce |
| `request` | a change, and answers owed to whoever asked | is it legitimate, aligned, worthy |
| `idea` | a classification — it becomes one of the others | is it even work yet |
| `inquiry` | **understanding, and more work items** | is it worth looking into |
| `change` | the work itself | do we still need it |

| | |
| --- | --- |
| **`defect`** | Something does not work, and nobody planned for it. A desired state was already supposed to hold and does not, which is a different shape from declaring a new one. **Judgeable now:** is it worth fixing? |
| **`request`** | **A change somebody outside asked for**, and you are accountable to them. That is the kind's whole point: a request may be declined, which ordinary work never is — you simply do not do ordinary work — and somebody is waiting until it is settled. It carries the heaviest vetting, and three separate questions: is it legitimate, is it aligned, is it worthy. |
| **`idea`** | Neither of the above, and **not judgeable yet**. A thought worth not losing, whose capture is not finished: somebody has to develop it before there is anything to evaluate, and what it becomes is one of the rows above or ordinary work. |
| **`inquiry`** | **Going to look.** A review, an audit, an investigation, a spike. It exists to gain understanding — to de-risk something, to find out what is there — and what comes out is work items, or a report that generates them. It changes nothing itself, and finding nothing still counts as done, because the looking was the work. |
| **`change`** | **None of the above.** Nothing broke, nobody asked, and it is formed enough to judge. Most of what a team builds. |
| *absent* | **Nobody has classified it.** Not a sixth state — a missing answer. |

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

## The inversion, recorded and not taken

There are two coherent arrangements of the same four states, and this bundle
ships the first.

| | values | blank means |
| --- | --- | --- |
| **A — shipped** | `defect` `request` `idea` `change` | nobody has looked |
| **B** | `defect` `request` `idea` `needs_triage` | it is a change |

**B is ergonomically better and A is safer, and that is the whole trade.**

Under B the common case is free: ordinary work says nothing, and only the
exceptions carry a value. That matches how the field will actually be used,
since most of what a team builds is a change.

Under A blank never asserts anything. Under B it asserts *somebody looked at
this and it is none of the other three* — which is false for every record that
arrives unlooked-at. **An issue synced from another tracker would silently read
as a change**, and a wrong classification is worse than a missing one because
nothing about it looks wrong. B answers that with an explicit value that
importers must write; A needs nothing written to stay honest.

**`needs_triage` rather than `undetermined`, if B is ever taken.** The other
three name what has to happen before the record can be judged — verify it,
answer them, develop it — and `needs_triage` is that shape: triage it.
`undetermined` names a state instead, leaving one value in the set answering a
different question from the rest. It is also the good version of naming the
value `issue`: it keeps *requires triage* and drops the conflation with a
container that holds defects, requests and questions alike.

It does not fix B's safety, only softens it. Blank still asserts *change* for
anything nobody looked at; `needs_triage` makes the write an importer has to
remember an obvious one rather than an admission.

The same reasoning settled the `workflow_status` default: a record was born
`idea` and asserted doubt about work somebody had just decided to do. **Absence
should make the smallest claim available**, and *nobody has looked* claims less
than *this is a change*.

**Re-open when real use says the field is empty anyway.** If people do not reach
for `--kind change` and most records sit blank, A is carrying no information at
the cost of a flag nobody types, and B is simply the truth about how it is being
used. That is a measurement, not an argument — count the blanks after a few
months.

**The test used to be different, and could not hold five.** An earlier version
sorted them by *what has to happen before the record can be judged* — verify it,
answer them, develop it, nothing. That derived the first four and cannot place an
inquiry, which is judgeable the moment it is written: *is this worth looking
into?* is answerable now. What distinguishes it is not what it needs first but
what comes out of it, which is why the test is now what each kind produces.

**`inquiry` is the only kind that predicts a relationship.** When it is written
it has no children; the kind claims it will produce some, and `sources` on those
children attests later that it did. Those are not redundant — one looks forward,
one records — and it makes this the one kind whose claim can be checked against
the corpus rather than believed.

**Finding nothing can be success here and is failure everywhere else.** A
*confirmed* defect that produces no fix is not delivered; an audit that finds no
problems can be. Both qualifiers matter — an unconfirmed defect producing no fix
was never a failure, and an audit can still fall short for reasons of its own.
That inversion is about completion, which is the thing this tool computes, and it
is the sharpest reason `inquiry` is a kind rather than a flavor of `change`.

**`review`, `audit`, `investigation` and `spike` are stored as `inquiry`.** They
are **instances** of it rather than synonyms for it, so unlike `bug` against
`defect` the alias loses a shade of meaning. Accepted, because the alternative is
worse: without it some records say *spike* and some say *inquiry*, and the filter
that earns the whole field stops finding half of them. What kind of looking it
was belongs in the title, where it fragments nothing.

**Why not `question`, `investigation` or `exploration`.** `investigation` is one
instance promoted over the others — a spike is not an investigation, and both are
inquiries. `exploration` is this project's word for the same idea one level down,
where it is a **unit**, and using it here would put one word at two altitudes.
`question` was the strongest rival: it names what you hold, which is the grammar
of `defect` and `request`, and *file this as a question* reads better than *file
this as an inquiry*. `inquiry` won on naming the finding-out rather than the
asking, which is what generates the work.

**Not settled beyond the concept.** English seems to handle this category as a
compound — *research question*, *open question*, *due diligence*, *feasibility
study* — and every single word is either one instance promoted or a vague
abstraction. **Re-open when a better one turns up.**

**`story` is not a kind.** It is a narrative template for *describing* work, not
a statement about where the work came from or what has to happen next. A team
that says *story* is relabeling the unit (`spec.md` §2.1), and ADR-0001 declined
the word as the unit's own name precisely because it carries that template.

**Internal against external is not the line; *are you answerable to somebody* is.**
A request from another department and one from a customer are both requests, and
`request` is the right word for both — it is external to the team even when it is
internal to the organization. Work somebody on the team decided to do is **not**
a request, and recording it as one would make an author role-play a requester.
That is the strain that set aside `request` and `ask` as names for the unit
itself (`open-questions.md` §16).

**A change is further along than a request the moment it is written**, because
your own team wrote it. Both still get vetted; the questions differ. A change is
asked *do we still need this*. A request is asked three: whether it is
legitimate, aligned and worthy — and somebody is waiting on the answer, which no
change carries.

**Reading `change` as *needs no vetting* is the misreading worth heading off.**
The lighter question is not an absent one.

**Absence now means what it says.** Before `change` existed, a blank field meant
either ordinary work or an unclassified record and nothing told them apart. It
means unclassified, and ordinary work says `change`.

**What would earn a sixth.** Something that produces none of a fix, an answer, a
classification, more work, or the work itself. Candidates that look like kinds
and are not: an **incident** is a defect plus urgency, and urgency is a different
axis; **debt** and **chores** are changes with a mood attached.

*A **question** was on that list, on the grounds that it is a request for
information. That was wrong, and it collapsed two things: a question somebody
asks you, whose answer is a decision you have the standing to make, and a
question you are asking, whose answer you can only get by doing work. The second
is an inquiry.*

**`issue` is not a kind.** It names where a record came from, not what it is — a
record arriving from an external tracker may be any of them, or none. That
belongs on a provenance axis, and putting it beside `bug` would import a source
system's vocabulary as though it were a distinction we need.

**Internal and external requests are the same kind.** Who asked is provenance,
and the tool treats them alike. An organization will not: an external request
may need a reply, a public status, a promise about time. **That difference is
ADR-0001's recorded trigger** for splitting requests from work items — an intake
population distinct from the people working the backlog, needing its own
lifecycle of answered, declined and duplicate. Until that population exists, one
kind and a provenance field is the smaller answer.

## The key is for finding, the path is for linking

**Both are true at once, and they answer different questions.** A wikilink
resolves against a path and breaks when a file moves. A person or an agent
searches by key and does not — the key travels with the record because it is
written into it.

That is the same split the decision records already run on, and the reason a
superseded decision stays findable after it moves into `archived/`. Work items
had only the path, so every reference was a slug derived from a title, and
changing a title meant either breaking inbound links or leaving a slug that no
longer matched the record.

**It is `key` rather than `id`, and the reason is the sentence above.** This
format says a record's identity *is* its path (`spec.md` §7.1, §4.1). A field
called `id` would claim that identity, and then two things would claim to
identify one record — which is the shape of failure this project keeps designing
against. `key` makes the smaller and true claim: a handle for finding, not an
identity. It also matches what the major trackers call the same thing, so nobody
has to be taught it.

*Considered and not taken (2026-09-04): `id`. Reopen if the format ever stops
treating the path as identity, which would make the objection disappear.*

**Four digits, matching the ADR numbers.** Padding exists so a lexical sort
matches a numeric one, which is what `ls`, git and an editor give you. It stops
working past 9999 — and by then nobody is reading a directory of ten thousand
work items by eye, so the property fails exactly where it had stopped being
worth anything. Padding buys legibility in the range you browse, and that range
is well inside four digits.

*Considered and not taken: five digits, on the argument that a sort break is
unfixable once keys are immutable. It is unfixable and it does not matter — and
if the order ever did matter at that scale, the fix is sorting numerically in
the tool, which touches no record and disturbs no key.*

**`WORK` is the only prefix, and it is written rather than derived.** A
repository that later wants its own three letters changes what gets written from
then on; nothing already on disk is renamed. A derived key would silently
rewrite the whole corpus the moment the setting changed, which is the failure
that made this worth being careful about.

**Allocated in one pass at creation**, from one sequence for the whole backlog,
with the accepted cost the ADR numbers carry: two branches can both claim the
next number and somebody renumbers on merge.

**A record written before keys existed has none**, and that is not an error. The
field is `recommended` rather than required for exactly that reason.

### Written and said as one string

`WORK-0002-lint-the-corpus` — the key and the slug joined, the way a decision's
filename joins its number and its slug. That is the form to use in prose, in a
commit message and out loud, and all of it resolves: the joined form, the key
alone, and the slug alone, with the key half matched case-insensitively.

**One identifier rather than two columns.** A reader wants one thing to copy, and
a listing with a key column leaves an empty cell on every record that carries
none — only a work item has a key, so an outcome's identifier is its slug and the
column is never blank.

**And it is the directory name.** `work-items/WORK-0002-lint-the-corpus/` — the
key leads so a listing sorts by it, and the slug follows so the directory still
reads as what the work is. That matches the decision records, where the number is
in the filename too.

**Migrating cost 32 link repairs, and it was done while there were 32.** Every
outcome and task carries `work_item: [[work-items/<directory>]]`, so a rename
touches every child of every work item — and that number only grows. `spec.md`
§7.1 is explicit that a record's identity is its path, so this changed what each
record is, once, deliberately, with the links repointed in the same commit.

**Three forms reach the same work item**: the directory, the slug half alone, and
the key. Requiring the long form everywhere would make the key a tax rather than
a handle, and the short forms are what people typed before it existed.

## What this deliberately does not declare

**Core fields** — `title`, `description`, `created`, `modified`, `stage` — are inherited and must not be restated here.

**Outcomes, waves, and tasks are not listed.** Membership lives on the member: a task names its work item, never the reverse. A field here would be a second copy of the same fact, and the two would disagree.

**`blocked` and `paused` carry no `field_type`** because the format has none that fits a small map. They are described rather than typed, which is legal — a Type Definition publishes intent rather than enforcing it, and the format does the same with `sources`.

> **Written from one record.** This describes `work items/first-usable-build/`, the only work item that exists. It is a description, not a prediction, and it should grow as real records need more — not in advance of them.
