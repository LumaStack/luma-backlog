# Format change requests

Changes this project asks of the [Luma Knowledge Format](https://github.com/LumaStack/luma-knowledge-format), with the reasoning behind each.

> **This project may drive the format, and is expected to.** It is the format's first consumer, the format is pre-1.0 and explicitly unstable, and design problems found here are the best evidence available about what the format is missing. **Until this project reaches 1.0, proposing a format change is the normal response to hitting a limit** — not a last resort after exhausting workarounds.
>
> That stops being true later. A mature consumer works around a stable format; an early one shapes it. The window is open now and should be used.

Each request states what is wanted, why, and whether this project is **blocked**, **shipping ahead** of the change, or **waiting**.

---

## 1. `type_namespace` — a bundle-level default for type names

**Status:** shipping ahead. Adopted in `SPEC.md` §4.1.

**Wanted:** a bundle MAY declare a default namespace on its root `index.md`. Unqualified `type` values resolve against it.

```yaml
lkf_version:    0.0.2
type_namespace: luma/backlog
```

A record then writes `type: task`, meaning `luma/backlog/task`. A fully qualified value is always legal and always wins.

**Why.** The format recommends namespacing types, which is right — `task` and `decision` are the names any other system will also want. But carrying the namespace in every record makes frontmatter noisier for no gain in the ordinary case, where **every record in a bundle shares one namespace.** Declaring it once is the same information in the place it actually varies.

**Conflicts must be loud.** If a short name could resolve to more than one type, that is an error a consumer reports — never a silent pick. Qualification is then required for the ambiguous names only, not for the whole corpus. Precedence rules and search orders were considered and are not wanted: quiet resolution is how the wrong type gets chosen and nobody finds out.

**And the bundle root probably needs a type.** These declarations have to sit on a record, and the format defines *Knowledge Bundle* as a term while giving its root no type name. This project writes `type: bundle`, which is invented — legal, since unknown types are tolerated, but invented. If the format is going to recognize bundle-level declarations at all, it may as well name the record that carries them.

**It has to be readable where a format consumer looks.** A tool that understands the format but not this one must be able to resolve `type: task` without parsing a private configuration file. In this project the declaration is authored in `config.yml` and **generated onto the bundle root `index.md`**, so there is one place to edit and one place the format can find it.

---

## 2. How a namespaced type resolves to a Type Definition path

**Status:** shipping ahead. Adopted as `_types/luma/backlog/task.md`.

The format puts Type Definitions in `_types/` and resolves them by type name, but never says how a namespaced name maps to a path. Mirroring the name is the obvious reading and is what this project does. It should be stated rather than left to each consumer to guess — two consumers guessing differently is the whole problem namespacing was meant to solve.

Worth noting alongside: the format's namespacing examples are all two-level, while this project uses three (organization, domain, type).

---

## 3. An optional evidence key on `actor_event`

**Status:** waiting; worked around locally.

**Wanted:** `verified` entries may carry what the evidence *was*, not only who confirmed and when.

**Why.** A verification event records attribution with no place for the artifact — which is precisely the unbacked assertion this design exists to distrust. `verified` is a core field and inheritance is add-only, so a consumer cannot add this itself.

**Worked around** by carrying evidence in a separate field on the outcome, at the cost of splitting one fact across two places. If the format adopts this, the local field becomes redundant rather than wrong.

---

## 4. `index.md` inside a directory may be that directory's record

**Status:** shipping ahead.

The format reserves `index.md` as derived navigation — a rebuildable cache. This project uses it as the **authoritative record for its directory**, which is the natural reading of "the document for this folder."

The two need not compete: an authoritative record can carry a **generated navigation section** within it, regenerated in place. That is strictly better than a separate cache file nobody edits, and it removes a reserved filename rather than adding one.

---

## 5. `log.md` — rename to `journal.md`, and specify its structure

**Status:** shipping ahead, and offering the specification.

The format reserves `log.md` for a directory's append-only history, newest first, and its roadmap lists the exact structure as *undecided, needs a decision*.

**This project has one**, arrived at from journals in daily use rather than from first principles — including from where those journals diverged from their own template. `SPEC.md` §5.5 has it: newest first, prepended, never rewritten, the newest entry acting as a resume pointer that marks everything below it historical.

**And a rename.** *Log* undersells what the file actually accumulates. The artifact is not an event stream — it is the memory of the work: what was decided, what was ruled out, what is still unknown. `journal` names that; `log` names a receipt.

---

## 6. Composite field types, or a way to declare a small map

**Status:** shipping ahead, undeclared.

The field-type vocabulary has no way to express a small structured value such as `{on: <date>, why: <text>}`. The roadmap defers user-defined composite types.

This project has two such fields already — `blocked` and `paused` — so they simply go undeclared, which is legal, since a Type Definition publishes intent rather than enforcing it. **The format does this to itself**: `sources` is documented as a bespoke shape for the same reason.

That is a workable answer and an honest gap. Worth revisiting once there is evidence about how often consumers actually want it.
