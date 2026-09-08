
❯ /clear

❯ whats next

⏺ Skill(backlog-next)
Successfully loaded skill

  Read 3 files, listed 1 directory, ran 2 shell commands

Now let me gather the actual state.

  Searched for 5 patterns, read 1 file, listed 1 directory, ran 18 shell commands

──[ Overview ]──────────────────────────────────────────────

Preparation considerations

○ WORK-0031 · Reshape the command surface
○ WORK-0053 · The tool reads a config filename this repository does not use
○ WORK-0050 · Review the backlog procedures

luma-backlog work-item list --status prepared

WORK-0031 needs no preparation — it is prepared already, and what it is waiting on is a selection nobody has made. WORK-0053 is a confirmed defect with a ten-minute fix. WORK-0050 has the only written body of the three; the other captures from this morning are empty templates.

To Do (0)

luma-backlog work-item list --status todo

In Progress (0)

luma-backlog work-item list --status in_progress

Risks

One commit on capture/session-findings, no upstream, no pull request. Everything written after 08:38 today exists in this working tree and nowhere else.

The second gate has never been used. No work item has ever been todo or in_progress — and WORK-0031 closed nine tasks while sitting at prepared. This project is the tool's first consumer, so a ladder the project itself routes around is evidence about the ladder.

Six of the seven newest captures — WORK-0051 through WORK-0056 — have empty ## The problem, ## What is being delivered, ## Out of scope and ## Constraints. WORK-0052 is "Quick capture has nowhere to put a long thought", so the defect documented itself five more times in the same burst.

WORK-0031's four outcomes are the only unverified ones in the corpus — 29 of 33 outcomes elsewhere carry verified:. None of the four is provable yet either: each verify_by names a task still todo.

WORK-0053 is real. internal/config/config.go:16 reads config/luma-corpus.yaml; .luma/config/luma-backlog.yaml sits beside it, near-identical and never read.

Last touched

WORK-0031 and the four defect captures, at 09:16 today — commit ddd8a4f, "Capture four findings that were noticed and never written down". WORK-0031's own modified reads 2026-09-06T16:51, a day stale, because closing a task or writing a journal line never touches the parent; the nine closed tasks and the thirty-odd journal lines below it are all newer than the stamp.

It is open, but you did not stop mid-edit — the working tree is clean and the last thing the session did was write down what it had noticed. That is a clean stop with an open record, not an interruption.

──[ Options ]───────────────────────────────────────────────

1. Select WORK-0031 through the second gate — move it to todo, then in_progress; the status is a claim about the present and has been false for a day, and it exercises a gate the tool ships and nobody has used.
2. Fill the six empty capture bodies — the context is a day old rather than a month, and WORK-0053 already proved one of them was worth writing.
3. Fix WORK-0053 — delete luma-backlog.yaml or make the tool read it; two near-identical configs where one is dead will drift, and the live one carries the ladder.
4. Take the next WORK-0031 task — eight remain; resolve-a-reference-the-way-a-person-writes-it unblocks one of the four outcomes on its own.
5. Cross something from the pile at the first gate — 32 captured against 10 unprepared, and seven arrived this morning while nothing moved up.

──[ Recommendation ]────────────────────────────────────────

Push the branch and open the pull request. Every option above adds work; this is the only one that removes risk, and it takes a minute. A day of records — the split policy, seven new work items, thirty journal lines — is currently one rm -rf from gone.

What would change it: if you have already pushed this from another checkout, take 1 — deciding whether WORK-0031 is being worked or waiting determines whether options 4 and 5 are even the same question.

✻ Brewed for 2m 54s

❯ for backlog-next i want risks to be formated in a list as bullets

  Listed 1 directory, ran 1 shell command

Confirmed local/backlog is authored here, not vendored — editing it is correct. Let me follow the bundle update procedure.

⏺ Skill(update-bundle)
Successfully loaded skill

  Read 1 file, ran 1 shell command

Now the procedure, including its worked example.

  Ran 6 shell commands

Now regenerate the harness adapters and audit.

  Ran 1 shell command

⏺ Skill(audit-bundle)
Successfully loaded skill

  Read 1 file, ran 4 shell commands

Done. The change landed in three places and the bundle is at 0.16.0.

What changed

- templates/next-report.md — the Risks slot in the shape block is bullets, plus two new rules: one bullet per concern with the checkable part in the first clause, and plain bullets rather than the state marks (a risk is not a record, so showing-records does not reach it). The "Last touched is prose, not a listing" rule became "the only prose in the overview", which is now true and wasn't going to stay accurate otherwise.
- procedure/backlog-next.md — the same rule stated in the Risks section, and the worked example converted so it demonstrates the shape rather than contradicting it. Its existing list needed a "What to look for:" lead-in, since it is examples of what to notice sitting directly under a rule about formatting.
- BUNDLE.md — 0.16.0 with a version entry. Minor rather than patch: it changes what the report looks like, which is more than a clarification. Stage stays draft, survival unset.

luma-foreman apply regenerated the adapters and both indexes; luma-foreman bundle index . --check passes and a link audit over the bundle is clean.

The Risks section from earlier, re-rendered

▎ Risks
▎
▎ - One commit on capture/session-findings, no upstream, no pull request — everything written after 08:38 today exists in this working tree and nowhere else.
▎ - The second gate has never been used: no work item has ever been todo or in_progress, and WORK-0031 closed nine tasks while sitting at prepared. This project is the tool's first consumer, so a ladder it routes around is evidence about the ladder.
▎ - Six of the seven newest captures — WORK-0051 through WORK-0056 — have empty bodies. WORK-0052 is "Quick capture has nowhere to put a long thought", so the defect documented itself five more times in the same burst.
▎ - WORK-0031's four outcomes are the only unverified ones in a corpus of 33, and none is provable yet — each verify_by names a task still todo.
▎ - internal/config/config.go:16 reads config/luma-corpus.yaml; .luma/config/luma-backlog.yaml sits beside it, near-identical and never read. WORK-0053, confirmed.

One thing I left alone: the commentary under the worked example says "Fifty-seven commits, 9 to 23, never — each one is checkable", but the example it points at says fifty-six. A one-word drift in the paragraph arguing for checkability — worth fixing, but it predates this change and I didn't want to bundle it in silently.

Nothing is committed.

✻ Worked for 2m 56s

❯ i think i want what's next to be fast and punchy,
  but i don't want to lose what the current /backlog-next does
  it's awesome, and i think i will make it more awesome (and slower) over time
  so what should we call the current /backlog-next if we want to reserve the future /backlog-next for something that's more simple and fast

Worth separating two things before picking a word, because one of them decides the other.

Are these two depths of one question, or two questions?

If fast-next is just the Recommendation division printed alone, they're one procedure at two depths and the honest shape is a flag — next and next --full — not two names. But I don't think that's what you're describing. The fast one answers what do I do right now. The current one answers where am I and then offers an opinion at the end. The recommendation is the smallest part of it: everything above the first rule is orientation, and that's the part you said you want to make more awesome.

That matters for naming, because it means the current one should not be named around choosing at all. Any name in the next / pick / up-next family keeps it tangled with the command you're reserving, and the two will fight over every "what's next" the user types.

So: name it for position.

- backlog-orient — a verb, like every other procedure in the bundle (capture, refine, move, verify, journal, show). Names position rather than choice, so it never competes with next for the same request. And it survives the growth you're planning: more risks, reading journals, checking other repos, noticing unpushed work — all of that is orientation, so you won't have to rename it again in three months.
- backlog-standup — the word a person would actually say, and its content maps exactly onto what's already there: what you last touched, what's in flight, what's blocking. Costs you an agile-ceremony association and a daily-cadence implication the thing doesn't have.
- backlog-sitrep — the most accurate genre label. Position, what changed, what's a risk, what you intend. Jargon, and it'd need spelling out once under your own initials rule.
- backlog-survey — neutral and accurate, but bloodless, and it drops the fact that the thing ends with an opinion.

Don't use brief. It's the obvious candidate and it's a trap — it reads as the short one, which is the other command.

I'd take backlog-orient, with backlog-standup as the runner-up if you'd rather have the word somebody says out loud than the one that files correctly.

One cost of the split to decide now rather than discover: skills trigger on their description, and both of these match "what's next", "where were we", "what should I pick up". If the boundary isn't written deliberately into both descriptions, they'll compete — and the failure is silent, you just get the wrong one. Probably: next owns the bare question, orient owns "catch me up", "where do things stand", and session start.

✻ Baked for 53s

❯ i know.  backlog-surface ?  backlog-survey is decent.  give 10 more potions that are like surface, survey, or orient

On surface: it names half the job — bringing up what's below, which is the risks section. It doesn't name the position half (what you last touched, what's in flight), and that's the half a returning maintainer opens it for. Two collisions worth knowing: this harness already ships clarify-surface-findings, and your own prose uses surface generically — the splitting policy says "the signal existed and nothing surfaced it." Evocative, but it's one section wearing the whole report's name.

Ten more, grouped by which part of the job they name.

Where you stand — closest to orient

1. backlog-bearings — get your bearings. Noun, so it breaks the verb convention, but it is exactly what a returning maintainer wants and nothing else in the vocabulary is near it.
2. backlog-situate — places the reader. Precise; slightly stiff to type.
3. backlog-recon — scout the ground before committing. Right genre, punchy, and honest that it precedes action. Military borrow.
4. backlog-map — lays out the terrain. Short, no collisions, no opinion in it — which is a mismatch, since yours ends with one.

Looking over the whole thing — closest to survey

5. backlog-sweep — covers everything and misses nothing; carries the thoroughness you're planning to add. Slight pull toward cleanup.
6. backlog-canvass — go over every part and gather what's there. Unusual enough to be unambiguous, obscure enough to need explaining once.
7. backlog-scan — plain and clean, but it promises fast, which is the other command's job.
8. backlog-examine — flat but accurate, and it implies looking closely rather than quickly.

And forms a judgment — the half survey drops

9. backlog-appraise — look, then say what it's worth. The only one here that carries the recommendation in the word.
10. backlog-reckon — a reckoning: take account of where things stand. Has weight, and it's the one that sounds like the report you actually wrote. Faintly archaic.

Two to rule out, so nobody proposes them later: backlog-status collides head-on with workflow_status, the field this whole model turns on. backlog-overview names one of the report's own three divisions, so the whole can't be called after the part.

If you want the one that survives growth: orient or bearings — both name position, so everything you add later is still the same thing. If you want the one that admits it ends with an opinion: appraise. survey remains a perfectly good safe pick; its only flaw is that it promises a look and delivers a look plus a verdict.

✻ Baked for 40s

❯ yeah backlog sweep is a name i like but i think that will be very expenseive and do EVERYTHIGN
  so scan is pretty good.
  scan might be fast when capared to a sweep.

That resolves my objection — the scale is internal to the family, not absolute. And scan has the second sense working for it: not skim, but examine systematically, the way a scan at a hospital is thorough and a sweep of a building is exhaustive. Same metaphor, one intensity apart, so learning one lets you guess the other.

next sitting outside the pair is right too. It's a question, not a method — which correctly signals it isn't just a third intensity.

The boundary needs to be checkable, or sweep becomes "scan but I tried harder." There's already evidence about where the line falls, from the report I gave you an hour ago: I ran gh pr list and read internal/config/config.go:16 to confirm WORK-0053. Neither is scanning a backlog. So:

- scan — the corpus and git. Records, frontmatter, journals, git log. Everything the backlog knows about itself.
- sweep — leaves the corpus. Reads the code, checks the forge, verifies that a claim in a record is still true, confirms outcomes against reality. Expensive because it goes and looks.

That line is checkable, it explains the cost difference rather than asserting it, and it means the two aren't competing — sweep is a superset by reach, not by effort.

Two things it implies, both fine but worth knowing:

- The current procedure is not quite scan as written. It tells you to check whether work is blocked and whether outcomes can be verified, both of which pull outside the corpus. Renaming it is a rename plus a small subtraction — those bits get marked as sweep's.
- sweep doesn't exist. I'd rename to scan now and record sweep as a reserved name with that boundary written down, rather than shipping a name with nothing behind it.

Still outstanding from before: next and scan will both match "what's next" unless their descriptions draw the line — next owns the bare question, scan owns catch me up, where do things stand, and session start.

Want me to do the rename and write the boundary into the bundle?

✻ Worked for 30s

❯ so scan, sweep, and next are all great
  but i think we can do better

  next should be quick and punchy
  then there should be one thats like an overview but it doesn't look at everything, it tries to ride the middel
  and then sweep should be the thing that goes and makes sure everything is right, all of it

Your tier 3 just changed the axis, and it's a better one. It isn't about reach — it's impression versus guarantee. next gives an answer, the middle gives a read, sweep gives a proof. That supersedes the corpus/outside-the-corpus line I drew; sweep isn't expensive because it goes further afield, it's expensive because it checks.

So the middle name should make partiality the point, not the limitation.

survey — I'd take this one. Its defining property is exactly what you described: a survey samples deliberately to produce a picture, and nobody has ever expected a survey to be exhaustive. It carries "doesn't look at everything" without apologizing for it, which scan doesn't quite — scan sounds like it tried to see everything and went fast. And survey/sweep are the same kind of word at two densities: both cover ground, one samples it, one clears it.

Others that put partiality-by-design in the word:

- rounds — a doctor's round visits each patient, forms a picture, and runs no tests. That is precisely tier 2 against tier 3, and rounds/sweep are both movements through a space. Cost: plural noun, breaks the verb convention.
- sounding — drop a line at points and learn the depth of the whole. The most precise of these about partial by design, still true. Cost: obscure, and backlog-sound is ambiguous.
- pulse — one small measurement that reveals systemic state. Punchy. Cost: noun, and it reads a little dashboard.
- recon — scouts, doesn't confirm; sweep confirms. A genuinely matched pair in the domain it's borrowed from. Cost: two military borrows in a three-name set makes the whole vocabulary read tactical, and next won't play along.
- glance — honest, but it undersells an overview and crowds next.

Two to rule out: sketch fits the meaning well but docs/design/sketches.md already owns that word here, and it owns it with a warning attached. skim describes the same partiality as carelessness rather than design.

next · survey · sweep is the set I'd pick. The middle one is the only word in the family that means incomplete on purpose, which is the thing you're trying to name.

✻ Brewed for 48s

❯ what do get out of a scan, that's what it should be called, scanning is what it's doing, what is the result

A scan's output is a picture of where you are — you know your position and what's around you, and you've checked nothing.

Named for that result:

- bearings — the strongest. Bearings are literally what you get from taking observations, and "get your bearings" is what a maintainer returning after a break is doing. It also covers the whole report rather than half of it: where you are, what's around you, which way you're pointed. Cost: plural noun, breaks the verb convention the other procedures follow.
- readout — what an instrument gives you after a scan. Unusual, unambiguous, no collisions anywhere in this repo. Cost: machine-flavored, and it implies numbers more than judgment.
- picture — plain English, exactly right, slightly soft as a command name.

Three that mean precisely the right thing and can't be used:

- fix — in navigation, a fix is the position you derive from sightings. Perfect word, unusable in software.
- position — same meaning, but this project already uses position for rank ordering, down to big.Rat. Hard collision.
- state — collides with desired_state on every outcome.

One consequence worth seeing before you pick: naming the middle for its result makes sweep the odd one, since sweep names a method. Its result is the all-clear — nothing's wrong, everything checked out. You can either accept the inconsistency (there's no short noun for "a guarantee") or notice that next is also a result, and the set is then two results and one method.

I'd take bearings. It's the only candidate that names the whole thing you get rather than one section of it, and it's the word for what somebody opening this at 9am actually wants.

✻ Crunched for 36s

❯ yeah keep trying what about backlog-report and ...  im have been waiting for things like this to surface

report is legible and nobody has to learn it — but it's the genre, not this tier. next produces a report, sweep produces a report; naming the middle one report claims the whole family for one member. It also says nothing about the partiality that's the point, and templates/next-report.md already holds the word inside the bundle.

More in that vein — plain nouns naming the artifact you get:

- rundown — "give me the rundown." The most natural thing a person actually says for this, and it carries brisk-but-covered-the-ground. No collision anywhere.
- digest — condensed by definition. Names the artifact and the condensation in one word, which is the thing report can't do. Slight pull toward email-newsletter in software contexts.
- briefing — worth resurrecting: I ruled out brief because it reads as the short one, but the noun form doesn't. A briefing is the document, not the adjective, and everyone knows what one contains. Slightly formal.
- dispatch — a report filed from the field by someone who went and looked. Nice, faintly journalistic.
- vitals — the small set of readings that tell you whether the patient is fine, where sweep is the full workup. Punchy, and the pairing with sweep is exact.
- roundup — a gathering of the notable items. Close to rundown, more collection-flavored.
- standing — from where things stand, which is how the procedure already describes itself in its own first line.

