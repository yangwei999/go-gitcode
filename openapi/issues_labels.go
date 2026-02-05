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

// ListRepoIssueLabels 获取仓库所有任务标签
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-labels
func (s *IssuesService) ListRepoIssueLabels(ctx context.Context, owner, repo, page string) ([]*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/labels", owner, repo)
	query := &url.Values{"page": []string{page}, "per_page": []string{"100"}}
	req, err := newRequest(s.api, http.MethodGet, urlStr, query, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var labels []*Label
	resp, err := s.api.Do(ctx, req, &labels)
	return labels, successCreated(resp), err
}

// CreateRepoIssueLabel 创建仓库任务标签
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-repos-owner-repo-labels
func (s *IssuesService) CreateRepoIssueLabel(ctx context.Context, owner, repo string, newLabel *Label) (*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/labels", owner, repo)
	req, err := newRequest(s.api, http.MethodPost, urlStr, newLabel, RequestHandler{t: Form})
	if err != nil {
		return nil, false, err
	}

	label := new(Label)
	resp, err := s.api.Do(ctx, req, label)
	return label, successCreated(resp), err
}

// UpdateRepoIssueLabel 更新一个仓库的任务标签
//
// api Docs: https://docs.gitcode.com/docs/apis/patch-api-v-5-repos-owner-repo-labels-original-name
func (s *IssuesService) UpdateRepoIssueLabel(ctx context.Context, owner, repo, originalName, newName, color string) (*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/labels/%s", owner, repo, originalName)
	req, err := newRequest(s.api, http.MethodPatch, urlStr, &url.Values{"name": []string{newName}, "color": []string{color}}, RequestHandler{t: Form})
	if err != nil {
		return nil, false, err
	}

	editedLabel := new(Label)
	resp, err := s.api.Do(ctx, req, editedLabel)
	return editedLabel, successModified(resp), err
}

// DeleteRepoIssueLabel 删除一个仓库任务标签
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo-labels-name
func (s *IssuesService) DeleteRepoIssueLabel(ctx context.Context, owner, repo, name string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/labels/%s", owner, repo, name)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}
	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// AddLabelsToIssue 创建Issue标签
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-repos-owner-repo-issues-number-labels
func (s *IssuesService) AddLabelsToIssue(ctx context.Context, owner, repo, number string, labelNameList []string) ([]*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s/labels", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPost, urlStr, labelNameList)
	if err != nil {
		return nil, false, err
	}

	var havingLabels []*Label
	resp, err := s.api.Do(ctx, req, &havingLabels)
	return havingLabels, successCreated(resp), err
}

// RemoveLabelsFromIssue 删除Issue标签
//
// api Docs: https://docs.gitcode.com/docs/apis/delete-api-v-5-repos-owner-repo-issues-number-labels-name
func (s *IssuesService) RemoveLabelsFromIssue(ctx context.Context, owner, repo, number, label string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s/labels/%s", owner, repo, number, label)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// GetIssueLabels 获取企业某个Issue所有标签
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-enterprises-enterprise-issues-issue-id-labels
func (s *IssuesService) GetIssueLabels(ctx context.Context, owner, issueID, page string) ([]*Label, bool, error) {
	urlStr := fmt.Sprintf("enterprises/%s/issues/%s/labels", owner, issueID)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}})
	if err != nil {
		return nil, false, err
	}

	var havingLabels []*Label
	resp, err := s.api.Do(ctx, req, &havingLabels)
	return havingLabels, successGetData(resp), err
}
