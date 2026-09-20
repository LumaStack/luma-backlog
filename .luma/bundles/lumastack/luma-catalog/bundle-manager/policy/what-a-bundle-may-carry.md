---
type: policy
type_version: "0.0.1"
title: What a bundle may carry
description: Somebody else's writing in a bundle — what adoption does to their licence, the five routes, and the one that is the default. Present the choice; never make it silently.
matches:
  - topic: putting somebody else's writing, code or data into a bundle
  - topic: deciding whether to vendor an external document or link to it
---

# What a bundle may carry

## The fact everything else follows from

**Adoption is redistribution.**

`foreman get` copies a bundle's directory into somebody else's repository. That
is publication, on their behalf, into a repository whose licence you do not know
and whose owner never saw the question. **Every obligation riding on those files
rides in with them.**

So the question is never *may I read this* or *may I use this*. It is:

> **May I hand this to strangers, forever, and make them responsible for it?**

Almost every licence says yes to the first two. The third is where the cost is,
and it is a cost somebody else pays.

**Linking is not redistribution.** A URL in a document carries no licence, no
attribution duty and no obligation for an adopter. This is the whole reason the
default below is the default.

## The default

**Carry only what you wrote. Point at everything else.**

One licence in the repository, no `NOTICE` file, no per-directory attribution,
and nothing an adopter inherits without knowing. It is the only route where a
bundle's licence file tells the whole truth about the bundle.

**Depart from it deliberately, with the user's agreement, and write down why.**
The rest of this policy is how to have that conversation and what the
alternatives actually cost.

## Never decide this silently

**Present the choice to the user before taking any route but the default.** They
own the repository, they own its licence, and the decision binds people who are
not in the room.

An agent facing this decision **states four things and then recommends**:

1. **What the source is**, and how much of it is in question.
2. **What its licence is** — found and read, not assumed from the project's
   reputation. A missing `LICENSE` file is a finding, not a detail.
3. **The routes that are open**, with what each means in plain terms — what
   *the user's repository* takes on, and what *every adopter* takes on.
4. **The recommendation**, which is the default unless something specific
   argues otherwise.

**Say what an adopter inherits, in the user's own terms.** *If we carry this,
everyone who adopts this bundle is redistributing it too, and they take on the
attribution requirement whether or not they notice* is a sentence a person can
decide on. *It is CC BY 4.0* is not.

## The routes

### 1. Point and cache — the default

**Link to the source. Carry none of its text. Cache a copy machine-locally so
reading it is cheap.**

> **In plain terms:** you are not copying anything, so no licence attaches to
> anything. Your repository stays single-licence. Adopters inherit nothing.
> A URL is not distribution.

**Choose it unless something specific argues otherwise.** Always available,
regardless of what the source's licence says — even *all rights reserved* — and
it is the only route with no downstream cost at all.

**What it costs:** the substance does not arrive with the repository, so a fresh
clone on a machine that has never fetched it has a URL and nothing else. The
cache closes most of that gap; see below. It cannot close the reproducibility
gap, and the source can change or vanish underneath you.

### 2. Restate it in your own words

**Write the rules yourself, in your own prose and your own order, citing the
source for anything a reader may want to check.**

> **In plain terms:** copyright protects the *way something is written*, not the
> facts or rules it states. A rule you wrote out yourself is your writing, under
> your own licence, and adopters inherit nothing.

**The line to respect:** restating rules in your own words and your own
structure is ordinary and safe. **Reproducing somebody's structure
section-by-section in reworded sentences is a derivative work**, however much
the wording moved, and it carries their licence and needs their attribution. If
your document's outline is theirs, you took the route below and should say so.

**What it costs:** it is now yours to maintain, and it can be quietly wrong in
ways a reader cannot detect — they are reading your summary while believing
they have the source's position. And a distillation hides what the reader did
not know to ask for: **you cannot seek a rule you do not know exists.** Prefer
restating rules you have actually applied over rules you have only read.

### 3. Quote a little

**Take a short excerpt — a definition, a rule, an aphorism — attributed to its
author, inside your own prose.**

> **In plain terms:** short quotations attributed to their source are ordinary
> practice and normally attract no licence obligation. Very short phrases carry
> thin copyright at best.

**Choose it when the exact words matter** and a link cannot do the job — a
tiebreaking maxim, a definition your rule turns on, a warning stated better than
you would state it.

