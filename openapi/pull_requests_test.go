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
	"github.com/opensourceways/go-gitcode/testdata"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPullRequests(t *testing.T) {
	client, mux, _ := mockServer(t)

	getPR(t, client, mux)
	updatePR(t, client, mux)
	listPRLinkingIssues(t, client, mux)
	getAllPRCommits(t, client, mux)
	getPRChangeFiles(t, client, mux)
	getPROperationLogs(t, client, mux)
	mergePR(t, client, mux)
}

func getPR(t *testing.T, client *APIClient, mux *http.ServeMux) {

	want := new(PullRequest)
	_ = testdata.ReadTestData(t, testdata.PullRequests, want)

	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	got, ok, err := client.PullRequests.GetPullRequest(context.Background(), owner, repo, number)
	assert.Nil(t, err)
	assert.True(t, ok)
	d1, _ := json.Marshal(*want)
	d2, _ := json.Marshal(*got)
	assert.Equal(t, d1, d2)
	assert.Equal(t, (*string)(nil), got.ClosedAt.ToString())
}

func updatePR(t *testing.T, client *APIClient, mux *http.ServeMux) {

	want := new(PullRequest)
	_ = testdata.ReadTestData(t, testdata.PullRequestsClosed, want)
	number1 := "12"
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s", owner, repo, number1)
	mockResponse(t, mux, urlStr, want)

	prState := "closed"
	got, ok, err := client.PullRequests.UpdatePullRequest(context.Background(), owner, repo, number1, &PullRequestRequest{
		State: &prState,
	})
	assert.Nil(t, err)
	assert.True(t, ok)
	d1, _ := json.Marshal(*want)
	d2, _ := json.Marshal(*got)
	assert.Equal(t, d1, d2)

	want1 := new(PullRequest)
	_ = testdata.ReadTestData(t, testdata.PullRequestsOpen, want1)
	number1 = "13"
	urlStr = fmt.Sprintf("/repos/%s/%s/pulls/%s", owner, repo, number1)
	mockResponse(t, mux, urlStr, want1)

	prState1 := "open"
	got1, ok, err := client.PullRequests.UpdatePullRequest(context.Background(), owner, repo, number1, &PullRequestRequest{
		State: &prState1,
	})
	assert.Nil(t, err)
	assert.True(t, ok)
	d3, _ := json.Marshal(*want1)
	d4, _ := json.Marshal(*got1)
	assert.Equal(t, d3, d4)
}

func listPRLinkingIssues(t *testing.T, client *APIClient, mux *http.ServeMux) {

	want := new([]*Issue)
	_ = testdata.ReadTestData(t, testdata.PullRequestsLinkingIssues, want)
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/issues", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.PullRequests.ListPullRequestLinkingIssues(ctx, owner, repo, number, page)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}
}

func getAllPRCommits(t *testing.T, client *APIClient, mux *http.ServeMux) {

	want := new([]*RepositoryCommit)
	_ = testdata.ReadTestData(t, testdata.PullRequestsCommits, want)
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/commits", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.PullRequests.ListPullRequestCommits(ctx, owner, repo, number)
	assert.Nil(t, err)
	assert.True(t, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}
}

func getPRChangeFiles(t *testing.T, client *APIClient, mux *http.ServeMux) {

	want := new([]*CommitFile)
	_ = testdata.ReadTestData(t, testdata.PullRequestsFiles, want)
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/files", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.PullRequests.GetPullRequestChangeFiles(ctx, owner, repo, number)
	assert.Nil(t, err)
	assert.True(t, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}
}

func getPROperationLogs(t *testing.T, client *APIClient, mux *http.ServeMux) {
	want := new([]*PullRequestOperationLog)
	_ = testdata.ReadTestData(t, testdata.PullRequestsOperateLog, want)
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/operate_logs", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.PullRequests.ListPullRequestOperationLogs(ctx, owner, repo, number, "asc", page)
	assert.Nil(t, err)
	assert.True(t, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}
}

func mergePR(t *testing.T, client *APIClient, mux *http.ServeMux) {

	want := new(PullRequestMergedResult)
	_ = testdata.ReadTestData(t, testdata.PullRequestsMerge, want)
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/merge", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.PullRequests.MergePullRequest(ctx, owner, repo, number, "merge")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, *want, *got)
}
