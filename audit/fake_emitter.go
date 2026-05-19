// Copyright 2026 Zero Day AI, Inc.
// Licensed under BUSL-1.1; see LICENSE.

// Package audit provides FakeEmitter — records audit events for assertion
// in tests. Slice 5.4.
package audit

import (
	"context"
	"sync"
)

// Event is the recorded shape of an audit emission.
type Event struct {
	Type    string
	Subject string
	Tenant  string
	Object  string
	Details map[string]string
}

// FakeEmitter records every Emit call.
type FakeEmitter struct {
	mu     sync.Mutex
	events []Event
}

// NewFakeEmitter constructs an empty recorder.
func NewFakeEmitter() *FakeEmitter {
	return &FakeEmitter{}
}

// Emit records the event.
func (f *FakeEmitter) Emit(ctx context.Context, e Event) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.events = append(f.events, e)
	return nil
}

// Events returns recorded events in emit order.
func (f *FakeEmitter) Events() []Event {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]Event, len(f.events))
	copy(out, f.events)
	return out
}

// Reset clears recorded events.
func (f *FakeEmitter) Reset() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.events = nil
}
