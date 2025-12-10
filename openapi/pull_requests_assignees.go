// Copyright (c) Huawei Technologies Co., Ltd. 2025. All rights reserved.
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

// Package openapi unified interface for GitCode OpenAPI
package openapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// UpdatePRAssignees 指派用户审查 Pull Request
//
// api Docs: https://docs.atomgit.com/docs/apis/post-api-v-5-repos-owner-repo-pulls-number-assignees
func (s *PullRequestsService) UpdatePRAssignees(ctx context.Context, owner, repo, number, assignees string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/assignees", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPost, urlStr, &PullRequestRequest{
		Assignees: &assignees,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// RemovePRAssignees 取消用户审查 Pull Request
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo-pulls-number-assignees
func (s *PullRequestsService) RemovePRAssignees(ctx context.Context, owner, repo, number, assignees string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/assignees", owner, repo, number)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, &url.Values{"assignees": []string{assignees}}, RequestHandler{t: Query})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// ResetPRAssigneesStatus 重置 Pull Request审查 的状态
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo-pulls-number-assignees
func (s *PullRequestsService) ResetPRAssigneesStatus(ctx context.Context, owner, repo, number string, all bool) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/assignees", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPatch, urlStr, &PullRequestAssigneesRequest{
		ResetAll: &all,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// UpdatePRReviewers 指派用户评审Pull Request
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-repos-owner-repo-pulls-number-approval-reviewers
func (s *PullRequestsService) UpdatePRReviewers(ctx context.Context, owner, repo, number, reviewers string, reviewerMode bool) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/reviewers", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPost, urlStr, &PullRequestRequest{
		Reviewers:     &reviewers,
		ReviewersMode: &reviewerMode,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// RemovePRReviewers 取消用户评审Pull Request
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo-pulls-number-approval-reviewers
func (s *PullRequestsService) RemovePRReviewers(ctx context.Context, owner, repo, number, reviewers string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/reviewers", owner, repo, number)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, &PullRequestRequest{
		Reviewers: &reviewers,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}
