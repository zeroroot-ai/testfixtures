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

// Package tenant provides FakeClient — a record-and-replay double for
// the per-tenant DB client wrapper used by gibson + tenant-operator.
// Slice 5.4.
package tenant

import (
	"context"
	"sync"
)

// Query is a recorded query invocation.
type Query struct {
	TenantID string
	SQL      string
	Args     []any
}

// FakeClient records query invocations + can be configured to return
// specific row results.
type FakeClient struct {
	mu      sync.Mutex
	queries []Query
}

// NewFakeClient returns an empty recorder.
func NewFakeClient() *FakeClient {
	return &FakeClient{}
}

// QueryContext records the query and returns nil. Real assertion of the
// recorded queries happens via Queries() inspection.
func (f *FakeClient) QueryContext(ctx context.Context, tenantID, sql string, args ...any) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.queries = append(f.queries, Query{TenantID: tenantID, SQL: sql, Args: args})
	return nil
}

// Queries returns recorded queries in invocation order.
func (f *FakeClient) Queries() []Query {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]Query, len(f.queries))
	copy(out, f.queries)
	return out
}

// Reset clears recorded queries.
func (f *FakeClient) Reset() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.queries = nil
}
