---
type: procedure
title: Create a bundle
description: Scaffold a new bundle, decide where it belongs, and get it to a publishable state. Use when asked to create, start, or extract a bundle.
---

# Create a bundle

## 1. Is it one bundle?

**The test: would anyone adopt half of it?** If yes, it is two bundles. If
adopting half would leave someone with rules and no procedure for following
them, it is one.

Bundles have no dependencies, so splitting is cheap now and expensive later —
two bundles that need each other cannot say so.

## 2. Decide where it belongs

See [[where-a-bundle-belongs]]. The short version: a project and a catalog are
both valid first homes, and which is right depends on how many adopters you
already know about. **When you cannot tell, start in the project** — not because
it is better, but because it is the cheaper mistake to correct.

## 3. Scaffold

Copy [the bundle template](../templates/bundle.md) to `<bundle>/BUNDLE.md` and
create the directories you need — see [[organizing-a-bundle]]. Create only the
directories that will have contents; an empty `policy/` is noise.

## 4. Fill in the manifest

```yaml
type: bundle
title: <org>/<catalog>/<name>
version: 0.1.0
stage: draft
consumers: [project]
description: <one line — what this holds and who it is for>
```

- **`version`** starts at `0.1.0`, not `1.0.0`. The conventions in a new bundle
  are extracted from one place's practice at best.
- **`stage`** is `draft`, and **the line is not optional in practice.**
  Omitting it declares `unknown` — *nobody has said* — which a reader cannot
  distinguish from nobody having thought about it. `draft` is the true answer on
  the day a bundle is written: its maintainers are developing it for their own
  use, and its shape can reverse without notice. Anything higher is a claim
  somebody makes deliberately, later; see step 8.
- **`consumers`** — `project`, `organization`, or both. Both is right when the
  same content is wanted at either level by different adopters; that is not the
  publisher's call to force.
- **`title`** — the bundle's full published ID; the generated index renders it
  as the heading.
- **The way in is not a field.** Declare `matches: eager` on the document a
  reader must open first — it surfaces the moment the bundle is in play, and
  the claim travels on the document it is about.
- **`description`** is what a consumer reads when deciding whether to adopt.

### Then ask how long it is meant to last

**`stage` and `survival` answer different questions, and one bundle needs
both asked.** Stage says what happens when the shape changes; survival says
what happens when the thing ends. A new bundle is `draft` on the first axis
almost always, and can honestly be anywhere on the second.

| the answer | what it means for a new bundle |
| --- | --- |
| **`probationary`** | written to find out whether it earns its keep. Many do not, and nobody should fall in love with it. |
| **`intended`** | meant to be kept, nothing promised. **The default, and the ordinary answer.** |
| **`promised`** | something will go on answering this, whatever shape it takes. A commitment to the problem, not to this content. |

**Write the field only when the answer is not `intended`.** The default is what
absence already says, so `survival: intended` is a line that adds nothing —
whereas `probationary` and `promised` each tell an adopter something they cannot
infer. **Ask every time; write it sometimes.**

`draft` + `probationary` and `draft` + `promised` are both ordinary and mean
opposite things: *finding out whether this is worth having* and *committed to
the problem with no idea yet what the answer looks like*. Neither is expressible
on the stage ladder alone, which is why the second question is asked at all.

## 5. Write the entry point first

Whatever a reader must understand before anything else makes sense. **Say what
surfaces it** — `matches: topic: …` for a subject, `matches: command: …` for a
moment — so it arrives when the work matches and costs nothing before then.

**Reach for `matches: eager` as rarely as you can.** It asks for a permanent
seat in every adopter's context, in every session, forever. A document that says
nothing is available on request, which is the right default for almost
everything: its name and line are in the index, so nothing is missed out of
ignorance, and its body waits to be asked for.

## 6. Add types only if they change something

Declare a Type Definition when a consumer must **validate, load, or behave**
differently — not because a distinction reads well.

If the bundle uses `procedure` or `policy`, copy the definition from a bundle
that already has it. Bundles are self-contained: carry your own copy.

## 7. Audit before publishing

Run [[audit-bundle]]. A bundle published with a broken link or an unquoted
frontmatter wikilink is one every adopter copies.

## 8. Publishing does not promote it

**A bundle stays `draft` until somebody decides otherwise**, and putting it in a
catalog is not that decision. Being reachable by people who did not write it
makes the question live — *is this safe for them* rather than *is this working
for us* — but nothing answers it on the bundle's own behalf.

**Heavy use by its own authors does not promote it either.** An author
exercising their own draft is testing it, which is what a draft is for.

So the question belongs to whoever maintains it, asked at the moment the bundle
becomes reachable by somebody else, and *no, still a draft* is a legitimate
answer to publish. Moving it later is [[update-bundle]]'s business.
