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
	"github.com/opensourceways/go-gitcode/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {

	client, mux, _ := mockServer(t)

	want := new(User)
	_ = testdata.ReadTestData(t, testdata.User, want)
	mockResponse(t, mux, "/user", want)

	got, ok, err := client.User.GetUserInfo(context.Background())
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, *want, *got)

	username := "user4"
	mockResponse(t, mux, "/users/"+username, want)
	got, ok, err = client.User.GetUserInfoByUsername(context.Background(), username)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, *want, *got)
}
