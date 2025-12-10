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
package openapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListOrgDashboard 获取组织看板列表
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-org-owner-kanban-list
func (s *DashboardService) ListOrgDashboard(ctx context.Context, owner, page, status string) (*DashboardView, bool, error) {
	urlStr := fmt.Sprintf("org/%s/kanban/list", owner)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}, "status": []string{status}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	view := new(DashboardView)
	resp, err := s.api.Do(ctx, req, view)
	return view, successGetData(resp), err
}

// AddIssueToDashboard 添加Issue或者Pull Request到看板
//
// api Docs: https://docs.gitcode.com/docs/apis/post-api-v-5-org-owner-kanban-id-add-item
func (s *DashboardService) AddIssueToDashboard(ctx context.Context, owner, kanbanId, repo, issueNumber string) (bool, error) {
	urlStr := fmt.Sprintf("org/%s/kanban/%s/add_item", owner, kanbanId)
	issueNumberVal, _ := strconv.Atoi(issueNumber)
	req, err := newRequest(s.api, http.MethodPost, urlStr, &DashboardRequestContent{
		Repo:            &repo,
		IssueNumberList: []int{issueNumberVal},
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successCreated(resp), err
}
