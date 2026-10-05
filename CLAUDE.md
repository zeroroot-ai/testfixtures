# testfixtures — CLAUDE.md

> **Workflow rules:** see [`zeroroot-ai/.github` → `AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md) — canonical for branching / commits / PRs / releases / merging. Conventional Commits MANDATORY. Never push to main. Never force-push.

This file is the per-repo addendum. Workspace-wide concerns live in the workspace `CLAUDE.md`; architectural decisions in ``docs/adr/`` (local docs → `adr`).

## TL;DR

Test fakes shared across the workspace (`github.com/zeroroot-ai/testfixtures`). One fake today: `fga.FakeStore`, an in-memory OpenFGA-like tuple store. Imported as a test dependency by gibson. Not a runnable binary.

## Architecture

One package per domain (`fga/`). A fake lives here only while a consumer in another repository reads it. `make lint-unwired` measures the surface with ast-checks, and each entry in `.unwired-baseline.txt` names its consumer on a `#` line. The 2026-10-02 burn-down (#10) deleted `authz`, `audit`, `tenant` and `spiffe`: no org repository read them, and the interfaces they mirrored do not exist in gibson. **Public** repo under the Elastic License 2.0, and a gibson test dependency. The Elastic License 2.0 is source-available, not open source, so the permissive tier (`sdk`, `adk`, `setec`, `ast-checks`) must not import it.

## Regen commands

No proto/codegen. Standard Go build:

```bash
make build
make check                 # fmt vet test-race lint-unwired lint-unwired-selftest
make lint-unwired-write    # re-measure #10; put the # consumer lines back before you commit
```

## Gotchas

- The scanner counts reads inside this module only. A consumer in gibson is not a read it can see, so the baseline names the consumer instead. Grep gibson `origin/main` before you delete an exported symbol.
- `make lint-unwired-write` drops the `#` reason lines. Put them back.

## Links

- Org-level workflow: [`AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md)
- Workspace map: workspace `CLAUDE.md`
