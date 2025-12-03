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
	"time"
)

func TestPullRequestsComments(t *testing.T) {
	client, mux, _ := mockServer(t)
	createPRComments(t, client, mux)
	getPRComments(t, client, mux)
	updatePRComment(t, client, mux)
	deletePRComment(t, client, mux)
}

func createPRComments(t *testing.T, client *APIClient, mux *http.ServeMux) {

	want := new(SimpleComment)
	_ = testdata.ReadTestData(t, testdata.PullRequestsCreateComment, want)
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/comments", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.PullRequests.CreatePullRequestComment(ctx, owner, repo, number, &PullRequestCommentRequest{
		Body: "fgujhgasd",
	})
	assert.Nil(t, err)
	assert.True(t, ok)

	assert.Equal(t, *want, *got)
}

func getPRComments(t *testing.T, client *APIClient, mux *http.ServeMux) {

	want := new([]*PullRequestComment)
	_ = testdata.ReadTestData(t, testdata.PullRequestsComments, want)
	number := "5"
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/%s/comments", owner, repo, number)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.PullRequests.ListPullRequestComments(ctx, owner, repo, number, page, "pr_comment", "")
	assert.Nil(t, err)
	assert.True(t, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}

	assert.Equal(t, "2024-12-11T14:38:33+08:00", *got[0].UpdatedAt.ToString())
	assert.Equal(t, time.Date(2024, time.December, 11, 14, 38, 33, 0, time.Local), got[0].UpdatedAt.ToTime())
}

func updatePRComment(t *testing.T, client *APIClient, mux *http.ServeMux) {
	commentID := "1274612"
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/comments/%s", owner, repo, commentID)
	mockResponse(t, mux, urlStr, nil)

	ctx := context.Background()
	ok, err := client.PullRequests.UpdatePullRequestComment(ctx, owner, repo, commentID, "gfgfds")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func deletePRComment(t *testing.T, client *APIClient, mux *http.ServeMux) {
	commentID := "64345"
	urlStr := fmt.Sprintf("/repos/%s/%s/pulls/comments/%s", owner, repo, commentID)
	mockResponse(t, mux, urlStr, nil)

	ctx := context.Background()
	ok, err := client.PullRequests.DeletePullRequestComment(ctx, owner, repo, commentID)
	assert.Nil(t, err)
	assert.True(t, ok)
}
