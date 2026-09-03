# Type Definition template

Copy the block below into `_types/<name>.md`. **Copy the block, not this file**
— a template carrying `type: type_definition` would define a type called
`CHANGE-ME` in this bundle.

**Declare a type only if a consumer must validate, load, or behave differently
because of it.** A type that changes none of the three is a label, and it costs
a name every future bundle has to avoid.

```yaml
---
type: type_definition
defines: CHANGE-ME
fields:
  a_field:
    field_presence: required
    field_type: text
    desc: "what it holds, in a few words"
---
```

- **`obligation`** — `mandatory` · `recommended` · `optional` · `deprecated`
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
