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

// CreatePullRequestComment 提交 pull request 评论
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-repos-owner-repo-pulls-number-comments
func (s *PullRequestsService) CreatePullRequestComment(ctx context.Context, owner, repo, number string, comment *PullRequestCommentRequest) (*SimpleComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/comments", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPost, urlStr, comment)
	if err != nil {
		return nil, false, err
	}

	addedComment := new(SimpleComment)
	resp, err := s.api.Do(ctx, req, addedComment)
	return addedComment, successCreated(resp), err
}

// ListPullRequestComments 获取某个Pull Request的所有评论
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-pulls-number-comments
func (s *PullRequestsService) ListPullRequestComments(ctx context.Context, owner, repo, number, page, direction, commentType string) ([]*PullRequestComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/comments", owner, repo, number)
	query := &url.Values{"page": []string{page}, "per_page": []string{"100"}}
	if commentType != "" {
		query.Set("type", commentType)
	}
	if direction != "" {
		query.Set("direction", direction)
	}
	req, err := newRequest(s.api, http.MethodGet, urlStr, query, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var comments []*PullRequestComment
	resp, err := s.api.Do(ctx, req, &comments)
	return comments, successGetData(resp), err
}

// UpdatePullRequestComment 编辑评论
//
// api Docs: https://docs.gitcode.com/docs/apis/patch-api-v-5-repos-owner-repo-pulls-comments-id
func (s *PullRequestsService) UpdatePullRequestComment(ctx context.Context, owner, repo, commentID, body string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/comments/%s", owner, repo, commentID)
	req, err := newRequest(s.api, http.MethodPatch, urlStr, &PullRequestComment{
		Body: &body,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// DeletePullRequestComment 删除评论
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo-pulls-comments-id
func (s *PullRequestsService) DeletePullRequestComment(ctx context.Context, owner, repo, commentID string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/comments/%s", owner, repo, commentID)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}
