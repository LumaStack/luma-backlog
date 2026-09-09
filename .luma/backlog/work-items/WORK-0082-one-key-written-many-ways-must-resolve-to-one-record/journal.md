# Journal — One key written many ways must resolve to one record

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-09

measured before capturing: case already works (work-0074, WoRk-0074 both resolve), everything else fails — WORK-74, WORK-0000000074, WORK---74, 'WORK 74', bare 74
the cause is one line: NormalizeKey upper-cases and returns, so WORK-74 is recognized AS A KEY and then compared as a STRING against WORK-0074. the number is never parsed
two different failures hide in one list — WORK-74 and WORK-0000000074 match keyPattern and still fail, which is repairing a promise already made; WORK---74 and 'WORK 74' do not match at all, and accepting them widens what a key IS, which is a new promise
normalizing to WORK-0074 is nearly free because FormatKey already renders it; the lowercase collapsed form would need TWO representations, a comparison form and a display form, since no record on disk is spelled work-74
third option for preparation: normalize to a PARSED value, prefix and number, compare those, render with FormatKey only at the edges — that makes 'compared as strings' impossible by construction rather than by discipline
the sweep should search for places that compare a key and never call NormalizeKey — there are only two callers, both in load.go, so listing callers finds nothing and inverting the search is the point
