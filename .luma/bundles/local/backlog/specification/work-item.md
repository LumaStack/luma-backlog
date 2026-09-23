# Work Item Specification

## Goal

We want work item data and prose that is:
1. optimized - for project success; using the latest research
2. maintainable - aggressively eliminate unnecessary overhead for upkeep
3. verifiable - easy for humans and agents to prove as completed
4. reviewable - easy for humans to review and improve
5. executable - easy for AI to consume and execute upon
6. minimal - reduce down to the lowest number of most impactful units of information

## Success criteria for work items

A work item should contain the minimum information required for a competent human or AI agent to:
- understand why the work exists
- determine what must change
- know what not to change
- understand the risks
- objectively determine when it is done
- verify and prove the work is complete
- (sometimes) know how to measure success upon delivery

All work items should strive to follow INVEST:
— Independent
- Negotiable
- Valuable 
- Estimable
- Small
- Testable

Active work items should strive to follow SMART:
- Specific:  Define your goal clearly by answering who, what, where, when, and why.
- Measurable:  Include exact numbers, amounts, or metrics to track your progress.
- Achievable:  Make sure the goal is realistic given your current skills and resources.
- Relevant:  Ensure the goal matters to you and aligns with your broader priorities.
- Time-Bound:  Set a clear deadline or target date to create urgency.

## Layers

Separate information into three layers.

| Layer | Purpose | Examples |
|-------|---------|----------|
| Metadata | Querying, routing, automation, reporting | Kind, Priority, Parent |
| Specification | Understand the work | Problem, Result / Desired State, Context, Constraints, Acceptance Criteria (Outcomes) |
| Execution state | Manage the work | Status, Assignee, Sprint |

Rule of thumb: If you expect to filter, sort, automate, route, aggregate, or report on something, make it metadata. Otherwise, keep it in the body as an expected section. That prevents the classic failure mode of 30 custom fields.

## Metadata

For most software/product organizations, I would start with approximately 8 meaningful metadata attributes.

| Field                |    Required? | Why                                                           |
| -------------------- | -----------: | ------------------------------------------------------------- |
| **Kind**             |          Yes | Defect / Change / Inquiry / Epic / etc.                       |
| **Title**            |          Yes | Human + AI identifier                                         |
| **Parent**           |      Usually | Connects execution to larger outcome                          |
| **Rank**             |          Yes | Backlog order                                                 |
| **Priority**         |      Ideally | Helps surface winners when backlog contains competing work    |
| **Assignee**         | When started | Accountability                                                |
| **Status**           |          Yes | Workflow state                                                |
| **Tags**             |     Optional | Ad-hoc classification                                         |

### Advanced metadata

Only use when necessary for business outcomes:

- Story points
- Due date
- Tags/labels
- Component/area
- Reporter
- Fix version
- Environment
- Target release
- Team
- Reviewer
- QA owner
- Risk
- Business value score
- Confidence
- Effort category
- Etc.

Those can exist when your process genuinely consumes them, but shouldn't automatically be mandatory.

### Test for adding metadata

> What decision, query, automation, or report becomes unusable without this field?  
> And how much do we actually need that thing?

## Body / Description

    ## Problem
    Why does this work need to exist?

    ## Result
    What should be true when this is complete?

    ## Context (optional)
    What does someone need to know to do this work?  
    Only information needed to understand or execute the work; that they won't find in our knowledge base.

    ## Out of scope (optional)
    What aren't we doing?
    Explicitly not included in this work.

    ## Constraints (optional)
    What rules must the solution obey?
    Technical, product, compatibility, compliance, security, or other restrictions on the solution.

Order matters: Problem → Result → Context → Out of Scope → Constraints.  For example, do not put Out of Scope before Context because Context establishes enough understanding for the reader or agent to correctly interpret the boundary. 

### Prefer Problem → Result over traditional user stories

Let's avoid requiring:

    As a [user], I want [thing], so that [reason].

It can be useful, but making every work item fit that syntax creates filler. The research literature recognizes the classic user-story structure of role, goal and benefit, combined with conversation and acceptance criteria.

But consider:

    As a developer, I want the Redis client upgraded so that we're running a supported version.

