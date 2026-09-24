# Changelog — `work-item`

The history of the `work-item` Type Definition, newest first, keyed by the
type's own `version`. History from before it declared one is in this
repository's git history — no retroactive backfill.

## 0.0.2

- **`former_keys` added** — optional, a list of text, holding the keys a record
  formerly answered to and no longer does, oldest first. Written by a key
  migration rather than by hand.
- **Nothing to do.** Adding a field is not breaking: a consumer that has not
  learned it reads the record exactly as before, and a record that has never
  been migrated has no list. Declared `optional` rather than `recommended` for
  that reason — raising an obligation is the change that breaks while looking
  additive.
- Its purpose is to make a rename survivable: the old key keeps resolving, so a
  reference held somewhere we cannot edit is not broken by our migration.

- **`0.0.2` rather than `0.1.0`, deliberately.** `change-a-shared-type` says to
  bump the minor for an added field, and the minor of `0.0.1` is `0.1.0`. That
  was declined: `outcome`, `task` and `exploration` are all still `0.0.1`, and
  leaving `0.0.x` says the shape has stopped moving. **Exiting `0.0.x` is a
  decision for whoever controls the type, not an increment that falls out of a
  change** — and nobody had made it.

## 0.0.1

- Versioning begins: the type declares its own `version`, independent of
  the bundle's. Shipped 2026-09-19 with this repository's migration to LKF
  v0.0.21, which moved every Type Definition to the folder shape — this
  definition now lives at `type_definitions/work-item/DEFINITION.md`, with
  this changelog beside it. The contract is otherwise unchanged.