**What it costs:** the boundary is a judgment call, and it moves with how much
you take and how central it is. **A quotation that could substitute for reading
the source is not a quotation any more** — it is route 4 without the paperwork.
Keep it to what your own argument needs.

### 4. Carry it, and pass the licence on

**Vendor the text into the bundle, with its licence, its attribution, and a
notice of any changes.**

> **In plain terms:** the file stays under *their* licence inside your
> repository — you cannot relicense it. Your repository becomes multi-licensed
> and needs a notice saying which parts are which. **And every project that
> adopts the bundle is redistributing their work too**, so every adopter takes
> on the same attribution duty, in a repository that may have a different
> licence entirely and an owner who never agreed to any of it.

**What it requires**, and all of it travels to adopters:

- **the licence text** available to anyone who receives the files
- **attribution** — author, source, a link, and the copyright notice
- **a statement of what you changed**, if you edited, trimmed or reordered it —
  required by Apache-2.0 §4(b) and CC BY §3(a)(1)(B) alike
- **a `NOTICE` file's contents carried forward**, if the upstream has one
  (Apache-2.0 §4(d))
- **the snapshot recorded** — which version or commit you took, since nothing
  in the format will tell you it has gone stale

**Choose it when the text must be present offline or pinned**, and the source's
licence permits it, and the user has agreed to the downstream cost.

**Prefer it strongly when the licence matches your own.** Vendoring an
Apache-2.0 document into an Apache-2.0 catalog still creates obligations, but
they are obligations the repository already meets.

**What it costs:** every adopter carries what you carried. A snapshot also goes
stale silently — and unlike a link, it will confidently keep saying whatever it
said the day you took it.

### 5. Leave it out

**Decide the bundle does not need it.**

> **In plain terms:** no permission needed, because nothing was used.

**Choose it when there is no licence at all** — no `LICENSE` file, no stated
terms — because the default is *all rights reserved* and you have no permission
to copy anything. Route 1 is still open: **linking to something is always
allowed**, whatever its terms.

Also the honest answer when a source turns out to add little that the others
do not already say.

## Reading the licence before you decide

**Find the actual licence text. Do not infer it from the project's reputation
or from what a similar project uses.**

Check, in this order: a `LICENSE` file in the source repository, a licence
statement on the page or site serving the content, and the project's
documentation. **A site-wide statement covers content served on that site**,
which is weaker than a repository licence but is usually what documentation
has.

| what you find | carrying it |
| --- | --- |
| **MIT, BSD** | permitted; keep the copyright and permission notice with it |
| **Apache-2.0** | permitted; licence, attribution, changes stated, `NOTICE` carried |
| **CC BY (3.0, 4.0)** | permitted; attribution, link, changes stated. **No ShareAlike** — you may license your own additions as you like |
| **CC BY-SA** | permitted, and **anything derived from it must be released under CC BY-SA too**. This is how a second, spreading licence enters a repository. Point instead |
| **CC BY-NC, CC BY-ND** | **do not carry.** `NC` conflicts with adopters who are commercial; `ND` forbids the trimming and reformatting vendoring involves |
| **GPL, AGPL** | **do not carry** without a deliberate decision far above this policy's level |
| **no licence file** | **no permission.** Link only |
| **unclear or contradictory** | treat as no licence until somebody resolves it |

**Documentation and code in one source can carry different licences.** Go's
documentation is CC BY 4.0 while its code samples are BSD — so a vendored copy
of one page carries two licences. Check whether the source says this before
assuming one covers the file.

## How to cache, so a fresh session gets it right

An agent has no memory of what it did earlier, and a pointed-at document is
useless if reading it is expensive. **A bundle that points at something should
say exactly how to cache it**, in the policy that points.

**Where.** One directory per bundle, named by the bundle's **full published ID**:

```
~/.cache/luma/luma-foreman/bundles/<org>/<catalog>/<bundle-name>/<source>.<ext>
```

```
~/.cache/luma/luma-foreman/bundles/lumastack/luma-catalog/command-line-interface/clig.dev.md
```

**The full ID, because a bare name is not unique.** Bundle names are unique
within a catalog and nowhere else — two catalogs may each publish a
`command-line-interface`, pointing at different documents and caching them
differently. The ID that already distinguishes them is the one `foreman get`
takes, so use all of it.