Three more collisions to know about, since they're words people reach for: take is reserved — ADR-0008 is about taking a task, and take/release/steal are named there. view collides with record-view and with §11.2's views. snapshot is already how session-checkpoint describes itself in this harness.

Of everything so far, the two doing the most work are rundown and digest — both name the artifact and both say condensed on purpose without apologizing. rundown is the one somebody would say out loud; digest is the one that files more precisely.

I'd put backlog-rundown at the top now, ahead of bearings. bearings names what you get psychologically; rundown names what actually lands on the screen, and it's the only candidate that's both plain and specific.

✻ Baked for 40s

❯ backlog-status

Stronger than I gave it credit for, on two counts I skipped. Structurally it fits exactly — status would be a verb-only top-level command sitting beside board, which is already the shipped shape under §9.1. And it passes the result test better than anything else: what you get from a scan is the status.

The objection is the one this project already ruled on. From the 0.10.0 entry, on why ref became name:

▎ borrowing a word the surrounding system has claimed is the mistake that set aside change for a kind and committed for a rung.

Here the surrounding system is the tool itself. workflow_status is the field the whole model turns on, --status is a shipped flag, and docs/workflow-status.md is normative. In this corpus status already means one record's rung on the ladder — and the command would mean the whole backlog's situation. Those two senses would appear within a few lines of each other in the same help output.

The concrete failure: somebody types luma-backlog status WORK-0031 expecting prepared, because that is what status means everywhere else in the tool.

If you want it anyway, the mitigation is real: status never takes a record argument. No positional, ever — it's the whole backlog or nothing, like board. That narrows the ambiguity to a wrong guess the tool can correct in one error message, which is survivable.

One casualty from checking: standing is out too. §5.2 has standing conditions — an outcome that must hold continuously — so the word is spoken for.

I'd still not take it, on the project's own precedent rather than on taste. rundown remains my pick: it's the plainest word left that nothing else in this corpus has claimed. But status is the most conventional name available and the cost is a single predictable misuse, so it's a defensible call — say the word and I'll write it.

✻ Crunched for 45s

❯ backlog pipelines

This one's your own vocabulary rather than a borrow, which is the opposite of the status problem — docs/workflow-status.md coins the preparation pipeline and the work pipeline and they're normative. That's a real point in its favor.

Two things kill it, both checkable against that document.

captured is not in a pipeline. workflow-status.md draws it in its own box above the first gate — "a pile that may or may not become work." Preparation considerations, which is the section the report ends with, draws mostly from captured. So the name excludes one of the four sections by the project's own model. Same failure as overview: naming the whole after some of its parts.

And it points at the wrong system. Outside this repo, pipeline means CI/CD to essentially everyone, and this repo has a .github/ directory. backlog pipelines reads as show me the build. That's worse than status's ambiguity — status at least resolved to a narrower sense of the same thing, this resolves to a different system entirely.

There's also the thread's own test: it names a structure, not a result. Pipelines are a thing that exists; the report is a thing you get.

But it's a good command that isn't this one. A pipelines view — the two pipelines, their rungs, what's sitting in each, nothing else — is coherent, cheap, and closer to fast-next than to the middle tier. Worth capturing as its own work item rather than spending on this name.

Still rundown, with status as the defensible-but-costly conventional pick.

✻ Worked for 29s

❯ how is status a problem?

Weaker than I made it sound. git status is the counterexample I should have led with: git reports per-file statuses inside a command called status, it's the most-typed command in software, and nobody has ever been confused by it. Two senses of a word at two different levels coexist fine — that's normal, not a defect.

luma-backlog status WORK-0031 also isn't the failure I claimed. git status <path> scopes the report rather than printing a field, so a scoped report is the conventional reading, and a useful one.

What actually survives is narrow: gh faced this exact choice and split the words. It ships gh pr status for the report and --state for the filter. This project shipped --status for the filter — and it's the project that restyled its help output on gh's model, so that's the closest reference point available.

So the real question isn't whether status is a good command name. It is. It's whether --status is the right flag name once status is a command, and the honest answer is that living with both is fine — one is a top-level verb-only command, the other is a filter on list. They appear in different sections of help and nothing forces a reader to hold them at once.

The alternative — renaming --status to --state for gh parity — is a breaking change to a shipped flag, and WORK-0031 already carries a pile of those. Not worth opening for this.

Revised: backlog-status is a good pick. It's the most conventional name available, it fits §9.1 as verb-only beside board, it names the result rather than the method, and its cost is one flag that spells the same word for a narrower thing. I was leaning on the ref precedent harder than it deserves — ref was ambiguous at the same level, a record's identity versus a git object. This isn't.

Say the word and I'll rename it, with sweep recorded as the reserved third tier.

✻ Brewed for 41s

❯ rundown is good too

❯ let's go with backlog rundown, i like that a lot

Doing the rename. First, everything that points at it.

  Searched for 1 pattern, ran 2 shell commands

Now BUNDLE.md — the live entries, the two live wikilinks in history, and a version entry.

  Ran 1 shell command

Remaining hits are code spans in older version entries — factual about what it was called then, and not links. Regenerating.

  Ran 3 shell commands
2 skills available

One live work item still points at the old filenames. Adding a journal line rather than rewriting the record.

  Ran 3 shell commands

Done. /backlog-rundown is live — the skill list above shows it registered, and backlog-show now cross-references it.

What moved

- procedure/backlog-next.md → procedure/backlog-rundown.md, retitled Run down where the work stands. Both [[next-report]] links repointed.
- templates/next-report.md → templates/rundown.md, so the template names the artifact the way listing.md and record-view.md do.
- procedure/backlog-show.md — four cross-references, including the one in its description, which is what keeps the two skills from competing.
- BUNDLE.md at 0.17.0, with the reasoning: the three tiers, the impression/guarantee axis, why rundown beat the method-names, and the list of words that were right and already claimed — status, state, position, standing, view, take.

Two calls I made

rundown keeps every trigger for now, including the bare "what's next". Handing that phrase to a command nobody has built would break the most common way in just to reserve a name. The version entry records that the boundary gets drawn when next ships, so nobody later reads it as an oversight.

Marked breaking but shipped minor — two documents renamed. Nothing outside this project adopts local/backlog, and every inbound link is repointed in the same commit, which is what §1 of update-bundle asks for.

One thing I didn't rewrite. WORK-0050's body has a table naming backlog-next and next-report with their line and commit counts — a measurement taken at a time, and rewriting records to match a later rename is exactly what WORK-0022 exists to decide. I added a journal line on WORK-0050 instead, saying what the files are called now and that the counts still hold.

Link audit over the bundle is clean, luma-foreman bundle index . --check passes, and apply removed the old adapter. Nothing committed.

Three things surfaced in this conversation that aren't written down anywhere yet: fast next, sweep, and a pipelines view. Want those captured?

✻ Worked for 2m 21s

❯ yeah i want next, sweep, and pipelines to be captured somewhere where they won't get forgotten or lost, but in a way that says these are reserved, potential, but not certain

⏺ Skill(backlog-capture)
Successfully loaded skill

  Read 1 file

Thoughtful capture, so the corpus check comes first.

  Ran 2 shell commands

❯ sweep might turn into doctor or cleanup

❯ pipelines might become unnecessary

Both of those land in what gets written — doctor/cleanup is a real fork for sweep, and the corpus agrees with you about pipelines. Here's what exists before I write anything.

Conflicts

- WORK-0046 · Evaluate the conditions the tool names (captured)
  This is sweep's diagnostic half, already specified. It builds check — and check is already a top-level verb-only command in the shipped surface (ADR-0006, WORK-0031's table), alongside init, board, config. So sweep is not a free name in the way rundown was: there is a spec'd command sitting in that space. Your doctor instinct is the same instinct — doctor and check are the same command in different vocabularies, and one of them is written down.

Overlaps

- WORK-0002 · Lint the corpus (captured)
  Records against their type definitions, wikilinks resolving, cross-references. The other half of "make sure everything is right", and its own text already uses the word — "a sweep found one with no meaning attached to its absence."
- WORK-0043 · Work the backlog from a terminal (captured)
  Delivers --column, sorting and grouping. .luma/config already defines columns that map one-to-one onto the ladder, so this is pipelines — you were right that it may become unnecessary, and it's closer to "already captured under another name" than to "might not be wanted".
- WORK-0016 · Ask the backlog what is open (captured)
  The predicate a fast next needs. The seam: WORK-0016 makes everything except closed askable; next decides what to do with the answer.

What I'd write — two records, not three

┌───────┬─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│       │                                                                                                                                     │
├───────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Title │ A fast next, without the rundown                                                                                                    │
├───────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Kind  │ idea                                                                                                                                │
├───────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Body  │ The pick and nothing else — no listings, no risks, no last-touched. Reserved at local/backlog 0.17.0 when backlog-next became       │
│       │ backlog-rundown. rundown holds the bare "what's next" trigger until this exists.                                                    │
└───────┴─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

┌───────┬─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│       │                                                                                                                                     │
├───────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Title │ An exhaustive sweep that verifies rather than reports                                                                               │
├───────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Kind  │ idea                                                                                                                                │
├───────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│       │ The tier above rundown: goes and confirms the records are true rather than reading what they claim. Name and shape both open — may  │
│ Body  │ become doctor, may become cleanup, may turn out to be WORK-0046's check grown a tier, may split into diagnose-and-repair. Recorded  │
│       │ because the tier is wanted, not because the command is.                                                                             │
└───────┴─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

idea on both, deliberately — it is the kind for a thought nobody can judge yet, and paired with captured ("might become work; nobody has decided") the ladder already says reserved, potential, not certain without inventing a mechanism for it.

For pipelines I'd add a line to WORK-0043 rather than a third record — it is that work item's --column under a different name, and a separate record would put a seam where work falls through. A note there saying the framing was considered and is absorbed keeps it findable without pretending it's distinct.

Two records and a note on WORK-0043 — or do you want pipelines standing on its own anyway?

✻ Churned for 1m 26s

❯ next will definitley exist but im not sure in what capacity
  other then i want it to be high quality but fastish
  so it can't do lots of expensive considerations

will definitely exist, shape unknown is a rung, not a kind — unprepared is defined as "will become work, and nothing has been worked out yet." I'll capture at captured as the procedure requires and flag the gate separately, since crossing it is your call and not a courtesy I extend on the way out.

  Read 1 file, ran 3 shell commands

Now the bodies — the empty-template problem I flagged earlier.

  Ran 2 shell commands

Now the pipelines note on WORK-0043, as a journal line rather than a body edit.

  Ran 2 shell commands

Captured.

○ WORK-0057 · A fast next, without the rundown — kind change
○ WORK-0058 · An exhaustive sweep that verifies rather than reports — kind idea

luma-backlog work-item list --status captured

Both landed at captured, which is the rung that already means what you asked for — written down so it is not lost; might become work, nobody has decided. No new mechanism needed.

Both have written bodies, not empty templates. That was the risk I flagged this morning — six of the seven newest captures have blank ## The problem — and writing two more of those while naming it as a risk would have been absurd. Each carries the evidence, the open questions, and links to the records it sits beside.

What went in that you said today: WORK-0057 records will definitely exist, capacity unknown, and the fast-but-high-quality constraint — with the observation that the real open question is whether fast means reading less or reasoning less, since those are different trades and nothing decides it yet. WORK-0058 records doctor and cleanup as live candidates and notes they aren't the same command — doctor diagnoses, cleanup repairs — so if both are wanted that's two work items, and the reciprocal test decides.

pipelines is a journal line on WORK-0043, not a record. Its --column work plus grouping is that view — the configured columns map one-to-one onto the ladder's rungs. The line carries a re-open trigger rather than a rejection: revisit if the two pipelines need something columns can't express, such as the gates between them. Cheap to promote to a record if you'd rather it stood alone.

I also journaled on WORK-0046 that WORK-0058 may be its check grown a tier rather than a new command — that's the one finding from the corpus check worth acting on, since check is already in the shipped surface and sweep isn't a free name the way rundown was.

One thing left for you. "It will definitely exist, nothing is worked out" is the definition of unprepared, not captured — that's the first gate, and you've effectively already made the call. I didn't move it, because the capture procedure is explicit that crossing a gate is a person's decision and not a courtesy extended on the way out. Say the word and WORK-0057 goes to unprepared.

Nothing committed.

✻ Churned for 1m 44s

❯ move WORK-0057 to unprepared

⏺ Skill(backlog-move)
Successfully loaded skill

  Read 1 file, ran 2 shell commands

Moved.

○ WORK-0057 · A fast next, without the rundown — captured → unprepared

luma-backlog work-item list --status unprepared

That's the first gate crossed, on your call that it will definitely exist. What it commits you to is working out what it is — not doing it; that's the second gate, and it's still above WORK-0057. I journaled the reasoning on the record first, so a future session finds why it crossed rather than just that it did.

One thing to look at. It got rank: 020.0010.000 — ordinal 020 for unprepared, then first position. It re-enqueued at the front of eleven items, not the back, and the procedure is explicit that a status change re-enqueues at the back because "a position in a queue does not travel with you when you leave the queue."

Two readings, and I can't tell which from here: the other ten may never have been ranked, in which case this is the documented behavior — the journal on WORK-0031 says unranked records sort last and ranked ones lead, so the first record ever ranked at a status lands ahead of them. Or the re-enqueue is computing the wrong end. The listing above is consistent with the first, since the other ten are in key order rather than a chosen one.

Worth ten minutes to confirm, and it's the kind of thing WORK-0058 exists to catch. Want me to check?

✻ Worked for 26s

❯ i want to work on backlog move, do we have a work item for that

⏺ Skill(backlog-show)
Successfully loaded skill

  Searched for 1 pattern, ran 1 shell command

One record names it, and it explicitly excludes doing it.

○ WORK-0050 · Review the backlog procedures (captured)

grep -rl 'backlog-move' .luma/backlog/work-items/

That's a grep rather than a command — there is no search over record contents, which is a gap rather than a limit and is the second time today it would have helped.

WORK-0050 puts backlog-move first in what is actually unreviewed:

▎ backlog-move — the longest judgment in the set. It carries both selection gates, which is where the whole ladder model lives, and nobody has read it since it was written.

Written in one pass, two commits, never reopened.

But it stops short of the work you're describing. Its Out of scope is explicit:

▎ Rewriting them. This decides what a good one is and which fall short; the rewriting is whatever work items come out of it.

So WORK-0050 is the inquiry, and working on backlog-move is one of the work items it would produce. Nothing covers that yet.

Its premise moved today, which matters for how you start. WORK-0050 argues that has been run is probably the whole definition of reviewed, and records that only backlog-show and backlog-next had ever been run against the real corpus. This session ran backlog-capture, backlog-journal and backlog-move for the first time — and backlog-move immediately produced a concrete finding, the rank landing at the front of eleven records rather than the back. That's exactly the evidence WORK-0050 says nobody has.

Three ways in, and they're different jobs:

1. Capture a work item for the backlog-move work directly — you know what you want to change, and WORK-0050 catches up later.
2. Work it under WORK-0050 — settle what reviewed means first, since backlog-move is its top candidate and the definition is the thing that record says is missing.
3. Chase the rank finding — narrower than either, and it's a real defect or a real documentation gap sitting in the procedure you want to work on.

What's driving it — is there something specific in backlog-move you want to change, or is this the review?

✻ Baked for 40s

❯ i want to create a work item and do an experiment, so we should use the journal heavily so we can learn as we go

  Listed 1 directory, ran 2 shell commands

exploration is a built command, and §7.2.1 settles the record type: an investigation lives in explorations/ inside a work item, "or as a work item in its own right when the investigation is the work." That's this. So a work item is correct, and explorations can hang off it later if a question gets big enough to need its own file.

Corpus check — one relation, no duplicates.

Overlaps
- WORK-0050 · Review the backlog procedures (captured)
  Names backlog-move first in what is unreviewed, and needs a definition of a good procedure it doesn't have. The seam: WORK-0050 decides what good means across all seven; this one goes and gets the evidence from one of them. Its Out of scope hands this over explicitly — "the rewriting is whatever work items come out of it."

What I'd write

┌─────────────┬───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│             │                                                                                                                               │
├─────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Title       │ Find out what backlog-move needs by running it                                                                                │
├─────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Kind        │ inquiry — the tool stores spike and investigation as this                                                                     │
├─────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Description │ an experiment: use backlog-move for every status change and journal every friction point as it happens, rather than reading   │
│             │ it and forming an opinion                                                                                                     │
└─────────────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

The experiment, and why it's shaped this way. WORK-0050 argues that has been run is probably the whole definition of reviewed, and backlog-move had never been run against the real corpus until an hour ago. It ran once today and immediately produced a finding — the rank landing at the front of eleven records. That is one data point from one invocation, which is the argument for doing this deliberately rather than hoping.

CLAUDE.md names three things that promote prose to a command, and two are reachable here: friction where it appears, which is what heavy journaling captures, and divergence where we provoke it — same instruction, two agents, diff the records. I'd propose the first arm now and hold divergence as a second arm once the prose's weak points are known, since provoking divergence blind is expensive.

