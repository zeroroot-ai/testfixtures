#!/usr/bin/env bash
# Proves `make lint-unwired` can fail (#10). A guard that cannot fail is worse
# than none. The script copies the tracked tree to a temp dir, adds an exported
# type nothing reads, runs the scanner against the committed baseline, and
# requires the failure to name that type and its field.
set -euo pipefail

root=$(cd "$(dirname "$0")/.." && pwd)
version=${UNWIRED_VERSION:?set UNWIRED_VERSION, the Makefile exports it}
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

git -C "$root" ls-files -z --cached --others --exclude-standard | tar --null -C "$root" -cf - --files-from=- | tar -C "$tmp" -xf -
mkdir "$tmp/zzprobe"
cat > "$tmp/zzprobe/probe.go" <<'GO'
package zzprobe

// Probe is read nowhere. The self-test plants it and expects the scanner to fail.
type Probe struct {
	Field int
}
GO

set +e
out=$(cd "$tmp" && go run "github.com/zeroroot-ai/ast-checks/cmd/unwired@$version" -dir . -baseline .unwired-baseline.txt 2>&1)
rc=$?
set -e

if [ "$rc" -eq 0 ]; then
	echo "FAIL: lint-unwired passed with an unread declaration planted" >&2
	printf '%s\n' "$out" >&2
	exit 1
fi
for want in 'type[[:space:]]+zzprobe\.Probe\b' 'field[[:space:]]+zzprobe\.Probe\.Field\b'; do
	if ! grep -Eq "$want" <<<"$out"; then
		echo "FAIL: lint-unwired failed but did not name $want" >&2
		printf '%s\n' "$out" >&2
		exit 1
	fi
done
echo "ok  lint-unwired fails on a planted unread declaration and names it (exit $rc)"
