---
type: work-item
key: WORK-0091
title: A project-wizard bundle that sets a new project up by asking
workflow_status: captured
kind: idea
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T05:15:06Z'}
description: a new bundle called project-wizard that helps people set up a new project. it asks whether this is a project or an organization, then walks the common things a project will want installed by asking questions — will you be using git (installs git-workflow), will two or more agents or people be working in parallel (installs git-worktrees), what language will the project use (installs coding best practices), and so on.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T05:15:06Z'}
---

# A project-wizard bundle that sets a new project up by asking

## The idea

**A new bundle called `project-wizard`** that helps people set up a new project.

**It asks whether this is a project or an organization**, and then goes through
the common things a project will want installed, by asking questions:

- **Will you be using git?** → installs `git-workflow`
- **Will two or more agents or people be working in parallel?** → installs
  `git-worktrees`
- **What language will the project use?** → installs coding best practices
- and so on.

---

*Everything above is the maintainer\'s, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## In the capturing agent\'s own words

**A wizard that turns questions about the work into a set of adopted bundles**,
so somebody starting a project does not have to read fifteen descriptions and
guess which apply. The questions are about what the project *is and does*, never
about the bundles — nobody should have to know a bundle exists to get it.

## Ideas and concerns from the capturing agent

### The first question already has data behind it

**Every bundle declares `consumers`**, and across the fifteen adopted here it is
already split: **ten are `[project, organization]` and four are `[project]`.**

So *is this a project or an organization* does not need new metadata — it needs
the wizard to filter on a field that already exists. **That is the cheapest part
of this and it works today.**

### The other questions have nothing behind them

**Nothing says which question selects a bundle.** `git-workflow` does not
declare *ask whether they use git*; `git-worktrees` does not declare *ask
whether work runs in parallel*.

**Which is the design fork, and it is the whole of this idea:**

- **The wizard holds the mapping** — a list of questions and the bundles each
  implies. Simple, works immediately, and **goes stale every time the catalog
  grows.** A new bundle is invisible until somebody edits the wizard.
- **Each bundle declares what selects it** — a field saying *adopt me when this
  is true of the project*. The wizard becomes a reader rather than a list, a new
  bundle is offered the day it is published, and **the knowledge sits with the
  thing it describes**, which is the pattern this estate uses everywhere else.

**The second is right and costs a format change**, which is
`docs/format-requests.md` territory rather than something the wizard decides
alone.

### It will find holes in the catalog

**There is no coding-best-practices bundle.** The language question has nothing
to install today, for any answer. The same will be true of most questions worth
asking.

**That is a feature of building the wizard rather than an obstacle to it** — the
questions somebody wants to answer are a better map of what the catalog is
missing than a list of bundles somebody thought of.

### What it must not become

**A questionnaire that installs everything.** `INDEX.md` opens with *"open a
bundle\'s index when its line matches the work, and not before"* — adoption has a
cost, and a wizard optimising for coverage would hand every project fifteen
bundles it has to read past.

**Nor a one-time gate.** Projects change: the second agent arrives months after
`init`. Re-running it later and being offered only what is newly relevant is
probably the more valuable half, and the harder one.

### Where it belongs

**`lumastack/luma-catalog`, not here.** This is the third idea in a row captured
into luma-backlog that belongs to another project ---
[[work-items/WORK-0090-session-save-and-session-end-and-moving-a-record-to-the-project-it-belongs-in]]
is the record for that problem, and this is more evidence for it.

## References

- `.luma/bundles/INDEX.md` --- what a project is offered today, and the *open it
  when the line matches* rule a wizard has to respect.
- `.luma/bundles/*/BUNDLE.md` --- the `consumers` field, which already answers
  the first question.
- `docs/format-requests.md` --- where a *what selects this bundle* field would
  be asked for.
- [[work-items/WORK-0090-session-save-and-session-end-and-moving-a-record-to-the-project-it-belongs-in]]
