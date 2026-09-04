---
type: bundle
title: local/backlog
version: 0.7.1
stage: draft
consumers: [project]
description: The record types this project defines and the procedures for writing them — what luma-backlog knows about its own corpus, kept where the tool can read it.
---

# local/backlog

This project keeps its backlog in `.luma/`, which means it is both the tool and
the tool's first corpus. The types that corpus conforms to have to live
somewhere a harness can reach, and until now they lived in two places — the
types in a bundle named for the project, the procedures in `.claude/skills/`.
Neither location survived being looked for.

**It is local rather than published because nothing here is settled enough to
adopt.** `work-item` reached its name on 2026-09-02 and the formation bar it
implies is deliberately undecided. A bundle in a catalog is a promise to
somebody who did not write it; this is a promise to ourselves, and the `local/`
namespace is what says so.

**What earns promotion out of here is a second consumer.** Until a project that
is not this one wants these types, publishing them would be guessing at what
they need.

## What is here

**Types**

- `_types/work-item` — the backlog unit. Named by decision record
  ADR-0001 in `.luma/records/decisions/`, which is worth reading before
  proposing a rename.
- `_types/outcome` — the condition that must hold for a work item to be done.
  True or false, never a task in disguise.

**Procedures**

- [[backlog-new]] — scaffolds a record so every one is shaped the same way
  regardless of which agent wrote it.
- [[backlog-journal]] — writes to a work item's journal, which is the only
  memory a session leaves behind.

## Version

`0.7.1` — **a work item is written as `WORK-00002-lint-the-corpus`.** Key and
slug joined, the way a decision's filename joins its number and slug, and all
three forms resolve — joined, key alone, slug alone.

That replaces the key column added in `0.7.0`, which left an empty cell on every
outcome and task. One identifier column instead: a work item reads as the joined
form and everything else as its slug, so the column is never blank.

**It is a reference and not a path.** The directory is still the slug alone, and
whether it should carry the key is recorded as open — it would match the decision
records and sort by number, and it would rename every work item, break the
`work_item` link on every outcome and task, and change what each record is.

Patch: how a record is displayed and addressed. Nothing on disk moved, and the
`key` field is unchanged.

`0.7.0` — **a work item carries a key: `WORK-00002`.**

**The key is for finding, the path is for linking**, and both are true at once —
the same split the decision records already run on, and the reason a superseded
decision stays findable after it moves. Work items had only the path, so every
reference was a slug derived from a title, and changing a title meant breaking
inbound links or leaving a slug that no longer matched.

**`WORK` is the only prefix, written into the record rather than derived**, so a
repository that later wants its own changes what gets written from then on and
nothing already on disk is renamed. A derived key would rewrite the corpus the
moment the setting changed.

Allocated at creation from one project-wide sequence, after the existence check —
so asking twice for the same title returns the first record and burns no number,
which is the trap the decision numbering had to be rescued from.

Minor: a recommended field added. A record written before keys has none, and
that is not an error.

`0.6.2` — **an inquiry gains understanding, not only work.** De-risking
something, or finding out what is there, is the point of a spike whether or not
any work falls out of it — so *its only product is more work items* was too
narrow, and the type now says understanding first and the work items that follow.

**And the completion inversion is qualified in both directions.** It read *a
defect that produces no fix is not delivered; an audit that finds no problems
is*. Now: a **confirmed** defect, and an audit that finds nothing **can be**
delivered. An unconfirmed defect producing no fix was never a failure, and an
audit can still fall short for reasons of its own.

Both corrections came from the maintainer and reached `workflow-status.md` first;
this carries them into the type, which is the normative source and had drifted to
being the less careful of the two.

Patch: descriptions sharpened. No value, rule or alias changed.

`0.6.1` — **`request` was described wrongly, and the correction adds a column.**
It said a request produces *an answer you already have the standing to give*,
which came from an earlier framing where a request was a yes-or-no decision. It
is not. A request is **a change somebody outside asked for**, and it differs on
two counts: you are accountable to them, and it has to be evaluated for whether
it is legitimate, aligned and worthy.

That surfaced what the kinds are *for* when scanning: each now records **what it
needs vetted**. A request carries the heaviest — is it legitimate, aligned,
worthy — and a change the lightest and a different question, *do we still need
it*, because your own team wrote it and it is already further along. Both are
vetted; reading `change` as *no vetting* is the misreading worth heading off.

Patch: a description corrected and a column added. No value, rule or alias
changed.

`0.6.0` — **`inquiry` joins the kinds: work that exists to create more work.**
Spikes, experiments, surveys, assessments, examinations, reviews, probes, audits
and investigations are all inquiries, and what comes out is work items, or a
report that generates work items. An inquiry changes nothing itself, so finding
nothing still counts as done.

**The test the kinds are sorted by changed with it.** It was *what has to happen
before the record can be judged*, which derived the first four and cannot place
an inquiry — that is judgeable the moment it is written. It is now **what each
kind produces**, which separates all five.

`review`, `audit`, `investigation` and `spike` are accepted and stored as
`inquiry`. Unlike `bug` against `defect` these are instances rather than
synonyms, so the alias loses a shade of meaning — accepted, because the
alternative fragments the filter that earns the field.

