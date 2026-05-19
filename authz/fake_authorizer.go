// Copyright 2026 Zero Day AI, Inc.
// Licensed under BUSL-1.1; see LICENSE.

// Package authz provides FakeAuthorizer — a test double that satisfies
// the shape of gibson's internal authz.Authorizer interface without
// importing internal packages (which is forbidden from outside the gibson
// module).
//
// Go's structural typing means callers can use FakeAuthorizer wherever
// an interface with the same method set is expected. This avoids the
// import-internal-package problem while still giving tests a
// record-and-replay fake.
//
// Slice 5.4 of the production-readiness epic (gibson#175 → board #16).
// Closes the deploy#230 Sub-slice A precondition.
package authz

import (
	"context"
	"sync"
)

// Decision is the answer FakeAuthorizer returns. Mirrors the shape of
// gibson's internal authz Decision type.
type Decision int

const (
	// DecisionAllow grants the operation.
	DecisionAllow Decision = iota
	// DecisionDeny denies the operation.
	DecisionDeny
)

// CheckCall records a single Check invocation. Used for assertion
// patterns like "the handler called Check with this subject + relation".
type CheckCall struct {
	Subject    string
	Relation   string
	ObjectType string
	ObjectID   string
}

// FakeAuthorizer records every Check call + returns a configurable
// Decision. Default decision is Allow; override per-test via
// SetDefault or via PerCall responses.
type FakeAuthorizer struct {
	mu       sync.Mutex
	calls    []CheckCall
	deflt    Decision
	perCall  map[string]Decision // key: "<relation>:<objectType>"
	failNext error
}

// NewFakeAuthorizer returns a fake with default Allow.
func NewFakeAuthorizer() *FakeAuthorizer {
	return &FakeAuthorizer{
		deflt:   DecisionAllow,
		perCall: make(map[string]Decision),
	}
}

// SetDefault sets the default Decision for Check calls.
func (f *FakeAuthorizer) SetDefault(d Decision) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.deflt = d
}

// SetDecision overrides the Decision for a specific (relation, objectType).
// Other call shapes still use the default.
func (f *FakeAuthorizer) SetDecision(relation, objectType string, d Decision) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.perCall[relation+":"+objectType] = d
}

// FailNext arms the fake to return err on the next Check call (then clears).
func (f *FakeAuthorizer) FailNext(err error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.failNext = err
}

// Check records the call and returns the configured decision.
func (f *FakeAuthorizer) Check(ctx context.Context, subject, relation, objectType, objectID string) (Decision, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = append(f.calls, CheckCall{
		Subject: subject, Relation: relation, ObjectType: objectType, ObjectID: objectID,
	})
	if f.failNext != nil {
		err := f.failNext
		f.failNext = nil
		return DecisionDeny, err
	}
	if d, ok := f.perCall[relation+":"+objectType]; ok {
		return d, nil
	}
	return f.deflt, nil
}

// Calls returns a copy of the recorded Check calls in invocation order.
func (f *FakeAuthorizer) Calls() []CheckCall {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]CheckCall, len(f.calls))
	copy(out, f.calls)
	return out
}

// Reset clears recorded calls + restores defaults. Useful between subtests.
func (f *FakeAuthorizer) Reset() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = nil
	f.failNext = nil
	for k := range f.perCall {
		delete(f.perCall, k)
	}
	f.deflt = DecisionAllow
}