One caution on journaling heavily, because it's the failure mode of exactly this plan. spec.md is explicit: "Has an obvious home — a task, an outcome, a decision, an exploration? Put it there. Otherwise the journal, now." The journal is for learnings in flight; a conclusion belongs in a decision record and an investigation in an exploration. Journaling everything is a way to accumulate material without ever settling anything, and the record should say so as a constraint rather than discover it in three weeks.

Write it as above? And how far up the ladder — it lands at captured, and doing it now means both gates, the second of which this project has never used.

✻ Crunched for 1m 9s

❯ the experiment is how should we handle ad hoc work
  where i know i want to:
  - act fast
  - i have a general idea of what i want, i understand a vague direction
  - i want to plow towards a solution instead of jumping through hoops, i want to take shortcuts - the experiment should help answer what are shortcuts you can opt for and which are always a bad idea
  - i want to figure out what the work is as i go along.
  - this is basically a spike or exploration that doesn't product work items, it produces a better version of what i laready have.
  - and walking through it or using it is how i want to fix it,  not record my observations to another work item and do that work item, which is hte more proper approach - and generally how it should be done with all things equal

That reframes it — backlog-move is the vehicle, not the subject. The question is when it's legitimate to fix the thing in flight instead of routing through capture → refine → select. Checking the corpus against that.

  Ran 1 shell command

Overlaps

- WORK-0029 · Separate quick capture from thoughtful capture (unprepared)
  The same question, one level down and already answered once — for capture only. Its principle is the one you're generalizing: "Quick capture is measured by latency, not by quality. A quick capture that produces a better record at the cost of two extra turns has failed. It is a different procedure with a different success condition — not the same one with the reading skipped." Seam: WORK-0029 gives one procedure a fast mode; this asks when work skips the ladder entirely.
  Worth knowing — it looks largely delivered. backlog-capture now has both modes with the quick rules written out, and nobody closed it.

Nothing duplicates. Nothing conflicts.

What I'd write

┌─────────────┬───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│             │                                                                                                                               │
├─────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Title       │ How ad hoc work should be done                                                                                                │
├─────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Kind        │ inquiry                                                                                                                       │
├─────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│             │ act fast, vague direction, plow toward a solution taking shortcuts, work out what the work is as you go — a spike that        │
│ Description │ produces a better version of what already exists rather than work items. Using the thing is how it gets fixed, not recording  │
│             │ observations into another work item and doing that. backlog-move is the vehicle.                                              │
├─────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Delivers    │ which shortcuts are safe, which are always a bad idea, and a better backlog-move                                              │
└─────────────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

---

Below is mine, kept separate.

The project already works this way and has never written down when it's allowed. Everything done today — the rename, the version bump, the Risks change, the cross-reference repointing — was ad hoc. No work item, no outcomes, no gate. It produced a real result and there is no record of the policy that permitted it, because there isn't one.

Which means one of this morning's risks was probably misread. I flagged WORK-0031 closing nine tasks while sitting at prepared as a status that had stopped being true. Under your framing that may not be a defect at all — it may be ad hoc work behaving correctly against a ladder with no rung for it. Worth resolving inside the experiment rather than assuming either reading.

It's self-demonstrating, which is rare and worth exploiting. This is ad hoc work about ad hoc work, so the journal is the instrument and the specimen at once. Nothing else in the corpus has that property.

One methodological point, and it's the only rigor I'd argue for. Journal each shortcut as it's taken — what it skipped, before you know whether it worked. Otherwise every shortcut gets written up as justified in hindsight and the experiment produces a story rather than evidence. That's a rule about when the line is written, not how many, so it costs nothing against acting fast.

And one hoop I'd keep. This needs a single bounding outcome, because improve backlog-move is the walk-forever kind that when-a-work-item-splits says absorbs tasks until nothing closes. Something like: a written answer exists naming which shortcuts are safe and which never are. One outcome, one sentence — the rest of refinement can go.

Small mismatch worth noting: the tool defines inquiry as "understanding, and the work items that follow", and you've said explicitly this one produces no work items. The kind still fits; its gloss doesn't.

Write it as above, with that one outcome?

✻ Cogitated for 1m 16s

❯ partly about routing through capture refeind select
  and partly about makign outcomes and tasks an afterthought rather than something i need ot flesh out properly, because doing that would probably product a worst result, in this case, because i do not know what i want until i see it

  Ran 2 shell commands

WORK-0032 already ranks exactly what you're describing, and you wrote it:

▎ worst to best: 1. Writing no outcomes. 2. Writing outcomes after the fact, and still proving them. 3. Writing outcomes ahead of execution. All three are allowed.

So the mode is permitted, with one condition attached — and still proving them. That's a better bound than the hoop I proposed, and it's yours. Dropping mine. Writing the record rather than proposing a fourth time, since another round of agreement is the thing being tested.

  Searched for 1 pattern, ran 3 shell commands

○ WORK-0059 · How ad hoc work should be done — captured, kind inquiry

The record carries both halves as you framed them: the gates, and outcomes-and-tasks-as-afterthought — with the stronger claim stated as the claim, that up-front specification would produce a worse result here, not just a slower one, because what's wanted isn't known until it's seen.

I dropped the bounding outcome I argued for. WORK-0032 already ranks writing outcomes after the fact, and still proving them second of three and allows it explicitly. That's a better bound than mine and it's already decided, so inventing a hoop on top of it would have been the behavior under test. The condition rides along instead as a constraint: written late is fine, written late and never proven is goalpost fitting — and close … delivered already refuses without outcomes, so the tool enforces it rather than prose asking.

Two shortcuts are journaled as taken, not in hindsight — writing the record instead of a fourth agree-cycle, and shipping with no outcomes. That's the one discipline the experiment needs: the line goes down before anyone knows whether the shortcut worked.

Three things went in that you'd want on the record. The project has been working this way all day — the rename, the version bump, the Risks reformat, all ad hoc with no work item and a real result, and no policy anywhere permitting it. This morning's WORK-0031 risk is flagged as probably misread: work happening at prepared may be ad hoc working correctly rather than a status gone stale. And WORK-0029 looks delivered and unclosed — it's the same question already answered for capture alone.

Status is captured. Doing it means both gates, and the second has never been used in this project — which makes WORK-0059 either the first thing through it, or the first thing to prove it shouldn't have to be. That's the experiment's first real question and I'd rather you answer it than have me pick.

Nothing committed. Where do you want to start on backlog-move — the rank finding, or open the file and read it?

✻ Crunched for 1m 43s

❯ this is like editor mode,
  i am editing a book, and coming up with all of the outcomes is impossible because i dont know what i need to edit until i see it.
  how do we handle this mode, what do we call it, how do we pivit our guidelines, etc etc etc, when is this acceptable, when should it be discouraged, it is dangerous if overused or used in the wrong way, or discided as editor mode when really im just trying to break free and make a mess because i am impatient or trying to undermine the system.

  i want to journal each and every interaction we have during this experiement

The editor analogy resolves something. An editor can't enumerate the edits before reading — but they always have a brief: cut it to 200 pages, make the tense consistent, this has to read for a general audience. The brief is a criterion, not an enumeration. You don't know which sentences you'll cut; you know you're cutting for length.

So editor mode isn't outcomes-absent. It's outcomes-as-criteria. And that gives your danger case a test.

  Ran 2 shell commands

Now the record — the questions become scope, not answers.

  Ran 1 shell command

In the record: an Editor mode section carrying your analogy and your danger, and a What it has to answer section holding your five questions as questions — name, when acceptable, when discouraged, how the guidelines pivot, how it's told apart from impatience. Not answers. Writing answers now is the up-front specification the mode exists to avoid.

The sharpest line is yours and it's the framing for the whole thing: this has to produce a test somebody can fail, not a permission somebody can claim. A mode you can declare is a mode you can abuse; a test you can fail isn't.

Journal-every-interaction is now a constraint on the record, with the reason: filtering for what seems worth keeping would apply the judgment under test to the evidence for it. Marked as an instrument log, deliberately denser than spec.md §5.5 asks for, so nobody later reads this journal as the house style. Six lines in so far, three of them this turn.

Three hypotheses are journaled as hypotheses, so they're falsifiable rather than decisions I slipped in:

- Editor mode is outcomes-as-criterion, not outcomes-absent. An editor can't enumerate the edits but always has a brief — cut to 200 pages, fix the tense. The criterion is statable before starting even when the edits aren't.
- Three tests separate it from thrashing: can you state the criterion, is the substrate reversible, are you editing something that exists rather than writing something new. Failing any one is the tell.
- Reversibility is doing quiet work in why this feels safe. Editor mode over git-tracked prose is nearly free to undo. The same mode over a published bundle, a corpus migration, or anything that leaves the repository is a different risk class, and the answer probably shouldn't generalize across it.

The second test is the one I'd expect to break first — "editing something that exists" is clean for backlog-move and gets murky the moment an edit turns into a rewrite.

Ready to open backlog-move. Do you have a criterion for it, or is finding out what the criterion should be part of what we're testing?

✻ Cooked for 1m 14s

❯ ok good analagy, capture it,
  and i am not editing for lenght, what am i editing for
  all i know is i dont like it and i want to make it better
  and asking me how i want to make it better will get in the way of me making it better,
  what do we do with that
  do we force the user to pause and think
  or do we allow them to move forward and explore,
  i think our system has to be flexible but also discourage bad behavior
  our system also needs to be able to recognize there are different modes/strategies/approaches and what approach we are using
  it needs to be able to name it when it recognizes it or when the user requests it
  it should help us evaluate if we are abusing it
  and it needs to be able to apply different guidelines based on the situation

That falsifies my hypothesis, and the falsification is the useful part — you don't have a brief, you have dissatisfaction. Journaling it as a miss rather than quietly revising it.

  Ran 2 shell commands

Now reading backlog-move properly, since a diagnosis is what an editor produces before edits — and asking you for one first is the move that just failed.

  Ran 4 shell commands

Captured, and the record now carries both things you sharpened.

On the brief: dissatisfaction is not the absence of a standard — it's a standard that isn't articulable yet. Knowing something is wrong before you can say why is the ordinary condition of editing, and a system demanding the why first forbids the reading that would produce it. So the test moved: not whether the criterion precedes the work, but whether it arrives. Finish the edit and still not be able to say what it was for — that's thrashing, diagnosed after, which is the only time it's diagnosable.

On force-a-pause versus let-them-move: ruled out the gate, since anything that stops the work to extract a specification destroys what it was protecting. The system already owns the right idiom — observed, never refused — what list does with skips and duplicates and what WORK-0046 specifies for conditions. Never block; name it, and report when it looks like abuse.

Your mode-recognition requirement is bigger than this work item — recognize the mode, name it on request or on sight, judge whether it's being abused, apply guidelines per mode. That's a mechanism with its own definition of done. Not splitting: WORK-0059 has no outcomes yet, so there's nothing to split against. Recorded with a trigger — split when its outcomes get written and the mode machinery needs a different one.

And the experiment caught me. I ended last turn asking whether you had a criterion for backlog-move. That question was the hoop, one turn after I described it. Journaled as evidence rather than quietly dropped.

So I read the file instead of asking. Diagnosis, worst first:

- It teaches a vocabulary a decision in force replaced. ADR-0007 is provisional — in force — and renames delivered to completed, adds rejected, and explicitly drops abandoned. backlog-move still teaches delivered/abandoned and spends a paragraph arguing the canceled vs abandoned distinction ADR-0007 deleted. It's pinned to the shipped binary while contradicting the decision that supersedes it.
- Line 107 says a move re-enqueues at the back. WORK-0057 landed at 020.0010.000 — the front of eleven. Doc and behavior disagree, and only running it showed that.
- It says journal before closing, never before crossing a gate — though the gate is the decision the whole ladder turns on. WORK-0057's crossing got journaled on instinct, not because the procedure asked.
- Nothing says to verify a move landed, which is why the rank anomaly survived until I happened to look.
- Sending back and reopening have no command form, and reopen may not be built — WORK-0020 is still captured.

29 journal lines. Next I start editing, beginning with the disposition table, since teaching a superseded vocabulary is the one actively doing damage.

✻ Brewed for 2m 4s

❯ i think different organizations and projects will want different strategies to be tolerated, so we need a menu of strategies or modes and orgs/proejcts should be able to configure what they tolerate and when/how/who can use them

  ok so that's enough preflight analysis, let's jump in...

  later on we should circle back and figure out how modes/strategies/whatever work in the grander scheme of things.
  for now let's focus on walking through the experiment and then later we can ponder once we have more real world experience.

  i want to walk through the backlog-move skill and make it better
  when i say that what should happen first
  where should this half baked work item go, what checks and balances should get put in place here
  this is something we should capture and answer later

  let's take a first best guess for now
  i think that hsould move the work into either preparing or directly into in progress
  it's hard to say since the person is going straight towards doing stuff,
  is it in progress because we are changing things already or is it preparing because we have no definition of done and we've skipped past all the considerations of who should be involved, and basically all preflight checks

  so now im editing reading as an editor.  what should have hapened already?
  i may ask you to make changes or i might just make them myself and then tell you to look, isn't this fun?  :-P
Interrupted · What should Claude do instead?

❯ i think different organizations and projects will want different strategies to be tolerated, so we need a menu of strategies or modes and orgs/proejcts should be able to configure what they tolerate and when/how/who can use them

  ok so that's enough preflight analysis, let's jump in...

  later on we should circle back and figure out how modes/strategies/whatever work in the grander scheme of things.
  for now let's focus on walking through the experiment and then later we can ponder once we have more real world experience.

  i want to walk through the backlog-move skill and make it better
  when i say that what should happen first
  where should this half baked work item go, what checks and balances should get put in place here
  this is something we should capture and answer later

  let's take a first best guess for now
  i think that hsould move the work into either preparing or directly into in progress
  it's hard to say since the person is going straight towards doing stuff,
  is it in progress because we are changing things already or is it preparing because we have no definition of done and we've skipped past all the considerations of who should be involved, and basically all preflight checks

  so now im editing reading as an editor.  what should have hapened already?
  i may ask you to make changes or i might just make them myself and then tell you to look

Jumping in. Capturing the modes menu, journaling the rest, and moving the record — then the answer.

  Ran 2 shell commands

○ WORK-0060 · A configurable menu of working modes — captured
◐ WORK-0059 · How ad hoc work should be done — captured → in_progress

in_progress, not preparing — and the reason is the answer to your ambiguity. in_progress is a claim about the present, and the present is that a live procedure is about to change. Sitting at preparing while editing production files would be the record lying in the other direction.

The pull toward preparing is real but it's the wrong instinct: preparing is not a penalty box. Dropping to a lower rung to account for skipped preflight uses the ladder to express guilt instead of position, and then the rung means something other than where the work is.

Both gates crossed in one move — and that's the first time this project has ever used the second gate. It used it by skipping it, which is better evidence than a clean crossing would have been.

What should have happened already — and mostly did.

- A baseline. Git has it. That's the reversibility test passing, and it's why editing is cheap here in a way a migration wouldn't be.
- Read the thing before proposing changes. Done — and it produced the ADR-0007 finding, which asking you wouldn't have.
- Know what already governs it. ADR-0007 is provisional, in force, and contradicts the file. Found before touching it, which is the one check that would have been expensive to skip.
- The criterion recorded as absent, so its arrival is checkable later. In the journal.
- The status honest. This is the only one that hadn't happened. Now it has.

Two things that should have happened and can't. Nobody owns it — ADR-0008 shipped no take, so ownership isn't expressible. And nothing recorded who should have been involved; on a one-maintainer project that's invisible, which is exactly the failure mode CLAUDE.md warns about — multi-actor failures produce no friction here, so they won't be found by working this way.

35 journal lines. Opening the disposition table now — delivered/abandoned against ADR-0007's completed/rejected. Make changes yourself whenever you'd rather; tell me to look and I'll read the diff.

✻ Crunched for 58s

❯ ok but im basically going rouge,
  so to me it is acceptable and ideal,
  and to a governing body or observer (compliance, leadership, security, etc) i might be causing problems
  so we need to allow this but also surface it to observers so they can intervene, only when it's an actual problem
  we don't have a way to do that yet, but we still need to record it all the same

Right — and the gap is precise: the journal records it, nothing indexes it. An observer doesn't know to open WORK-0059.

  Ran 2 shell commands

Recorded three ways, since the mechanism doesn't exist yet.

In the record, where an observer looks first — a Declared: this work went around the system section, above everything else. It's written for someone who wasn't here, and it doesn't defend the choice: both gates skipped in one move, no outcomes so nothing can say this is done or drifted, no review and no second actor, the corpus check skipped on WORK-0060, a convention invented with no decision behind it. It says plainly that nobody has reviewed any of it — same category as WORK-0003, which already records the Go merged unreviewed rather than leaving it implicit.

