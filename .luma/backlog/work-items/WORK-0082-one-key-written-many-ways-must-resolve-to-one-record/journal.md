# Journal — One key written many ways must resolve to one record

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-20

outcomes accepted by the maintainer; scope settled in the same exchange: prefixes follow Jira Cloud rules (^[A-Z][A-Z0-9]{1,9}$ — decided for WORK-0102 and adopted here for recognition), the sloppy separators from the capture are IN because the capture lists them, bare 74 stays OUT — deferred, reopen if typing bare numbers is what people do. it collides with decision numbers
built as the third option from preparation: ParseKey returns (prefix, number), every comparison goes through it via SameKey, rendering goes through FormatKeyAs — 'compared as strings' is now impossible by construction
the sweep, inverted as planned — every site that touched a key, and what each needed: load.go Resolve compared it.Key() == NormalizeKey(ref) as strings → SameKey; load.go scoped() same fix; duplicate.go Duplicates bucketed by the RAW stored key, so WORK-74 and WORK-0074 on two records read as two keys and the WORK-0014 detector was blind to exactly the collision ADR-0003 repairs → buckets by NormalizeKey; key.go highestKey pattern-matched stored keys, so an unpadded key's number was invisible to allocation and would be REUSED (WORK-0040's failure from a different cause) → ParseKey; create.go FormatKey and app/view.go Key() display only, untouched
NormalizeName grew the same canonicalization for the key half of a joined name, single-dash separator only — a name comes off a directory listing, not out of prose, so sloppy separators stay out of it
Jira Cloud narrows as well as widens: X-1 stops being a key (below the two-char minimum). nothing in this corpus uses a one-letter prefix, cost measured at zero
surprise worth keeping: the whole suite passed first run including cli — nothing anywhere depended on the old string comparison succeeding for sloppy spellings, because sloppy spellings never resolved at all

---

## 2026-09-09

measured before capturing: case already works (work-0074, WoRk-0074 both resolve), everything else fails — WORK-74, WORK-0000000074, WORK---74, 'WORK 74', bare 74
the cause is one line: NormalizeKey upper-cases and returns, so WORK-74 is recognized AS A KEY and then compared as a STRING against WORK-0074. the number is never parsed
two different failures hide in one list — WORK-74 and WORK-0000000074 match keyPattern and still fail, which is repairing a promise already made; WORK---74 and 'WORK 74' do not match at all, and accepting them widens what a key IS, which is a new promise
normalizing to WORK-0074 is nearly free because FormatKey already renders it; the lowercase collapsed form would need TWO representations, a comparison form and a display form, since no record on disk is spelled work-74
third option for preparation: normalize to a PARSED value, prefix and number, compare those, render with FormatKey only at the edges — that makes 'compared as strings' impossible by construction rather than by discipline
the sweep should search for places that compare a key and never call NormalizeKey — there are only two callers, both in load.go, so listing callers finds nothing and inverting the search is the point
