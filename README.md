# testfixtures

Canonical test fakes shared across the workspace (fakeAuthorizer, fakeFGAStore, fakeAuditEmitter, fakeTenantClient, fakeSPIFFEBundle). Each fake compile-time-asserts its real-interface conformance.

Internal Go module under the zero-day-ai workspace. See [`zero-day-ai/.github` → `AGENTS.md`](https://github.com/zero-day-ai/.github/blob/main/AGENTS.md) for workflow conventions (branching, PRs, releases, agent merge autonomy).

## Status

Bootstrap repo. Initial implementation lands via the corresponding production-readiness slice on board #16. Until then, this README + LICENSE + Makefile contract are the only contents.

## Install

```bash
go get github.com/zero-day-ai/testfixtures@latest
```

## License

[BUSL-1.1](./LICENSE).
