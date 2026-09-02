# Design Principles

The goals and values behind this project, and the reasoning underneath them.

> **These are leanings, not laws.** This project is intended to sit alongside a separate system that drives agents and runs workflows — the intelligence layer. **Where the boundary between the two falls is genuinely undetermined**, and it is expected to be discovered by building both, not by deciding in advance.
>
> Nothing here fixes how much this tool knows, infers, or enforces. Where a principle sounds absolute, read it as a direction with reasoning attached — the reasoning is the durable part, and the position is expected to move. When a design question is unclear, these are what it is argued against; when one turns out to be wrong, it is revised rather than worked around.
>
> [Luma Knowledge Format](https://github.com/LumaStack/luma-knowledge-format) is under rapid development and it is expected that this project will need to request changes to the specification.

This project assumes a world where a backlog is worked **continuously and concurrently**, by people and agents at the same time, and where the way agents work changes far faster than the backlog itself does.

## What this is

**The place work is represented, held, and manipulated.** It is where tasks and deliverables live, where their state can be seen, and where humans and agents change them. That includes visual representation as a first-class concern, not an afterthought — a board is one of the primary ways people will meet this tool.

**What it is not** is the thing that decides what happens next. Driving agents, running workflows, sequencing effort, and judging quality are expected to live in a separate system. That system is expected to change rapidly. This one is expected to hold still.

How much intelligence belongs on this side of that line is open. "Dumb store" is not the goal, and neither is a workflow engine. But the *method* for deciding is not open: almost no capability sits wholly on one side, so the question is not which side it goes on but **where the line runs through it** — the tool holding a neutral mechanism, and a team supplying the opinion it carries. `spec.md` §5.0 states that test and the two rules that bound it.

## The interface is the contract

Every actor — human, agent, or automation — works through the same interface. There is no privileged path, no internal API that the shipped commands lack, and nothing reachable through a view that a command cannot do.

Because other systems will be written against this surface, it should be treated as a public API: **stable, versioned, and additive.** Output shapes are as much a part of the contract as command names, and breaking either is a breaking change.

## Plain text is the system of record

The backlog is markdown files in a git repository, conforming to the [Luma Knowledge Format](https://github.com/LumaStack/luma-knowledge-format). The files *are* the system and the authoritative source of truth. Any index, cache, or database is derived, and can be deleted and rebuilt without loss.

- **Readable and editable with no tools.** Every operation this tool performs should be possible by hand in an editor, and nothing should be corrupted by someone doing so.
- **Diffable and mergeable.** Changes land as clean diffs and survive ordinary git workflows.
- **Portable.** The backlog outlives this tool.

The reason the work lives in git at all is that it gives agents context they cannot otherwise get. That purpose outranks tidiness when the two conflict.

## Concurrent access is the normal case

Many actors touching the backlog at the same time is the expected condition, not an edge case to harden against later.

> **This is about access, not execution.** Whether two *tasks* may be worked at the same time is a separate question with the opposite default — work is sequential unless something says otherwise, because forgetting to declare an ordering should cost speed rather than correctness (`spec.md` §4.5). Concurrency here means many actors reaching the same records; parallelism there means many tasks running at once. The words are kept apart deliberately.

- **Independent work should proceed independently.** Actors not touching the same thing should not wait on each other.
- **Contention should be bounded** to the records actually being changed.
- **Conflicting writes should be detected and surfaced rather than silently resolved.** A caller that would clobber a change it never saw should be told, and the end user made aware.
- **Concurrent-safe by hand, too.** A human editing in an editor while agents work is ordinary use.

## Completion is evidenced, not asserted

A unit of work is considered complete when the criteria attached to it pass with recorded evidence — derived by counting, rather than read from a field somebody set. An unbacked assertion of completion is the claim this design is most interested in distrusting.

This is not the same as forbidding an explicit act of closing. Deciding *when* to close is a judgment call that belongs to a human or an explicitly permitted agent; whether the evidence supports it is arithmetic. Keeping those separate is the point — the act is theirs, the check is the tool's.

It follows that criteria should be phrased so a check returns true or false, and the structure should make that discipline easy rather than optional.

## Derive what can be defended

The tool captures provenance faithfully — who acted, when, as human, agent, or process — and derives what follows from it: whether a criterion has evidence, whether the set of them passes, who confirmed what.

**How far past that it should reason is an open boundary.** Judging whether evidence is any good, ranking quality, or deciding when something has grown stale enough to distrust are the kinds of calls that need an intelligence this tool may or may not end up having. The current lean is to answer honestly and completely and leave conclusions to the caller — but this is the least settled principle here, and the one most likely to move.

## Preserve what you don't understand

Unrecognized fields, unknown types, and unfamiliar values are round-tripped untouched. This is inherited from the format, and it is also load-bearing architecture: **it is how other systems store their own state.** Something upstream may annotate any record with anything it needs, and this tool will neither interpret it nor lose it.

For the same reason the tool is permissive by default. It does not reject a record for a missing field, an unknown type, or an unresolved link. Strictness should be opt-in rather than the posture.

## Boring, on purpose

This tool should change slowly. Stability is a feature of infrastructure, and the volatile parts of the system have deliberately been placed elsewhere — as far as is currently understood.

That is paid for in a small surface area and reluctance about features that encode a particular way of working. Reluctance, not refusal: a capability that makes the tool markedly more useful may well be worth the churn it invites. The question is asked every time, not answered once.

**Dependency restraint means something specific**, and the two kinds are not comparable:

- **Anything the user must install** is what restraint is really about, and the target is **none**. A tool that needs a runtime installed first has already lost people before it does anything, and every version of that runtime becomes a support burden. This is a hard aim, and the reason the terminal tool is a single compiled binary.
- **Libraries compiled into that binary** cost maintenance and supply-chain attention, not user friction. They still deserve a reason on the record — but a library that removes real work is usually worth it, and treating it like the first kind is a mistake.

## Where the edges are today

Current positions, held loosely, and expected to be revisited as the boundary with the intelligence layer becomes clearer:

- **Deciding what to work on next.** Prioritization and readiness feel like they belong upstream.
- **Running or prompting agents.** This tool holds no prompts and no model configuration.
- **Enforcing process.** The tool may carry a gate a repository declares, and ships none of its own. It refuses only what a caller's own record contradicts — see `spec.md` §5.0. What remains open is whether a declared gate can be overridden (`open-questions.md` §6).
- **Being an organization's system of record.** External trackers may own the work; this stays a repository-local representation that gives agents context.
