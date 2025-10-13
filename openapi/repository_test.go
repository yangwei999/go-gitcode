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

func TestRepository(t *testing.T) {
	client, mux, _ := mockServer(t)

	getRepoContributors(t, client, mux)
	getRepoContentByPath(t, client, mux)
	createOrGetRepo(t, client, mux)
	deleteRepo(t, client, mux)
	getOrUpdateRepoMode(t, client, mux)
	getRepoCustomRoles(t, client, mux)
	getRepoTree(t, client, mux)
	getRepoFiles(t, client, mux)
	updatePRSetting(t, client, mux)
	updateRepoSetting(t, client, mux)
}

func getRepoContributors(t *testing.T, client *APIClient, mux *http.ServeMux) {

	var want []*Contributor
	_ = testdata.ReadTestData(t, testdata.RepositoryContributors, &want)
	urlStr := fmt.Sprintf("/repos/%s/%s/contributors", owner, repo)
	category := "authors"
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		assert.Equal(t, r.URL.Query().Get("type"), category)
		_ = json.NewEncoder(w).Encode(want)
	})

	got, ok, err := client.Repository.GetRepoContributors(context.Background(), owner, repo, "authors")
	assert.Nil(t, err)
	assert.True(t, ok)

	for i := range want {
		assert.Equal(t, *want[i], *got[i])
	}
}

func getRepoContentByPath(t *testing.T, client *APIClient, mux *http.ServeMux) {
	want := new(RepositoryContent)
	_ = testdata.ReadTestData(t, testdata.RepositoryFileContent, want)
	filePath := "2.txt"
	branch := "master"
	urlStr := fmt.Sprintf("/repos/%s/%s/contents/%s", owner, repo, filePath)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		assert.Equal(t, r.URL.Query().Get("ref"), branch)
		_ = json.NewEncoder(w).Encode(want)
	})

	ctx := context.Background()
	got, ok, err := client.Repository.GetRepoContentByPath(ctx, owner, repo, filePath, branch)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, *want, *got)
}

func createOrGetRepo(t *testing.T, client *APIClient, mux *http.ServeMux) {
	want := new(Repository)
	_ = testdata.ReadTestData(t, testdata.RepositoryCreate, want)
	urlStr := fmt.Sprintf("/orgs/%s/repos", owner)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		if r.Method == http.MethodPost {
			_ = json.NewEncoder(w).Encode(want)
		} else {
			_ = json.NewEncoder(w).Encode([]Repository{*want})
		}

	})

	public := 1
	autoInit := true
	ctx := context.Background()
	got, ok, err := client.Repository.CreateOrgRepo(ctx, owner, &RepositoryRequest{
		Name:          want.Name,
		Path:          want.Path,
		Description:   want.Description,
		Private:       want.Private,
		Public:        &public,
		AutoInit:      &autoInit,
		DefaultBranch: want.DefaultBranch,
	})
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, *want, *got)

	got1, ok, err := client.Repository.ListOrgRepo(ctx, owner, "all", page)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, 1, len(got1))
	assert.Equal(t, *want, *got1[0])
}

func deleteRepo(t *testing.T, client *APIClient, mux *http.ServeMux) {
	urlStr := fmt.Sprintf("/repos/%s/%s", owner, repo)
	mockResponse(t, mux, urlStr, nil)

	ctx := context.Background()
	ok, err := client.Repository.DeleteOrgRepo(ctx, owner, repo)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func getOrUpdateRepoMode(t *testing.T, client *APIClient, mux *http.ServeMux) {
	want := RepositoryPermissionMode{
		Model: 1,
	}
	urlStr := fmt.Sprintf("/repos/%s/%s/transition", owner, repo)
	mockResponse(t, mux, urlStr, want)

	got, ok, err := client.Repository.GetRepoPermissionMode(context.Background(), owner, repo)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, want, *got)

	ok, err = client.Repository.UpdateRepoPermissionMode(context.Background(), owner, repo, "2")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func getRepoCustomRoles(t *testing.T, client *APIClient, mux *http.ServeMux) {
	var want []*CustomRepoRoles
	_ = testdata.ReadTestData(t, testdata.RepositoryCustomRoles, &want)
	urlStr := fmt.Sprintf("/repos/%s/%s/customized_roles", owner, repo)
	mockResponse(t, mux, urlStr, want)

	got, ok, err := client.Repository.GetRepoCustomRoles(context.Background(), owner, repo)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range want {
		assert.Equal(t, *want[i], *got[i])
	}

}

func getRepoTree(t *testing.T, client *APIClient, mux *http.ServeMux) {
	var want RepositoryTree
	_ = testdata.ReadTestData(t, testdata.RepositoryTree, &want)
	branch := "dev"
	urlStr := fmt.Sprintf("/repos/%s/%s/git/trees/%s", owner, repo, branch)
	mockResponse(t, mux, urlStr, want)

	got, ok, err := client.Repository.GetRepoTrees(context.Background(), owner, repo, branch, page, "1")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, want, *got)
}

func getRepoFiles(t *testing.T, client *APIClient, mux *http.ServeMux) {
	fileName := "1.txt"
	want := []string{"dir1/1.txt", "dir2/subDir1/1.txt", "1.txt"}
	urlStr := fmt.Sprintf("/repos/%s/%s/file_list", owner, repo)
	mockResponse(t, mux, urlStr, want)

	got, ok, err := client.Repository.GetRepoFileList(context.Background(), owner, repo, "v1", fileName)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range want {
		assert.Equal(t, want[i], got[i])
	}

}

func updatePRSetting(t *testing.T, client *APIClient, mux *http.ServeMux) {
	urlStr := fmt.Sprintf("/repos/%s/%s/pull_request_settings", owner, repo)
	mockResponse(t, mux, urlStr, nil)

	ok, err := client.Repository.UpdateRepoPullRequestSetting(context.Background(), owner, repo, nil)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func updateRepoSetting(t *testing.T, client *APIClient, mux *http.ServeMux) {
	urlStr := fmt.Sprintf("/repos/%s/%s/repo_settings", owner, repo)
	mockResponse(t, mux, urlStr, nil)

	ok, err := client.Repository.UpdateRepoSetting(context.Background(), owner, repo)
	assert.Nil(t, err)
	assert.True(t, ok)
}
