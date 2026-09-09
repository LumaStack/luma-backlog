---
type: work-item
key: WORK-0082
title: One key written many ways must resolve to one record
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:09:57Z'}
description: WORK-0074, WORK-74, work-74, WORK-0000000074, 'WORK      74', WORK---74 and WoRk-74 should all be the same key in the internal engine. wants tests and a sweep to prove it. one form is the most normal — downcase, spaces to dashes, collapse runs of dashes, strip leading zeros — but normalizing to WORK-0074 may be better, so that if the normalization is ever printed it is already the correct form.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:09:57Z'}
---

# One key written many ways must resolve to one record

## The problem

**All of these name the same work item and only two of them find it:**

```
WORK-0074   WORK-74   work-74   WORK-0000000074   "WORK      74"   WORK---74   WoRk-74
```

**We should be normalizing keys**, and there should be **tests and a sweep** to
prove the engine treats them as one.

## Which normal form

**One option is the most normal**: downcase, replace spaces with dashes, reduce
runs of dashes down to one, and strip leading zeros on the number. That gives
`work-74`.

**The other is to normalize to `WORK-0074`** — so that if we ever output the
normalization, it is already the correct form.

## What is being delivered

**Normalization in the engine, tests that pin every spelling above, and a sweep
that finds the places comparing keys without it.**

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### Measured, not assumed

**Run against this corpus on 2026-09-09:**

| written | resolves |
| --- | --- |
| `WORK-0074` | ✔ |
| `work-0074` | ✔ |
| `WoRk-0074` | ✔ |
| `WORK-74` | ✘ |
| `work-74` | ✘ |
| `WORK-0000000074` | ✘ |
| `WORK---74` | ✘ |
| `WORK 74` | ✘ |
| `74` | ✘ |

**So case is already handled and nothing else is.**

### The cause is one line, and it is narrower than the list suggests

**`internal/corpus/key.go:46` — `NormalizeKey` upper-cases and returns.** Its own
comment says so: *"NormalizeKey upper-cases a key so `work-00002` finds
`WORK-00002`."*

**`keyPattern` is `^([A-Z]+)-(\d+)$`.** That splits the failures into two kinds
that need different fixes:

- **`WORK-74` and `WORK-0000000074` match the pattern and still fail.** They are
  recognized *as keys*, upper-cased, and then compared **as strings** against
  `WORK-0074`. **The number is never parsed.** This is the actual defect: the
  tool identifies something as a key and then declines to find it.
- **`WORK---74` and `WORK 74` do not match at all**, so they fall through and
  are treated as somebody's slug. Accepting them is a widening of what a key
  *is*, which is a different decision from comparing two keys correctly.

**Worth keeping those apart when scoping.** The first is repairing a promise
already made; the second is making a new one.

### The second form is the better one, and it is nearly free

**`FormatKey` already exists and already renders it** —
`internal/corpus/key.go:35`, `fmt.Sprintf("%s-%04d", KeyPrefix, number)`. So
normalizing to `WORK-0074` is *parse the number, call the function that is
already there*.

**The lowercase collapsed form would cost more, not less.** `work-74` matches no
record on disk — every stored key is `WORK-0074` — so choosing it means keeping
**two** representations, a comparison form and a display form, and remembering
which one is in hand at every boundary. **The maintainer's second instinct is
right and the reason is stronger than aesthetics: there is only one form that is
correct to print, so making it the only form removes a class of bug rather than
a keystroke.**

**A third option is worth putting on the table at preparation**: normalize to a
**parsed value** — prefix and number — rather than to a string, compare those,
and render with `FormatKey` only at the edges. That makes *"compared as strings"*
impossible by construction rather than by discipline, which is the failure mode
this record is about.

### The sweep has a second job nobody asked for

**Key comparison is not only in resolution.**
[[work-items/WORK-0014-detect-two-records-holding-one-key]] shipped duplicate
detection, and
[[records/decisions/ADR-0003-a-colliding-key-is-repaired-by-appending]] defines
the repair. **If two records ever carried `WORK-74` and `WORK-0074`, the
detector has to see one key and not two** — and whether it does today is
untested.

**Allocation has the same exposure.** `highestKey` reads keys back out of the
corpus with the same pattern; a key spelled unusually would be invisible to it,
and the next allocation would reuse the number.
[[work-items/WORK-0040-a-decision-number-is-reused-when-its-file-is-absent]] is
the same failure from a different cause.

### What the sweep should actually look for

**Not `NormalizeKey` callers — there are only two**, both in `load.go`. **The
finding is the places that compare a key and never call it.** That inverts the
search and is the reason a sweep is worth more here than a fix.

## Out of scope

**Resolving a record by path.**
[[work-items/WORK-0023-refer-to-a-record-by-the-path-a-person-would-type]] covers
`WORK-0017/outcomes/<slug>`, which is a different reference problem.

**Whether the key prefix is configurable.** The pattern accepts any `[A-Z]+`,
and nothing here changes that.

## Constraints

- **A key is meant never to change** (`internal/corpus/key.go`, and ADR-0003's
  appending repair). Normalization must not rewrite what is stored — it is about
  how a reference is *read*, never about editing records.
- **Padding is four digits and deliberately breaks past 9999**, which the code
  comment already accepts: *"it fails exactly where it had stopped being worth
  anything."* Normalization must not re-pad a five-digit key into something
  shorter.

## References

- `internal/corpus/key.go:35,40,46` — `FormatKey`, `IsKey`, `NormalizeKey`.
- `internal/corpus/load.go:240,302` — the only two callers.
- [[work-items/WORK-0014-detect-two-records-holding-one-key]] — duplicate
  detection, which compares keys.
- [[records/decisions/ADR-0003-a-colliding-key-is-repaired-by-appending]] — what
  a collision means.
- [[work-items/WORK-0023-refer-to-a-record-by-the-path-a-person-would-type]] —
  the other reference problem.
