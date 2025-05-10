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

func TestOrganization(t *testing.T) {
	client, mux, _ := mockServer(t)

	// ListOrgMember
	listOrgMember(t, client, mux)
	// AddOrgMember RemoveOrgMember
	addOrRemoveOrgMember(t, client, mux)
}

func listOrgMember(t *testing.T, client *APIClient, mux *http.ServeMux) {
	members := new([]*User)
	_ = readTestdata(t, userTestDataDir+"org_member.json", members)
	urlStr := fmt.Sprintf("/orgs/%s/members", owner)
	mockResponse(t, mux, urlStr, members)

	result, ok, err := client.Org.ListOrgMember(context.Background(), owner, "all", page)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range result {
		d1, _ := json.Marshal(*(*members)[i])
		d2, _ := json.Marshal(*result[i])
		assert.Equal(t, d1, d2)
	}
}

func addOrRemoveOrgMember(t *testing.T, client *APIClient, mux *http.ServeMux) {
	username := "user12"
	urlStr := fmt.Sprintf("/orgs/%s/memberships/%s", owner, username)
	mockResponse(t, mux, urlStr, nil)

	ok, err := client.Org.AddOrgMember(context.Background(), owner, username, "push", "")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = client.Org.RemoveOrgMember(context.Background(), owner, username)
	assert.Nil(t, err)
	assert.True(t, ok)
}