In the journal, with a grep-able prefix. Five lines now open with the literal string shortcut taken, and unproven claims with hypothesis, untested. That's the interim convention — no mechanism needed, a future check can find them, and WORK-0037 covers migrating the shape once there's a real one. I journaled that inventing it is itself a shortcut: a settled convention is supposed to be a decision record here, and nothing outside WORK-0059 knows this one exists.

○ WORK-0061 · Surface work that went around the system to observers — captured

The framing that matters is yours and it's the hard half: intervene only when it's an actual problem. A mechanism that flags every skipped gate is a linter everyone silences, and then the rogue work is genuinely invisible rather than merely unindexed. That's the same trap observed, never refused exists to avoid — the difficulty isn't detecting the shortcuts, it's ranking them, and detection with no ranking is worse than nothing because it looks like coverage.

The honest limit on all of this: one maintainer and one agent means the multi-actor checks produce no friction, so working this way can't discover them. That's in the declaration too — CLAUDE.md already says friction alone won't find them here, and this experiment can't be the thing that does.

Back to the disposition table.

✻ Baked for 59s

❯ yes so we should let the user know they are breaking protocol
  if they say yeah i know lets keep going
  we should have "modes" that we look for that give them a path forward
  if we can't find one then we put the brakes on
  and we start capturing stuff so it's easy to surface later that protocol was broken, and during a work item retro we can determine if that was good or bad, if our policy was wrong or needs new considerations/conditions or special circumstances for edge cases - and how to handle those without creating too much context bloat
  if we do find a mode, we work with the user to pick it, so they are informed what is happening, and we record it, so we can learn later if that was a good decision and how we are working (it is good to know what modes people work in, that is a learning in and of itself)
  disclaimer, mode is the term im using now, it's not decided
  questions?

Six, and the first is the one I'd want answered before anything gets written.

  Ran 1 shell command

1. The brake is strongest exactly when the menu is emptiest. Day one, almost nothing matches, so almost everything brakes — which teaches people to route around the system entirely and produces the invisible rogue work this exists to catch. Does there need to be an unnamed mode that always matches and records itself as unnamed? That keeps the work visible at the cost of the brake almost never firing.

2. Who detects the breach? If it's an agent noticing, that's prose compliance — and CLAUDE.md says measured compliance with prose-only rules runs far below a guarantee. This experiment already produced the evidence: I asked the hoop question one turn after describing it. If detection is the tool's job it's a check condition and lands in WORK-0046's territory, which is a very different build.

3. Consent comes before the mode is picked. Warn → "yeah I know" → find mode → pick it together. That's two interactions in a mode whose whole point is speed. Can it be one? "You're skipping the second gate — that looks like editor mode. Take it, or something else?" Same information, one turn.

4. What does "protocol" cover? Gates and missing outcomes a tool can check. Discuss before writing it cannot. If the menu covers only the checkable half it's narrow; if it covers prose rules it can't be enforced, only noticed. Which half is this for?

5. "Without too much context bloat" is the binding constraint and I think it forces the architecture. A system that learns by accumulating conditions pays for every rule in every session, forever. The escape is that a mode's rules load only while that mode is active — modes as lazily loaded rule sets, not additions to the base guidelines. Is that the intent? It's the only version I can see that survives ten retros.

6. What does a mode attach to? A work item already carries kind; mode is a different axis. Working WORK-0059 in editor mode for two hours and normally afterward is real — which suggests mode is per-interaction, not a field on the record. That would explain why the journal has been the right instrument all along rather than a workaround.

And one dependency nobody's flagged: the retro is load-bearing here and doesn't exist. WORK-0051 is captured with an empty body. Everything in your step 5 — was the call good, was the policy wrong, what edge cases need conditions — happens in a mechanism that hasn't been designed. If the retro stays hypothetical, modes get recorded and never evaluated, which is the accumulate-without-learning failure.

Noted that mode is provisional. All eight journaled, 43 lines now.

✻ Worked for 1m 2s

❯ all great questions, capture them.  now let's proceed

  Searched for 1 pattern, ran 2 shell commands

Correcting my own finding first — the ten other unprepared records carry no rank at all, so WORK-0057 landed at the back of one ranked record. The code and the doc agree; the surprise was mine.

  Ran 2 shell commands

Captured — the protocol and all eight questions are now WORK-0060's body, which also fixes its empty template. It holds your six steps verbatim before they get argued with, the questions as questions, and the retro flagged as load-bearing-and-nonexistent.

First real edit made. backlog-move's disposition section now teaches ADR-0007's vocabulary: completed · canceled · rejected · superseded. delivered and abandoned are gone.

The paragraph I replaced was arguing for a distinction ADR-0007 deleted — canceled vs abandoned. What survives is the distinction it actually makes: cancelled is we wanted this once, rejected is we never did — and only a person can say which, which is exactly why rejected earns a slot and abandoned doesn't. ADR-0007's test is the good line and the procedure now carries it: the enum carries what the record cannot.

I left the binary's lag visible rather than papering over it. A blockquote says --reason delivered|abandoned still ships, names the todo task that fixes it, and says to use the settled vocabulary and translate at the command line. The reasoning is the one CLAUDE.md already states — a decision in force outranks the implementation — and teaching the shipped spelling is how a dead vocabulary survives in people's heads.

And my rank finding was wrong. Corrected in the journal rather than quietly dropped. But the smaller version is real and now in the file: the back is not always the bottom of the listing — unranked records sort after ranked ones, so arriving where nothing is ranked reads as first. It looks like a bug the first time; it took reading internal/app/status.go to find out it isn't. That paragraph exists because I got it wrong, which is the cheapest kind of documentation.

Journaled as a shortcut too: I reported a defect from one observation without reading the code. Moving fast put the error in the diagnosis rather than the edit — the cheaper place for it to land, but worth knowing that's where this mode's mistakes go.

Still open in the file: the top-of-file invocation still shows --reason, gates have no journaling step, no verification that a move landed, and Sending work back describes a reopen the tool may refuse. Those next.

✻ Crunched for 1m 48s

❯ EDITOR:
  im looking at

  Move a work item along the workflow

  luma-backlog set <ref> workflow_status=<status>
  luma-backlog work-item close <ref> --reason <disposition>

  to me this could confuse agents and people
  as a maintainer i understand that when closing those things must always happen together (which by the way we should fix in the CLI so these are a transaction inside the CLI) but anyways, as a reader i might assume that this applies to all moves, not just closing.

  this is a very good thing to say:
  Status and rank are written together, always. You never write rank yourself.

  but it's so randomly placed in the document, i wonder if there needs to be rules section or something. whatever agents like these days.  guardrails or whatever is most effective for today's models.

  ---
  let's start there.
  so there are two things moving separately right now
  1. my request - all the stuff i typed above is how i would work throught the document, let's actually do that.  let's name that level "EDITOR:"
  2. separately - we are evaluating this experiement and talking at a meta process level.  let's name that level "META:"
  ---

  META: should we be doing an "exploration", which is a major feature i haven't had time to define yet, should we be creating tasks and outcomes now, or should we just journal as we go and forget about tasks and outcomes.  maybe we each turn should turn into a task,  but then we have 100s of tasks and some of them are half baked thoughts that will constantly correct and conflict as i flip flop on what i want.

Two channels, both answered.

EDITOR: you're right that the block over-promises — the second line only applies to closing, and Closing already shows it again fifty lines later. And the rank rule being stated at the top and in What this never does is the same rule in two places, which is the real defect.

  Ran 1 shell command

EDITOR — what changed. Top block is one command now, with a line saying it's every move except closing, so the reader isn't left inferring scope. Always true sits right after the H1: three invariants, each with the failure it prevents, and an explicit note that none of them is advice. What this never does is deleted — everything in it either moved up or was already said.

On format: I went with a named section near the top, bold leads, failure named per rule. That's what survives in this bundle's other procedures and it's what a model actually reads before acting. Happy to try a stricter shape — numbered, or a table — if you want to see it side by side.

Your CLI aside may already be satisfied. internal/app/status.go routes every status change through one function, and close calls it rather than setting the field, so there's no path producing a record where status and rank disagree. Journaled as confirm, don't build rather than acted on.

---

META: Journal. Not tasks, and not outcomes yet.

Tasks are the wrong unit and your instinct about why is exactly right. A task is an attempt at part of the work — something someone could pick up. A turn isn't that. Turn-as-task gives you a transcript wearing a work breakdown, and the flip-flopping is fatal: tasks are meant to be completed, so two contradicting each other is a defect. The journal is append-only and never curated, so "I wanted X, then I wanted Y" is correct data there and broken data in a task list. That difference is the whole answer.

Outcomes: not yet, and WORK-0032 already licenses that. Rank two — written after the fact, still proven. The criterion hasn't arrived. Writing them now is the guessing you rejected two hours ago.

Exploration: yes eventually, no right now — and the reason is the interesting one. The distinction that matters is the journal is chronological, an exploration is topical. The journal answers what happened; an exploration answers what did we work out about X. Flip-flopping belongs in the first. A worked-through conclusion belongs in the second.

So an exploration earns its place the moment one specific question gets deep enough that the journal can't hold it — what a rules section should look like for agents is a live candidate already. But creating one now means defining the feature under time pressure, mid-experiment, which is the worst moment to design anything. Better: journal until a question demands a file, and let that be the feature's first real use. That's this project's own bootstrap order — lead with the thing, backfill the mechanism — applied to the mechanism itself.

51 journal lines. Next in the file: the gates have no journaling step, nothing verifies a move landed, and Sending work back describes a reopen the tool may refuse.

✻ Cogitated for 1m 14s

❯ META: i want to journal this strategy im using where we have 3 voices present 1. agent, 2. editor, 3. meta.  i like this and want to remember it for later as a learning.  we can coin it voices or profiles or personas or something, this should be something i can say and then we kick into this strategy and i dont need to explain how it works, and if im a newby the agent teaches me how it works when it recognizes this is the best tool

  Ran 2 shell commands

META: Journaled, and it changed WORK-0060 more than I expected.

This is the second instance of your own mode system. Editor mode was one; this is two. That matters — one instance is an anecdote, two is a menu, and WORK-0060 was speculative until now because it had nothing to generalize from. Both are recorded there as worked examples.

And it added a step the protocol didn't have. You want to name it and have it start with nothing explained, and have the agent teach a newcomer when it recognizes the mode would help. Nothing in your six steps does that — a mode isn't just a permission and a rule set, it's a thing with an explanation attached, offered on recognition rather than on request. That's now a seventh step, sitting between find a mode and pick it with them.

On what it buys, since that's what makes it teachable: without labels, "this is confusing" is ambiguous between critique-the-document and observe-the-process — so either the work stalls to discuss process, or the process observation dies inside the work. The label lets both run in one turn. The bigger effect is that it makes noticing cheap: a process observation normally costs a derailment, with a prefix it costs three characters. That's why so much has been surfaced in the last hour.

On the name — one sharpening that affects the choice. These are registers, not speakers: either party can use any of the three, and I've used all three today. So voices slightly over-attributes, since it implies each channel belongs to somebody. personas collides with what the word already means around models. I'd take channels — it carries concurrent and non-interfering, which is the actual mechanic — with tracks as the runner-up. Left undecided in the record, with the reasoning, so you can overrule it cheaply.

Worth noting the count: you're running an experiment about ad hoc work, and it has now produced a mode system, an observer requirement, and a second mode — none of which were the subject. That's the mode working, and it's also the thing that would worry an observer.

60 journal lines. Back to backlog-move when you are — gates have no journaling step, nothing verifies a move landed.

✻ Cooked for 1m 9s

❯ EDITOR: i think this document is going to turn into some moves are trivial and some have checks and balances.
  let's find out.  interview me to figure out what happens in between each gate

  META: after we are done im going to copy and paste the full session conversation into a document that we can save, and then analyise where are journaling needs work and what it missed or didn't do well.  and this should be an option that i can invoke at any time in the future.  as part of /session command set or /backlog command set

META first, quickly, so it's not lost.

  Ran 1 shell command

○ WORK-0062 · Review journal quality against the session transcript — captured

The method is the part worth keeping: a journal is a lossy compression of a session, and loss can't be measured from inside the compressed artifact — only against the original. Nothing in this project keeps the original today, so the transcript is an input the corpus has never had. Flagged that it operates on the journal (backlog) using the transcript (session-scoped), so it's a backlog procedure needing something session-manager owns. Also flagged: WORK-0051 is your retro skill with an empty body and an undecided name, and this is a concrete instance of what a retro would do — it may belong inside WORK-0051 rather than beside it. Your call, later.

---

EDITOR: the interview. Five questions, and I've grounded each in what the file says now so you're reacting rather than generating.

1. preparing → prepared. The table says "it is worked out", which is circular — it's the only move with no prose anywhere in the document. When you make that call, what do you actually check? Is it "outcomes exist", or something more like "I could hand this to someone else and they wouldn't come back with questions"?

2. unprepared → preparing. Is this genuinely just I started, or does something happen — deciding who's involved, checking nothing else is already shaping it, reading what's related? I'd guess trivial, and I'd rather have that confirmed than assumed, because "most of the ladder is bookkeeping" is the claim your hypothesis is testing.

3. todo → in_progress. Right now it's described purely as a claim about the present. Is there anything to check before claiming it — how much is already in flight, whether you have capacity, whether something else will collide with it? Or is the only check the honesty of the claim itself?

4. Backward moves. Sending prepared back to preparing currently owes nothing — no reason, no journal line. Should it? A record silently sliding backwards loses the fact that somebody looked at it and judged it not ready, which seems like the most valuable thing about the move.

5. Skipping rungs. We did captured → in_progress in one move an hour ago, and the table has no row for it. Is skipping legitimate, and does it owe more than a normal move — an explanation, a mode declaration, an observer note? This is the one where your two threads meet.

And one thing I already have data on rather than a question: closing is not trivial and the document knows it — it's the only move that already demands a journal entry first. Whatever pattern comes out of your answers, closing is the worked example of the heavy end.

✻ Cogitated for 58s

❯ EDITOR: interview me going from top down,
  so captured -> unprepared is the first one

  Ran 1 shell command

EDITOR — move 1 of 7: captured → unprepared, the first gate.

The document covers why this gate exists well — the pile is honest, crossing commits you to working it out and not to doing it. What it says nothing about is what actually happens when someone crosses.

Live data from today: WORK-0057 crossed it. You said "next will definitely exist but I'm not sure in what capacity", I matched that to the rung, you said move it, I wrote a journal line and ran set. The only thing recording why it crossed is that journal line — and I wrote it on instinct. The procedure asked for nothing.

1. Is anything checked before crossing? Duplicates, whether something already covers it, whether it's blocked on something nobody's selected. Or is the pile assumed already-checked because capture did that work?

2. Does crossing owe a reason, recorded? Right now nothing captures why this item and not the twenty next to it. A month later unprepared is a list of things somebody chose, with no trace of the choosing. Is that a loss worth a required journal line, or is the selection self-evident from what's around it?

3. Who can cross it? Invisible on a one-maintainer project, and the reason I ask is that it's the first place your observer thread bites: if selection is the moment a project commits attention, an org may want that to be someone's call specifically. Is that real, or am I inventing a problem you don't have?

And one cheap one: does it go backwards? unprepared → captured isn't in Sending work back, which only covers the two lower pairs. Is de-selecting a thing that happens, or does something selected stay selected until it closes?

✻ Worked for 37s

❯ to become unprepared, we are saying
  this work went from something we may do -> to something we intend to understand better
  - we are saying we think this work is actually worth understanding and scoping (not doing)
  - we are at the very least committing to attempting to prepare it to find out more
  - does stage or any other LKF attributes matter at this point, it still can be considered a draft i think, we are not useing it yet so we don't need to force it to provisional.
  - i think the kind SHOULD be defined by now, it's not a showstopper but very few things should make it much further without being set.  so maybe this is a place we strongly encourage kind to be set if it's not already, or we make it easy to look at and change if they notice something (rather than wasting a turn asking)
  - this is where we might want to know if the user intends any exploratation - hmmmm - dont think so, i think that gets asked in the next step honestly, but im not sure yet.
  - some work will become unprepared and then later realized that it is not worth doing, but ideally not much, this step should filter out most of it, and observers should be able to review our reasoning and reverse it - and on good teams this will barely ever happen
  - ideas probably get changed to "change" or whatever at this point
  - the work item should be well formed enough that two independent readers can at least understand what problem it is trying to solve, they may have very different ideas on how to solve that problem, but the problem should not be hard to understand, otherwise this isn't formed enough; and maybe this should happen at the preping step, im not sure, i bet different orgs will have different opinions on this
    what else should we be considering?

  for context, in the next step "preparing":
  - we may break it down into more work items
  - we SHOULD (almost MUST) consider all parties that need to be involved and define when they are necessary and what they need to do (QA, legal, compliance, security, leadership, product, stakeholders, sometimes customers, etc) - but this should be enforced when it makes sense and easy to skip when it doesnt - for example on a single person project this almost never makes sense, unless we use agents to do the security and compliance analysis for us, and maybe a product agent that vets if this is the right thing to be working on with limited resources and capacity
  - it would be ideal if outcomes are defined and in some orgs they may want there to be a hard rule that prepared work items have fully fleshed outcomes - and on smaller teams or single person teams that might not matter so much as long as the outcomes are good before the work starts in honest

  Ran 2 shell commands

