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
	"net/http"
	"net/url"
)

// CreateIssueComment 创建Issue评论
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-repos-owner-repo-issues-number-comments
func (s *IssuesService) CreateIssueComment(ctx context.Context, owner, repo, number string, comment *IssueComment) (*IssueComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s/comments", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPost, urlStr, comment)
	if err != nil {
		return nil, false, err
	}

	addedComment := new(IssueComment)
	resp, err := s.api.Do(ctx, req, addedComment)
	return addedComment, successCreated(resp), err
}

// ListIssueComments 获取仓库某个Issue所有的评论
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-issues-number-comments
func (s *IssuesService) ListIssueComments(ctx context.Context, owner, repo, number, page, order, since string) ([]*IssueComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s/comments", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}, "order": []string{order}, "since": []string{since}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var comments []*IssueComment
	resp, err := s.api.Do(ctx, req, &comments)
	return comments, successGetData(resp), err
}