Minor: a value added and a test restated. No record has to change, and a kind
written under `0.5.0` still means what it meant.

`0.5.0` — **`new decision` conforms to the `decision-records` contract.** It
allocates the next `ADR-NNNN` from one project-wide sequence, writes the
`ADR-NNNN:` heading, and scaffolds Summary, Problem, Decision and Why in place
of the pre-bundle `Context / What was chosen` shape. `decided` and
`reopen_trigger` arrive present and empty.

The bundle is where that thinking ended up and the command was its first
iteration, so the command conforms rather than the reverse.

**Breaking for anything that guessed a decision's filename**, which is now
numbered rather than derived from the title alone. Asking twice still finds the
first record: the existence check matches on the slug rather than the path,
which the number would otherwise have defeated.

`0.4.3` — **the recorded inversion names its fourth value `needs_triage`.** The
other three say what has to happen before a record can be judged, and *triage
it* is that shape where *undetermined* was a state. It does not change B's
safety, only makes the value an importer has to write an obvious one.

Patch: wording inside an alternative that is not taken.

`0.4.2` — **the inversion is written down.** The same four states can be arranged
with `change` explicit and blank meaning *nobody looked*, which is what ships, or
with `undetermined` explicit and blank meaning *change*. The second is
ergonomically better and the first is safer, and the re-open condition is a
count rather than an argument: if most records sit blank anyway, the shipped
arrangement is carrying no information.

Patch: an alternative recorded, nothing changed.

`0.4.1` — **a blank kind is now avoided rather than merely allowed.** Creating a
work item without one prints the four values and creates the record anyway.

Capture has to stay free: denying a blank would make classification a toll on
intake, and a toll on intake is how things stop being written down. But blank
left frictionless is the path of least resistance, and then every idea arrives
unclassified and the field means nothing. So the tool names the better path and
lets you past, which is what `spec.md` §5.0 asks of everything that is not the
one refusal.

Patch: no field, value or rule changed. A record written before this is
unaffected, and nothing new is required of anyone.

`0.4.0` — **`change` joins the kinds, `defect` replaces `bug` as canonical, and
`bug` and `ask` become aliases.**

`change` is work that is none of the other three — nothing broke, nobody asked,
and it is formed enough to judge. It is defined by exclusion, which is why the
word reads weakly and why the search for a better one failed: a negatively
defined category has no positive noun. All work changes something, so it is at
least true. **Re-open when a better word turns up.**

**Absence changes meaning, and this is the breaking part.** A blank `kind` used
to mean ordinary work; it now means nobody has classified it, and ordinary work
says `change`. A record written under `0.3.x` with no kind now reads as
unclassified rather than as ordinary — no file is invalid, but a reader draws a
different conclusion from the same bytes.

`defect` is canonical because `spec.md` §2.1 puts the precise word where a
machine reads and the familiar one where a person types. `bug` and `ask` are
accepted on input and stored canonically, which makes importing from a tracker
that emits *bug* a relabel rather than a mapping exercise.

Minor rather than major under the pre-1.0 allowance, and said out loud because
the meaning of existing records moved.

`0.3.1` — **why there is no fourth kind, written down.** The three cover what has
to happen before a record can be judged, and ordinary work has already been
judged — a completeness check rather than a taste for short lists. Also records
that `story` is a template rather than a kind, that *is anybody owed an answer*
is the line rather than internal against external, and that a missing kind cannot
be told apart from ordinary work, which is accepted.

Patch: no field, value or rule changed. A reader who understood `0.3.0`
behaves identically.

`0.3.0` — **`work-item` gains `kind`, and its workflow vocabulary catches up
with what the tool ships.**

`kind` is `bug`, `request` or `idea`, and **absent means ordinary work**, which
is most of it. The test for declaring one is what has to happen before the record
can be judged — fix it, answer them, or develop it. `idea` sits upstream of the
other two: it is the only kind that changes, because developing an idea turns it
into a bug, a request, or ordinary work.

The `workflow_status` values were still the retired ladder — `idea, preparing,
ready` — two days after the corpus moved off them. That was a stale copy rather
than a decision, and it is corrected to the seven rungs in
`docs/workflow-status.md`.

Minor: new content, and nothing an existing record has to change. A record with
no `kind` was correct before and is correct now.

`0.2.0` — **`obligation` is now `field_presence`, and `mandatory` is
`required`.** The knowledge format renamed both in `luma-types` 0.10.0; these
types were written before that and kept the old spelling, so a consumer reading
`field_presence` found nothing here.

Same three presence levels, same meaning, and no field's presence was
strengthened or weakened — only what declares it. Minor rather than patch,
matching how `luma-types` shipped the identical rename: breaking for anything
parsing the old key, and the pre-1.0 allowance is what lets it travel as minor.
Stated here so it does not read as a mistake later.

**Nothing reads either spelling yet.** The tool does not consume these type
definitions, so the rename cost nothing to make and would have cost more the
longer it waited.

*Checked and unchanged:* `stage` stays `draft` — the audience has not moved, and
this is still developed by its maintainers for their own use. `survival` stays
undeclared, which reads as `intended`.

`0.1.0` — extracted rather than designed. The types and procedures existed and
were scattered; this gives them one address. Nothing about them changed in the
move, and the version says only that it has been done once.
