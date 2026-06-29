// Copyright 2026 Zero Day AI, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
