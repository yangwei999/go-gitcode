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
	"testing"
)

func TestIssueComments(t *testing.T) {

	client, mux, _ := mockServer(t)

	var want IssueComment
	_ = testdata.ReadTestData(t, testdata.IssuesComment, &want)

	// CreateIssueComment
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
