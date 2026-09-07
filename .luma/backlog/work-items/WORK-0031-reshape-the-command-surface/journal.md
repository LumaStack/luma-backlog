# Journal — Reshape the command surface

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-07

split at 23 tasks — five reading tasks went to WORK-0043 and rank repair to WORK-0022; what stays is shapes a decision already settled, what left was found by writing the procedures
the test that separated them: does a record say what this should look like? the reshape tasks all cite an ADR or this index; the five that left cite nothing because nobody had tried reading a backlog from a terminal before the procedures did
the rule for splitting is mostly not about splitting — three of the four reasons tasks keep arriving are defects in the outcomes, each already a condition in spec 5.2, and each fixed by editing one outcome
revising an outcome is not a smell: lifecycle 2.8 has a Redefine phase and work-item.drifted fires when work happens and no outcome is verified or revised — what is a smell is expansion at every boundary without governance
an outcome that can never be finally true is a standing condition in the wrong place, not a bad outcome — spec 5.2 is where those belong, and moving it beats splitting the work item around it
claimed across four reports that two outcomes were provable and never ran their checks — running them showed none of the four passes; the claim came from counting closed tasks, which is the substitution the outcome model exists to prevent
every-command-is-noun-then-verb fails on two checks, both for the same reason: show, set and list are top-level and none is in ADR-0006's verb-only list, and show and set are not parameterized by noun — the deferred cross-type decision is blocking an outcome, not just a task
a-record-is-addressed-by-the-path fails on emission, which is known, and on placement — scoped() went into internal/corpus because Resolve was already there, but the outcome says internal/app so every surface gets the same forms; that was decided by convenience and never noticed
this is work-item.drifted in the flesh — work happened, no outcome was verified or revised, and progress was reported from task counts; the condition that would have caught it is the one written as a task in the same session
the lesson is narrow and mechanical: an outcome is not provable until its verify_by has been run, and saying it is provable is itself a claim that needs the same evidence — read verify_by before reporting on an outcome, not after
redefined every-command-is-noun-then-verb: the old check asserted over every universal verb, a set nobody had settled, so no amount of building could satisfy it — a whitelist of required pairs is closed and cannot grow silently as verbs are added
six of ADR-0006's seven verb-only commands do not exist, so the old check demanded work this item never owed; each now has its own work item and three of them are inquiries because whether they should exist is genuinely open

## ▶ 2026-09-06