**The application segment is `luma-foreman`, in full and never shortened.** The
`luma-config` bundle settles this: XDG puts regenerable things under
`~/.cache/<org>/<application>/`, and **the slot after the organization holds an
application name** — so everything under `~/.cache/luma/` maps to a repository.
`bundles/` goes *inside* it, beside the `catalogs/` and `projects/` directories
already there.

**Do not put `bundles` in the application slot.** `~/.cache/luma/luma-foreman/bundles/…`
reads well and is wrong, and this is settled rather than arguable: foreman
cached at `~/.cache/luma/catalogs` until August 2026, a review found it, and it
moved for exactly this reason. **A plural noun there breaks the mapping for
everything under the organization**, and nothing reports it — a path nobody has
written to before is created on demand, and nothing notices the old one is
empty.

**One copy, shared by every tool.** What is cached is a public document fetched
raw, and nothing transforms it — so there is no version of this where two tools
should hold different copies. The application segment names where the cache
lives, not who may read it.

**A bundle with no catalog uses whatever ID it has**, which for an unpublished
one is local to its project. **Two projects holding same-named local bundles
will collide in one person's cache** — named here rather than solved, because
the failure is a stale document in a directory nobody published, and a rule to
prevent it would cost more than it saves. Publishing the bundle fixes it, since
publishing is what gives it a unique name.

**The base path is not this bundle's to invent.** The `luma-config` bundle
settles where machine-local paths go: XDG, so `~/.cache/<org>/<application>/`
holds anything regenerable, with the application name never truncated. **This
bends that by one segment** — `bundles/` is not an application, it is a sibling
namespace for things bundles own rather than programs do. Said plainly here
because the alternative was naming a program that owns none of it, and because
the fixed segment is what keeps a bundle named after a tool from claiming that
tool's cache directory.

A project organizing its caches differently substitutes its own base. The shape
is what matters — one directory per bundle, somewhere regenerable, never inside
the repository.

**Fetch the source, not the rendered page.** Most documents have a plain
markdown or HTML source behind a styled site; it greps well and costs a
fraction of the markup.

**Fetch to the file, never through the context window.**

```sh
curl -sSL -o ~/.cache/luma/luma-foreman/bundles/<org>/<catalog>/<bundle-name>/<source>.md <url>
```

`curl -o <path>` puts the bytes on disk at no token cost. Fetching a document
*into* a reply spends the whole thing in order to save it — which is the
mistake this instruction exists to prevent, and it is expensive exactly once
per session, every session, until somebody writes it down.

**Stamp it.** Write the date fetched and the source URL as the first line of the
file, so nothing can read the content without seeing how old it is.

**Refresh when the recorded date is not today, and only then.** That is "once
per session" expressed as something an agent can actually check — it cannot
remember, but it can compare a date.

**Never block on the network.** If a fetch fails, use what is cached and say how
old it is. A stale copy is worth far more than no copy.

**Never commit the cache.** It is derivable, and it is the licence boundary:
**a copy on one machine is nobody's business; a copy in a published repository
is distribution** and carries the licence with it. Committing the cache is
route 4 by accident, without any of the paperwork route 4 requires.

**What the cache buys, and what it does not.** Reading a file from disk costs
exactly what reading it from the web costs — the tokens are the content, not
the transport. So it buys working without a network, and **random access**:
coming back for one rule costs forty lines instead of the whole document. **It
does not make the first read cheaper**, and it is not permission to skip a read
the policy calls for.

## Record what you chose

**Say in `BUNDLE.md` which route the bundle took and why.** A bundle that points
looks identical to a bundle whose author never considered carrying, and the next
maintainer cannot tell them apart. One paragraph settles it permanently.

**If you carried something, record the snapshot** — the version or commit taken,
and the date. Nothing in the format signals that a vendored copy has gone stale
or been edited, so re-vendoring is a deliberate act rather than something a tool
reminds you about.

**If you departed from the default, record that as a decision** where this
project records decisions, not only in the manifest.

## This is not legal advice

**It is a working rule for a common situation, written to keep the obvious
mistakes from happening quietly.** Where real money, a real risk or an unclear
licence is involved, that is a question for a lawyer and not for a bundle.

**What this policy is confident about is smaller and still worth stating:**
adoption copies files into other people's repositories, obligations travel with
copies, and nothing travels with a link.
