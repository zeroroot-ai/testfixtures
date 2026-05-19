// Copyright 2026 Zero Day AI, Inc.
// Licensed under BUSL-1.1; see LICENSE.

// Package spiffe provides FakeBundle — a SPIFFE bundle source for tests.
// Returns configurable identities + JWKS-style bundle bytes. Slice 5.4.
package spiffe

import "sync"

// FakeBundle is a configurable SPIFFE bundle / SVID source.
type FakeBundle struct {
	mu       sync.Mutex
	identity string
	bundle   []byte
}

// NewFakeBundle constructs a bundle with the given SPIFFE ID + bundle bytes.
func NewFakeBundle(id string, bundle []byte) *FakeBundle {
	return &FakeBundle{identity: id, bundle: bundle}
}

// SPIFFEID returns the configured identity.
func (f *FakeBundle) SPIFFEID() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.identity
}

// Bundle returns the configured bundle bytes (typically JWKS).
func (f *FakeBundle) Bundle() []byte {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]byte, len(f.bundle))
	copy(out, f.bundle)
	return out
}

// SetIdentity rotates the configured identity (for tests that simulate
// SVID rotation).
func (f *FakeBundle) SetIdentity(id string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.identity = id
}
