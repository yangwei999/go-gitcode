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
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRepositoryBranches(t *testing.T) {
	client, mux, _ := mockServer(t)
	getRepoAllBranches(t, client, mux)
	getRepoOneBranch(t, client, mux)
	createRepoBranch(t, client, mux)
	createProtectedBranch(t, client, mux)
	addOrRemoveProtectedBranch(t, client, mux)
}

func getRepoAllBranches(t *testing.T, client *APIClient, mux *http.ServeMux) {
	var want []*Branch
	_ = readTestdata(t, reposTestDataDir+"repository_branches.json", &want)
	urlStr := fmt.Sprintf("/repos/%s/%s/branches", owner, repo)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.Repository.GetRepoAllBranch(ctx, owner, repo, page)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range want {
		assert.Equal(t, *want[i], *got[i])
	}
}

func getRepoOneBranch(t *testing.T, client *APIClient, mux *http.ServeMux) {
	var want Branch
	_ = readTestdata(t, reposTestDataDir+"repository_branch_detail.json", &want)
	urlStr := fmt.Sprintf("/repos/%s/%s/branches/%s", owner, repo, branch)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.Repository.GetRepoBranch(ctx, owner, repo, branch)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, want, *got)
}

func createRepoBranch(t *testing.T, client *APIClient, mux *http.ServeMux) {

	repo3 := "repo-branch-3"
	urlStr := fmt.Sprintf("/repos/%s/%s/branches", owner, repo3)
	mockResponse(t, mux, urlStr, nil)

	ctx := context.Background()
	ok, err := client.Repository.CreateRepoBranch(ctx, owner, repo3, branch, branch)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func createProtectedBranch(t *testing.T, client *APIClient, mux *http.ServeMux) {

	urlStr := fmt.Sprintf("/repos/%s/%s/branches/setting/new", owner, repo)
	mockResponse(t, mux, urlStr, nil)

	ctx := context.Background()
	ok, err := client.Repository.CreateRepoBranchProtectedRule(ctx, owner, repo, branch, permission)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func addOrRemoveProtectedBranch(t *testing.T, client *APIClient, mux *http.ServeMux) {

	urlStr := fmt.Sprintf("/repos/%s/%s/branches/%s/setting", owner, repo, branch)
	mockResponse(t, mux, urlStr, nil)

	ctx := context.Background()
	ok, err := client.Repository.UpdateRepoBranchProtectedRule(ctx, owner, repo, branch, "none")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = client.Repository.RemoveRepoBranchProtectedRule(context.Background(), owner, repo, branch)
	assert.Nil(t, err)
	assert.True(t, ok)
}
