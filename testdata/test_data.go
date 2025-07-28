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
package testdata

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	dir                       = string(os.PathSeparator) + "testdata" + string(os.PathSeparator)
	issue                     = dir + string(os.PathSeparator) + "issues" + string(os.PathSeparator)
	IssuesAddLabels           = issue + "issues_add_labels.json"
	IssuesComment             = issue + "issues_comment.json"
	IssuesCreateLabel         = issue + "issues_create_label.json"
	IssuesDeleteLabelFailed   = issue + "issues_delete_label_failed.json"
	IssuesHavingLabels        = issue + "issues_having_labels.json"
	IssuesLinkingPrs          = issue + "issues_linking_prs.json"
	IssuesListComments        = issue + "issues_list_comments.json"
	IssuesListLabels          = issue + "issues_list_labels.json"
	IssuesUpdate              = issue + "issues_update.json"
	pr                        = dir + string(os.PathSeparator) + "pr" + string(os.PathSeparator)
	PullRequests              = pr + "pull_requests.json"
	PullRequestsAddLabels     = pr + "pull_requests_add_labels.json"
	PullRequestsClosed        = pr + "pull_requests_closed.json"
	PullRequestsComments      = pr + "pull_requests_comments.json"
	PullRequestsCommits       = pr + "pull_requests_commits.json"
	PullRequestsCreateComment = pr + "pull_requests_create_comment.json"
	PullRequestsFiles         = pr + "pull_requests_files.json"
	PullRequestsLinkingIssues = pr + "pull_requests_linking_issues.json"
	PullRequestsMerge         = pr + "pull_requests_merge.json"
	PullRequestsOpen          = pr + "pull_requests_open.json"
	PullRequestsOperateLog    = pr + "pull_requests_operate_log.json"
	repo                      = dir + string(os.PathSeparator) + "repos" + string(os.PathSeparator)
	RepositoryBranchDetail    = repo + "repository_branch_detail.json"
	RepositoryBranches        = repo + "repository_branches.json"
	RepositoryContributors    = repo + "repository_contributors.json"
	RepositoryCreate          = repo + "repository_create.json"
	RepositoryCustomRoles     = repo + "repository_custom_roles.json"
	RepositoryFileContent     = repo + "repository_file_content.json"
	RepositoryMembers         = repo + "repository_members.json"
	RepositoryTree            = repo + "repository_tree.json"
	user                      = dir + string(os.PathSeparator) + "user" + string(os.PathSeparator)
	OrgMember                 = user + "org_member.json"
	User                      = user + "user.json"
	webhook                   = dir + string(os.PathSeparator) + "webhook" + string(os.PathSeparator)
	IssuesCreate              = webhook + "issues_create.json"
	IssuesNote                = webhook + "issues_note.json"
	PrCreate                  = webhook + "pr_create.json"
	PrNote                    = webhook + "pr_note.json"
	PushCode                  = webhook + "push_code.json"
)

func ReadTestData(t *testing.T, path string, ptr any) []byte {
	pwd, _ := os.Getwd()
	workDir := filepath.Dir(pwd)
	absPath, err := filepath.Abs(workDir + path)
	if err != nil {
		if t != nil {
			t.Errorf("finding [%s], err: %v", workDir+path, err)
		}
		return nil
	}
	if _, err = os.Stat(filepath.Clean(absPath)); !os.IsNotExist(err) {
		data, err := os.ReadFile(absPath)
		if err != nil {
			t.Error(path + " read failed")
			return nil
		}
		if ptr != nil {
			err = json.Unmarshal(data, ptr)
			if err != nil {
				_, _, line, _ := runtime.Caller(1)
				t.Errorf("code line: %d, error: %v", line, err)
			}
		}
		return data
	} else {
		if t != nil {
			t.Error(absPath + " not found")
		}
		return nil
	}
}
