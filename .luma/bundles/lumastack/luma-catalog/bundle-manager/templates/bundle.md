# Bundle manifest template

Copy the block below into `<bundle>/BUNDLE.md` and fill it in. **Copy the block,
not this file** — this file deliberately has no frontmatter of its own, because
a template carrying `type: bundle` would be a second manifest inside this
bundle, and every tool reading it would believe that.

```yaml
---
type: bundle
title: CHANGE-ME
version: 0.1.0
published: YYYY-MM-DD
stage: draft
consumers: [project]
description: One line — what this holds and who it is for.
---
```

- **`version`** — `0.1.0` for anything used in one place or none. `1.0.0`
  claims the shape has stopped moving.
- **`stage`** — `draft` on the day it is written, and **left in the block
  deliberately** so a new bundle declares something rather than nothing. Delete
  the line and the bundle declares `unknown`, which reads as *nobody has said* —
  indistinguishable from nobody having thought about it. Overwriting it with
  `provisional` or `stable` is a claim, and one somebody has to decide to make.
- **`survival`** — **not in the block, deliberately.** It defaults to `intended`,
  so writing `survival: intended` adds a line that says what silence already
  says. Add the field only to say something else: `experimental` when the bundle
  is out there to find out whether it earns its keep, `promised` when something
  will go on answering this whatever shape it ends up in. Ask the question every
  time; write the field only when the answer is not the default.
- **`consumers`** — `project`, `organization`, or both. Both when the same
  content is wanted at either level by different adopters.
- **`description`** — what a consumer reads when deciding whether to adopt.

There is no start-here field in the manifest. A document that should be read
before the rest declares `matches: eager` itself, and carries the claim with it
wherever it moves.

## Body

```markdown
# Bundle Name

Why this exists, in a paragraph. What goes wrong without it.

## What is here

- [[the-main-procedure]] — the procedure. Start here.
- [[a-policy]] — what this obliges.

## When these apply

Which documents bind, and when each one comes up. **Nothing here says when a
document loads** — that is computed from what it obliges and when it applies.

Keep an eye on any policy declaring `matches: eager`: it loads into every
session of every adopter, forever. That is the one expensive outcome, and it now
has to be asked for — a policy that says nothing is available on request, so
nobody buys the cost by forgetting a field.

## Version

Why this number.
```

The *What is here* section links each document with one line on why to open it.
It is not an inventory — the directory already lists the files.
