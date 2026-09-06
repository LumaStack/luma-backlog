---
type: work-item
key: WORK-0032
title: How goalpost fitting is discouraged without being prevented
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T14:20:00Z'}
---

# How goalpost fitting is discouraged without being prevented

## The problem

**We should discourage goalpost fitting without preventing it.**

It is a multi-layer approach:

- **The development lifecycle tooling should tell the model not to do it.** That
  is outside the scope of this tool.
- **The backlog should scold you when you do it**, so a model can correct
  itself, and so it surfaces as a problem to a person when it happens anyway.

**How it should work is not settled.** That is what this inquiry is for.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**Raised by hitting it.**
[[backlog/work-items/WORK-0018-extract-the-application-layer]] finished its work
and cannot close: it has no outcomes, so `close … delivered` is refused. The
obvious next move — write the outcomes now — is goalpost fitting, and there was
no signal to distinguish it from writing down a bar that was already stated in
prose before the work began.

**The scolding posture already exists.** `spec.md` §5.2 reports conditions and
refuses nothing, and states the tone: *"the tool states what it observed and
what that suggests — never what someone should have done."* This is a new
condition rather than a new mechanism.

**And the inverse condition is already there.** `work-item.drifted` catches
*work happened, but no outcome was verified or revised — the specification has
fallen behind reality.* Goalpost fitting is the same gap from the other side:
the specification racing to catch up with reality. They are a pair.

### The hard part: what is discouraged looks identical to what is encouraged

`spec.md` §2.4 says outcomes are **expected** to be *"tightened, split, and
rewritten as work reveals things"* — §4.4 separates `title` from
`desired_state` precisely so that can happen without breaking identity. So
revising an outcome mid-work is correct and must never be scolded.

From outside, both are *an outcome changed while work was underway*.

**What separates them is creation, not modification.** Tightening acts on an
outcome that already existed; fitting creates one that did not. The tool can see
that difference — `created` is on every record — and it is the only signal so
far that does not also fire on the behaviour the design wants.

### What a signal could be

| Signal | Catches | Fires wrongly on |
| --- | --- | --- |
| Outcome created after its work item began | declaring the bar after starting | an outcome legitimately discovered mid-work, which §2.4 also expects |
| Outcome verified moments after being created | writing a bar and ticking it | backfilling an old corpus, where both are recorded together |
| Outcome created **and** verified by the same actor | the doer setting their own bar | a solo maintainer, which is this project |

**None is clean, and the third is worth noting**: on a project with one person
and their agents, the honest signal fires constantly. A condition that always
fires is one people stop reading — which `spec.md` §5.2 already warns about for
duplicate keys.

### What must not happen

- **No refusal.** §5.0 permits refusing only what the caller's own record
  contradicts, and a late outcome contradicts nothing. It would also punish
  recording the bar at all, which is worse than recording it late.
- **No new field somebody has to set.** A signal that depends on being declared
  is one that will not be.
- **It must survive being ignored.** These fire rarely and matter once, which is
  the hardest kind of warning to write well.

## What this produces

A recommendation, and possibly a condition. **Concluding that the timestamps
cannot tell the two apart, and that this belongs entirely to the layer above, is
a complete result.**

## Related

[[backlog/work-items/WORK-0026-what-deserves-to-be-a-work-item]] — whether
`work-item.unarticulated` is the whole answer to records with no bar at all.

## References

- `docs/spec.md` §5.2 — the conditions, and the tone they are reported in.
- `docs/spec.md` §2.4, §4.4 — outcomes are expected to be rewritten.
- `docs/spec.md` §5.0 — what may be refused.
