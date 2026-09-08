# Journal — The backlog has no unit for a delivery

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-08

created from a live instance rather than from theory — the maintainer asked what has to be true before another project can use this tool, the answer was eight or ten pieces of work, and there was nowhere to record that as one thing except an eleventh work item
the maintainer ruled out dimensions before the research began, and spec 2.7 agrees in its own words: dimensions carry no mechanics of their own and for the MVP they classify and nothing more — a delivery has to say whether it is reached, which a classification axis cannot
milestone imports sequence and was rejected for it; epic and project do not, and can still be sequenced deliberately — the unit wanted is orderable but not ordered by construction
found while writing: spec 2.3 explains a wave as answering how many attempts a DELIVERY needs, so the specification already refers to this thing and has never defined it — a wave is the iteration and the thing iterated toward is unnamed
the fork in this record was answered before the research began: it is a DIMENSION, but a special one that carries metadata the system tracks — not a sixth record type, and not all dimensions gaining mechanics, but two classes of them
the reason it must stay a dimension is naming — people will call it a project, epic, milestone, release, initiative or phase and every one is right for somebody; a type called milestone forces a team that says release to rename a built-in, which is the interop mismatch 2.7 removed when project stopped being the unit name
AGENT: this project already has the mechanism one level down — workflow_status is configurable vocabulary carrying no meaning to the tool while the ladder's mechanics are the tool's; a tracked dimension is to epic and milestone what workflow_status is to todo and closed
containment corrected: a wave is INSIDE a work item, an attempt at a set of outcomes, and this is what a work item is inside of — opposite directions from the same record, and neither answers for the other
folded in: spec 2.7's 'under consideration' — whether project, epic and milestone ship as default-defined dimensions — now names this work item and says why the two are one decision; those three are exactly the candidates for being TRACKED, so defaulting them and defining them cannot be settled separately
WORK-0070 is the live instance — six pieces of work adding up to one delivery, recorded as a seventh work item because this research has not produced the unit yet; it says so on its own face rather than pretending to be a piece of work
LEADING ANSWER, maintainer: a work item should belong_to other work items and may belong to many — a relation rather than a unit, which replaces both the tracked dimension and the sixth-kind options
it gets 2.7's multi-axis property for free without declaring any axes, since belonging to several work items IS a record sitting on several at once; and the container's own title supplies the name, so no enum and no configuration ships a vocabulary
container-ness is DERIVED — a work item is a container if anything belongs to it — which needs no kind, no flag and nothing kept in sync, matching how outcome state and resolved_at already work
the sixth-kind option actually PASSES _types/work-item's test, which was a surprise: a container produces none of a fix, an answer, a classification, more work or the work itself; it fails on layout and on naming instead, since kind is an enum and shipping epic makes a team that says release file epics
maintainer floated WORK-EPIC-0001 for container keys — recorded as open with the argument against, that a key saying EPIC ships the vocabulary this approach exists to avoid, and a state mark is the cheaper answer since container-ness is already derivable
deepest open problem: 2.2 says workflow status is DECLARED, and a container's state is naturally DERIVED from its members — a release at in_progress while every member is captured is a lie nobody wrote on purpose
maintainer: epics may need to track work items from other projects or other backlogs — a delivery routinely spans projects
that breaks the cheapest part of the design: container-ness stops being DERIVABLE, because 'a container is anything something belongs to' only holds where everything that could belong is visible, and a foreign member pointing here is invisible from here
direction problem the local case does not have — belong_to on the member means a foreign member must know the container and the container learns nothing, while inverting it spans corpora and loses the free derivation; possibly both, derive locally and declare across, which is two mechanisms for one relation
same wall the dependency discussion hit earlier today: keys are per-corpus so a foreign reference needs a namespace that does not exist, spec 6.1 forbids a coordinator, and a foreign member reopening changes a container's completion with nobody touching the record — reading another corpus's state is a capability this tool has never had and now more than one thing wants it
