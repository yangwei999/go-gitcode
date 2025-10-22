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

import "encoding/json"

// PullRequest represents a GitHub pull request on a repository.
type PullRequest struct {
	ID             *int64          `json:"id,omitempty"`
	Number         *int64          `json:"number,omitempty"`
	State          *string         `json:"state,omitempty"`
	Title          *string         `json:"title,omitempty"`
	Body           *string         `json:"body,omitempty"`
	CreatedAt      *Timestamp      `json:"created_at,omitempty"`
	UpdatedAt      *Timestamp      `json:"updated_at,omitempty"`
	ClosedAt       *Timestamp      `json:"closed_at,omitempty"`
	MergedAt       *Timestamp      `json:"merged_at,omitempty"`
	Labels         []*Label        `json:"labels,omitempty"`
	User           *User           `json:"user,omitempty"`
	Draft          *bool           `json:"draft,omitempty"`
	CanMergeCheck  *bool           `json:"can_merge_check,omitempty"`
	Merged         *bool           `json:"merged,omitempty"`
	MergeAble      *bool           `json:"mergeable,omitempty"`
	MergeAbleState *MergeAbleState `json:"mergeable_state,omitempty"`
	MergedBy       *User           `json:"merged_by,omitempty"`
	PruneBranch    *bool           `json:"prune_branch,omitempty"`

	Head *PullRequestBranch `json:"head,omitempty"`
	Base *PullRequestBranch `json:"base,omitempty"`

	URL                 *string    `json:"url,omitempty"`
	HTMLURL             *string    `json:"html_url,omitempty"`
	IssueURL            *string    `json:"issue_url,omitempty"`
	StatusesURL         *string    `json:"statuses_url,omitempty"`
	DiffURL             *string    `json:"diff_url,omitempty"`
	PatchURL            *string    `json:"patch_url,omitempty"`
	CommitsURL          *string    `json:"commits_url,omitempty"`
	CommentsURL         *string    `json:"comments_url,omitempty"`
	ReviewCommentsURL   *string    `json:"review_comments_url,omitempty"`
	ReviewCommentURL    *string    `json:"review_comment_url,omitempty"`
	ReviewComments      *int       `json:"review_comments,omitempty"`
	Assignee            *User      `json:"assignee,omitempty"`
	Assignees           []*User    `json:"assignees,omitempty"`
	AssigneesNumber     *int       `json:"assignees_number,omitempty"`
	Testers             []*User    `json:"testers,omitempty"`
	TestersNumber       *int       `json:"testers_number,omitempty"`
	MaintainerCanModify *bool      `json:"maintainer_can_modify,omitempty"`
	AuthorAssociation   *string    `json:"author_association,omitempty"`
	NodeID              *string    `json:"node_id,omitempty"`
	RequestedReviewers  []*User    `json:"requested_reviewers,omitempty"`
	Milestone           *Milestone `json:"milestone,omitempty"`
}

// PullRequestBranch represents a base or head branch in a GitHub pull request.
type PullRequestBranch struct {
	Label *string     `json:"label,omitempty"`
	Ref   *string     `json:"ref,omitempty"`
	SHA   *string     `json:"sha,omitempty"`
	Repo  *Repository `json:"repo,omitempty"`
	User  *User       `json:"user,omitempty"`
}

type MergeAbleState struct {
	MergeRequestID          *int64               `json:"merge_request_id,omitempty"`
	State                   *bool                `json:"state,omitempty"`
	ConflictPassed          *bool                `json:"conflict_passed,omitempty"`
	ReviewersPass           *bool                `json:"approval_reviewers_required_passed,omitempty"`
	AssigneesPass           *bool                `json:"approval_approvers_required_passed,omitempty"`
	TestersPass             *bool                `json:"approval_testers_required_passed,omitempty"`
	ResolveDiscussionPassed *bool                `json:"resolve_discussion_passed,omitempty"`
	PRSetting               *MergeRequestSetting `json:"merge_request_switch,omitempty"`
}

type MergeRequestSetting struct {
	MergeMethod                               *string `json:"merge_method,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool   `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
}

