# Type Definition template

A Type Definition is a folder. Copy the block below into
`type_definitions/<name>/DEFINITION.md`, and start a `CHANGELOG.md` beside it.
**Copy the block, not this file** — a template carrying
`type: type_definition` would define a type called `CHANGE-ME` in this bundle.

```
type_definitions/CHANGE-ME/
  DEFINITION.md          # the contract — the block below. Required.
  CHANGELOG.md           # what changed in each version, and why
```

**Declare a type only if a consumer must validate, load, or behave differently
because of it.** A type that changes none of the three is a label, and it costs
a name every future bundle has to avoid.

```yaml
---
type: type_definition
type_version: "0.0.1"
defines: CHANGE-ME
version: "0.0.1"
fields:
  a_field:
    field_presence: required
    field_type: text
    desc: "what it holds, in a few words"
---
```

- **`version`** — the type's own version, independent of the bundle's. A new
  type starts at `0.0.1`; a bump means the contract changed and earns a
  `CHANGELOG.md` entry. Documents of this type copy it into their
  `type_version`.
- **`type_version`** — what every Document carries: the `version` of its
  type's definition it was written against. On a Type Definition it cites the
  `type_definition` contract itself. Bare `version` is what a document
  publishes; a qualified `*_version` is what it conforms to.
- **`field_presence`** — `required` · `recommended` · `optional` · `deprecated`
- **`field_type`** — `text` · `number` · `boolean` · `date` · `datetime` ·
  `semver` · `enum` · `wikilink` · `uri` · `actor` · `actor_event` ·
  `list of <type>`. **Omit it** when the shape cannot be expressed — that is
  legal, and the field stays discoverable.
- **`values`** — required when `field_type` is `enum`.

**Do not redeclare core fields** — `title`, `description`, `created`,
`modified`, `verified`, `stage`. They arrive from the
root, and inheritance is add-only, so a type cannot restate one to strengthen
it.

## Body

The body carries what the field table cannot:

- What a document of this type **is**, in a sentence.
- **Inherited fields that are load-bearing here** despite being optional at the
  root — the contract cannot say so, and the body is the only place to warn.
- **When to reach for this type** over a neighbouring one.
- **Value traps.** A `wikilink` in frontmatter must be quoted, or it parses as
  a nested array and silently never resolves.
