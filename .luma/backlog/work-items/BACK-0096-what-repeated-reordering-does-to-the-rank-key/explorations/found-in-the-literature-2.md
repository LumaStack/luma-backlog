---
type: exploration
type_version: "0.0.1"
title: Found in the literature, 2
work_item: '[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T20:24:53Z'}
---

# Found in the literature, 2

> **Literature search in isolation.** One agent, working in a directory containing the problem brief
> and nothing else, reported what already exists, with designing a scheme forbidden. **It could not see any other
> agent's answer, and no other agent could see this one** --- the isolation was
> the filesystem's rather than an instruction's.
>
> **Filed before any of the five was read**, so that reading order decided
> nothing. Unedited below this line; the workspace it was produced in has been
> deleted.

---

# Found in the literature — how this problem has been solved elsewhere

**Provenance.** Stage 3 of the brief in `PROBLEM.md`: a literature search, run
blind — no other file in this repository was read, no derivation of my own is
offered. Produced 2026-09-17 from web searches; every claim carries its source.
Where a property of a published scheme is *my arithmetic from its definition*
rather than a statement in a source, it is marked **[derived, unmeasured]** —
the brief asks for that honesty and it applies to a search too.

**The requirement numbers used throughout** are the seven from *What has to be
true*: (1) repair rare, (2) multi-card rewrites a handful of times ever,
(3) millions ahead of an immovable card, (4) millions behind one, (5) the front
of an immovable card appears inexhaustible, (6) rewrite once a decade at most,
(7) traffic on both sides of interior immovable cards forever.

---

## The one theorem that organizes everything else

Before the catalogue: the theory literature contains a result that explains
*why* requirement 7 is the one that decides this, and it is worth carrying as a
frame.

