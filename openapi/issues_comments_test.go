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

func TestIssueComments(t *testing.T) {

	client, mux, _ := mockServer(t)

	createIssueComment(t, client, mux)
	listAnIssueAllComments(t, client, mux)
}

func createIssueComment(t *testing.T, client *APIClient, mux *http.ServeMux) {
	var want IssueComment
	_ = testdata.ReadTestData(t, testdata.IssuesComment, &want)

	urlStr := fmt.Sprintf("/repos/%s/%s/issues/%s/comments", owner, repo, number)
	mockResponse(t, mux, urlStr, want)
	comment := "123987u41"
	got, ok, err := client.Issues.CreateIssueComment(context.Background(), owner, repo, "1", &IssueComment{
		Body: &comment,
	})
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, want, *got)
}

func listAnIssueAllComments(t *testing.T, client *APIClient, mux *http.ServeMux) {
	want := new([]*IssueComment)
	_ = testdata.ReadTestData(t, testdata.IssuesListComments, want)
	issueNumber := "8"
	urlStr := fmt.Sprintf("/repos/%s/%s/issues/%s/comments", owner, repo, issueNumber)
	mockResponse(t, mux, urlStr, want)

	ctx := context.Background()
	got, ok, err := client.Issues.ListIssueComments(ctx, owner, repo, issueNumber, page, "", "")
	assert.Nil(t, err)
	assert.True(t, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}

}
