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

// GetRepoAllBranch 获取项目所有分支
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-branches
func (s *RepositoryService) GetRepoAllBranch(ctx context.Context, owner, repo, page string) ([]*Branch, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/branches", owner, repo)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var branches []*Branch
	resp, err := s.api.Do(ctx, req, &branches)
	return branches, successGetData(resp), err
}

// GetRepoBranch 获取单个分支
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-branches-branch
func (s *RepositoryService) GetRepoBranch(ctx context.Context, owner, repo, branch string) (*Branch, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/branches/%s", owner, repo, branch)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	content := new(Branch)
	resp, err := s.api.Do(ctx, req, content)
	return content, successGetData(resp), err
}

// CreateRepoBranch 创建分支
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-repos-owner-repo-branches
func (s *RepositoryService) CreateRepoBranch(ctx context.Context, owner, repo, branch, createFrom string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/branches", owner, repo)
	req, err := newRequest(s.api, http.MethodPost, urlStr, &BranchRequest{
		CreateFrom: createFrom,
		BranchName: branch,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// CreateRepoBranchProtectedRule 新建保护分支规则
//
// api Docs: https://docs.gitcode.com/docs/apis/put-api-v-5-repos-owner-repo-branches-setting-new
func (s *RepositoryService) CreateRepoBranchProtectedRule(ctx context.Context, owner, repo, branch, permission string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/branches/setting/new", owner, repo)
	req, err := newRequest(s.api, http.MethodPut, urlStr, &BranchProtectedRuleRequest{
		PushPermission:  permission,
		MergePermission: permission,
		BranchName:      branch,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// UpdateRepoBranchProtectedRule 更新保护分支规则
//
// api Docs: https://docs.gitcode.com/docs/apis/put-api-v-5-repos-owner-repo-branches-wildcard-setting
func (s *RepositoryService) UpdateRepoBranchProtectedRule(ctx context.Context, owner, repo, branch, permission string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/branches/%s/setting", owner, repo, branch)
	req, err := newRequest(s.api, http.MethodPut, urlStr, &BranchProtectedRuleRequest{
		PushPermission:  permission,
		MergePermission: permission,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// RemoveRepoBranchProtectedRule 删除保护分支规则
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo-branches-wildcard-setting
func (s *RepositoryService) RemoveRepoBranchProtectedRule(ctx context.Context, owner, repo, branch string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/branches/%s/setting", owner, repo, branch)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// ListRepoAllTag 列出项目所有的tags
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-tags
func (s *RepositoryService) ListRepoAllTag(ctx context.Context, owner, repo, page string) ([]*RepositoryTag, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/tags", owner, repo)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var tags []*RepositoryTag
	resp, err := s.api.Do(ctx, req, &tags)
	return tags, successGetData(resp), err
}
