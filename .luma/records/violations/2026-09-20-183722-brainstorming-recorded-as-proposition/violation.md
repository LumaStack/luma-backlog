---
type: violation
violation_id: 2026-09-20-183722-brainstorming-recorded-as-proposition
violating_commit: 48bda9e
violating_actor: agent:claude-fable-5
occurred_at: 2026-09-20T18:34:21Z
noticed_at: 2026-09-20T18:37:22Z
noticed_by: human:maintainer
delivery: delivered
expectation: A record restates what a person said in a form they would recognize — brainstorming is never written down as a proposition, and thinking out loud is never upgraded to a stance taken.
policy: local/backlog procedure/backlog-capture 0.46.0
created_using: lumastack/luma-catalog/violation-records 0.6.0
---

# Brainstorming recorded as a proposition

The maintainer talked through a linking possibility — store `promoted_to`
only, derive the rest — as one of several options, thinking out loud. The
agent wrote "the maintainer proposed storing only `promoted_to`" into
BACK-0106's body and repeated the framing in ADR-0011. What was wanted was
the possibility recorded as a possibility.

## What was wanted

A record restates a person's words in a form they would recognize as theirs.
Brainstorming, talking through options, and "hmmm" are never written as
propositions, positions, or requests — a proposition carries commitment the
speaker did not give.

## Why it was not followed

**delivered** — the capture procedure's test was in context this session and
is exactly this: *"would they recognise this as what they said, or as what
somebody thought they meant?"* The agent applied it to descriptions and not
to its own analysis sections, where the maintainer's stance was paraphrased
with a stronger verb than the conversation carried. If this recurs, the rule
may need stating beyond the capture procedure — it binds anywhere a record
attributes a stance to a person.
