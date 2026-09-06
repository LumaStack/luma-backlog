# Journal — Extract the application layer

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-06

reshape split out to WORK-0031 — this change moves no golden file and that is its proof, where every item in the reshape moves them
the seven-eighths estimate for close.go was wrong — it lost two-thirds, 145 lines to 48; the property that mattered was that no judgment stayed in the adapter, and guessing at the arithmetic beforehand added nothing
view types were not planned and turned out to be forced — without them an adapter must hold a record, and the containment test cannot be written, so the test that IS the mechanism made the boundary real rather than nominal
the containment test was checked by breaking it deliberately, which found that the failure message needed to name the file and say where to go instead; a guard nobody has seen fail is a guard nobody knows works
system learning: splitting the reshape out was what made the golden files able to prove anything — a behaviour-preserving change and a behaviour-changing one in the same commit leave nothing able to attest to either
internal/backlog became internal/corpus and internal/policy became internal/guards after this closed — the outcome's verify_by names the old paths and is left as written, since it records what was checked at the time

## ▶ 2026-09-05

split out of WORK-0017 because the refactor is a prerequisite for the board rather than part of the specification, and can start while the specification is still argued
