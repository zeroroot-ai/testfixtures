// SPDX-License-Identifier: Elastic-2.0
// Copyright 2026 Zero Root AI

// Package fga provides FakeStore, an in-memory OpenFGA-like tuple store for
// tests. gibson's harness tests back a test Authorizer with it.
package fga

import (
	"context"
	"sync"
)

// Tuple is the (user, relation, object) triple, OpenFGA-style. It is the map
// key of the store, so all three fields take part in Check.
type Tuple struct {
	User     string
	Relation string
	Object   string
}

// FakeStore is an in-memory tuple store.
type FakeStore struct {
	mu     sync.Mutex
	tuples map[Tuple]struct{}
}

// NewFakeStore constructs an empty store.
func NewFakeStore() *FakeStore {
	return &FakeStore{tuples: make(map[Tuple]struct{})}
}

// Write adds a tuple to the store. Idempotent.
func (f *FakeStore) Write(ctx context.Context, t Tuple) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.tuples[t] = struct{}{}
	return nil
}

// Delete removes a tuple from the store. No-op if not present.
func (f *FakeStore) Delete(ctx context.Context, t Tuple) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.tuples, t)
	return nil
}

// Check returns true iff the tuple exists.
func (f *FakeStore) Check(ctx context.Context, t Tuple) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	_, ok := f.tuples[t]
	return ok, nil
}
