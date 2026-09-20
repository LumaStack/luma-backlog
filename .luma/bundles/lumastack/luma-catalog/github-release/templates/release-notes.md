# Release notes template

Copy the block below, fill it in, delete what does not apply. **Copy the block,
not this file.**

**Every heading here is conditional.** Delete the ones with nothing under them
rather than writing *none* or *nothing to do* — absence is the statement, and the
policy says so.

Title format: `vX.Y.Z` — the version alone, matching the tag exactly
Policy: `../policy/release-notes.md`

---

```markdown
> ⚠️ **Breaking.** <what stops working, in one line.> See **Upgrading** below.
```

<!-- The banner goes FIRST, above everything, and only when something breaks.
     A reader deciding whether to take this release needs that before any
     other word. Delete the whole line when nothing breaks — a warning that
     appears every time is a warning nobody reads. -->

```markdown
<One or two sentences: what this release is, and who should care.>

<!-- Required, and the only summary a release has — the title carries the
     version and nothing else, so a reader scanning the list opens this to
     find out whether the release affects them. Make the first line earn it. -->

## Upgrading from vX.Y.Z

<!-- The steps, in order.

     DELETE THIS HEADING when there is nothing to do. Not "nothing to do"
     written under it — no heading at all. Its absence says the same thing
     and costs the reader nothing.

     It sits above the change groups because it is what most readers came
     for. Not a copy of per-change migration notes: this is the whole
     upgrade in one place, written once the release is known. -->

## Added

<!-- New capability. Say what it is and why it exists — the reasoning is the
     part a diff cannot show. -->

## Changed

<!-- Existing behaviour that is now different. Anything requiring action also
     belongs in Upgrading above. -->

## Deprecated

<!-- Still works, discouraged, scheduled for removal. Name the release it goes
     away in if you know it. A removal whose first mention is under Removed
     never gave anyone a chance to act. -->

## Removed

<!-- Gone. What to use instead. -->

## Fixed

<!-- Bug fixes. What was broken, and for whom. -->

## Security

<!-- Vulnerabilities. Never file these under Fixed — this group exists so
     somebody scanning for "must I upgrade urgently" finds the answer in one
     place. Include severity, and whether exploitation was observed. -->

## Version category

<!-- Only when the number is not what the rules would obviously produce: a
     breaking change shipping as a patch pre-1.0, a large release that is only
     a minor, a skipped deprecation cycle. Unexplained, these read as mistakes
     later. Delete when the version is unremarkable. -->

## Known issues

<!-- Anything shipping broken, and what to do instead. Delete if none. -->

---

Full history in [`CHANGELOG.md`](https://github.com/OWNER/REPO/blob/main/CHANGELOG.md).

<!-- Absolute URL, not a relative path: these notes render on the releases
     page, where a relative link resolves against nothing. -->
```
