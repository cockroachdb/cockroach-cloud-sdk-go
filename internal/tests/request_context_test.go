// Copyright 2023 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cockroachdb/cockroach-cloud-sdk-go/v10/pkg/client"
)

func TestContextWithHeader(t *testing.T) {
	for _, tc := range []struct {
		name string
		ctx  context.Context
		want map[string]string
	}{
		{
			name: "sent with the request",
			ctx:  client.ContextWithHeader(context.Background(), "Cc-Example", "true"),
			want: map[string]string{"Cc-Example": "true"},
		},
		{
			name: "key is case-insensitive",
			ctx:  client.ContextWithHeader(context.Background(), "cc-example", "true"),
			want: map[string]string{"Cc-Example": "true"},
		},
		{
			name: "overrides a default header",
			ctx:  client.ContextWithHeader(context.Background(), client.CcClientHeader, "terraform"),
			want: map[string]string{client.CcClientHeader: "terraform"},
		},
		{
			name: "last write wins for the same key",
			ctx: client.ContextWithHeader(
				client.ContextWithHeader(context.Background(), "Cc-Example", "first"),
				"cc-example", "second",
			),
			want: map[string]string{"Cc-Example": "second"},
		},
		{
			name: "different keys accumulate",
			ctx: client.ContextWithHeader(
				client.ContextWithHeader(context.Background(), "Cc-One", "1"),
				"Cc-Two", "2",
			),
			want: map[string]string{"Cc-One": "1", "Cc-Two": "2"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			header := listClusters(t, tc.ctx)
			for name, want := range tc.want {
				values := header.Values(name)
				if len(values) != 1 || values[0] != want {
					t.Errorf("%s: got %q, want [%q]", name, values, want)
				}
			}
		})
	}
}

// Headers the SDK sets for itself are overridable per call, matching what
// AddDefaultHeader already allows client-wide.
func TestContextWithHeaderOverridesSDKHeaders(t *testing.T) {
	ctx := context.Background()
	want := map[string]string{
		"Authorization": "Bearer other-token",
		"Cc-Version":    "2000-01-01",
		"User-Agent":    "wrapping-client/1.0.0",
	}
	for name, value := range want {
		ctx = client.ContextWithHeader(ctx, name, value)
	}

	header := listClusters(t, ctx)
	for name, want := range want {
		if got := header.Get(name); got != want {
			t.Errorf("%s: got %q, want %q", name, got, want)
		}
	}
}

// The motivating case: one client sending the header on some calls and not
// others.
func TestContextWithHeaderIsolation(t *testing.T) {
	var got []http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = append(got, r.Header.Clone())
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"clusters":[]}`))
	}))
	defer server.Close()

	cfg := client.NewConfiguration("token")
	cfg.ServerURL = server.URL
	service := client.NewService(client.NewClient(cfg))

	plain := context.Background()
	withHeader := client.ContextWithHeader(plain, "Cc-Example", "true")

	for _, ctx := range []context.Context{withHeader, plain} {
		if _, _, err := service.ListClusters(ctx, &client.ListClustersOptions{}); err != nil {
			t.Fatalf("ListClusters: %v", err)
		}
	}

	if want := "true"; got[0].Get("Cc-Example") != want {
		t.Errorf("first call: got %q, want %q", got[0].Get("Cc-Example"), want)
	}
	if v := got[1].Get("Cc-Example"); v != "" {
		t.Errorf("second call: got %q, want no header", v)
	}
}