EDITOR — six more to consider, and the first two matter most.

1. Your two-readers test isn't prose — it's executable. That's CLAUDE.md's divergence driver exactly: same instruction, two agents, diff the records. Give two agents the work item, ask each what problem it solves, diff the answers. It's the first criterion on the whole ladder that can be run rather than asserted, which is a much stronger reason to put it at this gate than "it feels right here."

2. And nothing in the pile passes it today. Six of the seven newest captures have an empty ## The problem. The criterion you just named is already failed by the corpus it would govern — so adopting it means either a backfill or an explicit grandfather clause, and picking neither is how a rule becomes decorative.

3. The gate currently has no back-pressure. 32 captured, 11 unprepared, 0 preparing. Crossing costs nothing, so unprepared has become a second pile rather than a filter. If crossing commits you to attempting to prepare, and nothing is being prepared, then eleven commitments are outstanding and nobody's tracking that. A gate that only moves things between piles is ceremonial.

4. "No kind" and "kind: idea" are different failures here. The type definition says idea is "a classification — it becomes one of the others." So idea is explicitly transitional, and surviving this gate is a contradiction in terms — while an unset kind is merely unset. That argues for two different behaviors: nudge the missing one, but treat a surviving idea as a genuine inconsistency the tool can name.

5. Reversal has no vocabulary. close has four dispositions specifically so the record says why it ended. De-selecting has none — unprepared → captured can only be recorded as a bare status change. If observers are meant to review the reasoning and reverse it, the reversal itself needs to record its reasoning, and that's the same argument ADR-0007 made for close.

6. "Ideally not much" is a metric, and it'd be the ladder's first. Reversal rate out of unprepared measures selection quality directly. That's computable rather than eyeballed, and it belongs with WORK-0046's conditions. Worth noticing that you reached for a rate without meaning to — it's the first place this model implies a number.

One smaller thing: ten of the eleven unprepared records have no rank. Crossing decides something is worth understanding; nothing decides what to understand first. If the gate creates a queue, that queue isn't ordered yet.

On your two uncertainties — I'd keep the two-readers test at this gate, not preparing, precisely because it's runnable and cheap. It's a comprehension check, not a design check, and pushing it later means preparing starts on records nobody can read. Your instinct that orgs will differ is right, and that's a config knob rather than a decision.

✻ Crunched for 1m 12s

❯ EDTIOR: we need both of these as established policy:
  adopting it means either a backfill or an explicit grandfather clause
  we need to have some kind of work flow for both options and a lean towards backfill - ideally through mechanical, reproducable migration, but if not a model can do it for us with instructions.

  Ran 1 shell command

WORK-0037 already owns the backfill half thoroughly — five instances, the constraints, what a repair must do. What nothing covers is grandfathering, and that's the half that needs establishing. Writing it as a bundle policy, since it's a rule rather than work.

  Searched for 1 pattern, ran 3 shell commands

Caught one on my own edit — a bundle document must not link outside itself, and mine was the only one that did.

  Ran 1 shell command

EDITOR — established. local/backlog 0.18.0 adds policy/adopting-a-rule-the-corpus-does-not-meet.

The half that was actually missing was grandfathering. WORK-0037 already owns backfill completely — five real instances, the constraints, what a repair has to be able to do. Nothing anywhere covered the option of saying this applies from here. So the policy is mostly about the one nobody had written down.

The finding that makes it a policy rather than a note: doing neither is the default. Not choosing looks identical to adopting the rule, and produces a corpus visibly violating its own policy — at which point a reader concludes the rules are aspirational and reads the next one the same way. One unenforced rule devalues the others. That's what raises this above tidiness. The worked example is yours from twenty minutes ago: the two-readers test proposed the same week six records shipped with an empty ## The problem.

Lean backfill, with the preference order you gave. One addition I made beyond it: when a model does the backfill, the instructions are the artifact and get kept with the change. A model run nobody can repeat is a one-time edit wearing a migration's clothes — the instructions are the substitute for reproducibility, not a waiver of it.

Grandfathering is legitimate and must be explicit, in the same change that adopts the rule, naming which records are exempt and why. Unrecorded, it's indistinguishable from a rule nobody enforces. And a clause with no end creates a permanent second class of record, so if it's meant to expire, say what ends it.

Left open in the doc rather than invented: when grandfathering beats backfilling, and who decides. Cost is the obvious axis and probably not the only one — a rule about what a record means may be unbackfillable at any price, since nobody can reconstruct what an author intended.

And I broke a rule while writing one. The policy wikilinked out to WORK-0037; a bundle document must not point outside its bundle or it stops surviving being copied. Delinked to a prose reference, the way every other document there cites spec.md. Journaled as self-inflicted — it's the same failure the policy describes, one level up.

Index regenerated, links audit clean, policy registered as standby on a topic match so it fires when a rule is being adopted rather than loading every session.

Move 2 of 7 next — unprepared → preparing — unless you want to keep going on the gate.

✻ Sautéed for 2m 2s

❯ sorry i was asking to to cpature the need, not produce the answer.
  but ok, whatever for now.  i think i need a way to say i want you to record this vs i want your opinion and guidance vs i want you to do stuff about it
  let's move on to preparing

Misread that — you wanted the need recorded, not the answer produced. The policy is written and easy to pull if you'd rather it wasn't.

That's a real gap and I'm capturing it rather than solving it, since that's the lesson.

  Ran 2 shell commands

Recorded in WORK-0060, not designed. The useful part is the asymmetry: reading record as do burns real work and leaves an artifact nobody asked for; reading do as record costs one turn. The safe default when no verb is given is the cheaper misread — and mine went the other way.

---

EDITOR — move 2 of 7: unprepared → preparing.

You've already given a lot about what happens during preparing and what prepared should require. These are about the moment of entering it, which is the move itself.

1. Is entering preparing a claim about the present? workflow-status.md says it's "the only rung that describes activity rather than position", and that a record sitting there with nobody shaping anything is a signal worth seeing. That makes it the same kind of claim as in_progress — and the document warns about stale in_progress while saying nothing about stale preparing. Is I am working this out now the meaning, or is it this is next up for preparation?

2. Your own open question from last turn: does exploration intent get asked here? You said it probably belongs at this step but weren't sure. Now that you're at it — is does this need investigation before it can be scoped a question asked on entry, or something that emerges while preparing?

3. Who prepares it, and is it whoever selected it? Your parties-involved point covers who gets consulted during preparation. This is the narrower one: is the preparer assigned at this move, self-selected, or does nobody care?

4. What happens when preparing reveals it shouldn't be done? You said reversal out of unprepared should be rare and reviewable. This is the rung where it'd actually be discovered — you started working it out and found the answer is no. Does it go back to captured, or close as rejected without ever being worked?

✻ Sautéed for 55s

※ recap: We're running an experiment on ad hoc "editor mode" work (WORK-0059), improving the backlog-move procedure by walking it top-down and journaling everything. Next: you answer my four interview questions about the unprepared to preparing move.

❯ EDITOR:
  at it's most basic, preparing is the "in progress" of the preparation pipeline
  it where we make the work actionable
  but this stage might be the biggest collection of things that could happen
  and big orgs and little orgs will use it very differently
  we should plan around big orgs and then provide a way to trim and scale down for small ones - as well as single person teams (always a team because agents are the team members)

  some things i can think of, not everything:
  - product managers might have their own pipeline of inside of this: What problem are they experiencing? → Who else has it? → How painful is it? → How valuable is solving it? → Does it fit our strategy? → What's the cheapest good solution? → Is it more valuable than the other things we could build?
  - security officer might have their own pipeline inside or preparation: Understand → Classify → Threat Model → Assess → Control → Require → Verify
  - compliance officer: Classify → Identify → Map → Assess → Require → Evidence → Approve
  - legal: Understand → Identify → Analyze → Mitigate → Require → Approve
  - customer facing: Understand → Represent → Validate → Assess → Shape → Advocate
  - QA: Understand → Analyze → Plan → Specify → Test → Verify → Sign Off
  - operations/infrastructure: Understand → Assess → Design → Prepare → Provision → Validate → Operate
  - Engineer: Understand → Explore → Design → Estimate → Plan

  this is not a comprehensive or well formed list, it's a rough idea:

  Function    Fundamental question    Pipeline
  Product    Should we do it?    Discover → Validate → Assess → Prioritize → Commit
  Security    Can we do it safely?    Understand → Classify → Threat Model → Assess → Control → Require → Verify
  Legal    Can we legally do it?    Understand → Identify → Analyze → Mitigate → Require → Approve
  Customer Advocacy    Is this right for customers?    Understand → Represent → Validate → Assess → Shape → Advocate
  Compliance    Does it satisfy our obligations?    Classify → Identify → Map → Assess → Require → Evidence → Approve
  Infrastructure / Ops    Can we operate it reliably?    Understand → Assess → Design → Prepare → Provision → Validate → Operate
  QA    Does it actually work?    Understand → Analyze → Plan → Specify → Test → Verify → Sign Off
  Engineering    How do we build it well?    Understand → Explore → Design → Estimate → Plan → Build → Integrate → Verify → Deliver

  for the MVP we should keep this simple, but do know there will be many sub-workflows inside this workflow_status eventually
  bigger orgs will need all of it
  smaller ones will need little to none
  just because the sub-workflow pipelines exist doesn't mean they should get used for every work item. In fact, that would be a massive amount of noise. So part of this process is deciding which pipelines are getting activated. And then deciding do they need to be done before or after a work item is prepared? And what steps do they block? what can happen in parallel and what needs to be sequenced. and what depends on other pipelines. In this step, we will always figure out which work comes before and which work is dependent and which work will come after and who needs to be involved and who needs to be informed. Ideally, in most cases, the other pipelines should become their own work items that are attached to this work item.

  So to make this even simpler, it might be the easiest thing to do is to have it so you can create work items from other projects and block or assign them two steps in this work items workflow status. For example, we might say the legal work, LEGAL-0004 needs to be completed before XYZ. And most or all of that type of work should be defined as part of preparation. sometimes it will come later, but our goal should be to define or establish it all before a work item is marked prepared.

  From an MVP, we need to decide how much of the outcomes and tasks have to be written in this stage, how much we should encourage them to be written in this stage, and if they belong in this stage at all. or if there's a separate workflow status somewhere in here where that separates other people's work from the execution team's preparation.

  somewhere in this step is where estimation, scoping, requirements, prioritization, dependencies, risks, and things like that are (at minimum) established. and ideally well defined.

  So to get into the metadata, I think a work item can still be a draft. This is often where you may start to draft outcomes. A few tasks could come out of it, but generally tasks are created here. In some instances, we may want an exploration or a proper spec as part of this step. which means we needed to fill up some policy for: if a spec is appropriate. I don't know if we will always keep specs as optional or if they're just something we use in the beginning and eventually graduate away from. But to simplify, I think a spec will get used when users want more control or they don't trust the model to know how to do what it is they want to have happen. and ideal world outcomes should give us optimal results while holding all aspects of: the requirements, the spec, the definition of done, and the proof (test plan).

  It should be impossible for something of kind idea to make it past this step.

  What else are we not considering or I forgot to consider? I believe this is the most overwhelming workflow status. So we are gonna miss things on the first pass for sure.

Recording all of it first — this is the densest input the experiment has had.

  Ran 2 shell commands

EDITOR — nine things, worst first.

1. "Define or establish it all before prepared" is load-bearing and ambiguous. Established means the external work items have to close first — which in a big org is weeks, and makes preparing the rung where things go to die. Identified means the dependencies are named and prepared stays reachable. You wrote both words. Which one it is determines whether this rung is a stage or a graveyard.

2. Cross-project references break the key scheme. Keys are per-corpus, so LEGAL-0004 implies a foreign namespace that doesn't exist — and spec.md §6.1 forbids a coordinator. Worse, knowing whether LEGAL-0004 closed is a read across a corpus boundary, which is a capability this tool has never had. The whole "other pipelines become work items" idea rests on it, so it's the assumption to test first.

3. Blocking across corpora means readiness can regress with nobody touching the record. If LEGAL-0004 reopens, a prepared work item silently stops being prepared. Every status change on this ladder has been local until now; this is a new class of staleness.

4. A dismissed pipeline has to be recorded, not just an activated one. Otherwise nobody can tell later whether Security was assessed and ruled out or never thought of. That's precisely the canceled vs abandoned distinction ADR-0007 already settled, and its test applies unchanged: the record carries what it cannot reconstruct. Activation lists are cheap; dismissal lists are the ones that matter in a post-mortem.

5. Who decides which pipelines activate is unowned — and the failure mode is structural. The person choosing is often the one least qualified to know Security applies. A self-selected menu fails exactly where the expertise is absent, which is the small-org case this design is meant to serve. Agents standing in for those roles is your answer elsewhere; this is where it becomes load-bearing rather than nice.

6. "Plan around big orgs, then trim" conflicts with this project's own method. CLAUDE.md's bootstrap order leads with the thing and backfills the mechanism. Designing for the org you don't have is how you build something no small team can use — and you can't test any of it here. The reconcilable version: design the extension point, ship only the default. Worth deciding deliberately rather than discovering the conflict later.

7. You gave two different rules for idea in two turns. At the first gate: it probably resolves to change there. At preparing: impossible to get past this step. Both would end up in the document silently. The type definition says idea is "a classification — it becomes one of the others", which argues for the earlier one.

8. Estimation is the only thing in your establish-here list that decays. Requirements, risks and dependencies stay roughly true. An estimate written entering a long preparing is worthless by the exit. It may belong last rather than in the pile.

9. Your sub-pipelines-as-work-items preference is also the answer to a problem you didn't name. A work item carries one workflow_status. It cannot express Legal at step 3 while QA is at step 1 and Engineering hasn't started. Fanning out isn't just tidier — it's the only shape that fits the ladder, which is a much stronger argument for it than preference.

One more, smaller: spec-appropriateness probably tracks handoff distance, not just trust. Solo with an agent, you hold the spec in your head. Hand it to another team and you can't — so the spec is the artifact you need when the doer is not the definer. That scales with org size, which fits your framing better than trust does and is easier to write a policy against.

✻ Baked for 1m 35s

❯ And it's most basic form, I think preparation is where we scope work, break it down. define our definition of done, which is basically outcomes, and try to refine the outcomes so they are effective. which means they need to pass some kind of test. But it does not guarantee that all the outcomes are present. Although in an ideal world, it should.

  Ran 1 shell command

That reduction is the version that can actually ship. Four things on it.

The effectiveness test doesn't need inventing — it needs assembling. Three pieces already exist: spec.md §5.2 names outcome.unmeasured for an outcome with no verify_by; when-a-work-item-splits names the no-edge outcome that can never be finally true — "the corpus is clean", which absorbs tasks forever; and your two-readers test from the first gate covers comprehensibility. Has a check, has an edge, is understood by two people. Two of the three are already written down as conditions.

The test is also the exit condition, which makes it more important than it looks. Scoping and breakdown can run forever — nothing says when you've broken it down enough. Outcome refinement stops when the outcome passes. It's the only thing in preparing with a natural stopping point, so it's the candidate answer to when is preparing done as well as is this outcome any good.

"Doesn't guarantee all outcomes are present" has a consequence downstream. close … completed checks against the outcomes — so it's checking a set nobody ever claimed was complete. The outcomes hold really means the outcomes we happened to write hold, and nothing distinguishes a work item with three good outcomes from one with three good outcomes and a missing fourth. That's not an argument for requiring completeness; it's an argument that completed is a weaker claim than it reads as, and somebody should know that.

Effective and complete are two axes, and you may already have the field for the first. Outcomes carry stage — draft, provisional. That might already be the is this outcome any good yet marker rather than something new. WORK-0036 is the open question on whether stage is used correctly or removed, and this would be a real use for it rather than a field nobody exercises.

Nothing to decide there — recorded. preparing → prepared next, which is where the completeness question comes due.

✻ Sautéed for 38s

❯ And I forgot to include in that list of basic form. This is also where we define dependencies, which include both external teams and external work items. and consider who needs to be included and who needs to be informed.

  Ran 1 shell command

Three notes on that addition.

External teams and external work items are different kinds and you listed them together. LEGAL-0004 has a state the tool can check — did it close. "We depend on the platform team" has no state at all. One is a link that resolves; the other is a note that rots silently and nobody finds out until the work stalls. They probably need different treatment, and the second one is where a dependency quietly stops being true.

