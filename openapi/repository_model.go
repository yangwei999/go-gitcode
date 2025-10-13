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

// Repository represents a GitHub repository.
type Repository struct {
	ID                 *int64        `json:"id,omitempty"`
	Name               *string       `json:"name,omitempty"`
	Path               *string       `json:"path,omitempty"`
	FullName           *string       `json:"full_name,omitempty"`
	Description        *string       `json:"description,omitempty"`
	Homepage           *string       `json:"homepage,omitempty"`
	DefaultBranch      *string       `json:"default_branch,omitempty"`
	CreatedAt          *Timestamp    `json:"created_at,omitempty"`
	PushedAt           *Timestamp    `json:"pushed_at,omitempty"`
	UpdatedAt          *Timestamp    `json:"updated_at,omitempty"`
	HTMLURL            *string       `json:"html_url,omitempty"`
	MirrorURL          *string       `json:"mirror_url,omitempty"`
	Language           *string       `json:"language,omitempty"`
	Status             *string       `json:"status,omitempty"`
	Internal           *bool         `json:"internal,omitempty"`
	Fork               *bool         `json:"fork,omitempty"`
	ForksCount         *int          `json:"forks_count,omitempty"`
	NetworkCount       *int          `json:"network_count,omitempty"`
	OpenIssuesCount    *int          `json:"open_issues_count,omitempty"`
	StargazersCount    *int          `json:"stargazers_count,omitempty"`
	SubscribersCount   *int          `json:"subscribers_count,omitempty"`
	WatchersCount      *int          `json:"watchers_count,omitempty"`
	Size               *int          `json:"size,omitempty"`
	AutoInit           *bool         `json:"auto_init,omitempty"`
	Parent             *Repository   `json:"parent,omitempty"`
	Source             *Repository   `json:"source,omitempty"`
	TemplateRepository *Repository   `json:"template_repository,omitempty"`
	Organization       *Organization `json:"namespace,omitempty"`
	Topics             []string      `json:"topics,omitempty"`
	Archived           *bool         `json:"archived,omitempty"`

	// Additional mutable fields when creating and editing a repository
	Public            *bool   `json:"public,omitempty"`
	Private           *bool   `json:"private,omitempty"`
	HasIssues         *bool   `json:"has_issues,omitempty"`
	HasWiki           *bool   `json:"has_wiki,omitempty"`
	HasPages          *bool   `json:"has_pages,omitempty"`
	HasProjects       *bool   `json:"has_projects,omitempty"`
	HasDownloads      *bool   `json:"has_downloads,omitempty"`
	HasDiscussions    *bool   `json:"has_discussions,omitempty"`
	IsTemplate        *bool   `json:"is_template,omitempty"`
	LicenseTemplate   *string `json:"license_template,omitempty"`
	GitignoreTemplate *string `json:"gitignore_template,omitempty"`
}

type RepositoryCommit struct {
	ID          *string     `json:"id,omitempty"`
	Title       *string     `json:"title,omitempty"`
	Message     *string     `json:"message,omitempty"`
	SHA         *string     `json:"sha,omitempty"`
	Commit      *Commit     `json:"commit,omitempty"`
	Author      *CommitUser `json:"author,omitempty"`
	Committer   *CommitUser `json:"committer,omitempty"`
	Parents     *Commit     `json:"parents,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	URL         *string     `json:"url,omitempty"`
	CommentsURL *string     `json:"comments_url,omitempty"`
}

type Commit struct {
	SHA         *string     `json:"sha,omitempty"`
	Author      *CommitUser `json:"author,omitempty"`
	Committer   *CommitUser `json:"committer,omitempty"`
	Message     *string     `json:"message,omitempty"`
	Parents     *Commit     `json:"parents,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	URL         *string     `json:"url,omitempty"`
	CommentsURL *int        `json:"comments_url,omitempty"`
}

type CommitUser struct {
	Login *string    `json:"login,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`
	Date  *Timestamp `json:"date,omitempty"`
}

