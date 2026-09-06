---
type: decision
title: The command line is designed against clig.dev
decided: 2026-09-05
stage: provisional
reopen_trigger: a departure recorded here turns out to be an oversight rather than a choice, or clig.dev changes a rule this record relies on
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T23:10:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T01:00:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T00:24:45Z'}
---

# ADR-0006: The command line is designed against clig.dev

## Summary

The command line follows [clig.dev](https://clig.dev). This record holds the
departures, and the choices the guide deliberately leaves open — **ordering,
exit codes, prompting, and what a bare invocation does.**

## Problem

Adopting `lumastack/luma-catalog/command-line-interface` declares that this
project designs against clig.dev, and that **only a decision in force may
overrule it.** Nothing was in force: every decision record here is `draft`, and
`spec.md` is a specification rather than a decision.

An audit found four real disagreements, thirteen rules the specification is
silent on, and one place where **the tool contradicts its own specification** —
§9.1 says noun then verb; the binary does verb then noun. Precedent is not a
decision, so that drift had no standing either.

## Decision

### Ordering — noun then verb

`work-item close`, `outcome verify`, `task claim`. Verb-only where no noun
applies: `init`, `board`, `contract`, `config`, `check`, `log`, `serve`.

This confirms `spec.md` §9.1 and makes **the code wrong**, not the
specification. The reshape belongs inside
[[ADR-0004-every-interface-is-an-adapter-over-one-application-layer]]'s
refactor, so the command tree and the adapter rewrite are one breaking change
rather than two.

### Exit codes — six, not seven

`0` success · `1` unexpected · `2` usage · `3` not found · `4` conflict ·
`5` refused.

**`6` (already claimed) is dropped until claiming ships.** Claiming is unsettled
(`open-questions.md` §8 carries a live proposal to move claims out of records
entirely), and a code reserved for a feature nobody has designed is a promise
about a shape nobody chose. `spec.md` §9.9 makes adding one additive and
removing one breaking, so six is the reversible direction.

The test each survivor passes is **does an actor behave differently** — and `4`
is the one that matters most, being the only code where retrying is correct.

### Commands do not prompt

Missing input is a usage error naming what was needed. **No `--prompt`, no
`--no-input`, and no configuration key ships** — see *Prompting* below, which
records the design so it can be added rather than re-derived.

### A bare invocation opens the board when interactive

`luma-backlog` with no arguments opens the board (`spec.md` §11) **when a
terminal is attached**, and prints concise help otherwise. CLIG's help section
carries an explicit exception for programs that are interactive by default;
this invokes it, and adds the terminal condition the specification omits.

### Adopted from CLIG, absent from the specification

Terminal detection deciding human from machine output · `--plain` where
human output breaks parsing · formatted `--json` · color disabled by
`NO_COLOR`, `TERM=dumb`, `--no-color`, or a non-terminal · no animations
off a terminal · confirmation and `--dry-run` before bulk modification ·
the standard flag names `-n/--dry-run`, `-q/--quiet`, `-f/--force`,
`-a/--all` · help text leading with examples · a support path and a
documentation link in top-level help · configuration precedence · user
configuration under XDG base directories · a pager for long output, terminal
only · `-v` assigned deliberately rather than by accident.

## Departures, and why

| CLIG says | We do | Why |
| --- | --- | --- |
| **Prefer flags to args**; two args for different things is usually wrong | `work-item close <ref> <disposition>`, `outcome verify <ref> <verdict>`, `outcome assert <ref> <result>` | CLIG's own exception — *a common, primary action where the brevity is worth memorizing*. These are the most-typed commands in the tool, the second value is a small closed enum, and the shape is uniform across all three so one form is learned rather than three. |
| **Prompt for user input** when a value is missing | Usage error | See *Prompting*. |
| **Terminal detection** tells you whether a human is present | True for output, **false for input** | Agents driving this tool routinely run with a terminal attached. CLIG's heuristic held when it was written and does not hold for this tool's callers. So detection decides color, animation and paging — where a wrong guess is cosmetic — and never decides whether to block on input, where a wrong guess hangs a loop. |
| **Keep the name short**, a simple memorable word | `luma-backlog`, with `backlog` as an alias | `spec.md` §9a.2 — the full name keeps a future `luma <command>` dispatcher reachable at no cost, the way git finds `git-foo`. The short name ships as a symlink. Not a first-release concern either way. |

## Prompting, recorded for later

**Nothing here ships in the first release.** It is written down so that adding
it is an afternoon rather than an argument.

**Why it was not built now:** nothing in the output stream signals that a
question was asked — a prompt is text without a newline plus a blocked read,
indistinguishable to a caller from a slow command. A non-interactive caller has
no event to react to and waits until something times out. And nothing needs it
yet.

**Why it will probably be wanted:** these commands are wordy.
`work-item close WORK-0017 delivered --reason "…"` is a lot to type, several
commands take required enumerated values, and a person already in a terminal
closing one item should not have to open a full-screen application to be asked
a question. Prompting also teaches — *Close as? [delivered/rejected/canceled/
superseded]* conveys an enum better than a usage error does.

**The shape it would take:**

A **personal** setting, never a repository one — it changes how input is
gathered, never what is stored, which is exactly `spec.md` §8.4's test.

```toml
# ~/.config/lumastack/luma-backlog/config.toml
prompt = true          # absent or commented out means no prompting
```

**It must never appear in `[require]`.** A project able to mandate prompting
could hang every agent working that repository. The precedence document's own
test already excludes it; the exclusion is stated because the failure is severe
and silent.

| configuration | flag | no terminal | result |
| --- | --- | --- | --- |
| absent / `false` | — | — | no prompting |
| absent / `false` | `--prompt` | | prompt |
| absent / `false` | `--prompt` | ✓ | **usage error** |
| `true` | — | | prompt |
| `true` | — | ✓ | **no prompting, silently** |
| `true` | `--no-input` | | no prompting |

**The two terminal cases differ on purpose.** A configuration file states a
preference, so it degrades when it cannot be honored — erroring would break
every non-interactive invocation a person makes after setting it once. A flag
was said out loud, so it fails rather than being quietly ignored.

**Prompting must be uniform if it arrives** — *any required value not supplied
is prompted for* — never a list of commands that happened to grow one. CLIG:
*be consistent across subcommands.* Inconsistent prompting is worse than none,
because nobody can predict which invocation will stop and wait.

**What would bring it forward:** a specific command people find tedious enough
to wrap; evidence about how agents actually behave when a subprocess blocks on
input, which is reasoning here rather than measurement; or a convention for
signaling *I am waiting for input* that a non-interactive caller can detect.

### A title may be given either way

```
work-item new Fix the login timeout on mobile     # positional, words joined
work-item new --title "Fix the login timeout"     # precise
```

**Both, and an error if both are supplied.**

**The positional is the documented path**, and it absorbs multiple words — no
quoting, no field name, nothing to remember. That is not a nicety here: cheap
capture is load-bearing (`design-mvp.md` — *"make capture extremely cheap"*),
and `capture-is-one-command` is a **verified** outcome of WORK-0001, so a
flag-only form would be giving back a property this project already proved.

It is also what §9.0.1 requires — *"a command whose only path requires knowing
field names"* is ruled out, and `--title` is a field name. That section's
argument is about the moment a person takes the pen, *"least willing to spend
attention on a schema."*

**`--title` exists for the precise case**, which §9.0.1 also names: agents want
flags, and a title beginning with `-` or generated programmatically is safer
given explicitly.

**CLIG's objection is real in general and weak here.** It warns that positionals
become ambiguous once further input is added — but everything else `new` takes
is a named attribute (`--kind`, `--work-item`), and under noun-verb ordering
there is no plausible second positional. The failure it warns about mostly
cannot occur.

## Still open

**The configuration file format.** The estate's precedence model is adopted —
six layers, `[defaults]` and `[require]`, read per invocation, resolved at one
site. Whether the file is TOML is not this project's decision to make alone;
`.luma/config/` is shared across tools, and `spec.md` §8.1 argues for YAML on a
premise — that a repository should carry one format — already false here.

## Alternatives

| Candidate | Set aside because |
| --- | --- |
| **Verb then noun**, matching what the binary does today | Precedent, not a decision. Noun-first gives completion something to offer at every position and lets a new record type arrive without touching an existing command (`spec.md` §9.1). |
| **Keeping exit code `6`** | Deferred until claiming ships, at which point it is additive. |
| **`--agent` to disable prompting** rather than `--prompt` to enable it | Puts the hang risk on the default path — an agent that forgets the flag stalls silently. The dangerous behavior should require the explicit word. *Reopened only if prompting ships and opt-in proves unusable.* |
| **A title only as `--title`** | The more disciplined shape, and one obvious way to do it. Set aside because it requires amending §9.0.1, loses quote-free capture, and can be added later additively while removing a positional is breaking. |
| **Flags for the dispositions** — `close --as delivered` | Safer against enum churn, and we removed `abandoned` while designing this. Set aside for brevity on the most-typed commands; the enum is small and closed. *Reopened if the dispositions change again, which would be evidence the positional is too rigid.* |

## Revisit When

- A departure above turns out to be an oversight rather than a choice — which
  the adopted policy says is grounds to re-open, never grounds to quietly ignore.
- clig.dev changes a rule this record relies on.
- Prompting's triggers fire.
- Claiming ships, restoring exit code `6`.

## References

- `.luma/bundles/lumastack/luma-catalog/command-line-interface` — the policy
  this record answers to, and its step three, which is why this exists.
- `docs/spec.md` §9, §9a, §11 — the surface audited.
- `luma-config`'s configuration precedence — six layers, and the
  `[defaults]` / `[require]` split.