On a solo project the included list collapses to one person, but informed doesn't — future you, an observer, the agent in the next session. So the informed half is what earns this a place in the basic form rather than the maximal one. And it's the natural input to WORK-0061's observer mechanism: the people who needed to know are exactly the people a breach should surface to. That connects two threads that have been running separately.

Same shape, second sighting. Consider who needs involving and decide which pipelines activate both name an activity whose output nothing records. Consideration that leaves no trace is indistinguishable from omission — I said it last turn about dismissed pipelines, and it's now appeared twice in one rung. Twice makes it a shape rather than an oversight, and it's probably the thing preparing most needs from the record format: somewhere to put we thought about X and the answer was no.

preparing → prepared whenever you're ready — that's where outcome completeness comes due.

✻ Crunched for 37s

❯ It's not just external teams and external work items. It's also deliverables that we depend on, which could be documents or reports or designs or mock ups or security requirements, so on and so forth.

  Ran 1 shell command

Three kinds then, and the third has the nastiest properties.

Existence is checkable; adequacy isn't. A work item's closed is a claim somebody made. A file being present is not a claim about anything — the mockup can exist and be two revisions stale, the security requirements doc can exist and not cover your case. So a deliverable dependency looks resolvable in a way it isn't, which is worse than a team dependency that's honestly unresolvable.

And it drifts with no event. A work item reopening is at least an action taken on a record — something happened that could be noticed. A design doc gets edited and nothing signals it. Of the three kinds, this is the one that goes stale most quietly.

You usually need two links, not one. A deliverable is normally the output of somebody else's work item, and the two answer different questions: the work item says is it coming, the deliverable says is it here and does it serve. Depend on only the work item and you can't tell if what arrived is usable; depend only on the artifact and you can't tell if it's ever arriving.

This has already bitten this corpus, before it had a name. WORK-0037's instance 3: package renames left a closed work item's verify_by naming paths that no longer run — "and nothing could report which records referred to a shape that had moved." That's a deliverable dependency going stale, already logged as a migration problem. It's the same failure seen from the other end, which is worth knowing when you decide where this gets handled.

Recorded, not designed.

✻ Cogitated for 32s

❯ ok let's move on to prepared
  So at the prepared step, we should at the very least have all of the knowable unknowables defined. and this is a new concept, so I'll define it here. But we need to define https://www.theuncertaintyproject.org/tools/rumsfeld-matrix
  and handle each quadrant of it correctly
  I wanna get into that later and not right now. But at this point,
  we should know all the people involved
  when they need to be involved.
  what they need to deliver.
  where all the blockers will be, known ones
  We should have de risked the work by now. So that should be a suggestion. strong suggestion. The only exception is if this work is, in fact, derisking as an outcome.
  and we should say it is a smell if a work item derisks its own work. Usually, Derisking should be its own work item, generally an inquiry.
  So I know I'm missing some things, but to sum it up, at this point, we're saying the work is ready to be in queued, and the only thing stopping it from being in progress is scheduling and capacity.

  Ran 1 shell command

Rumsfeld matrix recorded as deferred, with the link, not designed.

Your summary line is the best definition of prepared in this whole conversation, and it's a test rather than a description. The only thing stopping it from being in progress is scheduling and capacity. That's falsifiable: ask someone to name a reason it can't start, and if any answer isn't scheduling or capacity, it isn't prepared. A checklist can be complete while something else still blocks — this can't. I'd lead the rung with it and let the list of people, timing, deliverables and blockers be the things you check because of it.

And it explains why the two gates feel different in kind. prepared is a claim about the record; todo is a claim about the world. Scheduling and capacity live nowhere on the work item, so prepared means no record-local blocker remains. That's a cleaner split than "we know what this is" versus "we're going to do it".

The same ambiguity has now appeared twice. "Where all the blockers will be, known ones" — identified, or cleared? If identified-but-unresolved, then something other than scheduling is stopping it and the summary is contradicted. This is the identical define-versus-establish question from preparing, one rung later, still open. Two sightings makes it the thing to settle before either rung can be written.

The fourth quadrant can't be handled here, and that's fine because it already has a home. Unknown unknowns are undefinable at prepared by construction — so prepared can never mean risk-free. The discovery of new unknowns during in_progress is lifecycle.md §2.8's Redefine phase, which already exists. Worth knowing before you design the matrix work: one quadrant is served by machinery you already have.

The de-risking smell is mechanically detectable — but only if outcomes carry a distinction they don't have. An outcome that reduces uncertainty versus one that delivers a state. A work item holding both is de-risking its own work, and that's a condition WORK-0046 could compute. Today the outcome type has no such field, so the smell can only be noticed by a reader. That's the difference between a rule and a guarantee, and it's cheap to close.

prepared → todo next — the second gate, and the one this project has never used honestly.

✻ Sautéed for 49s

❯ and really the only difference between this workflow status and the next one to do is scheduling is no longer a blocker. In To do, we are making a commitment, and the only blocker is capacity. short of there being some unexpected blocker or risk that became true.

  Ran 1 shell command

❯ for most enterprise teams putting in something in to do is the same as putting it into a sprint.

  Ran 1 shell command

That refinement generalizes into the best frame for the whole ladder: it's a monotone narrowing of the blocker set.

- captured — blocked by nobody having decided it's worth understanding
- unprepared — blocked by it not being understood
- preparing — blocked by understanding being underway
- prepared — blocked by scheduling and capacity
- todo — blocked by capacity alone
- in_progress — blocked by nothing but the work itself

Every move can then be tested identically: which class of blocker did this remove? That's a far better organizing idea than "most of the ladder is bookkeeping", which is the line currently in the file — and it makes the trivial-versus-heavy question you started with answerable rather than a matter of taste.

It also gives the second gate the affirmative test it lacks. The doc only says do not cross because preparation finished. The positive version is cross when scheduling is resolved — and now there's a reason it's the expensive gate that beats the one in the file: it's a promise, and a promise has a subject even when a solo project leaves it implicit.

"A risk that became true" is a transition the model can't name. It's the legitimate reason for todo → prepared, and the doc names only the weaker cause — "a todo item nobody will get to". Two distinct backward causes collapsed into one, and the interesting one is missing.

On the sprint anchor — it confirms the commitment reading and adds three things the ladder doesn't model:

- A sprint is time-boxed, so todo acquires an implicit expiry. Work not started by the boundary rolls over or comes back out. That makes todo the only rung with a natural lifetime, which fits it being the only one blocked by something that isn't a property of the record.
- A sprint is a batch. The commitment is to a set, not to each item — so the set means something no individual todo record carries. Open question whether the model needs the batch or whether rank within todo is sufficient.
- todo → prepared already has a name out there: descoped. Worth borrowing rather than inventing, since enterprise teams make that move constantly and will look for the word.

todo → in_progress next, unless the sprint batch question is worth pulling on first.

✻ Churned for 56s

❯ moving something to todo - is putting it into focus and saying We need to figure out a way to map this work to our available resources in in a way that best utilizes our capacity and produces the best quality.

  Ran 1 shell command

That corrects my framing — I had capacity as a blocker being waited on, and you're describing it as the moment the allocation problem becomes the live work. Four consequences.

It contradicts the ladder document. workflow-status.md says preparing is "the only rung that describes activity rather than position", and treats todo as a passive queue — "enqueued for action; should start when capacity allows." What you just described is an activity. So either that claim about preparing is wrong, or todo needs a different name for what it holds. It's a direct conflict between two normative documents, and it's the kind that survives for years because both sound right in isolation.

The mapping has two objectives that trade off. Best utilization and best quality pull apart — assigning the fastest hand maximizes throughput; assigning someone who has to learn it buys resilience. The document says nothing anywhere about how work maps to people, so there's no place for that judgment to live even as prose.

And the mechanism for the activity this rung is about doesn't exist. ADR-0008 settled that taking a task expires and owning a work item doesn't — and WORK-0031 records that take, release and steal don't ship. So todo is defined by an allocation the tool cannot record. That's the sharpest gap yet: not a missing rule, a missing verb for the rung's whole purpose.

Focus is scarce and distinct from capacity, which makes the size of todo a measurement rather than an accident. You can hold hours and still have no attention on a thing. Fifty items in todo is not a focus set — it's prepared with a different label. Zero, which is this project today, is the opposite failure: nothing is in focus at all. That gives the rung a health signal it currently has no way to express.

✻ Worked for 36s

❯ Side quest. This is meta. We should give an algorithm for assigning work to developers or assignees rather. there's gotta be a strategy that's well established in mathematics where you take the work that is the easiest and the work that is the most difficult and assign it to a spectrum of individuals who have different skill sets and abilities. thus that the work output is of the highest quality with a given combination.

META side quest — capturing it, and the math you're reaching for exists.

  Ran 1 shell command

○ WORK-0063 · An algorithm for assigning work to assignees — captured

The canonical form is the assignment problem, solved optimally by the Hungarian algorithm in cubic time, given a value matrix of person × work item. Add per-person capacity and it becomes the generalized assignment problem — NP-hard, but with well-studied heuristics. If both sides have preferences rather than a value matrix, it's stable matching instead, which optimizes for nobody wanting to defect rather than for total output.

Your specific intuition has a proof. Sort work by difficulty, sort people by skill, pair them in order — that's the rearrangement inequality, and it is the maximum when output is multiplicative in skill and difficulty. So it isn't folklore, it's a theorem.

But the catch is worth more than the algorithm. Sorted matching maximizes this period and degrades the next two ways:

- The skill distribution never changes. The hardest work always goes to the person already best at it, so nobody is ever stretched and the matrix you're optimizing over is frozen.
- The top person becomes a bottleneck. Their utilization approaches one, and queueing theory is unambiguous about what happens to waiting time as utilization approaches one — it grows without bound. The locally optimal assignment makes the system slower.

So the honest framing is explore-versus-exploit, not optimization. Assigning to the best hand is exploitation; assigning to someone who might learn it is exploration. The practical shape is mostly sorted, with deliberate stretch assignments — which is also exactly your two objectives from last turn, now with a reason they can't both be maximized.

Recorded, not solved. The open question on the record is the one that matters for this tool: whether it should implement an assignment, suggest one, or only record the inputs one would need — because the third is cheap and the first requires the tool to hold a skill model of people, which is a very different product.

Back to todo → in_progress when you are.

✻ Worked for 47s

❯ I was originally thinking of this idea for humans, but I think it also applies to agents where we can use cheaper agents to do easier work.

  Ran 1 shell command

That extension inverts most of what I said last turn, which makes it worth more than the original.

Agents don't learn from an assignment, so the trap disappears. The whole reason sorted matching degrades for humans is that people grow — assigning the hardest work to the best hand freezes the skill distribution. An agent's capability doesn't change by doing the work. So for agents, the rearrangement inequality is simply correct, not merely locally optimal. Which means a mixed team needs two assignment policies, not one — and the naive algorithm is right on one side and a trap on the other.

The queueing objection weakens too. Agents scale horizontally; ten instances of the expensive model run at once. The bottleneck that makes a top human's utilization dangerous isn't the same constraint.

But the objective function changes, not just the algorithm. A human costs roughly the same whatever they're assigned, so assignment maximizes quality. An agent's cost is per task and varies by orders of magnitude — so it becomes minimize cost subject to a quality floor, which is a budgeted knapsack, not the assignment problem. Different math for the same-looking question.

And the mistake people actually make: a cheap agent is only cheap if the output doesn't need checking. If it has to be reviewed by an expensive agent or a human, the true cost is cheap-plus-review, which can exceed doing it well once. The unit is cost to acceptable output, not cost per attempt.

This corpus has the instance already. WORK-0003 records every line of Go written by an agent and merged unreviewed — and "several bugs were found by running the tool rather than by its tests." That's cheap output with the verification cost deferred, not avoided. It's the pattern, already logged, before it had a name.

Difficulty isn't one scale across humans and agents. A mechanical refactor over two hundred files is tedious for a person and trivial for an agent; judging whether prose reads well to a newcomer is the reverse. A single difficulty matrix can't serve a mixed team, which breaks the clean formulation more than anything else here.

One thing you're already doing that pays off here: LUMA_BACKLOG_ACTOR records agent:<model>/<project> on every record. So the corpus already knows which model produced what — and "is work done by cheaper agents reopened more often" is answerable from provenance you're already collecting for a different reason. That's the cheapest version of this whole idea and it needs no algorithm at all.

✻ Crunched for 48s

❯ and by easier work. I mean, less complex work. And by part of work, I mean, more complex work or work that is more sensitive or critical.
Interrupted · What should Claude do instead?

❯ and by easier work. I mean, less complex work. And by harder work, I mean, more complex work or work that is more sensitive or critical.

  Ran 1 shell command

❯ or where quality matters more

  Ran 1 shell command

Four axes then, and they want four different responses.

┌─────────────┬────────────────────────────────┐
│    axis     │         what it wants          │
├─────────────┼────────────────────────────────┤
│ complexity  │ capability                     │
├─────────────┼────────────────────────────────┤
│ sensitivity │ authorization and independence │
├─────────────┼────────────────────────────────┤
│ criticality │ verification                   │
├─────────────┼────────────────────────────────┤
│ quality     │ judgment and iteration         │
└─────────────┴────────────────────────────────┘

Only complexity is answered by assignment at all. A better model doesn't make a security-sensitive change less sensitive, doesn't reduce what depends on a critical change working, and doesn't raise the standard something is held to. So "give the hard work to the expensive agent" is right on one axis and actively dangerous applied to the others — it concludes a capable agent can handle sensitive work alone, which is precisely the wrong inference.

This corpus already records that exact failure. WORK-0003: "The security-shaped parts were designed by the same agent that reviewed them." Capability was high and independence was zero. Sensitivity needed a second party, not a better first one.

And quality is the axis where cheap-plus-review breaks completely. Verification catches wrong; it does not catch bland. A correctness review is close to a decision procedure and mediocre output passes it cleanly. So the strategy that rescues complexity and criticality — use a cheap hand and check the result — cannot rescue quality. Quality has to be produced, not recovered.

Today's naming work is the clean instance. Choosing backlog-rundown broke nothing, had no criticality, and touched nothing sensitive — pure quality axis. A cheaper hand there would have failed silently, shipping a defensible name that was slightly wrong forever, and no review would have flagged it.

That's probably the practical rule hiding in all this: the axes differ in whether a bad outcome is loud or quiet. Complexity and criticality fail loudly, so cheap-plus-verification works. Sensitivity and quality fail quietly, and quiet failures are the ones you have to spend on up front.

✻ Cogitated for 54s

❯ I actually think it's five axes, complexity, sensitivity, which could also mean for, like, sensitivity of customers, criticality, which could have to do with uptime or how big the impact can be, quality, which I don't think needs definition. And then a fifth category, which is compliance and security, which means how much risk are we introducing to this becoming a massive lawsuit or damage of trust and reputation. .

  Ran 1 shell command

Five works, and the split you just made is the important one — separating how delicately this must be handled with people from how much legal and reputational exposure it creates. Those were collapsed in my version and they behave completely differently.

Axis 5 is the only one whose answer is a hard gate rather than an assignment. The other four fail in bounded, recoverable ways — rework, an outage, a slow reputational cost — so you can optimize against them. Compliance and security failure is unbounded. You don't optimize an unbounded downside; you refuse to proceed without the check.

And that answers something we left open hours ago. The mode protocol has put the brakes on, and my objection was that the brake fires hardest when the mode menu is emptiest, which trains people to route around the system. The brake is for axis 5. Not for skipped ceremony, not for missing outcomes — for compliance and security exposure. That makes it rare, defensible, and the one refusal nobody argues with, which is exactly what a brake needs to be to survive.

Axis 5 is also the only one where the record is itself a deliverable. Your own compliance pipeline lists Evidence as a step. For the other four the artifact is the work; here, proving it was done right is a separate output — and that's precisely what WORK-0061's observer mechanism has to serve. Two threads that have been running separately since this morning converge on this axis.

Customer sensitivity is where your standing solution breaks. A single-person team can have agents stand in for security, compliance, or product analysis. An agent cannot be the customer. Your own Customer Advocacy pipeline has Validate and Advocate steps that require actual customers, so this is the one axis that can't be filled by adding more agents — it can only be filled by talking to someone.

And the cheap MVP is smaller than it looks. Five ratings on a work item is enough to drive assignment, gating, and retro analysis with the tool implementing no algorithm whatsoever. Record the inputs, leave the judgment with people, and the data for evaluating whether that judgment was any good accumulates for free — same argument as the actor field already collecting which model did what.

✻ Brewed for 47s

❯ I don't see why an agent can't play as a customer advocate.

You're right — I overstated it at the role level. An agent can hold the customer's interest in the room, argue their side, and catch customer-hostile decisions, and most advocacy in practice is that rather than talking to anyone.

The limit is narrower and it's one step, not the role.

  Ran 1 shell command

❯ understand that is where they're weakest, but I think they can still role play the position.

