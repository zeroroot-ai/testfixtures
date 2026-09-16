# testfixtures

Canonical test fakes shared across the workspace (fakeAuthorizer, fakeFGAStore, fakeAuditEmitter, fakeTenantClient, fakeSPIFFEBundle). Each fake compile-time-asserts its real-interface conformance.

Go module under the zeroroot-ai workspace. See [`zeroroot-ai/.github` → `AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md) for workflow conventions (branching, PRs, releases, agent merge autonomy).

## Install

```bash
go get github.com/zeroroot-ai/testfixtures@latest
```

## License and history

Elastic License 2.0. See [LICENSE](LICENSE). Zero Root AI is the licensor. The Elastic License 2.0 is source-available, not open source.

Issue and pull request numbers cited in comments and documents dated before 2026-09-05 refer to the tracker before the history reset, archived offline. They do not resolve on GitHub.
