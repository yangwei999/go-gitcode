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

// CreatePullRequest 创建Pull Request
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-repos-owner-repo-pulls
func (s *PullRequestsService) CreatePullRequest(ctx context.Context, owner, repo string, prContent *PullRequestCreateRequest) (*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls", owner, repo)
	req, err := newRequest(s.api, http.MethodPost, urlStr, prContent)
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequest)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successCreated(resp), err
}

// GetPullRequest 获取单个Pull Requests
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-pulls-number
func (s *PullRequestsService) GetPullRequest(ctx context.Context, owner, repo, number string) (*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequest)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successGetData(resp), err
}

// ListPullRequests 获取仓库的Pull Request列表
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-pulls
func (s *PullRequestsService) ListPullRequests(ctx context.Context, owner, repo string, query *url.Values) ([]*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls", owner, repo)

	req, err := newRequest(s.api, http.MethodGet, urlStr, query, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var prs []*PullRequest
	resp, err := s.api.Do(ctx, req, &prs)
	return prs, successGetData(resp), err
}

// UpdatePullRequest 更新Pull Request信息
//
// api Docs: https://docs.gitcode.com/docs/apis/patch-api-v-5-repos-owner-repo-pulls-number
func (s *PullRequestsService) UpdatePullRequest(ctx context.Context, owner, repo, number string, prContent *PullRequestRequest) (*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPatch, urlStr, prContent)
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequest)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successCreated(resp), err
}

// ListPullRequestLinkingIssues 获取pr关联的issue
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-pulls-number-issues
func (s *PullRequestsService) ListPullRequestLinkingIssues(ctx context.Context, owner, repo, number, page string) ([]*Issue, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/issues", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var linkingPRList []*Issue
	resp, err := s.api.Do(ctx, req, &linkingPRList)
	return linkingPRList, successGetData(resp), err
}

// ListPullRequestCommits 获取某Pull Request的所有Commit信息
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-pulls-number-commits
func (s *PullRequestsService) ListPullRequestCommits(ctx context.Context, owner, repo, number string) ([]*RepositoryCommit, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/commits", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var commitList []*RepositoryCommit
	resp, err := s.api.Do(ctx, req, &commitList)
	return commitList, successGetData(resp), err
}

// GetPullRequestChangeFiles Pull Request Commit文件列表
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-pulls-number-files
func (s *PullRequestsService) GetPullRequestChangeFiles(ctx context.Context, owner, repo, number string) ([]*CommitFile, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/files", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var changes []*CommitFile
	resp, err := s.api.Do(ctx, req, &changes)
	return changes, successGetData(resp), err
}

// MergePullRequest 合并Pull Request
//
// api Docs: https://docs.gitcode.com/docs/apis/put-api-v-5-repos-owner-repo-pulls-number-merge
func (s *PullRequestsService) MergePullRequest(ctx context.Context, owner, repo, number, mergeMethod string) (*PullRequestMergedResult, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/merge", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPut, urlStr, &PullRequestRequestMerge{
		Method: mergeMethod,
	})
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequestMergedResult)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successCreated(resp), err
}

// ListPullRequestOperationLogs 获取某个Pull Request的操作日志
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-pulls-number-operate-logs
func (s *PullRequestsService) ListPullRequestOperationLogs(ctx context.Context, owner, repo, number, sort, page string) ([]*PullRequestOperationLog, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/operate_logs", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}, "sort": []string{sort}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var logs []*PullRequestOperationLog
	resp, err := s.api.Do(ctx, req, &logs)
	return logs, successGetData(resp), err
}
