# testfixtures — CLAUDE.md

> **Workflow rules:** see [`zeroroot-ai/.github` → `AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md) — canonical for branching / commits / PRs / releases / merging. Conventional Commits MANDATORY. Never push to main. Never force-push.

This file is the per-repo addendum. Workspace-wide concerns live in the workspace `CLAUDE.md`; architectural decisions in ``docs/adr/`` (local docs → `adr`).

## TL;DR

Canonical test fakes shared across the workspace (`github.com/zeroroot-ai/testfixtures`): `fakeAuthorizer`, `fakeFGAStore`, `fakeAuditEmitter`, `fakeTenantClient`, `fakeSPIFFEBundle`. Each fake compile-time-asserts conformance to its real interface. Imported as a test dependency; not a runnable binary.

## Architecture

Per-domain packages (`authz/`, `fga/`, `audit/`, `tenant/`, `spiffe/`) each export a fake plus a `var _ RealInterface = (*fakeX)(nil)` assertion so the fake breaks the build if the real interface drifts. **Public** repo under the Elastic License 2.0, and a gibson test dependency. The Elastic License 2.0 is source-available, not open source, so the permissive tier (`sdk`, `adk`, `setec`, `ast-checks`) must not import it.

## Regen commands

No proto/codegen. Standard Go build:

```bash
make build
make check        # fmt vet test-race
```

## Gotchas

- A fake silently passing a stale interface is the failure mode the compile-time assertions exist to prevent — never delete the `var _ Real = (*fake)(nil)` lines.

## Links

- Org-level workflow: [`AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md)
- Workspace map: workspace `CLAUDE.md`
- Open-core licensing: ``docs/architecture/open-core/`` (local docs → `architecture/open-core`)
- PR checklist: ``docs/agents/pr-checklist.md`` (local docs → `agents/pr-checklist.md`)
