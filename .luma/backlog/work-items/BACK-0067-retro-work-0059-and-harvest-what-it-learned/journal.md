# Journal — Retro WORK-0059 and harvest what it learned

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

MATERIAL FROM THE 2026-09-09 SESSION, for whoever runs this retro — it is now two sessions of evidence rather than one
the maintainer's summary judgement, recorded in their words rather than softened: 'we seem to have a pattern of creating noise more than creating signal'. they attribute it to the model and note they did not have this problem with the previous one. that is a claim the retro can test against the record rather than take on faith
four failures of the same shape, all caught by the maintainer and none by the tool: gates crossed in one burst without stopping; work started without confirming preparing was finished; twelve tasks worked and closed with none marked; and a decision record's reopen trigger cited as fired when it had not, on evidence that pointed the other way
the fourth is the one worth studying — it was not a process slip but a REASONING slip: an empty rung was read as evidence about the design, when the rung had been emptied by hand for an unrelated reason. cause and effect inverted, then documented confidently
and two outcomes were written wrong by the agent that wrote them, both caught only by reading every check rather than the first. one asserted a property spec 5.3.1 deliberately does not have; the other asked for a reverse listing that does not exist
counter-evidence the retro should weigh rather than ignore: the same session shipped transition, close --force, six warnings and three refusals, took two work items to completed with eight outcomes proven, and found four real defects in the process of doing it. the question is not whether there was signal — it is the ratio, and whether the churn was avoidable

## ▶ 2026-09-08

created out of WORK-0059 at the end of the session that produced it — the maintainer's observation was that many things learned in the experiment probably are not captured as work yet, which is the whole reason this exists
the reasoning that produced it, copied rather than moved per spec 4.8.1: ten work items came out of WORK-0059 while it ran, but those were the things somebody noticed AT THE TIME, and the ones nobody noticed are exactly the ones still only in prose
two inputs and they answer different questions — WORK-0059's journal is what the session thought was worth keeping while it ran, and evidence/session-capture.md is what actually happened including what the journal missed; reading only the journal inherits its blind spots, which is the whole reason the transcript was kept
WORK-0062 reads the SAME two inputs and asks what the journal lost rather than what we learned — same reading, two outputs, so they should be done in one pass or two thousand lines get read twice
WORK-0051 is the retro SKILL and is captured with an empty body; doing this one by hand first is the right order, because it gives that skill a real case to be written against rather than being designed on spec
WORK-0059's own wrap already gives a verdict on whether editor mode worked, so this is a harvest and not a re-litigation — the boundary matters because the verdict is the interesting reading and would eat the harvest if allowed to
