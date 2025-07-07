// Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package openapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	owner      = "111"
	repo       = "222"
	number     = "1"
	page       = "1"
	branch     = "master"
	permission = "push"
)

// setup sets up a test HTTP server along with a github.api that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the api method being tested.
func mockServer(t *testing.T) (client *APIClient, mux *http.ServeMux, serverURL string) {
	t.Helper()
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	handlerPath := "/api/v5/"
	apiHandler.Handle(handlerPath, http.StripPrefix(handlerPath[:len(handlerPath)-1], mux))

	// server is a test HTTP server used to provide mock api responses.
	server := httptest.NewServer(apiHandler)

	// api is the GitHub api being tested and is
	// configured to use test server.
	client = NewAPIClientWithAuthorization([]byte("1111111111"))
	uri, _ := url.Parse(server.URL + handlerPath)
	client.baseURL = uri

	t.Cleanup(server.Close)

	return client, mux, server.URL
}

func mockResponse(t *testing.T, mux *http.ServeMux, urlStr string, body any) {
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		if body != nil {
			err := json.NewEncoder(w).Encode(body)
			if err != nil {
				t.Errorf("mock response data error: %v", err)
			}
		}
	})
}
