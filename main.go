/*
Copyright 2023 The Cockroach Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cockroachdb/cockroach-cloud-sdk-go/pkg/client"
)

type requestPrinter struct{}

func (c *requestPrinter) RoundTrip(req *http.Request) (*http.Response, error) {
	body, _ := io.ReadAll(req.Body)
	fmt.Print(string(body))
	return nil, nil
}

func main() {
	c := client.NewClient(&client.Configuration{HTTPClient: &http.Client{Transport: &requestPrinter{}}})
	s := client.NewService(c)
	ul := &client.UsageLimits{}
	req := &client.UpdateClusterSpecification{
		Dedicated: nil,
		Serverless: &client.ServerlessClusterUpdateSpecification{
			UsageLimits: ul,
		},
		UpgradeStatus: nil,
	}
	req.Serverless.SetSpendLimit(nil, true)
	s.UpdateCluster(context.Background(), "cluster-id", req)
	req.Serverless.SetSpendLimit(nil, false)
	s.UpdateCluster(context.Background(), "cluster-id", req)
	// Output:
	// {"serverless":{"spend_limit":null,"usage_limits":{"request_unit_limit":0,"storage_mib_limit":0}}}
	// {"serverless":{"usage_limits":{"request_unit_limit":0,"storage_mib_limit":0}}}

}

/*
func main() {
	_ = client.NewClient(&client.Configuration{})
	fmt.Println("OK!")

}
*/
