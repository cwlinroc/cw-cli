---
locale: en
---
# Project Docs Structure

This file is the local authority on how docs are organized in this repository. Use it together with `docs-structure`; when the two conflict, follow this file.

## Target Layout

```text
.
├── README.md
└── docs/
    ├── CONTEXT.md      # glossary ONLY — shared terminology, no implementation detail
    ├── adr/            # Architecture Decision Records
    ├── design/         # designs under exploration
    ├── domain/         # business context: customers, SLAs, seasonality, domain rules
    ├── incident/       # blameless incident records
    ├── reference/      # technical reference + raw source material
    └── draft/          # gitignored scratch
```

## What Goes Where

- `README.md` — entry point and quick orientation.
- `docs/CONTEXT.md` — glossary and shared terminology, nothing else. No implementation detail, no specs, no scratch notes. See the Context Format template below.
- `docs/adr/` — decisions that are hard to reverse or hard to understand without context. See the ADR Format template below.
- `docs/design/` — designs and tradeoffs still being worked through.
- `docs/domain/` — business context: who uses this system, under what constraints, seasonal patterns, SLAs.
- `docs/incident/` — blameless incident records. See the Incident Format template below.
- `docs/reference/` — technical reference material and raw source documents worth keeping near the project. Store verbatim; optionally add a separate summary that cites it.
- `docs/draft/` — early thoughts, rough plans, personal scratch. Gitignored. Not team-facing.

## Naming Conventions

| Doc type | Pattern | Example |
|---|---|---|
| ADR | `YYYYMMDD-hhmm-<slug>.md` | `20260525-1430-use-postgres-for-write-model.md` |
| Incident | `YYYYMMDD-hhmm-<slug>.md` | `20260525-0910-payment-timeout-prod.md` |
| Draft | `Thoughts-<topic>.md` / `Problem-<topic>.md` / `PLAN-<topic>.md` | `PLAN-auth-rework.md` |
| Everything else | lowercase kebab-case | `checkout-flow.md` |

**Getting the timestamp:** Run the date-getting shell command yourself first (e.g., `Get-Date -Format "yyyyMMdd-HHmm"` in PowerShell or `date +%Y%m%d-%H%M` in bash). Only ask the engineer if command execution is unavailable or fails.

---

## Cross-Cutting Rules for All Commands

1. **Never auto-commit.** All output is a working-tree change for the engineer to review.
2. **Lazy folder creation.** Only create `docs/adr/`, `docs/incident/`, etc. when the first file in that folder is needed.
3. **ADRs and incidents are immutable.** Never edit original content. Corrections and follow-ups are appended as dated notes at the bottom of the file.
4. **Raw sources go in `reference/` verbatim.** When the engineer provides third-party API docs, spec excerpts, or other source material, offer to store the raw source (preferred) plus optionally a summary that cites it.
5. **Conflict with existing docs.** If new content contradicts an existing ADR, surface the conflict before writing. If confirmed outdated, append a dated correction note — do not rewrite the original.
6. **Default to `docs/draft/`.** When unsure whether a doc should exist, write to `docs/draft/` first and ask before graduating to `docs/`.
7. **Human intent before automation.** Ask when the target doc type is unclear.
8. **Author generated document content in the configured locale.** The embedded format templates are already localized; match their language when filling in surrounding prose.

---

## Format Templates

### ADR Format Template