type CommitPatch struct {
	Diff        *string `json:"diff,omitempty"`
	OldPath     *string `json:"old_path,omitempty"`
	NewPath     *string `json:"new_path,omitempty"`
	AMode       *string `json:"a_mode,omitempty"`
	BMode       *string `json:"b_mode,omitempty"`
	NewFile     *bool   `json:"new_file,omitempty"`
	RenamedFile *bool   `json:"renamed_file,omitempty"`
	DeletedFile *bool   `json:"deleted_file,omitempty"`
	TooLarge    *bool   `json:"too_large,omitempty"`
}

type CommitFile struct {
	SHA              *string      `json:"sha,omitempty"`
	Filename         *string      `json:"filename,omitempty"`
	Additions        *int         `json:"additions,omitempty"`
	Deletions        *int         `json:"deletions,omitempty"`
	Changes          *int         `json:"changes,omitempty"`
	Status           *string      `json:"status,omitempty"`
	Patch            *CommitPatch `json:"patch,omitempty"`
	BlobURL          *string      `json:"blob_url,omitempty"`
	RawURL           *string      `json:"raw_url,omitempty"`
	ContentsURL      *string      `json:"contents_url,omitempty"`
	PreviousFilename *string      `json:"previous_filename,omitempty"`
}

// RepositoryContent represents a file or directory in a github repository.
type RepositoryContent struct {
	Type *string `json:"type,omitempty"`
	// Target is only set if the type is "symlink" and the target is not a normal file.
	// If Target is set, Path will be the symlink path.
	Target   *string `json:"target,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	Size     *int64  `json:"size,omitempty"`
	Name     *string `json:"name,omitempty"`
	Path     *string `json:"path,omitempty"`
	// Content contains the actual file content, which may be encoded.
	// Callers should call GetContent which will decode the content if
	// necessary.
	Content         *string `json:"content,omitempty"`
	SHA             *string `json:"sha,omitempty"`
	URL             *string `json:"url,omitempty"`
	GitURL          *string `json:"git_url,omitempty"`
	HTMLURL         *string `json:"html_url,omitempty"`
	DownloadURL     *string `json:"download_url,omitempty"`
	SubmoduleGitURL *string `json:"submodule_git_url,omitempty"`
}

// Contributor represents a repository contributor
type Contributor struct {
	Contributions *int64  `json:"contributions,omitempty"`
	Name          *string `json:"name,omitempty"`
	Email         *string `json:"email,omitempty"`
}

// Branch represents a repository branch
type Branch struct {
	Name      *string           `json:"name,omitempty"`
	Commit    *RepositoryCommit `json:"commit,omitempty"`
	Protected *bool             `json:"protected,omitempty"`
}

type BranchRequest struct {
	CreateFrom string `json:"refs,omitempty"`
	BranchName string `json:"branch_name,omitempty"`
}

type BranchProtectedRuleRequest struct {
	PushPermission  string `json:"pusher,omitempty"`
	MergePermission string `json:"merger,omitempty"`
	BranchName      string `json:"wildcard,omitempty"`
}

type RepositoryRequest struct {
	Name              *string `json:"name,omitempty"`
	Description       *string `json:"description,omitempty"`
	Homepage          *string `json:"homepage,omitempty"`
	HasIssues         *bool   `json:"has_issues,omitempty"`
	HasWiki           *bool   `json:"has_wiki,omitempty"`
	AutoInit          *bool   `json:"auto_init,omitempty"`
	Private           *bool   `json:"private,omitempty"`
	Public            *int    `json:"public,omitempty"`
	DefaultBranch     *string `json:"default_branch,omitempty"`
	LicenseTemplate   *string `json:"license_template,omitempty"`
	GitignoreTemplate *string `json:"gitignore_template,omitempty"`
	Path              *string `json:"path,omitempty"`
	ImportUrl         *string `json:"import_url,omitempty"`
}

type RepositoryPermissionMode struct {
	Model int `json:"memberMgntMode,omitempty"`
}

type RepositoryPermissionModeRequest struct {
	Model int `json:"mode,omitempty"`
}

type RepositoryMemberPermission struct {
	Permission *string `json:"permission,omitempty"`
}

type Milestone struct {
	URL         *string    `json:"url,omitempty"`
	Number      *int       `json:"number,omitempty"`
	State       *string    `json:"state,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
	ClosedAt    *Timestamp `json:"closed_at,omitempty"`
	DueOn       *Timestamp `json:"due_on,omitempty"`
}