Versus:

  Problem
  Redis client 4.x reaches end of support next month.

  Result
  The service runs against Redis client 5.x without behavioral regressions.
  
The second version contains substantially more executable information. Consider user stories when the actor actually matters.

### Agents work best with constraints

For AI execution, constraints becomes particularly valuable. This tells an AI the boundaries of its solution space

Example:

    ## Constraints

    - Must remain backward compatible with API v2.
    - Do not modify the database schema.
    - Must support PostgreSQL 16+.
    - Existing authentication behavior must remain unchanged.

Without constraints, agents frequently produce a technically valid solution that is architecturally or operationally undesirable.

### Acceptance criteria as Outcomes

In our system, acceptance criteria are represented as individual Outcome items. We use the term Outcome because each item serves multiple purposes: it defines the desired state, establishes the acceptance criteria and definition of done, provides the basis for testing, and specifies what must be verified to prove completion.

Every Outcome carries a burden of proof: it must be objective, measurable, and verifiable. An Outcome is not complete merely because the work was performed; there must be sufficient evidence that the stated result has been achieved.

### Acceptance Criteria (Outcomes) should not carry implementation

**Bad — implementation tasks**

* [ ] Update `UserService`.
* [ ] Modify `getUser()`.
* [ ] Add an `if` statement.

These describe work to perform, not outcomes to prove.

**Better — observable outcomes**

* [ ] Requests for an existing user return the user.
* [ ] Requests for an unknown user return `404`.
* [ ] Disabled users return `403`.
* [ ] Existing API clients continue to work without modification.

These describe desired behavior, but leave some ambiguity about the exact conditions that constitute success.

**Best — measurable and verifiable outcomes**

* [ ] For every existing user in the supported test scenarios, `GET /users/{id}` returns HTTP `200` and the response contains that user's expected ID and profile data.
* [ ] For every tested user ID that does not exist, `GET /users/{id}` returns HTTP `404` with the documented error response.
* [ ] For every disabled user in the supported test scenarios, `GET /users/{id}` returns HTTP `403` and does not return the user's protected profile data.
* [ ] All existing API v2 compatibility tests pass without modification.

Each Outcome is objective, measurable, and independently verifiable. A reviewer or agent can determine whether it passed or failed from evidence rather than judgment. Now humans, tests, and AI all share the same contract.

Do not force numbers into every Outcome just to make it “measurable.” For software, measurable often means a binary observable condition—status code equals 404, response contains X, test suite has zero failures—not necessarily a numeric target like “99%.”

You can use Given/When/Then where it improves precision, but do not mandate Gherkin. Mandatory Gherkin quickly becomes ceremony.

### Handling empty body sections

When a required section in the body is empty, let's include a value of "TBD" so it's easy to audit, find, and update.

When new or draft:

    ## Problem
    TBD

    ## Result
    TBD

    ## Context

    ## Out of Scope

    ## Constraints

Once prepared or after (e.g. todo, in progress):

    ## Problem
    Users cannot...

    ## Result
    Users can...

    ## Constraints
    - Must remain compatible with API v2.

Here, Context and Out of Scope are omitted only after someone considered them and decided they weren't needed.

## What does not belong in work items

### Don't duplicate project knowledge 

This becomes very important once AI is involved.

Don't put this into every issue:

    Run npm test before submitting.
    Use TypeScript.
    Follow ESLint.
    Use React components from /components/ui.
    Tests belong under /tests.

The issue should contain delta context:

> What does the agent need to know about this work that it cannot reasonably discover from the repository and standing instructions?

## Different work Kinds should slightly modify the template

I would not use one giant template for everything.

Instead, establish a tiny universal grammar and specialize it. Some examples include:

### Change (Story, Feature)

    ## Problem

    ## Result

    ## Context 
    (optional)

    ## Out of scope 
    (optional)

    ## Constraints
    (optional)

### Change (Technical Task, Maintenance, Upgrades)

Don't force a fake "Problem" onto something like: Upgrade PostgreSQL 16 → 17.

    ## Objective
    What needs to change and why?

    ## Acceptance Criteria

    ## Constraints
    (optional)

    ## Context
    (optional)

