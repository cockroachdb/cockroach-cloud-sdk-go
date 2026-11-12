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
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/cockroachdb/cockroach-cloud-sdk-go/v10/pkg/client"
)

type timestampTransport func(*http.Request) (*http.Response, error)

func (f timestampTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestDatetimeQueryPrecision(t *testing.T) {
	for _, operation := range []string{"audit log cursor", "API key pagination snapshot"} {
		for _, timestamp := range []string{
			"2026-10-25T12:34:56.123456789Z",
			"2026-10-25T12:34:56.123456Z",
			"2026-10-25T12:34:56.123456789+05:30",
			"2026-10-25T12:34:56Z",
		} {
			t.Run(operation+"/"+timestamp, func(t *testing.T) {
				input, err := time.Parse(time.RFC3339Nano, timestamp)
				if err != nil {
					t.Fatal(err)
				}
				var got string
				cfg := client.NewConfiguration("test-token")
				cfg.HTTPClient = &http.Client{Transport: timestampTransport(func(r *http.Request) (*http.Response, error) {
					key := "starting_from"
					if operation == "API key pagination snapshot" {
						key = "pagination.as_of_time"
					}
					got = r.URL.Query().Get(key)
					return &http.Response{StatusCode: 200, Status: "200 OK", Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(`{}`)), Request: r}, nil
				})}
				service := client.NewService(client.NewClient(cfg))
				if operation == "audit log cursor" {
					_, _, err = service.ListAuditLogs(context.Background(), &client.ListAuditLogsOptions{StartingFrom: &input})
				} else {
					_, _, err = service.ListApiKeys(context.Background(), &client.ListApiKeysOptions{PaginationAsOfTime: &input})
				}
				if err != nil {
					t.Fatal(err)
				}
				actual, err := time.Parse(time.RFC3339Nano, got)
				if err != nil {
					t.Fatalf("invalid datetime query %q: %v", got, err)
				}
				if !actual.Equal(input) {
					t.Fatalf("query timestamp %q lost precision from %q", got, timestamp)
				}
				if got != timestamp {
					t.Fatalf("got %q, want %q", got, timestamp)
				}
			})
		}
	}
}
