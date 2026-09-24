---
type: exploration
type_version: "0.0.1"
title: Found in the literature, 1
work_item: '[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T20:24:53Z'}
---

# Found in the literature, 1

> **Literature search in isolation.** One agent, working in a directory containing the problem brief
> and nothing else, reported what already exists, with designing a scheme forbidden. **It could not see any other
> agent's answer, and no other agent could see this one** --- the isolation was
> the filesystem's rather than an instruction's.
>
> **Filed before any of the five was read**, so that reading order decided
> nothing. Unedited below this line; the workspace it was produced in has been
> deleted.

---

# Found in the literature: how ordered, concurrently-edited sequences are kept elsewhere

**Stage 3 output. Produced blind: no other file in this directory or repository was
read, and no stage-2 output was seen. Everything below comes from PROBLEM.md plus
public sources, cited inline.**

## How the search was run, and the shape of what it found

The problem in the brief — a total order over items, one write per move, no
coordinator, merge by `git` alone, an immovable interior item, unbounded lifetime —
is not one literature but four, which mostly do not cite each other:

1. **Application/database practice**: "user-defined ordering" of rows — sparse
   integers, floats, fractional/lexicographic keys, rationals. (Trello, Jira,
   GitLab, Figma, countless SQL threads.)
2. **Distributed collaborative editing**: sequence CRDTs (Treedoc, Logoot, LSEQ,
   WOOT, RGA, causal trees, YATA/Yjs, Fugue) and operational transformation.
3. **Theory**: the *order-maintenance* and *online list-labeling* problems
   (Dietz–Sleator 1987 through Bender et al. FOCS 2022), which carry the lower
   bounds that explain why the practical literature keeps rediscovering the same
   two unpleasant options.
4. **Version-control-native ordering**: patch theory (Darcs, Pijul), and the
   plain-file kanban tools that store order as line order in a manifest.

One result frames everything, so it goes first.

---

## The theory result that governs every positional scheme

