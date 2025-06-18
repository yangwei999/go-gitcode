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
	"net/http"
)

// GetUserInfo 获取授权用户的资料
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-user
func (s *UserService) GetUserInfo(ctx context.Context) (*User, bool, error) {
	req, err := newRequest(s.api, http.MethodGet, "user", nil)
	if err != nil {
		return nil, false, err
	}

	userInfo := new(User)
	resp, err := s.api.Do(ctx, req, userInfo)
	return userInfo, successGetData(resp), err
}

// GetUserInfoByUsername 获取一个用户
//
// api Docs: https://docs.gitcode.com/docs/apis/get-api-v-5-users-username
func (s *UserService) GetUserInfoByUsername(ctx context.Context, username string) (*User, bool, error) {
	req, err := newRequest(s.api, http.MethodGet, "users/"+username, nil)
	if err != nil {
		return nil, false, err
	}

	userInfo := new(User)
	resp, err := s.api.Do(ctx, req, userInfo)
	return userInfo, successGetData(resp), err
}
