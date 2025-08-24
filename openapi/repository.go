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
	"strconv"
)

// GetRepoContributors 获取仓库贡献者
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-contributors
func (s *RepositoryService) GetRepoContributors(ctx context.Context, owner, repo, category string) ([]*Contributor, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/contributors", owner, repo)
	var query url.Values
	if category != "" {
		query = url.Values{}
		query.Set("type", category)
	}
	req, err := newRequest(s.api, http.MethodGet, urlStr, &query, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var contributors []*Contributor
	resp, err := s.api.Do(ctx, req, &contributors)
	return contributors, successGetData(resp), err
}

// GetRepoContentByPath 获取仓库具体路径下的内容
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-contents-path
func (s *RepositoryService) GetRepoContentByPath(ctx context.Context, owner, repo, path, ref string) (*RepositoryContent, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/contents/%s?ref=%s", owner, repo, path, ref)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	content := new(RepositoryContent)
	resp, err := s.api.Do(ctx, req, content)
	return content, successGetData(resp), err
}

// CreateOrgRepo 创建组织仓库
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-orgs-org-repos
func (s *RepositoryService) CreateOrgRepo(ctx context.Context, owner string, repoContent *RepositoryRequest) (*Repository, bool, error) {
	urlStr := fmt.Sprintf("orgs/%s/repos", owner)
	req, err := newRequest(s.api, http.MethodPost, urlStr, repoContent)
	if err != nil {
		return nil, false, err
	}

	repo := new(Repository)
	resp, err := s.api.Do(ctx, req, repo)
	return repo, successCreated(resp), err
}

// ListOrgRepo 获取组织项目列表
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-orgs-org-repos
func (s *RepositoryService) ListOrgRepo(ctx context.Context, owner, repoType, page string) ([]*Repository, bool, error) {
	urlStr := fmt.Sprintf("orgs/%s/repos", owner)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}, "type": []string{repoType}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var repos []*Repository
	resp, err := s.api.Do(ctx, req, &repos)
	return repos, successGetData(resp), err
}

// DeleteOrgRepo 删除一个仓库
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo
func (s *RepositoryService) DeleteOrgRepo(ctx context.Context, owner, repo string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s", owner, repo)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// GetRepoPermissionMode 获取项目的权限模式
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-transition
func (s *RepositoryService) GetRepoPermissionMode(ctx context.Context, owner, repo string) (*RepositoryPermissionMode, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/transition", owner, repo)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	mode := new(RepositoryPermissionMode)
	resp, err := s.api.Do(ctx, req, mode)
	return mode, successGetData(resp), err
}

// UpdateRepoPermissionMode 更新仓库的权限模式
//
// api Docs: https://docs.gitcode.com/docs/apis/put-api-v-5-repos-owner-repo-transition
func (s *RepositoryService) UpdateRepoPermissionMode(ctx context.Context, owner, repo, mode string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/transition", owner, repo)
	m, err := strconv.Atoi(mode)
	if err != nil {
		return false, err
	}
	req, err := newRequest(s.api, http.MethodPut, urlStr, RepositoryPermissionModeRequest{
		Model: m,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// GetRepoCustomRoles 获取项目自定义角色
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-customized-roles
func (s *RepositoryService) GetRepoCustomRoles(ctx context.Context, owner, repo string) ([]*CustomRepoRoles, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/customized_roles", owner, repo)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var roles []*CustomRepoRoles
	resp, err := s.api.Do(ctx, req, &roles)
	return roles, successGetData(resp), err
}

// GetRepoTrees 获取仓库目录Tree
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-git-trees-sha
func (s *RepositoryService) GetRepoTrees(ctx context.Context, owner, repo, sha, page, recursive string) (*RepositoryTree, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/git/trees/%s", owner, repo, sha)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}, "recursive": []string{recursive}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	trees := new(RepositoryTree)
	resp, err := s.api.Do(ctx, req, &trees)
	return trees, successGetData(resp), err
}

// GetRepoFileList 获取文件列表
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-file-list
func (s *RepositoryService) GetRepoFileList(ctx context.Context, owner, repo, refName, fileName string) ([]string, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/file_list", owner, repo)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"ref_name": []string{refName}, "file_name": []string{url.QueryEscape(fileName)}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var fileList []string
	resp, err := s.api.Do(ctx, req, &fileList)
	return fileList, successGetData(resp), err
}

// UpdateRepoPullRequestSetting 更新 Pull Request设置
//
// api Docs: https://docs.gitcode.com/docs/apis/put-api-v-5-repos-owner-repo-pull-request-settings
func (s *RepositoryService) UpdateRepoPullRequestSetting(ctx context.Context, owner, repo string, prSetting *RepositoryRepoPullRequestSettingRequest) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pull_request_settings", owner, repo)
	req, err := newRequest(s.api, http.MethodPut, urlStr, prSetting)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// UpdateRepoSetting 更新仓库功能设置
//
// api Docs: https://docs.gitcode.com/docs/apis/put-api-v-5-repos-owner-repo-repo-settings
func (s *RepositoryService) UpdateRepoSetting(ctx context.Context, owner, repo string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/repo_settings", owner, repo)
	req, err := newRequest(s.api, http.MethodPut, urlStr, RepositoryRepoSettingRequest{
		GeneratePreMergeRef:            true,
		ForbiddenCommitterCreateBranch: true,
		ForbiddenDeveloperCreateBranch: true,
		ForbiddenDeveloperCreateTag:    true,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// UpdateRepoBaseSetting 更新仓库设置
//
// api Docs: https://docs.gitcode.com/docs/apis/patch-api-v-5-repos-owner-repo
func (s *RepositoryService) UpdateRepoBaseSetting(ctx context.Context, owner, repo string, repoContent *RepositoryRequest) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s", owner, repo)
	req, err := newRequest(s.api, http.MethodPatch, urlStr, repoContent)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}
