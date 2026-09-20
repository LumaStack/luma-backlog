---
type: policy
type_version: "0.0.1"
title: Command line style guide
description: The two templates a command's help follows — the bare tool and a command of its own — explained a piece at a time, plus the output conventions that go with them.
matches: eager
---

# Command line style guide

**Follow this and a new command looks like every other one.** The shape is
`gh`'s, because it is the one most people have already read.

There are two templates. The bare tool is what somebody sees first and is
mostly about *choosing*; a command's own help is about *using* it. They share
their parts and not their emphasis.

## The bare command

What `<tool>` alone prints, and `<tool> help`.

```
<tagline>

USAGE
  <tool> <command> <subcommand> [args]

CORE COMMANDS
  verb              What it does

<OTHER GROUP>
  verb              What it does

GETTING STARTED
  verb              What it does

FLAGS
  --help            Show help for command
  --version         Show version

EXAMPLES
  $ <tool> <command>

LEARN MORE
  <where to go next>
```

**A section with nothing in it is left out, not left empty.**

### The tagline

**One line saying what the tool is for, not what it does.** It is the first
thing anybody reads and often the only thing they read, so it is positioning
rather than description.

**Use the project's own words.** Whatever the README opens with belongs here,
word for word. Two copies of a positioning line drift apart, and the one in the
terminal is the one that gets seen and not maintained.

No command name, no version, no `usage:`.

### `USAGE`

```
USAGE
  <tool> <command> <subcommand> [args]
```

**The grammar, not an example.** Angle brackets for what the reader supplies,
square brackets for what is optional. Exactly one line.

At the top level it exists to teach the shape — that commands take
subcommands, that arguments come last. Real invocations go in `EXAMPLES`.

### `CORE COMMANDS`

```
CORE COMMANDS
  verb              What it does
```

**Core means reached often, in the ordinary loop of using the tool.** Not
important, not powerful — *frequent*. The test is whether somebody who uses the
tool weekly would type it in a normal week.

Three things keep a command out of core:

**It is run once.** Setting a project up is not part of the loop, however
essential it is.

**It is a different kind of act.** Governing what an agent may do is not the
same activity as managing knowledge, even when both are frequent.

**It is maintenance.** Migrations, repairs and one-off fixups belong in their
own group or nowhere.

**Order core by frequency, not alphabetically.** Alphabetical order serves
looking up a name you already know, which is the case a menu is least needed
for. The first two rows should be the pair somebody runs together most.

**This is a judgement, so make it a readable one.** An agent or a newcomer
choosing a command reads this grouping as a recommendation about where to
start — which is what it is. If a command is in core and nobody runs it, the
grouping is wrong rather than the reader.

### Groups beyond core

```
<OTHER GROUP>
  verb              What it does
```

**Name the group for what the commands are, not for what they are not.**
`RESTRICT` says what the section does; `OTHER COMMANDS` says only that somebody
gave up grouping.

**One command is a legitimate group** when it is genuinely a different kind of
act. A governance surface sitting alone under its own heading is clearer than
the same command buried among knowledge commands, and the heading is what says
so.

**Two or three groups total is usually right.** Past four, the grouping is
carrying more than a reader can hold and the tool probably has a shape problem
rather than a help problem.

### `GETTING STARTED`

```
GETTING STARTED
  verb              What it does
```

**What somebody runs before they have anything.** Setup, and `help` itself.
These are the commands whose whole audience is people who have not used the
tool yet.

It goes near the bottom despite being first in time, because the people
reading a menu repeatedly are not the people who need it.

### `FLAGS`

```
FLAGS
  --help            Show help for command
  --version         Show version
```

**At the top level, list only the flags that work with no command at all** —
in practice `--help` and `--version`. Anything that needs a command belongs in
that command's help.

**Describe what the flag changes, not what it is.** Short forms first where
they exist: `-g, --global`.

### `EXAMPLES`

```
EXAMPLES
  $ <tool> <command>
```

**Three to six, and together they are a first session.** Not a showcase of
capability — a loop somebody can type in order and watch work.

**Order them as they happen.** Set up, point at a source, take something, wire
it in, look at what you have. Somebody who runs all of them in sequence should
end up with the tool doing its job.

**Real commands only.** Never invent a flag or an argument to make an example
look richer; an example that does not run is worse than no example.

Prefix each with `$` so a copied block is obviously commands rather than
output.

### `LEARN MORE`

```
LEARN MORE
  <where to go next>
```

**Two lines. How to get help on one command, and where the project lives.**

**Keep it tight because it is the least valuable section on the page**, and
because everything in it rots — links move, docs sites are retired, tutorials
go stale. Two durable pointers beat five that need maintaining.

