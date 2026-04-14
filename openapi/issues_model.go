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
	"encoding/json"
	"net/url"
)

// Label represents a GitCode label on an Issue
type Label struct {
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
	Title string `json:"title,omitempty"`
}

func (l *Label) Form() *url.Values {
	return &url.Values{
		"name":  []string{l.Name},
		"color": []string{l.Color},
	}
}

type Issue struct {
	ID               *int64            `json:"id,omitempty"`
	HTMLURL          *string           `json:"html_url,omitempty"`
	Number           *string           `json:"number,omitempty"`
	IssueType        *string           `json:"issue_type,omitempty"`
	State            *string           `json:"state,omitempty"`
	IssueState       *string           `json:"issue_state,omitempty"`
	IssueStateDetail *IssueStateDetail `json:"issue_state_detail,omitempty"`
	Priority         *int              `json:"priority,omitempty"`
	Title            *string           `json:"title,omitempty"`
	Body             *string           `json:"body,omitempty"`
	User             *User             `json:"user,omitempty"`
	Assignee         *User             `json:"assignee,omitempty"`
	Repository       *Repository       `json:"repository,omitempty"`
	Labels           []*Label          `json:"labels,omitempty"`
	CreatedAt        *Timestamp        `json:"created_at,omitempty"`
	UpdatedAt        *Timestamp        `json:"updated_at,omitempty"`
	ClosedAt         *Timestamp        `json:"finished_at,omitempty"`
	ClosedBy         *User             `json:"closed_by,omitempty"`

	PullRequestLinks *PullRequestLinks `json:"pull_request,omitempty"`
	Milestone        *Milestone        `json:"milestone,omitempty"`
	DashboardList    []*DashboardItem  `json:"kanban_values,omitempty"`
}

type IssueStateDetail struct {
	Title  *string `json:"title,omitempty"`
	Serial *int    `json:"serial,omitempty"`
}

type PullRequestLinks struct {
	URL      *string    `json:"url,omitempty"`
	HTMLURL  *string    `json:"html_url,omitempty"`
	DiffURL  *string    `json:"diff_url,omitempty"`
	PatchURL *string    `json:"patch_url,omitempty"`
	MergedAt *Timestamp `json:"merged_at,omitempty"`
}

type IssueRequest struct {
	Repository    *string       `json:"repo,omitempty"` // 仓库地址
	Title         *string       `json:"title,omitempty"`
	Body          *string       `json:"body,omitempty"`
	Labels        *string       `json:"labels,omitempty"`   // 用逗号分开的标签
	Assignee      *string       `json:"assignee,omitempty"` // Issue负责人的 username
	State         *string       `json:"state,omitempty"`
	Milestone     *int          `json:"milestone,omitempty"`
	SecurityHole  *string       `json:"security_hole,omitempty"`  // 是否是私有issue(默认为false)
	IssueSeverity *string       `json:"issue_severity,omitempty"` // 优先级
	TemplatePath  *string       `json:"template_path,omitempty"`  // issue模板路径，即创建issue的类型
	IssueType     *string       `json:"issue_type,omitempty"`     // issue类型
	IssueState    *string       `json:"status,omitempty"`         // issue具体状态
	CustomFields  []CustomField `json:"custom_fields,omitempty"`  // 自定义字段（企业版支持）
}

type IssueComment struct {
	ID        *json.Number `json:"id,omitempty"`
	Body      *string      `json:"body,omitempty"`
	User      *User        `json:"user,omitempty"`
	CreatedAt *Timestamp   `json:"created_at,omitempty"`
	UpdatedAt *Timestamp   `json:"updated_at,omitempty"`
	Target    *Issue       `json:"target,omitempty"`
}

type CustomField struct {
	FieldName   string   `json:"field_name"`
	FieldValues []string `json:"field_values"`
}
