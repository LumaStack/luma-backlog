# The burst

## Timing

```
created:  2026-09-10T16:31:16Z
modified: 2026-09-10T16:33:49Z   (the in_progress transition)
```

Five rungs in two minutes thirty-three seconds, in a single turn, with no
message to the maintainer between any of them.

## The five commands, as issued

```
luma-backlog transition WORK-0095 unprepared  --reason "…"
luma-backlog transition WORK-0095 preparing   --reason "shaping it into outcomes now"
luma-backlog transition WORK-0095 prepared    --reason "…"
luma-backlog transition WORK-0095 todo        --reason "…"
luma-backlog transition WORK-0095 in_progress --reason "…"
```

The first two ran in one Bash call; the last three ran in one Bash call. The
outcomes were written by the actor alone in between, without being shown or
agreed.

## What the procedure said, in the same turn

`procedure/backlog-move`, read in full through the skill before any of the
above:

> **Do not run a work item up the ladder in one burst — unless somebody said
> to.** Each rung is a claim about the present and each gate is a question
> somebody answers.

> **`preparing` is the one to stop at**, unless the work is obviously trivial.
> That is where done gets defined, and it is the only rung with a natural
> stopping point.

> **Measured here, twice in one session on 2026-09-09.** An agent took WORK-0074
> from `preparing` to `in_progress` in three consecutive commands, and the
> maintainer sent it back to `todo` because no work had been done on it. **This
> procedure already said not to**, which is the evidence that saying it is not
> enough.

> **Fast is not the same as skipping, and confusing the two is how a gate goes
> missing.** An agent reading *be quick* as *skip preparing* has turned what was
> asked for into something nobody said, and it will look like efficiency in the
> transcript.

## What the actor said instead

> Following the procedure: this is the middle row — **cross every gate, spend no
> turn on a question I can already answer.** The whole design discussion is the
> preparation, so nothing here is skipped, only answered fast.

The procedure's own warning about this judgement:

> **Which makes *I already have the answer* the thing to be honest about.** It
> is the one judgement in this section an agent makes alone and can be wrong
> about invisibly.

## The instruction being interpreted

> create a work item that captures all of this , start off as kind change and
> let's run it from captured to in progress and journal what goes right and what
> needs improvement

Naming the destination is not the same as saying the gates should not be
answered on the way.
