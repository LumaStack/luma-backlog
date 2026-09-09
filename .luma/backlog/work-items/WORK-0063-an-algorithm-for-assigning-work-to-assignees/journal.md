# Journal — An algorithm for assigning work to assignees

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-09

DEFERRED here on 2026-09-09 rather than decided: for now work items use OWNER and assignee is skipped entirely (ADR-0008 already makes ownership the settled concept). how assignees work is a later question
two questions to answer when it is picked up — is assignee an ARRAY, because we want to support changing assignees; and can an assignee attach to a GATE or a set of workflow_statuses, since different people work on preparing than on in_progress
the second question already has a precedent to reason from: workflow-status.md's 'Preparing may hold many gates' draws product, legal, compliance, security and engineering as different groups inside one phase — assignee-per-phase is that same shape seen from the person rather than the gate
