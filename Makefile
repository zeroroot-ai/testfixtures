# Per-repo Makefile contract (per zeroroot-ai polyrepo convention).
# Targets: build / test / test-race / check / image (n/a here).

# ast-checks ships the per-declaration read counter behind #10. Pinned, because
# a floating version would change the count without a commit.
UNWIRED_VERSION ?= v0.4.0
export UNWIRED_VERSION

.PHONY: build test test-race check fmt vet lint lint-unwired lint-unwired-write lint-unwired-selftest

build:
	go build ./...

test:
	go test ./...

test-race:
	go test -race ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	@which golangci-lint >/dev/null 2>&1 || (echo "golangci-lint not installed"; exit 1)
	golangci-lint run

check: fmt vet test-race lint-unwired lint-unwired-selftest

# #10 is the standing tracker for declarations nothing reads. The baseline only
# shrinks, so adding one fails here. A consumer in another repository does not
# count as a read; it is named on a # line above the entry instead.
lint-unwired:
	go run github.com/zeroroot-ai/ast-checks/cmd/unwired@$(UNWIRED_VERSION) -dir . -baseline .unwired-baseline.txt

# Re-measure #10 and rewrite the baseline. -write drops the # reason lines, so
# put them back before you commit.
lint-unwired-write:
	go run github.com/zeroroot-ai/ast-checks/cmd/unwired@$(UNWIRED_VERSION) -dir . -baseline .unwired-baseline.txt -write

# Proves lint-unwired can fail, and that the failure names the declaration.
lint-unwired-selftest:
	bash scripts/check-unwired-fails.sh