**Never promise a route that does not exist.** If no subcommand answers
`--help`, do not offer `<command> <subcommand> --help`. Help that spends a
reader's trust before they find out is worse than help that says less.

## A command of its own

What `<tool> <command>` and `<tool> <command> --help` print.

```
<one line: what this is>

USAGE
  <tool> <command> <subcommand> [args]

<GROUP NAME>
  verb              What it does
  verb <arg>        What it does

FLAGS
  --flag            What it changes

EXAMPLES
  $ <tool> <command> <arg>

NOTES
  <what a reader must know and cannot guess>

EXIT CODES
  0 fine   1 <what refused means here>   2 could not run
```

**The first line says what the thing is** — lead with the verb and follow it
with what it manages: *Manage bundles — what this project has taken, and what
shape it is in.* No tagline; the tool already introduced itself.

**Groups are named for what the reader wants** — `READING`, `AUTHORING`,
`THE GATE` — rather than for core-ness, which is a top-level idea.

**Verbs are bare.** A row under `bundle` says `list`, not `bundle list`; the
heading already said it. Arguments go on the verb — `show <name>` — not into
the description.

**Descriptions start with a capital and a verb and take no full stop.** Align
them all to one column, including `FLAGS`, so the page has a single left edge
for prose.

**No `LEARN MORE`.** That is the top level's.

### `NOTES`

This is where things a reader cannot guess go, and where padding accumulates.
So it takes one test:

**A line stays if a reader can act on it.**

| | |
| --- | --- |
| *This command refuses a vendored copy* | changes what somebody types — keep |
| *`register nothing` marks a bundle landed and not wired* | changes what a value means — keep |
| *This command needs a network* | changes when it can be run — keep |
| *Checks your project against the baseline* | changes nothing — delete |

**Delete a failing line rather than rewording it.** The last row also names a
concept that exists nowhere, which is the worse half: a phrase that sounds
specific and points at nothing costs a reader more than a vague one, because
they go looking for it.

### `EXIT CODES`

**Every command states its own**, on one line, because what `1` means is local
to the tool: refused, behind, findings present.

**Local to the tool, not to the subcommand.** A subcommand states which codes
it can return; it does not get its own meanings for them — see
[[exit-codes]].

**`0` includes doing nothing.** *Nothing to do* is a successful outcome, and
reporting it as a failure makes every wrapper treat a settled state as a
problem.

**Keep `2` for could-not-run.** A tool that refused is working; one that could
not try is not, and a caller conflating them cannot tell a healthy tool from a
broken one.

## Beyond the help

The same eye applies to what a command prints when it runs.

**Print only what differs from the default.** A column showing the same value
on every row is wallpaper, and it hides the row where the value matters.

**Where a mark classifies state, make the fallback mean *something is wrong*.**
Name the specific states and let everything else fall to the *not working*
mark. The other way round means the next unanticipated failure prints as
healthy.

**Explain only the marks on the screen.** A legend covering states nothing is
in is one people stop reading.

**Name the fix for the cause, not for the mark.** Where one mark covers several
conditions, give the remedy for *this* one — a remedy that does nothing is
worse than none.

**A refusal names the literal command that resolves it** — the line, ready to
copy, not a description of what to type.

**Refuse rather than warn.** A warning after the act is useless, and one before
it needs a prompt, which breaks the moment anybody scripts the tool.

**Numbers beat markers where the number is the point**, right-aligned so `3`
and `12` line up under each other.

## Naming a command

**The verb names the goal, never the procedure.** `get`, not
`fetch-and-copy-and-write-a-receipt`. The user has a goal; the tool has the
procedure.

**One word means one thing across the whole tool.** Before reusing a word,
check what it already means somewhere else.

**A borrowed word carries its old promise.** `push` means *sent, no review*
everywhere else; `publish` in a self-service registry means *live now*. Borrow
the word only if the promise still holds.

**Required values are operands; optional ones are flags** — a flag reads as
optional however the help describes it. Do not infer an operand because there
happens to be one candidate: refuse, and print the candidates as literal
commands.

## Decoration

**Decoration is for a person.** Box drawing and marks are worth having in front
of a reader and are noise in a pipe, a log, or a screen reader — which
announces every glyph before every row.

**Use characters every monospace font has.** Box drawing and geometric shapes
qualify; anything needing a patched font does not.

**Check the width class before mixing glyph families.** Characters from
different Unicode blocks can render at different widths, and a column that
aligns on one terminal and not another is worse than one that never aligned.
Where the mark that means the right thing has the wrong width class, prefer
meaning.

## Worked examples

Four real outputs. Between them they cover the shapes most commands take, and
they are the fastest way to check whether something new fits.

### A noun's help

Nine verbs, so they are grouped. Verbs bare, arguments on the verb, one left
edge for all prose, `NOTES` carrying only what changes what somebody does.

