# Journal — How two workstations avoid colliding

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-05

the trilemma is the shape of it: sortable, dense, unique-by-construction — pick two. A counter drops uniqueness, a fraction or a gap drops density, an actor component drops sorting.
which property to give up depends on what a key is for — citation wants uniqueness and treats sorting as decoration; reading a backlog in order wants sorting. Nobody has said which.
leaning toward a fraction, and the table had it wrong: it judged every scheme as allocation, and a fraction is repair — nothing hands out 13.5 up front
the property that makes it work is that it does not cascade: renumbering to the next integer walks into whatever holds it and the fix runs through the corpus, where a fraction moves one record and nothing else
confirming no collision is detection rather than allocation, useful immediately, and independent of which scheme wins — a check for two records holding one key works against the counter shipped today
increment first, fraction only when incrementing would cascade — that keeps fractions rare and the corpus dense, which is the cost the table charged them
the safe test is whether the next number EXISTS, not whether anything references it: existence is checkable, references escape the corpus into commits and conversations, and decision-records already says assume you will miss one
spacing unsettled; the outer principle probably decides it — sequential decimals until one has to go between two of them, then bisect
better answer: pick a loser and append it to the end. One record moves, nothing cascades, no fractions, and the corpus stays dense integers.
correcting this record: the cascade claim assumed repairing to n+1. The repair is to append, and the end of the sequence is free by construction, so there was never a cascade to avoid.
the cost is that a key stops implying creation order, which is acceptable because created already records it — a key encoding order would be a second copy of a fact another field holds

## ▶ 2026-09-04

raised the same day the key shipped: spec.md §6.4 names deriving names from titles rather than counters as the mitigation, and the key is a counter — the path beside it is not
two questions tangled: colliding identifiers, and two people doing the same work. They may not want the same mechanism.
