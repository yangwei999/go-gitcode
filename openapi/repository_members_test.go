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
	"github.com/opensourceways/go-gitcode/testdata"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRepositoryMembers(t *testing.T) {
	client, mux, _ := mockServer(t)
	getRepoAllMember(t, client, mux)
	getRepoMemberPermission(t, client, mux)
	checkUserIsRepoMember(t, client, mux)
	addOrRemoveRepoMember(t, client, mux)
}

func getRepoAllMember(t *testing.T, client *APIClient, mux *http.ServeMux) {
	var want []*User
	_ = testdata.ReadTestData(t, testdata.RepositoryMembers, &want)
	urlStr := fmt.Sprintf("/repos/%s/%s/collaborators", owner, repo)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.Repository.GetRepoAllMember(ctx, owner, repo, page)
	assert.Nil(t, err)
	assert.True(t, ok)

	for i := range want {
		assert.Equal(t, *want[i], *got[i])
	}

}

func getRepoMemberPermission(t *testing.T, client *APIClient, mux *http.ServeMux) {

	user1 := "user1"
	user := "{\n    \"id\": 412,\n    \"login\": \"user1\",\n    \"permission\": \"admin\"\n}"
	urlStr := fmt.Sprintf("/repos/%s/%s/collaborators/%s/permission", owner, repo, user1)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(user))
	})

	ctx := context.Background()
	got, ok, err := client.Repository.GetRepoMemberPermission(ctx, owner, repo, user1)
	assert.Equal(t, nil, err)
	assert.Equal(t, [2]bool{true, false}, ok)
	assert.Equal(t, user1, *got.Login)
	assert.Equal(t, "admin", *got.Permission)

	user2 := "145123"
	msg := "{\"message\":\"404 Not Found\"}"
	urlStr = fmt.Sprintf("/repos/%s/%s/collaborators/%s/permission", owner, repo, user2)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(msg))
	})

	ctx1 := context.Background()
	got, ok, err = client.Repository.GetRepoMemberPermission(ctx1, owner, repo, user2)
	assert.Equal(t, msg, err.Error())
	assert.Equal(t, [2]bool{false, true}, ok)
	assert.Equal(t, User{}, *got)
}

func checkUserIsRepoMember(t *testing.T, client *APIClient, mux *http.ServeMux) {

	user1 := "user4"
	urlStr := fmt.Sprintf("/repos/%s/%s/collaborators/%s", owner, repo, user1)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	got, ok, err := client.Repository.CheckUserIsRepoMember(ctx, owner, repo, user1)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.True(t, got)

	msg := "{\"message\":\"404 Not Found\"}"
	user2 := "user5"
	urlStr = fmt.Sprintf("/repos/%s/%s/collaborators/%s", owner, repo, user2)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(msg))
	})

	ctx1 := context.Background()
	got, ok, err = client.Repository.CheckUserIsRepoMember(ctx1, owner, repo, user2)
	assert.Equal(t, msg, err.Error())
	assert.Equal(t, true, ok)
	assert.Equal(t, false, got)
}

func addOrRemoveRepoMember(t *testing.T, client *APIClient, mux *http.ServeMux) {
	addUser := "user9"
	repo1 := "repo11"
	urlStr := fmt.Sprintf("/repos/%s/%s/collaborators/%s", owner, repo1, addUser)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	ok, err := client.Repository.AddRepoMember(ctx, owner, repo1, addUser, "read")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = client.Repository.RemoveRepoMember(context.Background(), owner, repo1, addUser)
	assert.Nil(t, err)
	assert.True(t, ok)
}