Agreed — and role-playing the position has real value even ungrounded, because most customer-hostile decisions aren't subtle. An unvalidated advocate still catches the obvious ones, which are the majority. Nobody holding the position at all is much worse than someone holding it from priors.

So the fix isn't don't use an agent for it — it's record which one it was.

  Ran 1 shell command

And the model for that already exists in this project. ADR-0007 separates the doer's assertion from the checker's verdict — that's exactly this. An agent can assert the customer position; only contact can validate it. Two fields, not a permission question. So the record distinguishes advocated from validated, and nobody has to decide who's allowed to advocate.

Worth noticing that ADR-0007's shape has now answered three different questions today: dispositions at close, de-risking outcomes versus delivering ones, and this. The assertion-versus-verdict split generalizes well past outcomes, which is probably a finding about the decision rather than about any of the three.

The generalization that survives from my overstatement: for all five axes, the question isn't whether an agent can do the role — it's which steps need external contact or external authority. An agent can threat model but can't observe whether a control is actually deployed; it can map compliance obligations but can't approve, which is authority, not capability. Neither limit is about how good the model is, which is why "use a better agent" doesn't move them.

We're a long way from backlog-move. Want to keep going down this thread, or head back to todo → in_progress?

✻ Churned for 56s

❯ I wonder if there was things that had to happen to make something to do from a metadata standpoint. I think at that point, We should highly recommend it no longer be a draft. I think making it in progress forces it to be provisional and to do is where users should start getting a warning about draft being a problem.

  Ran 1 shell command

You've invented a reusable shape there: tolerate → warn → force. That's three steps of escalation along the ladder, and it needn't be about stage alone — the same pattern could govern kind being set, outcomes being present, and the five axes. Worth naming as the pattern rather than as a stage rule, so each future requirement doesn't invent its own escalation.

On "forces" — there's a version that doesn't refuse. The move can write the field rather than block on it. Moving to in_progress means it's no longer a draft, so the tool sets provisional the same way it already writes rank alongside workflow_status. Same guarantee, no refusal, and consistent with report, never refuse — which WORK-0022 states as a hard constraint citing spec.md §5.2 and §5.0. Given that the only place we've agreed a genuine block belongs is axis 5, I'd lean hard toward force-means-write here.

And stage is doing double duty, with the less important one getting the attention. Outcomes carry stage too. A work item at draft while in progress is untidy; an outcome at draft while in progress means you're building against a target nobody has settled. That's the one worth warning about. WORK-0036 is the open question on whether stage is used correctly at all, and this would be its first real use.

EDITOR note on the document itself: metadata obligations are piling up by rung — kind set by unprepared, idea gone by preparing, stage warned at todo and forced at in_progress. Right now each would land as a sentence inside its own section. That wants to be one table, or the rules become findable only by reading the whole file, which is how they drift.

One small thing: the warning has to fire at the transition, not persist. A draft warning that shows on every listing afterwards becomes the noise everybody learns to skip past — and then it's not a warning, it's decoration.

✻ Worked for 39s

❯ yes force should write the field, not refuse the move

  Ran 1 shell command

Recorded. The general form is broader than stage: when a move implies a field value, the move writes it. ADR-0005 already does exactly that for rank alongside workflow_status — status.go says nothing offers a way to write one without the other — so this generalizes an existing precedent rather than inventing one.

It also makes the tool's refusal surface nearly enumerable, which nothing has stated yet:

- close … completed without outcomes — ADR-0007
- axis-5 compliance and security gates

Everything else observes or writes. That's a much smaller surface than the ladder's prose currently implies, and it's a strong thing for the document to say out loud — a reader who knows the tool only refuses in two places will trust the rest of it differently.

Not writing it as an ADR — it's decision-shaped and that's your call to make, not mine to assume. Say the word and I'll record it properly.

Same for the document: I'm holding the edits until the interview finishes rather than writing them as we go. Three moves left — todo → in_progress, closing, and reopening — and several answers so far have revised earlier ones. Editing now would bake in versions that contradict what's still coming.

✻ Churned for 32s

❯ and we should establish that all changes two metadata and key fields should be visible. So users and agents have the ability to see that and correct it.

  Ran 1 shell command

That's what makes force-writes safe, and the two belong together as one position. Force-writes without visibility is just silent mutation — the rule you accepted a minute ago is only defensible with this one attached.

There are two visibilities and they need different mechanisms. At the moment — the command tells you what it wrote. After the fact — the history shows what was written and by whom. Git gives you the second completely and unforgeably; spec.md §5.5 already says git is the machine record. The first is currently absent.

Observed in this session, an hour ago. I ran set WORK-0059 workflow_status=in_progress and the output was updated backlog/work-items/WORK-0059-.../index.md. It never mentioned that it had also written rank: 060.0010.000. I found that out by grepping the file. That's precisely the failure this requirement guards against, and it happened while we were designing the requirement.

One tension with a decision in force. "And correct it" collides with set refusing the rank field — ADR-0005, and a shipped behavior. So there's a field the tool writes that you can see and cannot directly change. Correction does exist, through rank --before/--after/--top/--bottom, so the requirement probably wants to say there must be a path to correct it, not that every written field is settable. Otherwise it contradicts a decision you'd have to reopen.

And the highest-stakes case is already specified. ADR-0003 has the tool repairing a colliding key by appending — which rewrites a key, and a changed key breaks every inbound reference in the corpus. If exactly one write has to be loud, it's that one, and today nothing says it is.

✻ Worked for 36s

❯ ok let's move on to in progress. So this is where I'm a little torn or confused. I'm not sure if tasks are part of in progress and creating them is doing the work or if, in many cases, tasks should be formed ahead of in progress similar to how specs work. It's well established that creating a spec ahead of doing the work is how you have the best outcomes.

  Ran 1 shell command

The corpus already answers this, and the answer is both. when-a-work-item-splits says "discovery is what doing the work is for, and a work item that never gains a task was trivially scoped or nobody learned anything." WORK-0031 went from nine tasks to twenty-three, eleven of them found by doing it — and that record explicitly calls it "the reshape working as intended."

I think the dilemma exists because the spec argument is being applied to the wrong object. A spec describes what — the target. Tasks describe how — the route. Specs-before-work is well established precisely because the target shouldn't move. The route is supposed to change as you learn. So the argument transfers cleanly to outcomes, not to tasks:

Outcomes are the spec. Tasks are the route. Outcomes before the work — which you already believe, and already placed at preparing. Tasks planned before and revised during, which is a different claim and doesn't inherit the spec argument at all.

There is a floor, though, and it comes from your own definitions. prepared and todo both turn on mapping work to capacity, and capacity can't be assessed without knowing roughly what the work is. So enough breakdown to size it must precede todo regardless of what happens to the rest. That's the minimum, and it's derived rather than asserted.

And the before-versus-discovered ratio is measurable from data you already store — task created timestamps against the work item's move to in_progress. It's diagnostic in both directions: almost nothing discovered means the work was well understood or nobody was paying attention; almost everything discovered means under-prepared or genuinely exploratory. when-a-work-item-splits already discriminates these by timing rather than count — tasks still arriving late, after the work is understood, is one of its four causes.

So the shape of an answer: tasks before to size it, tasks during as discovery, and the timing of arrival is what tells you whether preparation worked.

One data point from right now: WORK-0059 is in_progress with zero tasks and seventy-odd journal lines. This session is currently running the most extreme version of the question you're asking — and it hasn't obviously failed yet, which is itself worth knowing.

✻ Crunched for 43s

❯ in many cases, the tasks are filling in as I spec. as well as a history. as well as to highlight attempts. places where we struggled. and opportunities for learning.

  Ran 1 shell command

The tool already agrees with you, and my "route" framing last turn was wrong.

luma-backlog task --help
An attempt at part of the work

That's the type's own definition — an attempt, backward-looking, not a plan item. Your reading is the specified one.

  Ran 1 shell command

And there's a real contradiction in the corpus, which is a good reason to feel torn. backlog-journal's What stays out lists "remaining work (tasks)" — forward-looking, a plan. The type definition says attempt — backward-looking, a record. Both are written down, in documents this project treats as authoritative. You're not failing to decide; the corpus says both things.

The actionable gap: a task cannot record why it ended. There's no task close — only set workflow_status=closed — and no disposition vocabulary at all. So if tasks are the history of attempts, the history loses the half that matters: this was tried and it didn't work. That's precisely the loss ADR-0007 gave work items four dispositions to prevent. Tasks got none, and nobody noticed because tasks were being read as plan items.

The struggle-and-learning half may not be the task at all — it may be the wave. spec.md §2.4: "an agent may attempt an outcome many times — attempt, probe, adjust — and those passes are transient", and §2.3 names the wave as the durable unit. WORK-0019 says it flatly: "until waves exist there is nowhere to hang the context an attempt had."

So there are plausibly three levels and only two are built:

- outcome — what must be true
- task — an attempt at part of the work
- the passes against an outcome — transient, and WORK-0019 wants them tied to the effort that produced them

That would explain why tasks feel overloaded: they're currently carrying planning, attempts, and the struggle record, and the specification says the third belongs somewhere that doesn't exist yet.

The cheap thing that unblocks the most: give tasks a disposition. It makes the history real without waiting for waves, and the vocabulary is already designed — ADR-0007 did the thinking.

✻ Cogitated for 59s

※ recap: You're walking through the backlog-move procedure to improve it while running an experiment on ad hoc "editor mode" work, journaling every turn into WORK-0059. You've interviewed through in_progress; next is closing, then reopening.

❯ capture the task disposition gap

  Ran 1 shell command

No existing record — only journal mentions in three work items.

  Ran 2 shell commands

○ WORK-0064 · A task cannot record why it ended — captured, kind change

Captured with a written body, not a template. Four things in it:

The gap — close is registered on work-item only, so a task ends with set workflow_status=closed and carries no disposition. Succeeded, tried and abandoned, overtaken, found unnecessary — all identical afterwards. ADR-0007's own test applies unchanged: the enum carries what the record cannot.

Why it went unnoticed — tasks have been read as plan items, and under that reading a disposition is pointless: a plan item is done or not done. Under the attempt reading it's the most valuable field on the record. That's the two-document contradiction from a minute ago doing real damage, not just being untidy.

What it blocks — spec.md §4.6 succession says "when an attempt does not succeed and another begins, a new record is created and links back." That presupposes the corpus knows an attempt failed. Nothing does. WORK-0024 is the other half of the same gap: it can't create the successor, this can't mark the predecessor. Neither record mentions the other today.

Why it's cheap — WORK-0019 wants the rich version and is blocked on waves, which are specified and unbuilt. A disposition on a task needs none of that, and the vocabulary is already designed. The open question is only whether a task wants ADR-0007's four, fewer, or different ones — and the thing to avoid is inventing a second unrelated vocabulary.

Back to in_progress whenever you're ready — we got as far as what tasks are for, and haven't touched what the move itself means.

✻ Worked for 59s

❯ ok let's move on to in progress

EDITOR — move 5: todo → in_progress. The document gives this three lines total: "says somebody is on it now… a record left in_progress across weeks is lying about what is happening." That's the whole section, and it's the rung where all the work actually happens.

1. Does entering it require an assignee? It's a claim that somebody is on it — and the record has no way to say who. ADR-0008 settled that taking a task expires and owning a work item doesn't, and WORK-0031 records that take, release and steal don't ship. So the rung's central claim has no subject. Invisible solo; not invisible the moment agents are doing the work and you want to know which one.

2. Can several things be in_progress at once — and should the move say something? backlog-rundown already carries an opinion: "Several things in_progress is itself the finding. More was started than gets finished." If that's true, the warning belongs at the move, not only in a report nobody may run.

3. What actually makes it stale? The doc says "across weeks", but when-a-work-item-splits is explicit that this project measures in sessions — "Rate is the sharper number, and the unit is the session… Count sessions, not elapsed time", because a work item here can be created, worked and closed inside one evening. Two documents, two clocks. Which one governs a stale in_progress?

4. Does in_progress have internal structure, or is it flat until close? §2.3 names the wave as the durable unit for passes at an outcome, and it's unbuilt. You described tasks accumulating as history and struggle. Is there a meaningful inside to this rung, or is everything that happens there recorded on the children rather than the rung?

5. What happens when work stops but isn't finished? Sending work back covers prepared → preparing and todo → prepared. It says nothing about leaving in_progress. Paused work has nowhere to go — you either lie by staying in_progress, or close it as something it isn't. Given you said the claim is about the present, this seems like the gap that would bite first.

✻ Worked for 26s

