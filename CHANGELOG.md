# Changelog

## [0.3.0](https://github.com/zeroroot-ai/testfixtures/compare/v0.2.0...v0.3.0) (2026-10-02)


### ⚠ BREAKING CHANGES

* **deadcode:** the authz, tenant, spiffe and audit packages are gone, and fga.FakeStore loses Tuples and Reset. gibson internal/engine/harness/authorize_testfixtures_test.go drops the tfxaudit import and the two lines that construct and discard the emitter when it bumps to this release.

### Bug Fixes

* **ci:** link-check checks only the Markdown a PR touched (.github v0.7.2) ([#9](https://github.com/zeroroot-ai/testfixtures/issues/9)) ([88593ce](https://github.com/zeroroot-ai/testfixtures/commit/88593ceb2aa78915846d45b862012f586199ca31))
* **ci:** pin every zeroroot-ai/.github reference to v0.5.1 ([#8](https://github.com/zeroroot-ai/testfixtures/issues/8)) ([e7b092b](https://github.com/zeroroot-ai/testfixtures/commit/e7b092b719f8bb46815c7e883bf384fab27cd545))
* **ci:** pin the org tree guards to a commit SHA ([#6](https://github.com/zeroroot-ai/testfixtures/issues/6)) ([c8c8654](https://github.com/zeroroot-ai/testfixtures/commit/c8c86542db7572a28d3ad943c6ef350fdd81a3dc))


### Miscellaneous Chores

* **deadcode:** delete the four fakes no consumer can use ([#14](https://github.com/zeroroot-ai/testfixtures/issues/14)) ([ba64970](https://github.com/zeroroot-ai/testfixtures/commit/ba649708bda6d679f4bdac69371e86d3ff9e9196))