help presentation follows gh, recorded on ADR-0006 rather than as a new decision — the record already holds the choices clig.dev leaves open, and this is one
docker was considered and rejected as the model: [OPTIONS] COMMAND advertises global options, and this tool has none in the spec, none registered in code, and --help/--version terminate rather than combine
the synopsis is already one line and tested; it gains <subcommand> when the tree nests, and [command] must stay bracketed because a bare call is real
two things ADR-0006 adopted on 2026-09-05 were never built — help text leading with examples, and a documentation link in top-level help; they are gh's EXAMPLES and LEARN MORE and are now on the restyle task
flag shorthands deferred until every command in the reshape exists — a letter scheme designed against a partial flag set gets redone; re-open when the last command lands and the full flag set can be seen at once
no test asserting identical flags across nouns: flag sets legitimately diverge (--kind is work-item only, --project decision only), so such a test would be weakened until it asserted nothing
the noun-verb split touches only internal/cli — app, corpus and record already take the record type as a string, so the noun moves from a positional argument to a constructor parameter and nothing below the adapter changes
cross-type listing has no home in noun-verb — spec 9.2 makes list per-noun and ADR-0006's verb-only set omits it, yet 18 of 21 list call sites in the suite are bare list; deferred rather than decided mid-refactor
while it is deferred the noun commands carry no list verb and top-level list is untouched — two ways to list one type would be a worse state than one inconsistency
the containment guard caught the first attempt — naming a record type in the adapter meant importing internal/corpus, which ADR-0004 forbids; the unit names are now re-exported through internal/app, which is what its failure message tells you to do
a noun command needs Args NoArgs or cobra takes an unregistered verb as a positional, prints help and exits 0 — task close looked like it worked; the same trap root.go already documents, one level down
cobra honors --help before validating arguments, so a negative test must probe without it: task close --help exits 0 whether or not close exists
show and set stay top-level with list, not for tidiness but because they reach records that are not creatable units — PROJECT.md is type luma/project and corpus.Units is what can be created; scoping them to a noun would strand it
key-scoped references resolve — WORK-0031/tasks/<slug>, the name form, the slug half, and with or without .md; scoping is tried only after exact matches and before prefix matches, so it adds a form without loosening any existing one
emission is the half still outstanding: the outcome asks for key-scoped paths to be emitted as well as accepted, and changing the path field in --json is a breaking contract change where adding a ref field is additive — not decided yet
replace-move-with-work-item-rank was one task holding five: ADR-0005 needs config to carry explicit sparse ordinals, a bisection, the command, a set refusal, and re-enqueueing on every status change — split so the size is visible rather than discovered halfway
workflow_status now accepts a mapping of status to ordinal as ADR-0005 asks, and still accepts a list, deriving ordinals ten apart — every config written before this is a list, including this repo's own, so refusing it would have been a migration nobody asked for yet
a mapping's ordinals are read from the yaml node rather than the decoded map, because ranging a Go map loses the document order that IS the ladder order
ordinals must ascend with the ladder and are refused at load if they do not — a ladder that reads one way and sorts another would put the board in an order nobody wrote down
the ordinal prefix has to be zero-padded like the position, which neither ADR-0005 nor spec 9.6 says — without it ordinal 100 sorts before 20 as text and the prefix that exists to make sorting work is the thing breaking it
positions use big.Rat rather than float — bisection halves, halving a finite decimal stays finite, so every value is representable exactly and precision extends by one digit at a time instead of rounding
spec 9.6's worked table is now a test, so the document and the code cannot drift; the squeezed row was the one worth having, since it is where precision extends past three decimals
listing every record type interleaved was never wanted — it was what the code did when the type was an argument; list is now work-item list, because the tool is called backlog and listing the backlog means listing work items
the tree filter narrows work items and never their children — asking for prepared work items and seeing only their prepared tasks would hide the ones nobody has started, which is usually why you looked
tree assembly lives in internal/app, not the adapter: the board wants the same thing a terminal does, and an adapter building it would be carrying judgment
--tree on list rather than a tree verb: a flag on a settled surface instead of a sixth verb needing ADR-0006 amended, and the single-item case was already served by task list -w X
help examples went stale the moment the tree changed and nothing caught it — journal's said luma-backlog journal, a command that no longer exists; worth a check that every Example in the tree actually parses
listings show key, status and title — TYPE repeated the command it was run under and the slug repeated the title, so both went; flat output went from 95 columns to 55
children are marked OUT and TASK rather than named: they have no key and their slug is their title in kebab case, so printing it put the same sentence on the row twice
tasks before outcomes — an outcome reads unverified for nearly the whole life of a work item, so leading with them puts constant text between the work item and what somebody came for; the argument that outcomes summarize status only holds at closing time
child keys were considered and dropped: OUT-014 is not something anyone says out loud, which is the reason work items have keys at all — recorded rather than left as a gap somebody rediscovers
rank without sorting is invisible — the command wrote a field nothing read, so listings now sort by it: ranked records in rank order, unranked last because a rank is a position somebody chose and an unplaced record should not outrank a considered one
records with no rank are skipped as neighbors rather than seeded — seeding a whole status on first rank would make a reorder a multi-record write, which is the thing decimal ordering keys exist to avoid
ranking against a record at another status is refused: rank orders within a status, so there is no position between them to compute
reversing a listing by default was rejected: piping gives the wrong end silently, and the board and terminal would render one request in opposite orders — the pager ADR-0006 already adopted is the answer to the scrolling it was meant to fix
rank is the default sort and stays it — a listing that does not follow work order is not a backlog; --sort covers created and updated, which are the other things people reorder by
first shorthand collision: -r is the listing convention for reverse but close already holds it for --reason; recorded on ADR-0009 as input rather than assigned, since the letters are deferred until every flag exists
settled on --reverse over git branch's --sort=-key: the sign-prefix form exists for multi-key orderings and we have three keys and one direction
direction goes in the sort key as --sort=-updated, not a separate --reverse: two flags for one ordering is the weaker model, and it extends to multi-key later without a second mechanism
the parsing objection against --sort=-key was false and nearly decided it — pflag takes the next argument as a value whether or not it starts with a dash; checked rather than assumed, after asserting it the other way
--reverse is not being built, so the -r collision recorded on ADR-0009 an hour ago dissolved; the note is corrected rather than left, since stale guidance in a record in force is worse than none
status and rank are written by one operation in internal/app that no caller can bypass — set routes a workflow_status assignment into it and close calls it rather than setting the field, so there is no path that produces a record where the two disagree
drift from a hand-edited vocabulary is observed and never refused, in the shape list already uses for skips: it names the record, the ordinal the status carries now, and the rank that says otherwise
ascending ordinals give ladder order, so a flat listing leads with captured and trails with closed — right for a board with columns, wrong for a terminal where the least actionable work ends up on top; belongs with --group rather than here
rank repair is not built — ADR-0005 names it as the third mechanism and the message tells you to re-set the status instead, which does the same thing one record at a time
the proposed grouping put decision, exploration, outcome and task under MANAGE WORK ITEMS while work-item sat in CORE — they are sibling record types, so the split is by what a reader wants: a verb, a record type, or the machinery
cobra's IsAvailableCommand is false for the help command by design, so any custom listing must say or eq .Name help or lose it silently — cobra's own template does; dropping the clause is how help vanished
InitDefaultHelpCmd declines to add help to a command with no subcommands yet, so calling it before registering them did nothing — it has to come last
the stale-example check is built and earned its keep immediately: the first version used cobra Find, which returns the root with arguments unconsumed rather than erroring, so it passed a command that does not exist
nine tasks were finished and none marked closed — the backlog said todo for everything while the work was in the commits; caught only by asking what was next
there is no task close: close is registered on work-item only, so a task is closed with set workflow_status=closed — worth deciding whether that is right or a gap
writing the procedures found four things the commands cannot do — new takes no --description, show gives fields but not outcomes or tasks, there is no --column though the config defines them, and nothing asks for everything except closed; that is the bootstrap order working rather than a detour
the two procedures that existed predated the binary and neither called it — backlog-journal still said there is no binary yet, and backlog-new carried a frontmatter copy telling people to write workflow_status: idea, a value the ladder lost at ADR-0002
backlog-journal's fixed entry template contradicted the spec it cited: 5.5 says headings are named after what they settle, not drawn from a template, and the command writes a bare date heading on purpose
