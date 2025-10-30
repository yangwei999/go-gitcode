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
)

// CreateIssue 创建Issue
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-repos-owner-issues
func (s *IssuesService) CreateIssue(ctx context.Context, owner string, issueContent *IssueRequest) (*Issue, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/issues", owner)
	req, err := newRequest(s.api, http.MethodPost, urlStr, issueContent)
	if err != nil {
		return nil, false, err
	}

	createdIssue := new(Issue)
	resp, err := s.api.Do(ctx, req, createdIssue)
	return createdIssue, successCreated(resp), err
}

// UpdateIssue 更新Issue
//
// api Docs: https://docs.gitcode.com/docs/apis/patch-api-v-5-repos-owner-issues-number
func (s *IssuesService) UpdateIssue(ctx context.Context, owner, number string, issueContent *IssueRequest) (*Issue, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/issues/%s", owner, number)
	req, err := newRequest(s.api, http.MethodPatch, urlStr, issueContent)
	if err != nil {
		return nil, false, err
	}

	updatedIssue := new(Issue)
	resp, err := s.api.Do(ctx, req, updatedIssue)
	return updatedIssue, successCreated(resp), err
}

// ListIssueLinkingPullRequests 获取 issue 关联的 pull requests
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-issues-number-pull-requests
func (s *IssuesService) ListIssueLinkingPullRequests(ctx context.Context, owner, repo, number string) ([]*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s/pull_requests", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var linkingPRList []*PullRequest
	resp, err := s.api.Do(ctx, req, &linkingPRList)
	return linkingPRList, successGetData(resp), err
}

// GetIssue 获取仓库的某个Issue
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-issues-number
func (s *IssuesService) GetIssue(ctx context.Context, owner, repo, number string) (*Issue, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	issue := new(Issue)
	resp, err := s.api.Do(ctx, req, issue)
	return issue, successGetData(resp), err
}

// ListRepoIssues 获取仓库的所有issue
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-issues
func (s *IssuesService) ListRepoIssues(ctx context.Context, owner, repo string, search *IssueQueryRequest) ([]*Issue, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	req, err := newRequest(s.api, http.MethodGet, urlStr, search, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var issues []*Issue
	resp, err := s.api.Do(ctx, req, &issues)
	return issues, successGetData(resp), err
}
