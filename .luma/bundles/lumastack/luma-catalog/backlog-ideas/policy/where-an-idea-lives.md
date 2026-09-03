---
type: policy
title: Where an idea lives
description: Choosing the scope — project, department or organization — and the default that applies when it is unclear.
matches:
  - topic: capturing an idea worth keeping
---

# Where an idea lives

```
.luma/backlog/ideas/<slug>.md        preferred
docs/ideas/<slug>.md                 where luma is not installed
```

**Preferred, not required.** `.luma/backlog/ideas/` is the right home because
the backlog tier is exactly what the practice needs — *what we intend and have
not done* — but **not every repository has luma installed, and an idea still
has to go somewhere.**

What actually matters is **one file per idea, in one consistent place**, with
frontmatter a reader and a tool can both use. None of that depends on the path.

**Never create `.luma/` in a repository that has not adopted it.** Bringing a
directory structure into somebody's project because you wanted to file an idea
is a change they did not ask for. Ask — and if the answer is no, use whatever
the repository already keeps such material in.

**Where luma is absent**, follow what is there rather than inventing: `docs/`
if that is where prose lives, or a sibling of whatever `IDEAS.md` you are
replacing. Consistency within a repository beats consistency across them.

## Choosing the scope

Ask **who would act on this**, not what it is about.

| `scope` | when |
| --- | --- |
| `project` | this repository would do it, and the idea dies if the repository does |
| `department` | several projects under one team would, and no other team cares |
| `organization` | it is about how the organization works, not how anything is built |

**When unclear, ask for guidance or lean towards `project`.** A project-scoped
idea is cheap to promote later; an organization-scoped one nobody owns is the
kind that goes stale in silence. Narrow and promotable beats broad and
orphaned.


## Choosing between repositories

`scope` says whose idea it is. **Which repository it lands in is a separate
question**, and the one that goes wrong in an organization with several.

**Route to where it gets built and committed.** An idea belongs next to the
backlog that will hold it. This is the primary test and it settles most cases
outright — concrete where the alternatives are hypothetical.

**Governing content is a tool idea; writing content is a content idea.** The
sharpest form of the rule, and the one that resolves the case people get wrong:

| | |
| --- | --- |
| *content should behave like this* · *this should be checked* · *this should be enforced* | **the tool that governs it** |
| *we should write one of these* · *this entry is missing* · *this needs updating* | **the repository that holds it** |

A blog is the clearest analogy. **The text formatter goes in the codebase, not in
the repository holding the text it formats.** *Write a post about X* goes with the
posts.

**Two ways this goes wrong, both seen:**

**Topic association is not a build site.** An idea *about* bundles feels like it
belongs with the bundles. It belongs wherever somebody would do the work. Asking
*what is this about* returns the wrong answer confidently, where *where would
somebody open an editor* returns the right one.

**Structural correctness of an artifact is not routing of an idea.** *This field
belongs on that type, and that type is defined in the format* can be entirely true
and still be the wrong home — because a field declaration is a paragraph, and the
behaviour that makes it mean anything is the idea. **The format supplies the
shape; the consumer supplies the behaviour**, and the consumer is where it goes.

**Check the ideas already filed before choosing.** The same shape has usually been
routed before, and a precedent on disk beats an argument in the moment. It is also
mechanical: grep the existing idea files for the pattern rather than reasoning from
first principles each time.

**Then to where the mechanism runs.** A capability that executes inside projects
belongs with the tool that runs there, whatever repository the idea was written
in. A repository that is structurally never present when the work happens cannot
own the answer, only the argument.

**Ownership beats eventual location.** An idea may end up shipping as a bundle,
a package, or somebody else's feature. It still lives where the decision to ship
it gets made.

**"It might become a bundle" is not a routing signal.** Nearly every feature idea
might, because that is how features get distributed. A test that fires on almost
everything does not discriminate — it just adds a step where a correct placement
gets talked out of. A catalog is a publication target, not a place to think. *The
exception is an idea about the catalog machinery itself.*

