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
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIssues(t *testing.T) {
	client, mux, _ := mockServer(t)

	var want Issue
	_ = readTestdata(t, issuesTestDataDir+"issues_update.json", &want)

	createIssue(t, client, mux, want)
	getIssue(t, client, mux, want)
	updateIssue(t, client, mux, want)
	getLinkedPR(t, client, mux)
}

func createIssue(t *testing.T, client *APIClient, mux *http.ServeMux, want Issue) {

	urlStr := fmt.Sprintf("/repos/%s/issues", owner)
	mockResponse(t, mux, urlStr, want)
	got, ok, err := client.Issues.CreateIssue(context.Background(), owner, &IssueRequest{})
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, want, *got)
}

func getIssue(t *testing.T, client *APIClient, mux *http.ServeMux, want Issue) {

	urlStr := fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repo, number)
	mockResponse(t, mux, urlStr, want)
	got, ok, err := client.Issues.GetIssue(context.Background(), owner, repo, number)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, want, *got)
}

func updateIssue(t *testing.T, client *APIClient, mux *http.ServeMux, want Issue) {

	urlStr := fmt.Sprintf("/repos/%s/issues/%s", owner, number)
	mockResponse(t, mux, urlStr, want)
	got, ok, err := client.Issues.UpdateIssue(context.Background(), owner, number, &IssueRequest{
		Repository: repo,
		Title:      "issue1",
	})
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, want, *got)

	urlStr = fmt.Sprintf("/repos/%s/issues/%s", owner, "2")
	errMsg := "{\n    \"error_code\": 403,\n    \"error_code_name\": \"FORBIDDEN\",\n    \"error_message\": \"no scopes:read_projects\",\n    \"trace_id\": \"33809e888a654b78bb2be8e7c97c9423\"\n}"
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, errMsg, http.StatusBadRequest)
	})
	got, ok, err = client.Issues.UpdateIssue(context.Background(), owner, "2", &IssueRequest{
		Repository: repo,
		Title:      "issue2",
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, Issue{}, *got)
	assert.Equal(t, errMsg+"\n", err.Error())
}

func getLinkedPR(t *testing.T, client *APIClient, mux *http.ServeMux) {
	prs := new([]*PullRequest)
	_ = readTestdata(t, issuesTestDataDir+"issues_linking_prs.json", prs)
	urlStr := fmt.Sprintf("/repos/%s/%s/issues/%s/pull_requests", owner, repo, number)
	mockResponse(t, mux, urlStr, prs)
	result1, ok, err := client.Issues.ListIssueLinkingPullRequests(context.Background(), owner, repo, number)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range *prs {
		d1, _ := json.Marshal(*(*prs)[i])
		d2, _ := json.Marshal(*result1[i])
		assert.Equal(t, d1, d2)
	}
}
