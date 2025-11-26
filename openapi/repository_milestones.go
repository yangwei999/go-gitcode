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

// GetRepoAllMilestones 获取仓库所有里程碑
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-repos-owner-repo-milestones
func (s *RepositoryService) GetRepoAllMilestones(ctx context.Context, owner, repo string, request *MileStonesRequest) ([]*MileStonesResponse, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators", owner, repo)
	req, err := newRequest(s.api, http.MethodPut, urlStr, request)
	if err != nil {
		return nil, false, err
	}

	var milestones []*MileStonesResponse
	resp, err := s.api.Do(ctx, req, &milestones)
	return milestones, successModified(resp), err
}
