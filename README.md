# testfixtures

Canonical test fakes shared across the workspace (fakeAuthorizer, fakeFGAStore, fakeAuditEmitter, fakeTenantClient, fakeSPIFFEBundle). Each fake compile-time-asserts its real-interface conformance.

Internal Go module under the zeroroot-ai workspace. See [`zeroroot-ai/.github` → `AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md) for workflow conventions (branching, PRs, releases, agent merge autonomy).

## Status

Bootstrap repo. Initial implementation lands via the corresponding production-readiness slice on board #16. Until then, this README + LICENSE + Makefile contract are the only contents.

## Install

```bash
go get github.com/zeroroot-ai/testfixtures@latest
```

## License

[BUSL-1.1](./LICENSE).

## License and history

Elastic License 2.0. See [LICENSE](LICENSE). Zero Root AI is the licensor.

Issue and pull request numbers cited in comments and documents dated before 2026-09-05 refer to the tracker before the history reset, archived offline. They do not resolve on GitHub.