❯ an assignee should be required,
  and paused/blocked will need its own status, because you can get blocked in any workflow status, not just in progress (although this has little to do with in progress, it's just where it will happen the most often)
  in progress should force stage: provisional
  im not sure if there should be some kind of analysis done to vet the outcomes before progress starts
  maybe that should be part of it, it shouldn't block but it should definitely strongly encourage you to stop and vet them if they aren't already in good shape (stable

  Ran 1 shell command

Blocked being possible at any status means it isn't a rung — it's a modifier. Every other status is a position on a line; blocked co-occurs with any of them. Make it a status and you lose where it was blocked from: a blocked preparing record and a blocked in_progress one become indistinguishable, and returning means remembering something the record no longer holds.

And there's a concrete cost, from a decision in force. ADR-0005 puts the status ordinal in the rank prefix, and every status change re-enqueues at the back. So a record that goes blocked and comes back loses its queue position twice. That's not a modeling preference — blocked-as-status actively destroys ordering somebody chose.

It isn't a spec.md §5.2 condition either: conditions are computed from the record, and blocked is declared by a person. So it's a third kind of thing the model doesn't have — a declared modifier — which is a real design question rather than a missing enum value.

Assignee-required is unimplementable today. ADR-0008 separates taking, which expires, from owning, which doesn't — and take, release and steal don't ship. The rule needs the verb before it can be adopted, so it's a dependency rather than something to write into the document now.

And it's the first real test of this morning's policy. WORK-0059 is in_progress right now with no assignee. Adopting assignee-required means choosing backfill or an explicit grandfather clause — the rule being proposed in the same conversation that produced the policy about adopting rules. That's a good sign the policy was worth writing, and a slightly uncomfortable one.

Your uncertainty about vetting outcomes is already answered by your own escalation pattern. Tolerate → warn → force, from two turns ago. Strongly encourage without blocking is the warn step. So this isn't a missing decision; it's an instance of one you already made — which is what the pattern was for.

One argument for putting the prompt exactly there, stronger than quality: before work starts, revising an outcome is free. After it starts, revising is Redefine — governed, and the place goalposts move. So entering in_progress is the last cheap moment to change an outcome. That's a structural reason for the nudge to live at this move and nowhere else.

✻ Cogitated for 46s

※ recap: You're walking through the backlog-move procedure to improve it, interviewing rung by rung and journaling everything into WORK-0059. You've finished in_progress. Next is closing, then reopening, and then the document rewrite.

❯ no way, when somehting is blocked, it stays where it's at, and when you unblock it, it stays where it's at.

Then we agree, and my objection dissolves — the rank cost only existed if blocking moved the record. It doesn't. So it's a flag on the record, not a rung on the ladder.

  Ran 1 shell command

Three consequences of it being a flag, all cheap to note now and expensive to discover later.

workflow_status stops being the whole answer to "where is this." A board or listing built on status alone shows a blocked item sitting in In Progress looking perfectly active. So every rendering has to carry the flag or it misrepresents the work — and that's a change to showing-records, not just to the model.

There's no mark for it. The set is ○ ◐ ✔ ↪ ⊘ ✘, and all six describe progress or an ending. Blocked is neither — it's a suspension of progress at a point. It needs a seventh mark, or a modifier drawn alongside the state mark rather than instead of it, which is a different rendering rule than anything in that policy today.

Blocked has to name what it's blocked on, or it's a flag with no route out. And you already built that vocabulary earlier today: an external work item, a team, or a deliverable. Same three kinds, same different resolution behaviors — a work item has a checkable state, a team has none, a deliverable exists without being adequate.

One useful asymmetry falls out of that: blocking is declared by a person, but unblocking can be observed. If something is blocked on WORK-0022 and WORK-0022 closes, the flag is stale — and that's computable as a §5.2-style condition. So the tool can notice a block has cleared without anyone telling it, which is the half that would otherwise rot.

in_progress → closed next, which is the only move the document already treats as heavy.

✻ Brewed for 35s

❯ maybe for glyphs block could use the same icons but turn from circles to squares

Elegant — shape carries blocked, fill still carries progress, so the flag doesn't cost you the state underneath.

○ WORK-0022 · not started          □ WORK-0022 · not started, blocked
◐ WORK-0031 · under way            ◧ WORK-0031 · under way, blocked
✔ WORK-0018 · delivered

  Ran 1 shell command

Four notes, and the first makes it cheaper than it looks.

Only two of the six can ever be blocked. ✔ ↪ ⊘ ✘ are all endings, and a finished thing can't be blocked. So the square family needs exactly two members, not six. It's a two-glyph addition, not a doubling of the vocabulary.

It extends the system's own logic rather than bolting on a second one. showing-records already uses shape-family for category and fill for progress within it — "unfinished work shares the circle; finished work changes shape." Squares add one more family under the same rule, which is why it reads as obvious rather than clever.

One thing to check before adopting: font coverage. That policy already made a font-driven call — the heavy ✔ and ✘ over the light pair, because "the light marks are thinner than the circles beside them, so a finished row reads as fainter than an unfinished one." ◧ (U+25E7) has noticeably worse coverage than ◐ (U+25D0), and a mark that falls back to a box is worse than no mark. Worth eyeballing in a real terminal before it's written down.

And shape is the right carrier rather than colour, which is worth stating explicitly since WORK-0043 already has colour the state marks as a task. Colour dies in a pipe, and a blocked record written to a file still has to read as blocked. Shape survives everything; colour can reinforce it.

Process catch: showing-records says its marks mirror command-line-interface's policy/ascii-styleguide, and that the section should be replaced by a pointer once that bundle is adopted at 0.3.0 or later — which is WORK-0056. So a new mark belongs upstream in the style guide, not added to a vendored section already scheduled for deletion. Otherwise the re-adoption silently drops it.

✻ Worked for 43s

❯ capture the glyph proposal and the upstream note

Blocked itself was never captured — the whole model came out of this conversation. Folding it in as the parent rather than making the glyph record point at nothing.

  Ran 2 shell commands

○ WORK-0065 · Blocked is a flag rather than a rung — captured, kind change

The glyph proposal is in it as a section, with the rendered pair, the two-glyph argument, the shape-over-colour reasoning, and the font-coverage check as an open item rather than a settled choice.

The upstream note went on WORK-0056, where it's actionable — a new mark added to the vendored showing-records section is silently dropped by the re-adoption that replaces it with a pointer. That record is the one that would cause the loss, so that's where the warning has to sit.

Two things I added that you didn't ask for, both flagged as open rather than decided:

Every listing has to render the flag or it lies. A board built on workflow_status alone shows a blocked item sitting in In Progress looking active. That makes this a change to how records are shown, not only to what they hold — which is a wider blast radius than the model change suggests.

Paused and blocked may not be the same thing. You raised them together — "paused/blocked will need its own status" — and they behave differently: blocked is external and waiting on something nameable; paused may be a choice with no blocker at all. Nobody has separated them, and the model above only covers blocked. Left as an open question on the record rather than resolved.

in_progress → closed is the last move, unless you want to double back.

✻ Cogitated for 50s

❯ let's move on to closed
  so closed shoudl make things stage: stable, since they are not supposed to change much after being closed
  and there should be some checks and balances
  all outcomes should be successful if they are trying to mark the work item successful
  all the outcomes should be proven
  all the tasks should be completed or superceded, and the superceded should be completed (unless it's part of another work item)
  we should do some journaling, and we should give guidelines on what's worth capturing

  Ran 1 shell command

Your first two checks are already decided and in force — ADR-0007 is stricter than I'd have guessed.

▎ work-item close <ref> completed is refused unless at least one live outcome exists and every one is proven. Unreadable, unverified, disproven and inconclusive are all the same answer: not proven.

  Ran 1 shell command

So outcomes need no new rule — they need backlog-move to say the rule, which it doesn't. Two details of it worth knowing:

It gates on verification, never on the assertion — because gating on the doer's claim would gate on the thing the design distrusts, and would let a doer close their own work. So "successful" means the checker's verdict, not the doer's. And --force is already designed to not touch outcomes: they stay as they are, completion still computes two of five, and the work item carries the forced close — so a reader sees a completed work item whose own arithmetic contradicts it, which is the truth.

On tasks I'd push back. spec.md §2.4 says a work item is judged on its outcomes and on nothing else. Requiring every task completed contradicts that — you can prove every outcome while abandoning half the attempts, and that's success, not a defect. Attempts that failed are exactly what tasks-as-history is for. I'd make it a warning, not a refusal.

And it's unimplementable today regardless. WORK-0064: a task has no disposition, so superseded can't be expressed. WORK-0024: succession can't create the successor, because idempotency wins over §4.6. Your transitive check — that the superseding task is itself completed — needs both to land first.

Your checks correctly apply only to completed. ADR-0007 has canceled, rejected and superseded closing freely and deliberately: "gating cancellation on completion would make it impossible to stop work because it was unfinished, which is the usual reason." That also answers closing something never started — reject it and nothing is checked.

One question on stage: stable — does it reach the outcomes? An unproven outcome marked stable would assert settledness about something unresolved. And a canceled work item at stable is settled without being successful — which is fine if stage tracks how settled the content is rather than how it ended, but that's worth saying out loud since the two readings diverge exactly here.

The journaling guidance already exists, and it's good — backlog-journal's test is relitigation risk: "anything that should not have to be argued a second time." What's missing is the close-specific version, which is retro material: what the estimate got wrong, whether the outcomes were the right ones. That's WORK-0051, still captured with an empty body.

✻ Worked for 1m 11s

❯ if people want to mark a work item successful without all of these thigns, we should make them "force" it
  and if they want to close it as canceled, then they should give a reason for canceling it and that should be recorded

Force already exists and is carefully designed — ADR-0007's --force leaves the outcomes untouched, so completion still computes two of five and the work item carries the forced close. A reader sees a completed record whose own arithmetic contradicts it, which is the point.

  Ran 1 shell command

On requiring a reason — the same argument reaches rejected. The four dispositions split by whether they point at evidence:

┌────────────┬────────────────────────────────────┐
│            │          evidence for why          │
├────────────┼────────────────────────────────────┤
│ completed  │ the outcomes                       │
├────────────┼────────────────────────────────────┤
│ superseded │ the link ADR-0007 already requires │
├────────────┼────────────────────────────────────┤
│ canceled   │ nothing but the prose              │
├────────────┼────────────────────────────────────┤
│ rejected   │ nothing but the prose              │
└────────────┴────────────────────────────────────┘

So if canceled needs a recorded reason, rejected needs one by identical logic — we declined this with no why is the same loss as we stopped this with no why. And the other two don't, because the record already carries their reason structurally.

A forced close needs a reason more than a cancel does. Forcing overrides a refusal — the most questionable action the tool offers. ADR-0007 makes the override visible in the arithmetic, but nothing anywhere records why it was forced. That's the one place a reader can see something was overridden and has no way to learn whether it was justified.

And that's exactly WORK-0061's event. A forced close is the went-around-the-system moment the observer mechanism exists to surface, and the required reason is its evidence. Those two threads connect rather than being separate features — which makes the reason field worth more than it looks.

One dependency: the prose reason doesn't exist yet. In the shipped binary --reason is the disposition. ADR-0007 makes the disposition positional and --reason prose — WORK-0031's make-close-take-its-disposition-positionally, still todo. So this rides on that task landing.

That's the ladder end to end. Want me to write the interview up into backlog-move now, or keep going on the backward moves and reopening?

✻ Cogitated for 38s

❯ ok sure, all the tasks should be resolved but they don't need to be successful.

  Ran 1 shell command

❯ the should not be open and ready to start

  Ran 1 shell command

"Open and ready to start" names the failure precisely — a task sitting at todo under a closed work item advertises available work that nobody can pick up. That's the honesty problem, not a completeness one, and it's a much better justification than "finish everything."

It splits into a part you can build now and a part you can't. No task is left open is checkable today. Every task says how it ended needs WORK-0064 — right now a task closes with set workflow_status=closed and "resolved" is indistinguishable from "resolved successfully."

Warn rather than refuse, on your own tolerate → warn → force ladder. spec.md §2.4 says a work item is judged on its outcomes and on nothing else, so a second hard gate for something the model says doesn't determine done-ness would contradict the refusal surface we just agreed — close … completed without proven outcomes, and axis-5.

And one trap worth naming, because ADR-0007 already hit it one level up. Do not auto-close dangling tasks when the parent closes. That invents a disposition nobody chose — exactly what ADR-0007 refuses to do with outcomes: "The tempting implementation marks outcomes verified so the arithmetic comes out clean. That destroys the record." Same shape, same answer: warn, and let a person say which it was.

262 journal lines. That's the ladder covered end to end — want the write-up into backlog-move now, or the backward moves and reopening first?

✻ Cooked for 42s

❯ rejected is for stuff in captued, once it's started it becomes a thing you cancel.
  you only reject it when you first look at it before moving it through the pipeline

That makes ADR-0007's distinction checkable, which it wasn't before.

  Ran 1 shell command

ADR-0007 defines the split by intent — rejected is "never a consideration", canceled is "yes, then changed our minds". That requires somebody to introspect honestly. Yours is a field lookup.

And the two are the same statement. Crossing the first gate is the act of considering it — that's what captured → unprepared means. So never a consideration and never crossed the first gate say the same thing, one as intent and one as an observable fact. That's a strict improvement: same distinction, checkable test.

Which makes the choice derivable. The tool can tell which disposition applies from where the record sits, so it can suggest the right one rather than accept a wrong one — nobody has to hold the rule in their head. And a rejected on something at preparing isn't a judgment call gone wrong, it's a factually false claim about history, which is a different class from a policy preference and argues for correcting it rather than allowing it.

This is an amendment to a decision in force, not a contradiction. ADR-0007 is provisional, and its disposition table would want the positional wording. Worth doing deliberately — the intent wording is what somebody would otherwise keep reasoning from.

One edge the rule has to answer: a record that crossed to unprepared, came back to captured, and is then closed. Current position says rejected; history says it was considered. Git holds the answer either way, so the rule needs to say whether it reads the current status or whether the record ever crossed. I'd guess ever-crossed, since that preserves the meaning rather than the bookkeeping — but it's the kind of thing that gets decided by accident in the implementation if nobody says.

✻ Crunched for 42s

❯ rejected means it came in, and you never put it intp unprepared, you denied it.
  that is different than attempting it and then deciding you don't want to continue.  that's canceled.

That answers the edge case I raised — "never put it into unprepared" is a claim about history, not about where it sits now. So a record that crossed, bounced back to captured, and then stopped is canceled. Ever-crossed governs, not current position.

  Ran 1 shell command

Two things fall out of "you denied it" that are worth having on the record.

Denial has an actor in a way cancellation needn't. You denied it implies somebody with standing to say no. Cancelling can be circumstantial — priorities moved, funding went, the thing it depended on died. So rejected carries an authority claim the other three dispositions don't, which matters the moment more than one person can close things.

And that closes a loop from the first gate. I asked back at captured → unprepared who's allowed to cross, and we moved on without settling it. Rejection is that same question from the negative side: whoever can cross the first gate is whoever can deny at it. One answer covers both, so it's cheaper to settle than it looked.

Structurally, rejected is the first gate's "no" recorded as a close — which is why it sits oddly beside the other three. Those end work that was underway; this ends work that never started. It's the ladder's only gate-refusal, and it's expressed as a terminal state rather than as an outcome of the gate. That's probably fine, but it explains why it kept needing a special rule.

That's the disposition vocabulary settled. Backward moves and reopening are the last unexplored corner of the document — or I can write up the interview now, which is a fair amount of material.

✻ Cooked for 32s

❯ the last move is reopen, something that happens when its closed in any way
  i wonder if things that are reopened should always create a new one, probably not, but it's something to consider later.
  for now we shoudl allow reopen to mutate,
  and then we probably need to clear some things.
  stage should to back to draft or provisional, definitely not stable
  i think we should leave any closed data on the object, don't wipe it away
  we should journal why it's getting reopened
  we should ask where it's getting reopened into because you can reopen into unprepared -> in progress.  but we should try not to skip steps so we shoudl recommend unprepared, todo, or in progress
  we might end up with a new owner, which makes me think owner should become an array, and the most recent entry is the owner

  Ran 1 shell command

Your "leave closed data on the object" is already how it works. ADR-0007 makes closed an append-only list of {by, at, as} — so reopening and re-closing appends a second entry and the whole close/reopen history is native. Nothing to preserve by hand.

And owner-as-array has a written precedent with the exact argument. WORK-0019 records that asserted was originally a single field overwritten on each attempt, and became a list because "a replaced field could not represent a second attempt and would have been the only place in the design that destroys history." That transfers to owner unchanged — it'd be the fourth instance of the same pattern after closed, asserted and verified, which is a strong sign it's right rather than novel.

On "we need to clear some things" — there's a principle that decides it without enumerating fields:

▎ Clear what claims the present. Keep what records the past.

stage is a claim about how settled the content is now, so it resets. closed, verified and owner are accounts of what happened, so they append. That covers every field including ones nobody's thought of yet.

Which answers a question hiding in the reopen: does it invalidate verification? No. If scope grew, the old proofs still stand. If the work turned out not to be done, then the proof was wrong — and ADR-0007 already has a verdict for that: disproven. So you append a new verdict rather than clearing an old one. Same logic, and it keeps the fact that somebody once believed it was proven, which is exactly what a reopen wants to know.

Your three re-entry points are coherent, and worth saying why so they don't read as taste: unprepared is the head of the preparation pipeline, todo the head of the work pipeline, in_progress the resume-now case. preparing and prepared are excluded because they're states you arrive at by working, not states you declare.

WORK-0020 is captured and is the home for all of it. And backlog-move already holds the mutate-versus-new-record judgment — "same work resuming keeps the history and the journal; new work that happens to rhyme should be its own record linking back" — which is consistent with allowing mutation now.

That's the whole ladder. Say the word and I'll write it up.

✻ Cooked for 50s

❯ yes and we should probalby allow preparing, because forcing someone to go to unprepared when they want to be preparing, that's the kind of extra steps that make people feel like the tool is too much red tape

You're right and my framing was too clever. Forcing unprepared on someone who is actively shaping it makes the record less true, not more disciplined — it says "queued for preparation" while preparation is happening.

  Ran 1 shell command

Which means the rule is the one that's been running through this entire interview: the status must be true. Any rung whose claim holds right now is a legal reopen target — including prepared, if nothing about the preparation actually changed. Closed as canceled for budget, budget comes back, the preparation is still good: prepared is the honest answer and forcing three moves to get there would produce three lies on the way.

Your red-tape complaint and the honesty test are the same principle, which is worth naming because it's the fourth time it's shown up today:

▎ Every step the tool forces should prevent a false record. A step that only enforces process is the red tape.

That's what's underneath observed, never refused; underneath force writes the field rather than refusing the move; and underneath the deliberately narrow refusal surface — close … completed without proven outcomes, and axis-5. Each one was decided separately today and they're all the same rule.

One real risk with reopening straight to prepared: the preparation may be stale. But staleness is observable from the record's own dates, so it's a warning rather than a block — tolerate → warn → force, again.

So the recommendation of unprepared / preparing / todo / in_progress stays useful as help, not as a gate. It's there so somebody who hasn't thought about it lands somewhere sensible, and anyone who has can go where the truth is.

Ready to write the whole thing up into backlog-move — it'll be a substantial rewrite rather than an edit, given how much of this the file doesn't currently say.

✻ Cooked for 40s

❯ very true: good point,  Closed as canceled for budget, budget comes back, the preparation is still good: prepared is the honest answer and forcing three moves to get there would produce three lies on the way.

  ok write all this up please
  make sure you include the full conversation, there was a lot in the beginning, middle, and end of this session that has massive value

Writing it up. backlog-move first — this is a rewrite rather than an edit, since most of the interview isn't in the file at all.
Interrupted · What should Claude do instead?

❯ wait commit what we have before doing the full write up review

  Ran 1 shell command

Committing before the rewrite. First the version entry backlog-move's corrections never got.

  Ran 6 shell commands

Committed, working tree clean.

3b4762b Capture what the ad hoc experiment turned up
31d7391 Make the procedures say what the decisions already settled
ddd8a4f Capture four findings that were noticed and never written down

Two commits, split by story rather than by file. The first is the bundle — the rundown rename, risks as bullets, the new adopting-a-rule policy, and backlog-move no longer teaching delivered/abandoned. The second is the corpus — WORK-0057 through WORK-0065, plus journal lines on WORK-0043, 0046, 0050 and 0056.

One thing I fixed on the way in: backlog-move's corrections had never been versioned. The rename was 0.17.0 and the policy 0.18.0, but the disposition fix and the Always true section were sitting unversioned in the working tree — which is exactly the drift this bundle polices. Bumped to 0.19.0 with an entry covering them, including the note that one of the three invariants exists only because I reported it as a defect and the code turned out to be right.

Both messages say why rather than what, per the convention. 284 journal lines on WORK-0059.

Three commits ahead of origin/main, nothing pushed — say the word if you want that, or a pull request.

Now the write-up. Given how much there is, I'd suggest it lands in two places: backlog-move rewritten from the interview, and a topical synthesis of the cross-cutting findings — the principle that every forced step should prevent a false record, the refusal surface, tolerate/warn/force, the blocker-narrowing frame. Those recur across rungs and would be buried inside any single section of the procedure.

✻ Crunched for 1m 25s