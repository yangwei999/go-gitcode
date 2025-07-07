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

// GetRepoAllMember 获取仓库的所有成员
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-collaborators
func (s *RepositoryService) GetRepoAllMember(ctx context.Context, owner, repo, page string) ([]*User, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators", owner, repo)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var members []*User
	resp, err := s.api.Do(ctx, req, &members)
	return members, successGetData(resp), err
}

// GetRepoMemberPermission 查看仓库成员的权限
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-collaborators-username-permission
func (s *RepositoryService) GetRepoMemberPermission(ctx context.Context, owner, repo, login string) (*User, [2]bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators/%s/permission", owner, repo, login)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, [2]bool{false, false}, err
	}

	user := new(User)
	resp, err := s.api.Do(ctx, req, user)
	respNormal := resp != nil && (resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusNoContent)
	return user, [2]bool{successModified(resp), respNormal}, err
}

// CheckUserIsRepoMember 判断用户是否为仓库成员
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-collaborators-username
func (s *RepositoryService) CheckUserIsRepoMember(ctx context.Context, owner, repo, username string) (bool, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators/%s", owner, repo, username)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return false, false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	respNormal := resp != nil && (resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusNoContent)
	return successModified(resp), respNormal, err
}

// AddRepoMember 添加项目成员或更新项目成员权限
//
// api Docs: https://docs.gitcode.com/docs/apis/put-api-v-5-repos-owner-repo-collaborators-username
func (s *RepositoryService) AddRepoMember(ctx context.Context, owner, repo, username, permission string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators/%s", owner, repo, username)
	req, err := newRequest(s.api, http.MethodPut, urlStr, &User{Permission: &permission})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// RemoveRepoMember 移除项目成员
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo-collaborators-username
func (s *RepositoryService) RemoveRepoMember(ctx context.Context, owner, repo, username string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators/%s", owner, repo, username)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}