**A contract is dumb about how things get used.** Format and specification
repositories hold shapes, not behaviour. An idea about what a consumer should
*do* with a field belongs with the consumer, even when the field is the format's.
A specification that encodes one consumer's behaviour is false the moment another
ignores it.

**Do not split an idea when half would be orphaned.** Entries often contain a
capability and an open question that belong in different places. Split only when
each half is independently buildable and independently valuable; otherwise keep
it whole and record where the other half would go.

**When two repositories both have a claim, ask which one loses the idea.** The
one where nobody would look for it is the wrong answer, however good the argument
for it.

## Sensitivity decides before scope does

**Ask what disclosure level the idea needs, before asking whose it is.** A
misplaced idea is an inconvenience; a published one is not recoverable.

**When unsure, assume the narrower level.** Being wrong toward restriction is
something somebody notices and fixes in a minute. Being wrong toward permission
is forked, cloned and cached before anybody looks. Route only to a repository
whose declared `disclosure_level` accepts what the idea needs.

**A sensitive idea heading for a public repository stops.** Say plainly what
would be exposed and that publishing does not reverse, and get an explicit second
confirmation. **The reverse never needs one** — an ordinary idea filed somewhere
internal is a mild inconvenience, and a guard that fires in both directions is
one people learn to click through.

**Provenance is not sensitivity.** An insight learned from running an internal
headquarters can produce an entirely public feature. Where knowledge came from
does not decide where the idea lives; only what the idea itself contains does.

**A capability and its output have different sensitivities.** *Help decide what
to build next* is harmless; a prioritized list of what an organization is about
to build is competitive information. Route on the idea, not on what running it
would produce — and say so in the file, because the distinction gets
re-litigated.

**A public repository never names an internal one.** Not in prose, not in a
cross-reference, not in a migration marker. Say *the organization's internal
headquarters* and give the path within it: obvious to anyone who owns it,
invisible to everyone else.

## When the tool is about the thing

A product about how organizations work produces ideas that read as both product
features and organizational knowledge, because they genuinely are both. The
ordinary tests blur, and the people closest to it get the most confused.

**Two questions cut it:**

- **Does this get built and committed here?** If yes, it is a feature, however
  organizational the subject.
- **Would it still matter if the organization threw out this tooling entirely and
  used a competitor's?** If yes, it is organizational knowledge that happens to
  be about tools.

The first is usually decisive. The second is for when it is not.

## The list of scopes will grow

`project`, `department` and `organization` are a **first guess**, made before
any real backlog met them. Some ideas will fit none of them, and some will fit
badly — belonging to a customer, a product line, a community, a piece of shared
infrastructure that no single project owns.

**When that happens, record it rather than forcing the nearest fit.** An idea
with no good home is evidence about this list, and it is the only evidence there
will be.

**The default is what puts that evidence at risk.** An awkward idea quietly
filed as `project` is indistinguishable from a well-placed one, so leaning on
the default without noting the discomfort is how a missing scope stays missing.
Lean on it, and say when you did.

## The file follows the scope

An idea scoped to the organization belongs in the organization's own repository,
not in whichever repository somebody happened to be working in. An idea captured
in the wrong place is one the people who would act on it never see.

**Do not let the location be decided by where you were standing.** That is the
most common way an idea ends up invisible, and it happens silently. Ideas belong
in the repository that would most naturally act upon them.

If an organization's headquarters is present, ask it — it has the breadth to
see where an idea belongs, which a single repository does not. *Not implemented
yet.*

**Asking is a person or an agent consulting it, never a tool depending on it.**
Foreman must work in a repository belonging to no organization at all, so
nothing in the toolchain may require a headquarters to be reachable. Absent one,
this step is skipped and the guidance above stands on its own.

## One idea or theme per file

Both are fine, and the test is whether they **rise and fall together**. Three
variations on the same underlying change are one file. Three unrelated
improvements that happen to touch the same subsystem are three separate files.

A themed file that starts collecting unrelated entries has become a second
`IDEAS.md` junk drawer, which is what this system exists to replace.

## Naming

`<slug>.md` — kebab-case, from the idea rather than from the area.

`cache-the-dependency-layer.md`, not `build-improvements.md`. The second is a
bucket, and buckets attract everything.