**Online list labeling / order maintenance.** The formal problem: assign each of
*n* items a label from a bounded universe so that label order equals list order,
minimizing how many items must be *relabeled* per insertion. It has been studied
since 1981; [Dietz & Sleator (1987)](https://www.cs.tau.ac.il/~haimk/adv-alg-2014/order-bender2014.pdf)
gave O(1)-amortized structures, [Bender, Cole & Zito (2002)](https://erikdemaine.org/papers/DietzSleator_ESA2002/paper.pdf)
simplified them, and [Bender, Conway, Farach-Colton, Komlós, Kuszmaul & Wein
(FOCS 2022)](https://arxiv.org/abs/2203.02763) broke the long-standing
O(log² n) relabeling bound with a randomized O(log^{3/2} n) structure — and
proved matching **lower bounds**: with labels of bounded length, relabeling
cannot be avoided, only amortized and spread. (Survey of bounds:
[Saks, "Online Labeling: Algorithms, Lower Bounds and Open Questions"](https://www.semanticscholar.org/paper/d365cfc45928443a7bdcfa4c2a2ef2ea7effe960);
tight Ω bound for the monotone case: [Dietz, Seiferas & Zhang](https://dx.doi.org/10.1137/S0895480100315808).)

The consequence, and it is a theorem, not an engineering opinion:

> **For any scheme where one bounded-size stored field sorts the column, there is
> no third option. Either the field's length grows without bound under repeated
> insertion at one spot, or items other than the moved one must periodically be
> relabeled.** The only freedom is where on that dial you sit and how local the
> relabeling is.

The brief's requirement 2 ("a rewrite of more than a few cards happens a handful
of times in the board's entire history") and requirement 7 (unbounded traffic
around an immovable interior card) sit on *opposite ends of that dial*. Every
approach below is some position on it — except the ones that abandon "one sorted
field" entirely (§D), which is exactly what the escape hatch looks like in the
literature too.

The list-labeling structures themselves are worth naming as a candidate, not
just a bound: they relabel **O(log n) amortized items per insert, always
adjacent to the insertion point** — i.e. "twenty cards, next to the problem,
reviewable" rather than "the whole board." That is precisely the brief's goal-3
shape of repair. What they give up: relabels are *frequent* (small, constant
background churn — many multi-card diffs per year, each tiny), so they fail the
letter of **requirements 1, 2 and 6** while arguably honoring their spirit
(no diff is ever large). They also assume a single mutator per structure;
concurrent relabeling without coordination is an open engineering problem
(a concurrent version exists but assumes shared memory, not git:
[Guo & Sekerinski 2022](https://arxiv.org/abs/2208.07800)).

---

## A. Positional keys in one field (the "fractional indexing" family)

The dominant industry answer. All members subdivide an interval or extend a
string; all are subject to the theorem above and to the brief's one measured
result (linear length growth at a fixed spot). They differ only in where they
put the pain.

### A1. Sparse integers with gaps

Leave room between ranks; insertion takes the midpoint; renumber when a gap
closes. This is BASIC line numbering — increments of 10, `RENUM` when you run out
([C64 Wiki: RENUMBER](https://www.c64-wiki.com/wiki/RENUMBER_(BASIC_3.5))) — and
it is what **GitLab** ships today for issue boards: a signed-integer
`relative_position` with an ideal gap of 500, new items at max+gap, and a
**background rebalancing job** when gaps close
([RelativePositioning concern](https://gitlab.com/gitlab-org/gitlab/-/issues/230953),
[gap-size discussion](https://gitlab.com/gitlab-org/gitlab/-/work_items/30352),
[production rebalancing incident](https://gitlab.com/gitlab-com/gl-infra/production/-/work_items/4584)).
The classic SQL treatment is [begriffs, "User-defined order in SQL"](https://begriffs.com/posts/2018-03-20-user-defined-order.html).

- **Gives up:** interior capacity. A gap of *g* absorbs ~log₂ *g* same-spot
  insertions (a gap of 500 absorbs ~8) before *neighbors must be renumbered*.
  Renumbering is routine, not exceptional — GitLab runs it as a job and has had
  production incidents from it.
- **Fails:** **5** and **7** immediately (workloads 1, 3, 4, 5 each close a gap
  in single-digit operations); then **1, 2, 6**, because the repair is the
  ordinary maintenance path. Meets 3 and 4 only until the integer range ends.
  Workload 9: two midpoints of the same gap are **equal — both valid,
  indistinguishable**, the silent failure the brief names.

### A2. Floating-point positions

Same scheme, doubles instead of integers. **Trello**: `pos` is a 64-bit float;
insertion averages the neighbors; when two positions come within ~0.0001 the
nearby cards are renumbered, worst case a whole list, "every ~20 moves or so in
degenerate cases" ([Trello developers on HN](https://news.ycombinator.com/item?id=10957165)).

- **Gives up:** everything A1 does, plus a hard ~50-halvings precision floor,
  after which averaging **returns a value equal to a neighbor's — order lost,
  silently** (the identical bug the brief reports having shipped).
- **Fails:** **5, 7** (≈50 same-spot ops per gap, ever), **1, 2, 6** (rebalance
  is routine). Workload 9 silently collides. The literature's own users report
  this scheme as the one you migrate *away* from.

### A3. Arbitrary-precision / lexicographic string keys ("fractional indexing" proper)

Keys are digit strings ordered lexicographically; between two keys there is
always another; appending at an end lengthens the key only logarithmically, but
insertion inside a fixed gap appends roughly **one digit per operation**.
Canonical exposition: [Figma, "Realtime editing of ordered sequences"](https://www.figma.com/blog/realtime-editing-of-ordered-sequences/)
(child ordering in the multiplayer tree); reference implementations
[rocicorp/fractional-indexing](https://github.com/rocicorp/fractional-indexing)
(from the Figma post's scheme), practical failure-mode walkthrough by
[Steve Ruiz (tldraw)](https://www.steveruiz.me/posts/reordering-fractional-indices),
survey of use in sync engines: [Liveblocks](https://liveblocks.io/blog/how-crdts-and-sync-engines-keep-realtime-lists-ordered-with-fractional-indexing).

- **Gives up:** bounded key size. No rewrite is ever *forced* — instead the key
  grows: the brief's own measurement (2,925 same-spot placements → 2,929
  characters) is exactly this family's known behavior.
- **Requirements:** passes **1, 2, 6** vacuously (nothing ever *must* be
  rewritten), passes **3, 4** genuinely (end-appends grow the integer part;
  a million appends is a short key). Fails **5** and **7** *in effect*: at 10⁶
  interior ops the stored value is ~10⁶ characters — unreadable (goal 7 gone),
  undiffable, and quadratic to compute with. The scheme dies by obesity instead
  of exhaustion.
- **Workload 9:** two actors bisecting the same gap generate **byte-identical
  keys** — valid, indistinguishable, silent. Known and documented; the standard
  patches are jitter (below) or unique suffixes (§B3).

### A4. LexoRank (Jira): string keys + scheduled whole-board rewrite

Atlassian's production variant: base-36 lexicographic rank plus a **bucket
prefix** (`0|`, `1|`, `2|`); when ranks reach a length threshold (128 chars,
hard alarm at 160) a **rebalance job rewrites every row into the next bucket**
([Managing LexoRank](https://confluence.atlassian.com/adminjiraserver/managing-lexorank-938847803.html),
[troubleshooting KB](https://support.atlassian.com/jira/kb/troubleshooting-lexorank-system-issues/),
[performance during rebalance](https://jira.atlassian.com/browse/JSWSERVER-13163)).

- **Gives up:** requirement 2 *by design* — the whole-collection rewrite is not
  an emergency, it is the amortization mechanism, run by a scheduled service.
  Instructive precisely because it shows what A3 turns into once someone insists
  on bounded key length: the theorem collects its fee as a global rewrite.
- **Fails:** **2** and **6** structurally (rebalance cadence is driven by usage,
  not by decade); in a git setting the rebalance commit is the unreviewable
  many-thousand-file diff the brief forbids, and any branch concurrent with a
  rebalance conflicts on *every* card.

### A5. Jittered fractional indexing (the concurrency patch)

Add random bits below the midpoint so concurrent same-gap inserts *probably*
differ: [rocicorp README](https://github.com/rocicorp/fractional-indexing),
[jittered-fractional-indexing](https://github.com/nathanhleung/jittered-fractional-indexing)
(30 bits of jitter → ~4.5% birthday collision at 10k concurrent inserts).

- **Gives up:** determinism, and only *shrinks* the silent-collision window —
  collisions remain possible, remain silent, and cost extra key length per
  insert. A probabilistic fix for the one failure the brief calls "the dangerous
  one" is the weakest fix the literature offers; the principled fix is
  identity-bearing keys (§B3).

### A6. Exact rationals (Stern–Brocot mediants)

Store position as a true fraction; insert the mediant; sort by p/q. Postgres
`pg_rational`, advocated in [begriffs](https://begriffs.com/posts/2018-03-20-user-defined-order.html)
and the [PostgreSQL wiki](https://wiki.postgresql.org/wiki/User-specified_ordering_with_fractions).

- **Gives up:** either boundedness or the fixed width: in 64 bits the
  denominators grow along Fibonacci under alternating insertion (~46 same-spot
  ops to overflow); with bignums it is A3 with worse sort ergonomics (p/q does
  not sort lexicographically as text — goal 5 drops to "needs a program").
- **Fails:** **5, 7** (fixed width) or goal 5 + goal 7 (bignum). Same-spot
  concurrent mediants are identical → workload 9 silent.

**Family verdict.** Everything in §A is the brief's "subdividing a finite
interval" result wearing different clothes, exactly as the brief's measured
result predicts. The industry's universal answer to the growth is *periodic
rebalance* — tolerable behind a database with a maintenance daemon, and almost
exactly what requirement 2 exists to forbid in reviewable git history.

---

## B. Sequence CRDTs: orders designed to merge without coordination

The collaborative-editing literature hit this problem's *concurrency* half
head-on twenty years ago: no coordinator, replicas edit independently, merges
must converge. Field guides: [Ian Duncan's CRDT dictionary](https://www.iankduncan.com/engineering/2025-11-27-crdt-dictionary/),
[archagon, "Data Laced with History"](http://archagon.net/blog/2018/03/24/data-laced-with-history/).
Two sub-families matter here, and they fail the brief differently.

### B1. Dense unique identifiers (Treedoc, Logoot, LSEQ)

Each element carries a variable-length path identifier — Logoot: a list of
(digit, replicaID, counter) triples; Treedoc: a binary-tree path; LSEQ: Logoot
with adaptive allocation. Identifiers are globally unique, totally ordered, and
allocated *between* any two others with no communication.
Sources: [LSEQ paper (Nédelec et al., DocEng 2013)](https://www.researchgate.net/publication/262162421_LSEQ_an_Adaptive_Structure_for_Sequences_in_Distributed_Collaborative_Editing);
Logoot (Weiss, Urso & Molli, ICDCS 2009) discussed therein.

- **What they solve that §A cannot:** **workload 9.** Because the replica ID and
  counter are *inside* the identifier, two actors inserting into the same gap
  produce **different identifiers that every replica sorts identically** — both
  valid, *distinguishable*, deterministic, convergent. No silence, no conflict.
  This is the literature's answer to the brief's "dangerous one," and it is the
  single most transferable idea found: **a position that does not embed identity
  cannot survive concurrent allocation; one that does, can.**
- **What they give up:** identifier growth — these are §A3 keys with identity
  attached, so the interval-subdivision arithmetic still applies. Logoot's ids
  grow linearly with depth of insertion (space linear in total inserts); LSEQ's
  adaptive boundary strategy gets *average* polylog growth but adversarial
  same-spot repetition (workloads 4, 5) still lengthens ids without bound.
  Known **interleaving anomaly**: two concurrently inserted *runs* can merge
  interleaved card-by-card — directly hits workload 8's "does the batch keep
  its relative order on arrival"
  ([Kleppmann et al., "The Art of the Fugue"](https://arxiv.org/pdf/2305.00583)).
- **Fails:** **5** and **7** in the same by-obesity sense as A3 (plus goal 7
  readability much sooner, since every path element carries a replica ID);
  passes 1, 2, 3, 4, 6.

### B2. Insert-after graphs (WOOT, RGA, Causal Trees, YATA/Yjs, Fugue)

The other half of the CRDT family stores **no position at all**: each element
carries a fixed-size unique ID and the ID of the element it was inserted
*after* (plus a Lamport timestamp/tie-break); the order is *derived* by walking
the resulting tree with a deterministic sibling rule. RGA uses timestamps;
Yjs/YATA and Fugue refine the sibling rule so concurrent runs never interleave
([Weidner, "Fugue: A Basic List CRDT"](https://mattweidner.com/2022/10/21/basic-list-crdt.html);
[Fugue paper](https://arxiv.org/pdf/2305.00583);
[Kleppmann, "Moving Elements in List CRDTs"](https://martin.kleppmann.com/papers/list-move-papoc20.pdf)
— the move-vs-insert distinction, directly relevant to cards that *move* rather
than die and reappear).

- **Why it matters here:** this family **passes all seven requirements**. There
  is no positional budget to exhaust: metadata per card is constant-size
  (an ID and a parent reference), a million insertions before, after, or beside
  an immovable card cost one small write each, forever, and nothing is ever
  renumbered. It is the only family the search found for which requirement 7 is
  simply a non-event.
- **What it gives up — and it is exactly the brief's goal 5:** the order is
  *computed*, not *sorted*. `sort` on a field cannot produce it; a shell user
  must chase parent pointers (the brief's named "last resort"). Two further
  costs: **tombstones/anchors** — an element that leaves must often persist as
  an anchor for things inserted relative to it (WOOT/RGA keep tombstones
  forever; workload 6's "does leaving free anything" answer is *worse than
  nothing frees*), and **concurrent same-spot placement merges silently but
  arbitrarily**: convergent and deterministic on every replica, yet the chosen
  relative order is a timestamp/ID tie-break no human sanctioned — not the
  brief's "silently *wrong*" divergence, but not its "loud conflict" either.
  In per-card-file form there is one bright spot the editing literature never
  gets: two branches *re-placing the same card* both rewrite that card's one
  file, and **git itself conflicts loudly** on it.
- **Fails:** none of the seven requirements; sacrifices **goal 5** (and part of
  goal 4's "loud"), which the brief pre-identifies as "the likeliest trade."

### B3. The bridge: identity-bearing *sortable* keys (position strings / Fugue-as-a-string)

Recent work collapses B1's sortability and B2's tree structure into one
artifact: [Weidner's position-strings](https://mattweidner.com/2023/04/13/position-strings.html)
and [list-positions](https://mattweidner.com/2024/04/29/list-positions.html)
(npm: [position-strings](https://www.npmjs.com/package/list-positions)) — keys
that sort lexicographically like fractional indices but embed replica identity
(global uniqueness → workload 9 solved deterministically), use Fugue's sibling
rule (no interleaving → workload 8 solved), and grow **logarithmically for
sequential runs** at either end rather than linearly. A same-direction variant
appears in [sqliteai/fractional-indexing](https://github.com/sqliteai/fractional-indexing)
(CRDT-flavored, base-62, collision-averse).

- **Gives up:** worst-case growth is still governed by the §-theorem — repeated
  adversarial placement into one interior gap (workload 5, and workload 4 if
  each promotion lands *between* the same two survivors) lengthens keys without
  bound, merely at a gentler slope than A3, and every key visibly carries a
  replica ID (goal 7 readability erodes early). No rebalance is ever forced;
  none is provided either — trimming keys would change identities, so a
  "repair" is a semantic migration, not a renumber.
- **Fails:** nothing outright at 10⁶-op scale; at 10⁸ same-gap operations it
  fails **5/7** the A3 way (by growth). It is the literature's closest existing
  artifact to the brief's goal set, and its documentation is candid that
  key length, not correctness, is the failure axis.

### B4. Operational transformation — ruled out by the constraints

The other classic collaborative-editing lineage (Google Docs). Transforms
concurrent operations against each other, which in every production deployment
requires a **central server holding a canonical operation order**
([OT vs CRDT comparison, Sun et al.](https://arxiv.org/pdf/1905.01518);
practitioner summary [here](https://medium.com/@sohail_saifi/building-collaborative-editing-the-battle-between-operational-transform-and-crdts-fdceb63c54ac)).
Fails the **no-coordination hard constraint** outright; also operates on
operation logs, not mergeable state files, so `git`-only merge has nothing to
transform with. Listed because its absence is informative: the half of the
field that assumed a coordinator produced nothing transplantable here.

---

## C. Ordering by identity + time (no allocated positions)

### C1. Timestamp-sortable identifiers (ULID, UUIDv7, Snowflake, Lamport/HLC)

128-bit IDs with the timestamp in the high bits sort lexicographically into
creation order: [ULID spec](https://github.com/extrawurst/ulid),
[ULID vs UUID (Baeldung)](https://www.baeldung.com/cs/ulid-vs-uuid),
[UUIDv7/ULID/Snowflake compared](https://www.authgear.com/post/time-sortable-identifiers-uuidv7-ulid-snowflake/).

- **What it buys:** workload 2 (append-only done column) at **exactly zero
  cost, forever** — arrival order *is* the sort, nothing is allocated, nothing
  can be exhausted, and concurrent appends get distinct IDs (random bits) that
  merge into a sane order. Requirement 4 free of charge.
- **Fails:** **7** (and 3, 5) totally as a sole mechanism — creation time
  cannot express "now between these two older cards"; any deliberate reorder
  needs a second mechanism. The literature uses these as the *tie-break inside*
  other schemes (RGA's timestamps, B3's suffixes), which is their right role
  here too.

### C2. Adjacency: "after card X" stored in the card

The database world's linked list (surveyed and rejected for query cost in
[begriffs](https://begriffs.com/posts/2018-03-20-user-defined-order.html)); the
CRDT world's RGA is exactly this plus a merge rule (§B2). Kept separate to name
its git-specific failure: with *bare* neighbor references and no tie-break
metadata, two branches inserting after the same card merge cleanly in git into
**two cards claiming the same predecessor** — order ambiguous, no conflict
raised (fails goal 4 silently) — and a card leaving the column dangles every
reference to it (repair becomes common, failing **1**). The CRDT dressing in
§B2 is precisely what fixes both.

---

## D. Keeping the order *outside* the cards

### D1. A manifest per column: order = line order in one small file

What nearly every plain-text kanban tool actually ships: the column is a file
(or a directory plus an index), and a card's position is *which line it is on*.
[Obsidian Kanban-style markdown boards](https://savolai.net/notes/edu-tech-blog/llm-text-files-obsidian-kanban-practical-project-management-for-developers/),
[personal-kanban](https://github.com/YJPL/personal-kanban),
[kotban](https://github.com/medavox/kotban) (column = directory),
[Backlog.md](https://news.ycombinator.com/item?id=44483530),
[Imdone](https://imdone.io/markdown-kanban-board),
[TrackDown](https://github.com/mgoellnitz/trackdown).

- **Against the requirements: passes all seven, vacuously.** There are no
  stored positions, so nothing grows, nothing exhausts, nothing is ever
  renumbered; a million cards around an immovable one are a million lines
  around an untouched line. Reordering is moving one line; `cat` *is* the
  order-recovery tool (better than goal 5's "best case" — not even a sort).
  Workload 8 is one reviewable diff; git merges concurrent insertions at
  *different* spots cleanly and **conflicts loudly** on the same spot — the
  exact goal-4 outcome the brief calls a good one.
- **What it gives up, named:** (1) **the card no longer carries its place** —
  the brief's opening premise, explicitly listed as renegotiable ("whether the
  thing being ordered is the right thing to order"); (2) **one contended file
  per column** — every reorder in a column touches the same file, so
  merge-conflict *frequency* rises even though each conflict is two readable
  lines; (3) **a second source of truth** — a card present in the directory but
  absent from the manifest (or vice versa) is the "degenerate state [that] must
  still be readable," requiring a stated append-missing/ignore-ghost rule and
  making small repairs *routine* (pressure on requirement **1**, the only one
  it strains); (4) a known unsoundness of git's textual merge itself: 3-way
  line merge can, in edge cases, reorder or misplace lines *without* raising a
  conflict — the founding criticism behind patch-theory systems
  ([Why Pijul](https://pijul.org/manual/why_pijul.html),
  [jneem's merge analysis](https://jneem.github.io/pijul/)) — so goal 4's
  "never silent" holds usually, not provably.

### D2. Patch theory: ordering as explicit graph edges (Darcs, Pijul)

Pijul models a file as a graph of content blocks whose **edges are the order**;
merges are pushouts, order between merged lines is provably preserved, and an
unresolvable order is a first-class *conflict state*, not a silent guess
([Pijul model](https://pijul.org/model/),
[Why Pijul](https://pijul.org/manual/why_pijul.html),
[Darcs patch theory](https://en.wikibooks.org/wiki/Understanding_Darcs/Patch_theory_and_conflicts)).

- **Relevance:** not adoptable — the brief fixes git as the merge engine — but
  it is the strongest statement in the literature of the brief's goal 4, and it
  identifies precisely where D1 leans on luck: git's merge does not *guarantee*
  order preservation, Pijul's does. If order ever moves into per-column files,
  the mergeable-by-construction version of that idea already exists and works;
  it just lives in the wrong VCS. Fails the hard constraint ("files live in
  git"), not any of the seven requirements.

---

## What did **not** surface — and the brief said absence would be informative

- **No scheme anywhere combines**: a single lexicographically sortable field,
  bounded human-readable width, no coordination, and unbounded same-gap
  insertion with no rewrites. The list-labeling lower bounds (§C-theory) say
  why: that combination is impossible, not undiscovered. Every production
  system inspected (Trello, Jira, GitLab, Figma) chose bounded-ish keys **plus
  routine rebalancing** — affordable behind their databases, and the exact cost
  requirement 2 makes unaffordable in reviewable git history.
- **No git-native project found allocates positional keys at all.** The tools
  that live where this problem lives (plain files, git merges) uniformly store
  order as **line order in a shared file** (§D1) and let git's line machinery
  be the merge algorithm. That an entire tool ecosystem independently landed
  there is the closest thing to field evidence the search produced.
- **The one idea that transfers into *any* allocated-key design regardless of
  family** (from §B): a stored position that does not embed actor identity
  cannot distinguish two concurrent allocations at the same spot — jitter only
  shrinks the window (§A5); identity closes it (§B1/B3). Whatever scheme stage
  2 produced, this is the test to run against its workload-9 answer.

---

## Requirements matrix

R1 repair rare · R2 big rewrites ≈ never · R3 millions ahead of a fixed card ·
R4 millions behind · R5 front looks inexhaustible · R6 rewrite ≤ once/decade ·
R7 unbounded traffic around an immovable interior card.
✔ pass · ✖ fail · ◐ passes formally but degrades (key growth) at the brief's 10⁶–10⁸ scale.

| Approach | R1 | R2 | R3 | R4 | R5 | R6 | R7 | Named sacrifice |
|---|---|---|---|---|---|---|---|---|
| A1 gapped integers (GitLab, BASIC) | ✖ | ✖ | ✖ | ✖ | ✖ | ✖ | ✖ | rebalancing is the maintenance path |
| A2 float midpoints (Trello) | ✖ | ✖ | ✖ | ✖ | ✖ | ✖ | ✖ | + silent precision collision |
| A3 fractional string keys (Figma) | ✔ | ✔ | ✔ | ✔ | ◐ | ✔ | ✖ | unbounded key growth; silent same-gap collision |
| A4 LexoRank buckets (Jira) | ✖ | ✖ | ✔ | ✔ | ◐ | ✖ | ✖ | whole-board rewrite institutionalized |
| A5 + jitter | ✔ | ✔ | ✔ | ✔ | ◐ | ✔ | ✖ | collision made rare, kept silent |
| A6 rationals (Stern–Brocot) | ✔ | ✔ | ◐ | ◐ | ✖ | ✔ | ✖ | Fibonacci overflow or bignum growth; loses text-sortability |
| B1 Logoot / LSEQ | ✔ | ✔ | ✔ | ✔ | ◐ | ✔ | ◐ | id growth; run interleaving (workload 8) |
| B2 RGA / Fugue / YATA | ✔ | ✔ | ✔ | ✔ | ✔ | ✔ | ✔ | **goal 5**: order needs a program; tombstones; arbitrary (deterministic) concurrent tie-break |
| B3 position strings (Weidner) | ✔ | ✔ | ✔ | ✔ | ◐ | ✔ | ◐ | slower but unbounded growth; ids in every key |
| B4 operational transformation | — | — | — | — | — | — | — | requires a coordinator: excluded by hard constraint |
| C1 timestamp ids (ULID/UUIDv7) | ✔ | ✔ | ✖ | ✔ | ✖ | ✔ | ✖ | append-only; cannot express interior placement |
| C2 bare "after X" adjacency | ✖ | ✔ | ✔ | ✔ | ✔ | ✔ | ✔ | silent sibling ambiguity; dangling anchors (fixed by B2) |
| C-theory list labeling | ✖ | ✖ | ✔ | ✔ | ✔ | ✖ | ✔ | frequent, small, *local* relabels (the "cheap usually" branch, with optimal bounds) |
| D1 column manifest (line order) | ◐ | ✔ | ✔ | ✔ | ✔ | ✔ | ✔ | card no longer self-contained; contended file; drift repair; git merge not provably order-safe |
| D2 patch theory (Pijul/Darcs) | ✔ | ✔ | ✔ | ✔ | ✔ | ✔ | ✔ | wrong VCS: excluded by "files live in git" |

Three rows satisfy requirement 7 for the life of the board without a rewrite:
**B2** (buy it by sacrificing goal 5 — the sortable field), **D1** (buy it by
sacrificing the card's self-containment), and **C-theory** (buy it by
sacrificing "rewrites ≈ never" in exchange for "rewrites always tiny and
adjacent, with proofs that nothing does better"). Every scheme that keeps *both*
a self-contained sortable field *and* bounded width sits in the top half of the
table and fails requirement 7 — which is the theorem, observed in the wild four
separate times (Trello, Jira, GitLab, Figma), each having independently
purchased its way out with the rewrite the brief forbids.

---

## Sources

Theory: [Dietz–Sleator lecture notes](https://www.cs.tau.ac.il/~haimk/adv-alg-2014/order-bender2014.pdf) ·
[Bender–Cole–Zito 2002](https://erikdemaine.org/papers/DietzSleator_ESA2002/paper.pdf) ·
[Bender et al., FOCS 2022](https://arxiv.org/abs/2203.02763) ·
[Saks survey](https://www.semanticscholar.org/paper/d365cfc45928443a7bdcfa4c2a2ef2ea7effe960) ·
[Dietz–Seiferas–Zhang lower bound](https://dx.doi.org/10.1137/S0895480100315808) ·
[Order-maintenance problem (Wikipedia)](https://en.wikipedia.org/wiki/Order-maintenance_problem) ·
[concurrent order maintenance](https://arxiv.org/abs/2208.07800)

Industry positional keys: [Figma blog](https://www.figma.com/blog/realtime-editing-of-ordered-sequences/) ·
[rocicorp/fractional-indexing](https://github.com/rocicorp/fractional-indexing) ·
[Steve Ruiz](https://www.steveruiz.me/posts/reordering-fractional-indices) ·
[Liveblocks](https://liveblocks.io/blog/how-crdts-and-sync-engines-keep-realtime-lists-ordered-with-fractional-indexing) ·
[jittered-fractional-indexing](https://github.com/nathanhleung/jittered-fractional-indexing) ·
[sqliteai/fractional-indexing](https://github.com/sqliteai/fractional-indexing) ·
[Trello pos on HN](https://news.ycombinator.com/item?id=10957165) ·
[Jira Managing LexoRank](https://confluence.atlassian.com/adminjiraserver/managing-lexorank-938847803.html) ·
[LexoRank troubleshooting](https://support.atlassian.com/jira/kb/troubleshooting-lexorank-system-issues/) ·
[LexoRank rebalance performance](https://jira.atlassian.com/browse/JSWSERVER-13163) ·
[GitLab RelativePositioning](https://gitlab.com/gitlab-org/gitlab/-/issues/230953) ·
[GitLab gap sizing](https://gitlab.com/gitlab-org/gitlab/-/work_items/30352) ·
[GitLab rebalance incident](https://gitlab.com/gitlab-com/gl-infra/production/-/work_items/4584) ·
[begriffs, user-defined order in SQL](https://begriffs.com/posts/2018-03-20-user-defined-order.html) ·
[PostgreSQL wiki: ordering with fractions](https://wiki.postgresql.org/wiki/User-specified_ordering_with_fractions) ·
[C64 RENUMBER](https://www.c64-wiki.com/wiki/RENUMBER_(BASIC_3.5))

CRDTs / collaborative editing: [LSEQ (DocEng 2013)](https://www.researchgate.net/publication/262162421_LSEQ_an_Adaptive_Structure_for_Sequences_in_Distributed_Collaborative_Editing) ·
[The Art of the Fugue](https://arxiv.org/pdf/2305.00583) ·
[Weidner, Fugue list CRDT](https://mattweidner.com/2022/10/21/basic-list-crdt.html) ·
[Weidner, position strings](https://mattweidner.com/2023/04/13/position-strings.html) ·
[Weidner, list-positions](https://mattweidner.com/2024/04/29/list-positions.html) ·
[Kleppmann, moving elements in list CRDTs](https://martin.kleppmann.com/papers/list-move-papoc20.pdf) ·
[archagon, Data Laced with History](http://archagon.net/blog/2018/03/24/data-laced-with-history/) ·
[CRDT dictionary](https://www.iankduncan.com/engineering/2025-11-27-crdt-dictionary/) ·
[OT vs CRDT (Sun et al.)](https://arxiv.org/pdf/1905.01518)

Identifiers & VCS-native: [ULID spec](https://github.com/extrawurst/ulid) ·
[ULID vs UUID](https://www.baeldung.com/cs/ulid-vs-uuid) ·
[time-sortable ids compared](https://www.authgear.com/post/time-sortable-identifiers-uuidv7-ulid-snowflake/) ·
[Why Pijul](https://pijul.org/manual/why_pijul.html) ·
[Pijul model](https://pijul.org/model/) ·
[jneem on merging](https://jneem.github.io/pijul/) ·
[Darcs patch theory](https://en.wikibooks.org/wiki/Understanding_Darcs/Patch_theory_and_conflicts) ·
markdown-kanban field examples: [personal-kanban](https://github.com/YJPL/personal-kanban) ·
[kotban](https://github.com/medavox/kotban) ·
[Backlog.md](https://news.ycombinator.com/item?id=44483530) ·
[Imdone](https://imdone.io/markdown-kanban-board) ·
[TrackDown](https://github.com/mgoellnitz/trackdown)
