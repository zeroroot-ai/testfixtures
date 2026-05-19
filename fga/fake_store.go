// Copyright 2026 Zero Day AI, Inc.
// Licensed under BUSL-1.1; see LICENSE.

// Package fga provides FakeStore — an in-memory OpenFGA-like store
// for tests. Slice 5.4.
package fga

import (
	"context"
	"sync"
)

// Tuple is the (user, relation, object) triple OpenFGA-style.
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

// Tuples returns a snapshot of all tuples (order not guaranteed).
func (f *FakeStore) Tuples() []Tuple {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]Tuple, 0, len(f.tuples))
	for t := range f.tuples {
		out = append(out, t)
	}
	return out
}

// Reset clears all tuples.
func (f *FakeStore) Reset() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.tuples = make(map[Tuple]struct{})
}