The problem of assigning ordered labels to a dynamic list is studied as
**online list labeling** (also *file maintenance* / *order maintenance*),
opened by [Dietz & Sleator, STOC 1987](https://www.semanticscholar.org/paper/Two-Simplified-Algorithms-for-Maintaining-Order-in-Bender-Cole/4b581c4a6905ea0adbf941b84079ed7709fd82f9)
and still active — a 2022 paper
([Bender et al., "Online List Labeling: Breaking the log²n Barrier"](https://arxiv.org/pdf/2203.02763))
improved a bound that had stood for forty years. The lower bounds are the part
that matters here
([Bulánek, Koucký & Saks](https://arxiv.org/pdf/1112.5636);
[survey](https://link.springer.com/chapter/10.1007/978-3-319-90530-3_3)):

> With labels drawn from a **fixed, polynomial-size space**, inserting n items
> forces **Ω(n log n) relabelings** of *other* items — and with a linear-size
> label space, Ω(n log² n). This is a proved lower bound, not an artifact of
> any particular algorithm.

Translated into the brief's vocabulary: **bounded-width keys, one-file-per-move,
and unbounded interior insertion cannot coexist. This is a theorem.** Every
scheme in the catalogue below is a choice of which leg to give up:

| leg given up | what that looks like | families that chose it |
| --- | --- | --- |
| bounded key width | the stored value grows | fractional indexing, Logoot/LSEQ, Fugue, rationals |
| one file per move | periodic renumbering of neighbours | gapped integers, Trello, LexoRank, order-maintenance structures |
| key orders directly | order is *derived* by a program from identity + references | RGA / Yjs / linked lists |

There is no published fourth option, and the lower-bound literature says there
cannot be one. A candidate claiming all three at once is claiming to have
refuted a STOC lower bound.

---

## Family 1 — Gapped integers, renumber when full

**What it is.** Positions 10, 20, 30…; insertion takes the midpoint; when a gap
closes, renumber. The oldest practice in the field — BASIC line numbers,
`/etc/rc.d` prefixes — and also the *current* practice in the brief's exact
medium: [Backlog.md](https://github.com/MrLesk/Backlog.md/), a git-native
markdown task board, stores an integer `ordinal` field in each task's YAML
frontmatter for custom column order. The
[begriffs survey of user-defined order in SQL](https://begriffs.com/posts/2018-03-20-user-defined-order.html)
opens with this approach and its cost: moving one row can force updating every
subsequent row.

**What it gives up.** Renumbering is routine, not exceptional; a busy gap
renumbers its neighbourhood constantly, and in git that is exactly the
unreviewable many-file ordering diff.

**Requirements failed: 2, 5, 6, 7** (renumbering is the ordinary path, and it
is the forbidden rewrite; 3 and 4 are "met" only by invoking it). Workload 9:
two actors compute the same midpoint — silent identical values.

## Family 2 — Floating-point midpoints (Trello)

**What it is.** Trello gives each card a `pos` — a 64-bit float; new cards get
max+1024 at the bottom, min/2 at the top, the average in between, and when two
values get too close "those cards and some nearby ones are all re-numbered",
worst case a whole list, roughly every ~20 adversarial moves
([Trello engineers on HN](https://news.ycombinator.com/item?id=10957165)).

**What it gives up.** A double survives only ~50 successive midpoints before
two values are equal ([Figma's note](https://madebyevan.com/algos/crdt-fractional-indexing/));
after that, local renumbering — frequent, not rare.

**Requirements failed: 1, 2, 5, 6, 7.** The gap dies after ~50 same-spot
placements — workloads 4 and 5 kill it in an afternoon. Workload 9: identical
`pos`, silent. This family is what a server with cheap multi-row UPDATE can
afford; the git medium cannot.

## Family 3 — Arbitrary-precision fractional indexing (Figma, Notion-class apps)

**What it is.** The scheme Figma named "fractional indexing": every index a
fraction, insertion takes the midpoint, arbitrary precision so it never runs
out ([Evan Wallace, Figma blog](https://www.figma.com/blog/realtime-editing-of-ordered-sequences/);
practical string form in [dgreensp's `fractional-indexing`](https://dev.to/sonim1/fractional-indexing-implementing-drag-and-drop-ordering-and-avoiding-index-collisions-g3),
[Rust `fractional_index`](https://docs.rs/fractional_index/latest/fractional_index/),
[Liveblocks' explainer](https://liveblocks.io/blog/how-crdts-and-sync-engines-keep-realtime-lists-ordered-with-fractional-indexing)).
Keys are digit strings under lexicographic order — plain `sort` works, one key
per item, one write per move, no rewrite ever.

**Growth.** Appends at either end grow keys logarithmically (a counter in a
higher digit), so workloads 1–4 at the *ends* are cheap: ~10 chars at 10⁶,
~14 at 10⁸ **[derived from the base-62 encoding, unmeasured]**. But repeated
insertion into the *same interior gap* adds ~one digit per insertion — this is
precisely the brief's measured family (2,925 placements → 2,929 characters).
At 10⁸ interior placements the key is ~10⁸ characters. Figma's own page says
it plainly: "index length can become long in pathological scenarios… typically
harmless in practice" — a survivable position for a design tool with rebalance
available, not for a board that may never rewrite.

**Concurrency.** Two peers inserting at the same spot generate **identical
keys**; Figma's mitigations are **random jitter** appended to each fraction
plus last-writer-wins with peer-ID tiebreaks, and even then "if two peers both
simultaneously insert a run of objects at the same location, the resulting
objects may be interleaved"
([madebyevan, CRDT: Fractional Indexing](https://madebyevan.com/algos/crdt-fractional-indexing/)).
Unsuffixed, workload 9 is the silent indistinguishable tie the brief fears;
jittered, it is silent but *probabilistically* distinguishable.

**Requirements failed: 5 and 7** — not by ever needing a rewrite, but because
the value's length is linear in same-gap traffic, which fails "appears
inexhaustible" in its practical form and destroys goal 7 (readable value) long
before 10⁸. Passes 1, 2, 3, 4, 6 to the letter.

## Family 4 — LexoRank: fractional indexing plus institutionalized rebalance (Jira)

**What it is.** Atlassian's production answer to Family 3's growth: base-36
fractional strings with a **bucket prefix** (`0|`, `1|`, `2|`). When strings
get long, a background **rebalance** walks every issue, assigning fresh
evenly-spaced ranks in the next bucket; the bucket prefix keeps old and new
ranks totally ordered during the migration. Rebalancing is *scheduled
maintenance*: at 128 characters Jira schedules a rebalance within 12 hours; at
160 it starts one immediately
([Managing LexoRank](https://confluence.atlassian.com/adminjiraserver/managing-lexorank-938847803.html),
[troubleshooting KB](https://support.atlassian.com/jira/kb/troubleshooting-lexorank-system-issues/)).

**What it gives up.** The rewrite. LexoRank does not avoid the many-row
renumbering; it makes it safe, incremental and automatic — on a database,
where a million-row UPDATE is invisible. In git the same operation is a
board-wide diff of ordering data, exactly what requirement 2 forbids, and its
frequency scales with traffic, not with the calendar.

**Requirements failed: 2 and 6** (rebalance is routine and whole-field; at the
brief's volumes it would run far more often than yearly). Meets 3, 4, 5, 7
*only through* those rebalances. Workload 9: same silent-tie behaviour as
Family 3. Its transferable idea is the **bucket/epoch prefix** — a way to make
a rewrite, when one ever does happen, mergeable and ordered against unrewritten
stragglers.

## Family 5 — True rationals: Stern–Brocot mediants (pg_rational)

**What it is.** Store position as an exact fraction and insert between a/b and
c/d at the **mediant** (a+c)/(b+d), walking the Stern–Brocot tree, which
enumerates every rational in lowest terms. Built as a PostgreSQL extension for
exactly this use case
([begriffs, "User-defined order in SQL"](https://begriffs.com/posts/2018-03-20-user-defined-order.html);
[pg_rational](https://github.com/begriffs/pg_rational);
[Simon Willison's summary](https://simonwillison.net/2018/Mar/21/user-defined-order-in-sql/)).

**Why it matters here — a direct qualification of the brief's one measured
result.** The brief measured *decimal midpoint* subdivision: linear growth,
2,929 characters for 2,925 placements. Mediants behave differently on that
same workload **[derived, unmeasured]**: placing repeatedly against a *fixed*
neighbour p/q grows the denominator by q per step, so n placements at one spot
cost a denominator ~n·q — 2,925 placements ≈ **4 digits**, 10⁸ ≈ 9 digits,
digits O(log n). The measured result is a fact about decimal subdivision, not
about subdivision as such. The true worst case is still bad: adversarial
*alternating* insertion (zig-zag down the tree) produces Fibonacci-growth
numerators — digits linear in n again — but none of the brief's ten workloads
has that shape; workloads 1, 3, 4 are monotone runs against a fixed card,
mediants' best case.

**What it gives up.** Lexicographic sortability: `a/b < c/d` is
cross-multiplication, not string order, so goal 5 falls from "plain sort" to
"a short pipeline that computes the quotient and `sort -g`" — the brief's
*acceptable* tier. Overflow of fixed-width integers is the degenerate state
(pg_rational uses 64-bit num/den and can fail a `between` call).

**Requirements failed: none outright for the listed workloads; 5/7 fail under
adversarial alternation** [derived]. Workload 9: the mediant is
*deterministic*, so two actors inserting between the same neighbours produce
**identical values — silent and indistinguishable**, the exact failure the
brief has been bitten by, unless a tiebreak component is added.

## Family 6 — Order-maintenance structures (the relabeling leg, done optimally)

**What it is.** Dietz–Sleator and successors
([Bender et al., "Two Simplified Algorithms for Maintaining Order in a List"](https://erikdemaine.org/papers/DietzSleator_ESA2002/paper.pdf))
keep integer tags in a fixed universe and *amortize* relabelings — O(log n)
amortized per insertion, O(1) with indirection; packed-memory arrays are the
same idea on disk. This is the theoretically optimal version of "renumber when
full": relabel a small neighbourhood, exponentially rarely per card.

**What it gives up.** Relabeling *is the mechanism* — the lower bounds above
say it cannot be avoided in this family. Each relabeled card is a rewritten
file; bursts touch O(log n)–O(polylog n) cards at once, uncoordinated actors
relabeling concurrently produce overlapping renumberings that git cannot
merge meaningfully.

**Requirements failed: 1 and 2** (repair/renumbering is frequent by design,
multi-card by design). Its value to stage 4 is as the *measuring stick*: if
rewrites are ever accepted, this literature says the minimum you must pay and
the smallest neighbourhood that can pay it.

## Family 7 — Growing positional identifiers with site IDs: Logoot and LSEQ

**What it is.** The first CRDT branch: each element carries a variable-length
**position identifier** — a list of levels, each level `(digit, siteID,
counter)` — totally ordered, allocated without coordination, never reassigned
([Logoot overview](https://github.com/t-mullen/logoot-crdt/blob/master/README.md);
[LSEQ paper](https://concordant.lip6.fr/uploads/ConcoRDanT/LSEQ-%20an%20Adaptive%20Structure%20for%20Sequences%20in%20Distributed%20Collaborative%20Editing.pdf)).
Order is total, one value per element, one write per move, **no rewrite ever**
— growth of the identifier is the entire cost.

**Directly relevant detail.** Logoot's identifiers grow badly under monotone
front- or back-editing; LSEQ's contribution is an allocation strategy per tree
level — **boundary+ favours insertion at the back, boundary− at the front**,
chosen randomly per level and remembered — giving empirically polylogarithmic
identifier growth whichever end the pressure comes from. This is the only
literature found that *explicitly engineers for the front/back budget
asymmetry* the brief describes. Growth at 10⁶–10⁸ is beyond anything
characterized in the papers (they measure ~10⁴–10⁵ edits); worst case remains
linear.

**What it gives up.** Goal 7 entirely — identifiers are multi-level tuples,
tens to hundreds of bytes, unreadable; goal 5 partially — a byte-encoding can
make them lexicographically sortable, but they are opaque. Workload 8:
**concurrent** bulk arrivals from two actors can interleave card-by-card — the
documented anomaly ([Kleppmann et al., PaPoC 2019](https://martin.kleppmann.com/papers/interleaving-papoc19.pdf))
affecting exactly Logoot and LSEQ; a single actor's batch keeps its order.

**Requirements failed: none.** Workload 9: the embedded (siteID, counter)
makes two concurrent same-spot placements **distinct, deterministically
ordered, and attributable** — no conflict, no silence about *which* write is
which; the merge is quiet but not indistinguishable. Site IDs must be unique
per actor — in a no-coordination setting that means random, an assumption to
name.

## Family 8 — Fugue / position-strings: the modern synthesis

**What it is.** Matthew Weidner's Fugue
([A Basic List CRDT](https://mattweidner.com/2022/10/21/basic-list-crdt.html);
[The Art of the Fugue, with Kleppmann](https://arxiv.org/pdf/2305.00583) —
proves *maximal non-interleaving*) and its packaged form
[position-strings](https://github.com/mweidner037/position-strings/blob/master/algorithm.md)
([explainer](https://mattweidner.com/2023/04/13/position-strings.html);
successor [list-positions](https://github.com/mweidner037/list-positions)).
Each position is a **single string**; **lexicographic order on the strings is
the list order** — plain `sort` on one field, the brief's goal-5 *best* tier.
Strings encode paths in a tree whose branch points ("waypoints") are labeled
with the writer's random replica ID; a writer extending its own run reuses its
waypoint and just increments a counter, so **monotone runs grow the string
logarithmically, not linearly** — the library's stated point of difference
from fractional indexing, along with **global uniqueness** and
**non-interleaving**.

**Behaviour against the brief's asks.** One value per card, one file per move,
no rewrite ever, no tombstones needed for ordering (each string is
self-contained — the ancestry is *inside* the value, so a departed neighbour
costs nothing and re-entry mints a fresh position: workloads 6 and 7 are
first-entry-priced). Workload 8: batches from one actor stay contiguous, and
concurrent batches cannot interleave (the Fugue theorem). Workload 9: two
actors at the same spot produce distinct strings under their own waypoints,
deterministically ordered, attributable by embedded replica ID — quiet, valid,
distinguishable. Benchmarked sizes: **10–100 characters** on real collaborative
editing traces; behaviour at 10⁸ operations in one gap is **not characterized
anywhere found** — the honest gap in this family's evidence, since worst-case
growth is still linear in tree depth.

**What it gives up.** Goal 7: the value is 10–100 characters of machine
gibberish (`r95sU223.B…`), sortable but not meaningful; migration of existing
numeric ranks needs a one-time mint (one big, but mechanical, diff); actors
need unique random IDs, and the randomness is what stands between workload 9
and a same-key tie.

**Requirements failed: none found** — with the caveat that 5/7 at 10⁸ rest on
extrapolated logarithmic growth, not on a measurement at that scale.

## Family 9 — Derived order: linked lists made merge-safe (RGA, Causal Trees, Yjs/YATA, WOOT, Treedoc)

**What it is.** The other CRDT branch abandons ordering-by-value: each element
stores a **unique ID plus a reference to the element it was inserted after**;
the sequence is *derived* by a deterministic tree walk, with concurrent
same-anchor children ordered by ID
([RGA](https://www.researchgate.net/publication/220379659_Replicated_abstract_data_types_Building_blocks_for_collaborative_applications);
[survey of WOOT/CT/RGA/YATA](https://arxiv.org/pdf/2310.18220);
[Treedoc, which needs a coordinated "flattening" epoch to shorten paths]).
This is the database folk-practice "predecessor pointer / linked list column"
with its merge problems actually solved. Yjs (YATA) and Automerge (RGA-family)
are the production implementations of local-first software.

**What it gives up.** Goal 5 at the *last resort* tier, in full: **no sort on
any stored field yields the order** — a program (or a nontrivial script
walking anchor references) must reconstruct it, which is the brief's "database
with a worse query language" cost, priced exactly. Two further documented
costs: **tombstones** — an element that leaves must leave a marker, or a
concurrent insertion anchored to it dangles (time/space overhead is the
standard criticism of this class); and **moves**: a move implemented as
delete+reinsert means two actors concurrently moving the same card produce
**two copies** — documented, with the fix (a move operation treating position
as a LWW register per element) in
[Kleppmann, "Moving Elements in List CRDTs", PaPoC 2020](https://martin.kleppmann.com/papers/list-move-papoc20.pdf).
That paper is the closest single match to "cards that move between columns" in
the entire literature and its failure mode (silent duplication under
concurrent moves) applies to *any* scheme here that models a move as
delete-then-insert, Families 3–8 included.

**Requirements failed: none** — IDs are small and fixed-width forever, interior
insertion is free forever, no rewrites. Everything is paid in goal 5 and in
tombstone bookkeeping.

## Family 10 — Operational Transformation

Order maintained by transforming operations against concurrent ones — requires
a central server (or total order broadcast) to sequence operations; Figma
rejected it as overkill and the OT/CRDT comparison literature is extensive
([Sun et al.](https://arxiv.org/pdf/1810.02137)). **Fails the no-coordination
hard constraint outright.** Listed only so stage 4 can say why it was excluded.

## Family 11 — The order lives in a manifest, and git merges it as text

**What it is.** Not a paper — a practice, widespread in exactly the brief's
medium: markdown kanban tools keep each column's order as a **list in a text
file**, cards carrying no rank at all
([Backlog.md](https://github.com/MrLesk/Backlog.md/) (board views),
[MarkdownTaskManager](https://github.com/ioniks/MarkdownTaskManager) — whose
docs discuss resolving `kanban.md` merge conflicts by hand,
[TODO.md kanban](https://www.producthunt.com/products/todo-kanban-board),
[kotban](https://github.com/medavox/kotban) — column-per-directory,
[KanbanBoardInMarkdown](https://github.com/MozaicWorks/KanbanBoardInMarkdown)).
A reorder is a one-line move in one file; the diff is the most readable of any
family (the line *is* the card's name); ordering the board without a program is
`cat`.

**Concurrency is git's line merge, with both its virtues and its documented
trap.** Two actors inserting at the same spot edit adjacent lines → a **loud
conflict a person resolves by reading two lines** — the *only* family whose
same-spot behaviour is the brief's stated goal-4 preference. But git merges
hunks that are far apart *silently*: a card **moved in both branches**
(delete+insert at two distant places) merges cleanly into **two copies or a
lost card** — the same move-duplication failure Family 9 documents, here
produced by `git` itself. It is silent but *detectable* by a uniqueness check
over the manifest, and the degenerate state (duplicate/missing lines) is still
readable. Interior insertion is free forever — line position carries no
budget.

**Requirements failed: none of the seven as written** — they are phrased about
cards, and no card is ever rewritten. What it abandons is the brief's opening
design sentence: *"whatever a card needs in order to know its place has to be
written inside that card"* — which the brief itself lists as renegotiable
("that the order lives in one field… or something stored alongside"). Costs to
name: the manifest is a per-column contention hotspot (conflict frequency
scales with concurrent reordering of the same column); a cross-column move
touches two or three files (the brief's all-or-nothing provision covers this);
manifest and card files can disagree (card exists, line missing) — an
invariant to detect, per the hard constraints.

## Family 12 — Do not store an order: compute it

[Taskwarrior's urgency](https://taskwarrior.org/docs/urgency/) orders tasks by
a configurable polynomial over task attributes (age, due date, priority, tags)
— there is no hand-ordered list at all, and the docs argue this positively:
sorting cannot express what a weighted score can. This is the literature's
standing answer to the brief's own open question *"whether the thing being
ordered is the right thing to order"*: if the order is *derivable* from
attributes people already set, the whole allocation problem dissolves —
nothing is consumed by placement, concurrent "reorders" become attribute edits
that merge field-wise. It fails the brief only if the order genuinely carries
information no attribute holds (pure "I want this one above that one"). Passes
all seven requirements vacuously; gives up manual ordering itself.

## Family 13 — Components, not schemes

- **Sortable time-ordered IDs** (ULID/KSUID/Snowflake): merge-safe,
  fixed-width, lexicographic creation-order keys — exactly the shape of the
  append-only *done* column (workload 2) and of "advancing lands at the back",
  but with no interior insertion at all: as a whole scheme they fail 7 on the
  first promotion. Relevant as the cheap half of a hybrid.
- **Order-preserving encodings** ([ELEN](https://github.com/ealmansi/elen),
  [bytekey](https://github.com/danburkert/bytekey)): make variable-length
  numbers sort correctly as plain strings — the substrate that could lift
  Family 5's rationals or any numeric scheme from "needs `sort -g` on a
  computed quotient" toward plain lexicographic sort, at the cost of an
  encoded (less readable) stored form.

---

## The cross-cutting findings

**1. The trilemma is proved, not folklore.** Fixed-width value, no multi-card
rewrites, unbounded interior insertion: pick two
([lower bounds](https://arxiv.org/pdf/1112.5636)). Families 1/2/4/6 rewrite;
3/5/7/8 let the value grow; 9/11 stop ordering by value. Stage 4's real
decision is *which leg*, and the brief's requirements 2+6 all but rule out the
rewriting leg — which the theory says forces growing values or derived order.

**2. On the brief's one measured result.** The measurement (linear growth,
2,925 → 2,929 chars) is correct *for midpoint/decimal subdivision* (Family 3)
— but "subdividing a finite interval cannot meet the requirements" is broader
than the literature supports. Mediant subdivision (Family 5) grows
logarithmically on that same monotone workload **[derived, unmeasured — worth
five minutes of measurement before stage 4 repeats the brief's fifth method
failure]**, and tree-path schemes (7, 8) are engineered to be logarithmic on
monotone runs. What *is* universally true: every subdivision scheme has *some*
adversarial pattern with linear growth — the difference is whether the brief's
actual workloads are that pattern (for decimal midpoints they are; for
mediants and waypoint trees they are not).

**3. Workload 9 splits the field three ways.** (a) *Silent and
indistinguishable*: deterministic allocators — decimal midpoints, mediants,
unsuffixed fractional strings — two actors produce byte-identical values; the
brief's own past failure. (b) *Silent but distinguishable and deterministic*:
anything embedding an actor ID in the value (Logoot/LSEQ, Fugue/
position-strings, Figma's jitter probabilistically) — the merge is quiet, both
writes valid and attributable, the order between them arbitrary but stable.
(c) *Loud*: only Family 11, where the same spot is the same lines of the same
file and git raises a conflict. The CRDT literature's core value — always
merge silently — is the opposite of the brief's goal-4 preference for loud;
the literature's defence is that its silent merges are *valid and
attributable*, never the indistinguishable tie. Stage 4 should decide whether
(b) satisfies goal 4 or only (c) does.

**4. Moves are the under-examined operation everywhere.** Nearly all of this
literature is about *insertion*; the brief's dominant operation is *moving an
existing card*. The one paper squarely about it
([Kleppmann 2020](https://martin.kleppmann.com/papers/list-move-papoc20.pdf))
shows move-as-delete+reinsert **duplicates cards under concurrent moves** in
any scheme — and git's text merge reproduces the same failure in Family 11.
Any stage-4 candidate should be asked the brief's workload-9 question a second
time, about two actors *moving the same card* rather than placing at the same
gap.

**5. The front/back asymmetry has prior art.** LSEQ's boundary+/boundary−
strategies are precisely "the two budgets are consumed differently", solved by
biasing allocation per level toward the end under pressure — evidence the
asymmetry is real enough that a literature grew around it, and a design any
positional candidate can borrow.

**6. Nobody in the literature runs this medium.** Every polished scheme
(Figma, Jira, Trello, CRDT libraries) assumes a program present at merge time
or a database beneath. The only prior art for *git-merges-the-files-alone* is
the markdown-kanban folk practice (Family 11) and integer ordinals in YAML
(Backlog.md, Family 1) — both naive by this brief's standards. The
combination of an academic-grade scheme with git-as-the-merge-engine appears
to be genuinely unoccupied territory; the closest published fit is Family 8
(a single lexicographically sortable string field, no rewrites, unique and
non-interleaved under concurrency), and its untested spot is growth at the
brief's 10⁸ scale.

## Summary table

| family | value at 10⁶ / 10⁸ same-gap ops | files per move | workload 9 | rewrite ever needed | requirements failed |
| --- | --- | --- | --- | --- | --- |
| 1 gapped ints | n/a — renumbers first | 1, or the column | silent identical | routinely | 2, 5, 6, 7 |
| 2 floats (Trello) | dead at ~50 ops | 1, or ~column/20 moves | silent identical | routinely | 1, 2, 5, 6, 7 |
| 3 fractional strings (Figma) | ~10⁶ / ~10⁸ chars | 1 | silent; identical unless jittered | never (grows instead) | 5, 7 (via length) |
| 4 LexoRank (Jira) | capped 254 chars | 1 | silent identical | scheduled, threshold-driven | 2, 6 |
| 5 mediants (pg_rational) | ~7 / ~9 digits [derived] | 1 | **silent identical** | on int overflow | none listed; 5/7 adversarial |
| 6 order-maintenance | small ints | O(log n) amortized | merge chaos | constantly, by design | 1, 2 |
| 7 Logoot/LSEQ | polylog, uncharacterized at 10⁸ | 1 | silent, distinct, attributable | never | none (goals 5/7 traded) |
| 8 Fugue/position-strings | 10–100 chars measured; log extrapolated | 1 | silent, distinct, non-interleaved | never | none found (goal 7 traded; 10⁸ unmeasured) |
| 9 RGA/Yjs derived | fixed-width ID forever | 1 (+tombstone) | distinct; moves can duplicate without move-op | never | none (goal 5 fully traded) |
| 10 OT | — | — | — | — | fails no-coordination |
| 11 manifest file | zero (line position) | 1 (2–3 cross-column) | **loud conflict**; move-vs-move can silently duplicate | never | none as written (per-card promise traded) |
| 12 computed order | zero | 1 (attribute edit) | field-wise merge | never | none (manual order itself traded) |

## Sources

- [Figma — Realtime Editing of Ordered Sequences](https://www.figma.com/blog/realtime-editing-of-ordered-sequences/) · [madebyevan — CRDT: Fractional Indexing](https://madebyevan.com/algos/crdt-fractional-indexing/) · [dgreensp fractional-indexing writeup](https://dev.to/sonim1/fractional-indexing-implementing-drag-and-drop-ordering-and-avoiding-index-collisions-g3) · [fractional_index (Rust)](https://docs.rs/fractional_index/latest/fractional_index/) · [Liveblocks explainer](https://liveblocks.io/blog/how-crdts-and-sync-engines-keep-realtime-lists-ordered-with-fractional-indexing)
- [Atlassian — Managing LexoRank](https://confluence.atlassian.com/adminjiraserver/managing-lexorank-938847803.html) · [LexoRank troubleshooting](https://support.atlassian.com/jira/kb/troubleshooting-lexorank-system-issues/)
- [Trello pos discussion (HN)](https://news.ycombinator.com/item?id=10957165)
- [begriffs — User-defined Order in SQL](https://begriffs.com/posts/2018-03-20-user-defined-order.html) · [pg_rational](https://github.com/begriffs/pg_rational) · [Willison summary](https://simonwillison.net/2018/Mar/21/user-defined-order-in-sql/)
- [Bulánek–Koucký–Saks lower bounds](https://arxiv.org/pdf/1112.5636) · [Online labeling survey](https://link.springer.com/chapter/10.1007/978-3-319-90530-3_3) · [Bender et al. 2022, breaking log²n](https://arxiv.org/pdf/2203.02763) · [Bender et al., Two Simplified Algorithms](https://erikdemaine.org/papers/DietzSleator_ESA2002/paper.pdf)
- [LSEQ paper](https://concordant.lip6.fr/uploads/ConcoRDanT/LSEQ-%20an%20Adaptive%20Structure%20for%20Sequences%20in%20Distributed%20Collaborative%20Editing.pdf) · [Logoot README](https://github.com/t-mullen/logoot-crdt/blob/master/README.md)
- [Kleppmann et al. — Interleaving anomalies (PaPoC 19)](https://martin.kleppmann.com/papers/interleaving-papoc19.pdf) · [Kleppmann — Moving Elements in List CRDTs (PaPoC 20)](https://martin.kleppmann.com/papers/list-move-papoc20.pdf) · [Weidner & Kleppmann — The Art of the Fugue](https://arxiv.org/pdf/2305.00583)
- [Weidner — Fugue: A Basic List CRDT](https://mattweidner.com/2022/10/21/basic-list-crdt.html) · [position-strings algorithm](https://github.com/mweidner037/position-strings/blob/master/algorithm.md) · [position-strings explainer](https://mattweidner.com/2023/04/13/position-strings.html) · [list-positions](https://github.com/mweidner037/list-positions)
- [RGA (Roh et al.)](https://www.researchgate.net/publication/220379659_Replicated_abstract_data_types_Building_blocks_for_collaborative_applications) · [CRDT approaches survey](https://arxiv.org/pdf/2310.18220) · [OT vs CRDT](https://arxiv.org/pdf/1810.02137)
- [Backlog.md](https://github.com/MrLesk/Backlog.md/) · [MarkdownTaskManager](https://github.com/ioniks/MarkdownTaskManager) · [kotban](https://github.com/medavox/kotban) · [TODO.md kanban](https://www.producthunt.com/products/todo-kanban-board) · [KanbanBoardInMarkdown](https://github.com/MozaicWorks/KanbanBoardInMarkdown)
- [Taskwarrior — Urgency](https://taskwarrior.org/docs/urgency/)
- [ELEN](https://github.com/ealmansi/elen) · [bytekey](https://github.com/danburkert/bytekey)