type PullRequestRequest struct {
	ID              int64       `json:"id,omitempty"`
	Title           string      `json:"title,omitempty"`
	Body            string      `json:"body,omitempty"`
	State           string      `json:"state,omitempty"`
	Labels          string      `json:"labels,omitempty"`
	MilestoneNumber string      `json:"milestone_number,omitempty"`
	Draft           string      `json:"draft,omitempty"`
	User            User        `json:"user,omitempty"`
	Target          PullRequest `json:"target,omitempty"`
}

// PullRequestCreateRequest 创建PR的请求参数
type PullRequestCreateRequest struct {
	Title               string `json:"title"`
	Head                string `json:"head"`
	Base                string `json:"base"`
	Body                string `json:"body,omitempty"`
	MilestoneNumber     string `json:"milestone_number,omitempty"`
	Labels              string `json:"labels,omitempty"`
	Issue               string `json:"issue,omitempty"`
	Assignees           string `json:"assignees,omitempty"`
	Testers             string `json:"testers,omitempty"`
	PruneSourceBranch   bool   `json:"prune_source_branch,omitempty"`
	Draft               bool   `json:"draft,omitempty"`
	Squash              bool   `json:"squash,omitempty"`
	SquashCommitMessage string `json:"squash_commit_message,omitempty"`
	ForkPath            string `json:"fork_path,omitempty"` // fork项目路径【owner/repo】，跨仓PR 必填。

}

type PullRequestRequestMerge struct {
	Method string `json:"merge_method,omitempty"`
}

type SimpleComment struct {
	Body string `json:"body,omitempty"`
}

type PullRequestComment struct {
	ID        *json.Number `json:"id,omitempty"`
	Body      *string      `json:"body,omitempty"`
	User      *User        `json:"user,omitempty"`
	CreatedAt *Timestamp   `json:"created_at,omitempty"`
	UpdatedAt *Timestamp   `json:"updated_at,omitempty"`
}

type PullRequestCommentRequest struct {
	Body     string `json:"body,omitempty"`
	Path     string `json:"path,omitempty"`
	Position string `json:"position,omitempty"`
}

// PullRequestOperationLog represents a comment in a GitHub DiscussionCommentEvent.
type PullRequestOperationLog struct {
	Project        *string    `json:"project,omitempty"`
	CreatedAt      *Timestamp `json:"created_at,omitempty"`
	DiscussionID   *string    `json:"discussion_id,omitempty"`
	ID             *int64     `json:"id,omitempty"`
	Content        *string    `json:"content,omitempty"`
	Action         *string    `json:"action,omitempty"`
	MergeRequestId *int64     `json:"merge_request_id,omitempty"`
	UpdatedAt      *Timestamp `json:"updated_at,omitempty"`
	User           *User      `json:"user,omitempty"`
}

type PullRequestMergedResult struct {
	SHA     *string `json:"sha,omitempty"`
	Merged  *bool   `json:"merged,omitempty"`
	Message *string `json:"message,omitempty"`
}

type PullRequestQueryRequest struct {
	State           *string `json:"state,omitempty"`
	Base            *string `json:"base,omitempty"`
	Since           *string `json:"since,omitempty"`
	Direction       *string `json:"direction,omitempty"`
	Sort            *string `json:"sort,omitempty"`
	MilestoneNumber *string `json:"milestone_number,omitempty"`
	Labels          *string `json:"labels,omitempty"`
	Page            *string `json:"page,omitempty"`
	PerPage         *string `json:"per_page,omitempty"`
	Author          *string `json:"author,omitempty"`
	Assignee        *string `json:"assignee,omitempty"`
	Reviewer        *string `json:"reviewer,omitempty"`
	MergedAfter     *string `json:"merged_after,omitempty"`
	MergedBefore    *string `json:"merged_before,omitempty"`
	CreatedAfter    *string `json:"created_after,omitempty"`
	CreatedBefore   *string `json:"created_before,omitempty"`
	UpdatedBefore   *string `json:"updated_before,omitempty"`
	UpdatedAfter    *string `json:"updated_after,omitempty"`
}
