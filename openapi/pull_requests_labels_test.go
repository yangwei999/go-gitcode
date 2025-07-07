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
	"strings"
	"testing"
)

func TestPullRequestsLabels(t *testing.T) {
	client, mux, _ := mockServer(t)
	addOrGetPRLabels(t, client, mux)
	removePRLabel(t, client, mux)
}

func addOrGetPRLabels(t *testing.T, client *APIClient, mux *http.ServeMux) {

	var want []*Label
	_ = testdata.ReadTestData(t, testdata.PullRequestsAddLabels, &want)
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/labels", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.PullRequests.AddLabelsToPullRequest(ctx, owner, repo, number, []string{"fa", "fsw"})
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range want {
		assert.Equal(t, *want[i], *got[i])
	}

	got, ok, err = client.PullRequests.GetLabelsOfPullRequest(context.Background(), owner, repo, number)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range want {
		assert.Equal(t, *want[i], *got[i])
	}
}

func removePRLabel(t *testing.T, client *APIClient, mux *http.ServeMux) {
	delLabels := []string{"fa", "fsw"}
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/labels/%s", owner, repo, number, strings.Join(delLabels, ","))
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	ok, err := client.PullRequests.RemoveLabelsFromPullRequest(ctx, owner, repo, number, delLabels)
	assert.Nil(t, err)
	assert.True(t, ok)
}