```
Manage bundles — what this project has taken, and what shape it is in.

USAGE
  luma-foreman bundle <command> [args]

READING
  list                          Every bundle this project carries
  show <name>                   One bundle's receipt and contents
  outdated                      Which have a newer version published

AUTHORING
  new <name>                    Start a bundle here, under local/
  index <dir>                   Generate a bundle's INDEX.md — --check verifies

INTENT
  set <bundle> <field> <value>  Record intent in the manifest
  unset <bundle> <field>        Back to the field's default
  migrate-manifest              Rewrite the record as MANIFEST.md, retiring
                                a legacy adopted.toml

FLAGS
  --to <project>                Work on a repository other than this one
  --help                        Show this

EXAMPLES
  $ luma-foreman bundle list
  $ luma-foreman bundle show git-workflow
  $ luma-foreman bundle new house-rules

NOTES
  `index` is an authoring act and refuses a vendored copy.
  `list`, `show` and `new` work offline; `outdated` reaches each bundle's
  catalog and does not.

EXIT CODES
  0 fine   1 something is wrong or behind   2 could not run
```

### A noun's help, when there are only a few verbs

Three verbs need no grouping, so there is one section rather than an invented
split. **The shape follows the noun, not a template.**

```
Manage catalogs — where this project's knowledge comes from.

USAGE
  luma-foreman catalog <command> [args]

COMMANDS
  list                Every catalog this project draws from
  show <name>         What a catalog publishes
  add <source>        Register one, so `get` needs no --from

FLAGS
  --to <project>      Work on a repository other than this one
  --help              Show this

EXAMPLES
  $ luma-foreman catalog list
  $ luma-foreman catalog add https://github.com/acme/catalog
  $ luma-foreman catalog show acme/catalog

NOTES
  `list` works offline; `add` and `show` reach the catalog and need a network.

EXIT CODES
  0 fine   1 refused   2 could not run
```

### `list` — many things, one line each

A mark for state, then the name, then facts about it. **Everything notable
lands in one column** so the eye finds it in the same place whether it is how
the thing loads or what is wrong with it.

```
lumastack/luma-catalog
  ├─ ● audit-records           0.10.1  3 skills
  ├─ ◐ backlog-ideas           0.14.1  3 skills   · edited
  ├─ ⊘ command-line-interface  0.1.0               · standby
  ├─ ○ decision-records        0.12.1              · missing
  └─ ● git-secrets             0.7.1   3 skills   · eager

local
  └─ ● widgets                 0.1.0   1 skill

20 bundles · 44 skills

  ◐ 1 not as recorded          luma-foreman inspect --rule adoption
  ⊘ 1 turned off               luma-foreman bundle set command-line-interface register
  ○ 1 recorded, not on disk    luma-foreman get lumastack/luma-catalog/decision-records
```

What to copy from it:

**A heading per group, rows tied to it.** The heading carries the shared part
of every row, so no row repeats it.

**A mark that is a residual.** `●` working, `⊘` deliberately off, `○` not
here — and `◐` for *anything else wrong*, so a failure nobody anticipated
reads as broken rather than fine.

**Nothing said where it is the default.** Every one of these is `offered`
unless tagged, and printing that word twenty times would bury `eager`, the one
that costs something.

**A count, right-aligned, only where there is one.** `3` and `12` line up.

**A legend only for marks on the screen**, each naming the command that fixes
*that* state — not a general one that fits all of them.

**A total, in words.** `20 bundles`, never `20 bundle(s)`.

### `show` — one thing, in depth

The inverse of `list`: one subject, so there is room to explain rather than
tag.

```
lumastack/luma-catalog/git-workflow  0.8.1
  How changes get integrated — merge commits rather than squash or rebase, the
  repository settings that make it true, and how to prove a change landed.

  source     https://github.com/LumaStack/luma-catalog
  commit     ce13c21e65900542c1570a6afdf903d8ac4fbf73
  at         .luma/bundles/lumastack/luma-catalog/git-workflow
  copy       ok

  documents  4
    INDEX
    policy/merge-commits
    policy/proving-work-landed
    procedure/configure-merge-settings
```

What to copy from it:

**Name and version on the first line**, description under it. No section
headings — there is one subject, and headings would be scaffolding around a
single thing.

**Labels in a fixed column, values in another.** Lowercase labels, because they
are field names rather than headings.

**Say a state in words here.** `list` marks `ok` with a glyph because it has
twenty rows; `show` has one and can spend the word.

**Prose where `list` uses tags.** The same fact — an edited copy, a missing
one — is a tag in a list and a sentence here. Neither is the other's
abbreviation; they are written for different readers.

**Omit a row rather than print it empty.** A bundle written here has no source
and no commit, so it says so in one line instead of showing two blank fields.
