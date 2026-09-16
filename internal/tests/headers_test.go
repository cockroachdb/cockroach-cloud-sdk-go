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

// listClusters calls an endpoint against a test server and returns the headers
// the server received.
func listClusters(t *testing.T, opts ...client.ConfigurationOption) http.Header {
	t.Helper()

	var got http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Clone()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"clusters":[]}`))
	}))
	defer server.Close()

	cfg := client.NewConfiguration("token", opts...)
	cfg.ServerURL = server.URL

	service := client.NewService(client.NewClient(cfg))
	_, _, err := service.ListClusters(context.Background(), &client.ListClustersOptions{})
	if err != nil {
		t.Fatalf("ListClusters: %v", err)
	}
	return got
}

// setCcClient stands in for a wrapping client overriding the default once it
// holds the configuration.
func setCcClient(key, value string) client.ConfigurationOption {
	return func(cfg *client.Configuration) {
		cfg.AddDefaultHeader(key, value)
	}
}

func TestCcClientHeader(t *testing.T) {
	for _, tc := range []struct {
		name string
		opts []client.ConfigurationOption
		want string
	}{
		{name: "default", want: client.CcClientSDK},
		{
			name: "wrapping client overrides the default",
			opts: []client.ConfigurationOption{setCcClient(client.CcClientHeader, "terraform")},
			want: "terraform",
		},
		{
			name: "override under a non-canonical key",
			opts: []client.ConfigurationOption{setCcClient("cc-client", "ccloud_cli")},
			want: "ccloud_cli",
		},
		{
			name: "last write wins",
			opts: []client.ConfigurationOption{
				setCcClient("cc-client", "first"),
				setCcClient(client.CcClientHeader, "second"),
			},
			want: "second",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			values := listClusters(t, tc.opts...).Values(client.CcClientHeader)
			if len(values) != 1 || values[0] != tc.want {
				t.Errorf("got %q, want [%q]", values, tc.want)
			}
		})
	}
}

// The Cc-Client default leaves the other configured headers alone.
func TestHeadersCoexist(t *testing.T) {
	header := listClusters(t,
		setCcClient(client.CcClientHeader, "terraform"),
		client.WithVanityName("my-org"),
		client.WithUsername("someone"),
	)

	for name, want := range map[string]string{
		client.CcClientHeader: "terraform",
		"Cc-Version":          client.ApiVersion,
		"Cc-Vanity-Name":      "my-org",
		"Cc-Username":         "someone",
	} {
		values := header.Values(name)
		if len(values) != 1 || values[0] != want {
			t.Errorf("%s: got %q, want [%q]", name, values, want)
		}
	}
}

// A default header replaces the value the request set for itself rather than
// appending a second one.
func TestDefaultHeaderReplaces(t *testing.T) {
	header := listClusters(t, func(cfg *client.Configuration) {
		cfg.AddDefaultHeader("Accept", "application/custom")
		cfg.AddDefaultHeader("User-Agent", "wrapping-client/1.0.0")
	})

	for name, want := range map[string]string{
		"Accept":     "application/custom",
		"User-Agent": "wrapping-client/1.0.0",
	} {
		values := header.Values(name)
		if len(values) != 1 || values[0] != want {
			t.Errorf("%s: got %q, want [%q]", name, values, want)
		}
	}
}