type CustomRepoRoles struct {
	ID          *string    `json:"role_id,omitempty"`
	AccessLevel *int       `json:"access_level,omitempty"`
	Name        *string    `json:"role_name,omitempty"`
	Description *string    `json:"role_description,omitempty"`
	Type        *string    `json:"role_type,omitempty"`
	MemberCount *int       `json:"member_count,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}

type RepositoryTree struct {
	Sha  *string              `json:"sha,omitempty"`
	Tree []*RepositoryContent `json:"tree,omitempty"`
}

type RepositoryRepoPullRequestSettingRequest struct {
	DisableMergeBySelf                        *bool   `json:"disable_merge_by_self,omitempty"`
	AddNotesAfterMerged                       *bool   `json:"add_notes_after_merged,omitempty"`
	CanReopen                                 *bool   `json:"can_reopen,omitempty"`
	AllowLiteMergeRequest                     *bool   `json:"is_allow_lite_merge_request,omitempty"`
	ApprovalRequiredReviewersEnable           *bool   `json:"approval_required_reviewers_enable,omitempty"`
	ApprovalRequiredReviewers                 *int    `json:"approval_required_reviewers,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool   `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	DisableSquashMerge                        *bool   `json:"disable_squash_merge,omitempty"`
	AutoSquashMerge                           *bool   `json:"auto_squash_merge,omitempty"`
	MergeMethod                               *string `json:"merge_method,omitempty"`
	SquashMergeWithNoMergeCommit              *bool   `json:"squash_merge_with_no_merge_commit,omitempty"`
	MergedCommitAuthor                        *string `json:"merged_commit_author,omitempty"`
	ApprovalApproverIds                       *string `json:"approval_approver_ids,omitempty"`
	ApprovalRequiredApprovers                 *int    `json:"approval_required_approvers,omitempty"`
	ApprovalTesterIds                         *string `json:"approval_tester_ids,omitempty"`
	ApprovalRequiredTesters                   *int    `json:"approval_required_testers,omitempty"`
	LiteMergeRequestPrefixTitle               *string `json:"lite_merge_request_prefix_title,omitempty"`
	CloseIssueWhenMrMerged                    *bool   `json:"close_issue_when_mr_merged,omitempty"`
}

type RepositoryRepoSettingRequest struct {
	GeneratePreMergeRef            bool `json:"generate_pre_merge_ref,omitempty"`
	ForbiddenCommitterCreateBranch bool `json:"forbidden_committer_create_branch,omitempty"`
	ForbiddenDeveloperCreateBranch bool `json:"forbidden_developer_create_branch,omitempty"`
	ForbiddenDeveloperCreateTag    bool `json:"forbidden_developer_create_tag,omitempty"`
}

type RepositoryTag struct {
	Name       *string `json:"name,omitempty"`
	Message    *string `json:"message,omitempty"`
	Commit     *Commit `json:"commit,omitempty"`
	ZipballURL *string `json:"zipball_url,omitempty"`
	TarballURL *string `json:"tarball_url,omitempty"`
}

// FileRequest 文件操作请求参数
type FileRequest struct {
	Content *string `json:"content"`
	Message *string `json:"message"`
	Branch  *string `json:"branch,omitempty"`
	Sha     *string `json:"sha,omitempty"`
}

// FileCommitResponse 文件操作响应
type FileCommitResponse struct {
	Content *RepositoryContent `json:"content,omitempty"`
	Commit  *RepositoryCommit  `json:"commit,omitempty"`
}
