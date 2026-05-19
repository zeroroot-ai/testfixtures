// Copyright 2026 Zero Day AI, Inc.
// Licensed under BUSL-1.1; see LICENSE.

package authz

import (
	"context"
	"errors"
	"testing"
)

func TestFakeAuthorizer_DefaultAllow(t *testing.T) {
	f := NewFakeAuthorizer()
	d, err := f.Check(context.Background(), "alice", "viewer", "tenant", "acme")
	if err != nil || d != DecisionAllow {
		t.Errorf("default should be (Allow, nil); got (%v, %v)", d, err)
	}
	if len(f.Calls()) != 1 {
		t.Errorf("expected 1 recorded call; got %d", len(f.Calls()))
	}
}

func TestFakeAuthorizer_PerCall(t *testing.T) {
	f := NewFakeAuthorizer()
	f.SetDecision("admin", "tenant", DecisionDeny)
	d, _ := f.Check(context.Background(), "alice", "admin", "tenant", "acme")
	if d != DecisionDeny {
		t.Errorf("admin:tenant should be Deny; got %v", d)
	}
	d, _ = f.Check(context.Background(), "alice", "viewer", "tenant", "acme")
	if d != DecisionAllow {
		t.Errorf("viewer:tenant should be Allow (default); got %v", d)
	}
}

func TestFakeAuthorizer_FailNext(t *testing.T) {
	f := NewFakeAuthorizer()
	sentinel := errors.New("authz unavailable")
	f.FailNext(sentinel)
	_, err := f.Check(context.Background(), "", "", "", "")
	if !errors.Is(err, sentinel) {
		t.Errorf("expected sentinel error; got %v", err)
	}
	// Next call should succeed (FailNext clears after one use)
	d, err := f.Check(context.Background(), "", "", "", "")
	if err != nil || d != DecisionAllow {
		t.Errorf("after FailNext, next call should succeed; got (%v, %v)", d, err)
	}
}

func TestFakeAuthorizer_Reset(t *testing.T) {
	f := NewFakeAuthorizer()
	f.Check(context.Background(), "", "", "", "")
	f.SetDefault(DecisionDeny)
	f.SetDecision("x", "y", DecisionDeny)
	f.Reset()
	if len(f.Calls()) != 0 {
		t.Errorf("Reset should clear calls")
	}
	d, _ := f.Check(context.Background(), "", "", "", "")
	if d != DecisionAllow {
		t.Errorf("Reset should restore default Allow; got %v", d)
	}
}
