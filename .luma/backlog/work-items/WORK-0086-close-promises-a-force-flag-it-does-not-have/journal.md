# Journal — Close promises a force flag it does not have

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

BUILT 2026-09-09: close --force exists. it proceeds past every refusal, announces which ones on stderr, writes a FORCED line per override to the journal, and never touches the outcomes — so a forced completed close leaves 0 proven of 1 live, which backlog-move calls the truth
WORK-0086-close-promises-a-force-flag-it-does-not-have captured → unprepared: the defect is real and stated in one sentence: three messages advertise a flag that was never built
outcomes written AFTER the fix, which is backwards and worth saying: they were derived from this record's own problem statement, written on 2026-09-09 before any code, rather than from the code — otherwise this is goalpost fitting with extra steps (WORK-0032)
all ten checks run against a scratch corpus rather than asserted. one needed a case I had not built — a close forced past TWO refusals, to prove it writes two journal lines rather than one summary — and the first attempt only had one refusal to override
the flag audit turned up four names that are not registered flags, and all four are legitimate: --top and --bottom in a comment explaining why they were rejected, --use-hold and --your as example text in journal's -- escape message. a crude grep finds them, and reading them is what tells them apart
NOT COVERED, deliberately: backlog-move says never force without the owner's approval and record who answered. no field can check that — there is no owner and nothing authenticates one — so it stays a rule for the caller rather than a guarantee, and this record does not claim it

## ▶ 2026-09-09

found 2026-09-09 while wiring force-tracking on transition: close's refusal message names --force, backlog-move devotes four paragraphs to how it behaves, and the flag was never built
so the refusal it carries is currently absolute — a completed close over an unmet outcome cannot be done at all, where the design says it can be done and recorded. that is a stricter tool than anybody chose
same shape as WORK-0075: a guarantee table asserting behaviour the binary does not have. this one is a message promising an escape hatch that does not exist, which is worse than silence — it tells the user to type something that fails