### Defect

A bug has different information requirements:

    ## Problem

    ### Expected
    What should happen?

    ### Actual
    What happens instead?

    ### Reproduction
    Minimum reliable reproduction.

    ## Context
    Logs, screenshots, affected versions, etc.

For bugs, Expected / Actual / Reproduction are much higher-value than forcing a generic feature template. Atlassian's own bug-template examples similarly emphasize environment, reproduction, actual result and expected result.

### Inquiry (Spike, Investigation, etc.)

This one should be fundamentally different:

    ## Question
    What are we trying to learn or decide?

    ## Context
    Why do we need the answer?

    ## Deliverable
    What concrete output concludes the investigation?

    ## Constraints
    (optional)

The acceptance criterion for a spike isn't "software works." It may be:
    
    Document recommendation, alternatives considered, evidence, risks, and proposed next step.

## Epics

Epics should describe business outcomes, not giant specifications.  Epics should lean heavily on their children to explain the details of the work, we want to avoid creating duplicate text or metadata that can easily rot over time.

I'd keep Epics extremely clean:

    ## Problem
    What larger problem are we solving?

    ## Result
    What meaningful change should result?

    ## Success Measures
    How will we know the initiative succeeded?

    ## Scope
    What is included?

    ## Out of Scope
    What explicitly isn't?

    ## Context
    (optional)

Individual work items then contain the implementation-level acceptance criteria and context.

### The Breakdown

    Epic         WHY / outcome
    |_ Change    WHAT / behavior
       |_ Task   HOW / execution

That hierarchy is also excellent for AI because context becomes progressively more concrete.

## Strongly consider Component / Area

There is one custom metadata field I think may earn its keep: `Component / Area`

Avoid free-form labels for this.

Instead use a controlled vocabulary such as:

    API
    Web
    Authentication
    Billing
    CLI
    Infrastructure
    Observability
    Documentation

This becomes extremely useful for:

- ownership
- filtering
- AI context
- reporting
- routing
- CODEOWNERS-like automation
- architectural analysis

## Be cautious with Priority

You probably need Priority, but define exactly what it means.

A common mistake is allowing:

    Highest
    High
    Medium
    Low
    Lowest

with no semantics.

Then almost everything becomes meaningless.

I would define it around scheduling urgency, for example:

| Priority     | Timing              | Meaning |
| ------------ | ------------------- | ------- |
| **Critical** | Now                 | Requires immediate attention; interrupt planned work to address it. |
| **High**     | As Soon As Possible | Address at the next reasonable opportunity, ahead of Normal work. |
| **Normal**   | Planned             | Address through normal planning and backlog prioritization. |
| **Low**      | Opportunistic       | Do when surplus capacity permits; may be deferred without significant consequence. |

Then use backlog rank to distinguish items within a priority. Priority and Rank shouldn't attempt to encode the same thing.

## AI changes what "good hygiene" means

Traditionally work tickets were optimized for humans reading them.

Now we should optimize for three consumers simultaneously:

    Human
    AI agent
    Automation

That favors:

- Structured headings over arbitrary prose.
- Explicit acceptance criteria over implied expectations.
- Controlled metadata over free-form labels.
- Links/references over copied documentation.
- Observable outcomes over implementation prescriptions.
- Explicit constraints over tribal knowledge.
- Small independent work items over giant specifications.

## MVP

Deliberately begin with a small schema:

| Category          | Fields             |
| ----------------- | ------------------ |
| Identity          | Kind, Title        |
| Hierarchy         | Parent             |
| Classification    | Area               |
| Planning          | Priority           |
| Execution         | Status, Assignee   |
| Flexible taxonomy | Labels (optional)  |
| Specification     | Body / Description |

And make Body structured rather than creating five separate custom fields.

Body headings give the AI nearly the same semantic structure while remaining much easier for humans to edit and Jira administrators to maintain. You can promote something into a custom field later only if you discover that you need to query/report/automate against it.

## Sources
- https://www.sciencedirect.com/science/article/pii/S0950584925003180
- https://link.springer.com/article/10.1007/s00766-016-0250-x