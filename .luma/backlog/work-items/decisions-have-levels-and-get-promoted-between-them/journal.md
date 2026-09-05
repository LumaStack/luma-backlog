# Journal — Decisions have levels and get promoted between them

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-05

spec.md §4.8.1 already covers work item to project — copies never moves, promoted_from, and the two records having different jobs. The new parts are the third level, the keys, and the guidelines.
under copy semantics the key worry may dissolve: the promoted record is a NEW record with its own key at its own level, and the original keeps the key it always had
the body above a project is a GitHub organization, a company, a department, a steering committee — and if its decisions live in its own repository, promotion crosses repositories and a wikilink does not
position: the tool creates work item decisions and can promote them; whether it creates project decisions directly is open, and if it does it is a flag rather than an inference
found while recording it: nothing lets you say which level a decision is — new decision derives it from the working directory, so the level is decided by where you were standing, which where-an-idea-lives names as the failure that happens silently
all three ADRs here were created from the repository root and are correctly project-level; none of them was placed on purpose
the working directory is a good signal for a person and a constant for an agent, which is why every decision here is project-level — in an agent-first tool that is the signal not to lean on
which work item and which level are two questions; cwd can answer the first and answers the second only by accident
precedent for requiring it: new task already refuses without a work item, as a usage error rather than a judgment about content