```md
# {Short title of the decision}

{1–3 sentences: what's the context, what was decided, and why.}

That's it. An ADR can be a single paragraph. The value is in recording *that* a decision was made and *why* — not in filling out every section.

## Optional sections

Only include these when they add genuine value. Most ADRs won't need them.

- **Status** — `proposed | accepted | deprecated | superseded by YYYYMMDD-hhmm-<slug>` — useful when decisions are revisited.
- **Considered options** — only when the rejected alternatives are worth preserving.
- **Consequences** — only when non-obvious downstream effects need to be called out.

## When to offer an ADR

All three must be true:

1. **Hard to reverse** — the cost of changing your mind later is meaningful.
2. **Surprising without context** — a future reader will look at the code and wonder "why on earth did they do it this way?"
3. **The result of a real trade-off** — there were genuine alternatives and you picked one for specific reasons.

If any of the three is missing, skip the ADR.

### What qualifies

- **Architectural shape.** "We're using a monorepo." "The write model is event-sourced."
- **Integration patterns between components.** "These two services communicate via domain events, not HTTP."
- **Technology choices that carry lock-in.** Database, message bus, auth provider, deployment target. Not every library — just the ones that would take a quarter to swap out.
- **Boundary and scope decisions.** "Customer data is owned by the Customer context; other contexts reference it by ID only."
- **Deliberate deviations from the obvious path.** "We're using manual SQL instead of an ORM because X." Anything where a reasonable reader would assume the opposite.
- **Constraints not visible in the code.** "We can't use AWS because of compliance requirements." "Response times must be under 200ms because of a partner API contract."
- **Rejected alternatives when the rejection is non-obvious.** If you considered GraphQL and picked REST for subtle reasons, record it — otherwise someone will suggest GraphQL again in six months.

## Immutability

ADRs are immutable. Do not edit an ADR's original content after it is written. If a decision is revisited:

- Append a dated note to the existing ADR acknowledging the change.
- Write a new ADR for the new decision.
- Update the existing ADR's **Status** to `superseded by YYYYMMDD-hhmm-<new-slug>`.
```

### Context Format Template (Glossary)

```md
# Context

{One or two sentences: what domain or system this context describes and why the terminology matters here.}

## Language

**Order**:
A confirmed intent to purchase placed by a Customer.
_Avoid_: Purchase, transaction

**Invoice**:
A request for payment sent to a Customer after delivery.
_Avoid_: Bill, payment request

**Customer**:
A person or organization that places Orders.
_Avoid_: Client, buyer, account

## Rules

- **Be opinionated.** When multiple words exist for the same concept, pick one and list the others as aliases to avoid.
- **Flag conflicts explicitly.** If a term is used ambiguously across the codebase, call it out under a "Flagged ambiguities" subheading with a clear resolution.
- **Keep definitions tight.** One or two sentences maximum. Define what it *is*, not what it does.
- **Only include terms specific to this domain.** General programming concepts (timeouts, error types, retry logic) do not belong even if the project uses them heavily. Before adding a term, ask: is this concept unique to this domain, or general? Only the former belongs.
- **Show relationships.** Express cardinality or ownership where obvious (e.g., "An Order contains one or more LineItems.").
- **Group by subheadings** when natural clusters emerge. A single flat `## Language` section is fine for a cohesive domain.

## Example

Dev: "So when a user clicks 'buy', we create an Order?"
Expert: "Yes — an Order is confirmed intent to purchase. The Cart becomes an Order at checkout."
Dev: "And the line items... are those OrderItems?"
Expert: "LineItems. Items is too vague — it could mean anything."
Dev: "What about the person buying? User or Customer?"
Expert: "Customer once they've placed an Order. Before that we don't have a word for them."
```

### Incident Format Template

```md
# {Short title of the incident}

**Occurred:** YYYYMMDD-hhmm
**Level:** while-dev | while-test | while-production

{1–2 sentence summary of what happened.}

## Symptoms

{What the engineer observed. What looked wrong. What errors or alerts fired.}

## Timeline

{Best-effort sequence of events. Mark uncertain entries with "approx:" and anything the engineer described from memory with "user reported:".}

## Observations

{Facts confirmed during the investigation — things you can point to in logs, code, or output. Keep this separate from conclusions.}

## Conclusion

{Root cause as best understood. If unknown, say so explicitly and list what was ruled out. An honest "root cause unknown — ruled out X and Y" is more useful than a confident wrong answer.}

## Contributing factors

{What made this possible or worse? Only include when non-obvious.}

## What we learned

{Things a future engineer should know when they see similar symptoms. Not action items — those live elsewhere.}

## Related

{Links to related ADRs, design docs, or other incidents. Optional.}

## Immutability

Incident records are immutable. Do not edit the original content after it is written.
Later corrections or follow-up findings are **appended** at the bottom as a dated section:

### Follow-up (YYYYMMDD-hhmm)

{What was learned, corrected, or resolved after the original entry was written.}
```
