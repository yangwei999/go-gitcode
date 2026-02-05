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
	"errors"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewRequestError(t *testing.T) {

	client, _, _ := mockServer(t)

	msg1 := "1231412"
	mockErr := errors.New(msg1)
	patch := gomonkey.ApplyFunc(newRequest, func(api *APIClient, method, urlStr string, body any, handlers ...RequestHandler) (*http.Request, error) {
		return nil, mockErr
	})

	defer patch.Reset()

	targetLabels, ok, err := client.Issues.ListRepoIssueLabels(context.Background(), owner, repo, page)
	compare(t, ok, err, mockErr, targetLabels)

	repo11 := repo
	issue11 := "issue1"
	issue, ok, err := client.Issues.UpdateIssue(context.Background(), owner, "2", &IssueRequest{
		Repository: &repo11,
		Title:      &issue11,
	})
	compare(t, ok, err, mockErr, issue)

	pr, ok, err := client.Issues.ListIssueLinkingPullRequests(context.Background(), owner, repo, number)
	compare(t, ok, err, mockErr, pr)

	commentID := "2567"
	comment := "fajhgdahjksghj"
	issueComment, ok, err := client.Issues.CreateIssueComment(context.Background(), owner, repo, number, &IssueComment{
		Body: &comment,
	})
	compare(t, ok, err, mockErr, issueComment)

	labels, ok, err := client.Issues.ListRepoIssueLabels(context.Background(), owner, repo, page)
	compare(t, ok, err, mockErr, labels)

	mockLabelName := "test label"
	username := "user1"

	result1, ok, err := client.Issues.CreateRepoIssueLabel(context.Background(), owner, repo, &Label{Name: mockLabelName, Color: "#fff"})
	compare(t, ok, err, mockErr, result1)

	result2, ok, err := client.Issues.UpdateRepoIssueLabel(context.Background(), owner, repo, mockLabelName, mockLabelName, "#000000")
	compare(t, ok, err, mockErr, result2)

	ok, err = client.Issues.DeleteRepoIssueLabel(context.Background(), owner, repo, mockLabelName)
	compare(t, ok, err, mockErr, nil)

	result3, ok, err := client.Issues.AddLabelsToIssue(context.Background(), owner, repo, number, []string{mockLabelName})
	compare(t, ok, err, mockErr, result3)

	ok, err = client.Issues.RemoveLabelsFromIssue(context.Background(), owner, repo, number, mockLabelName)
	compare(t, ok, err, mockErr, nil)

	result4, ok, err := client.PullRequests.GetPullRequest(context.Background(), owner, repo, number)
	compare(t, ok, err, mockErr, result4)

	prState := "open"
	result5, ok, err := client.PullRequests.UpdatePullRequest(context.Background(), owner, repo, number, &PullRequestRequest{
		State: &prState,
	})
	compare(t, ok, err, mockErr, result5)

	result6, ok, err := client.PullRequests.ListPullRequestLinkingIssues(context.Background(), owner, repo, number, page)
	compare(t, ok, err, mockErr, result6)

	result7, ok, err := client.PullRequests.CreatePullRequestComment(context.Background(), owner, repo, page, &PullRequestCommentRequest{
		Body: comment,
	})
	compare(t, ok, err, mockErr, result7)

	result8, ok, err := client.PullRequests.AddLabelsToPullRequest(context.Background(), owner, repo, number, []string{mockLabelName})
	compare(t, ok, err, mockErr, result8)

	ok, err = client.PullRequests.RemoveLabelsFromPullRequest(context.Background(), owner, repo, number, []string{mockLabelName})
	compare(t, ok, err, mockErr, nil)

	result9, ok, err := client.Repository.CheckUserIsRepoMember(context.Background(), owner, repo, username)
	compare(t, ok, err, mockErr, nil)
	assert.False(t, result9)

	result10, ok, err := client.Repository.GetRepoAllMember(context.Background(), owner, repo, page)
	compare(t, ok, err, mockErr, result10)

	result11, okAndNg, err := client.Repository.GetRepoMemberPermission(context.Background(), owner, repo, username)
	assert.Equal(t, [2]bool{false, false}, okAndNg)
	assert.Equal(t, (*User)(nil), result11)
	assert.Equal(t, err, mockErr)

	result12, ok, err := client.Repository.GetRepoContributors(context.Background(), owner, repo, "")
	compare(t, ok, err, mockErr, result12)

	result13, ok, err := client.User.GetUserInfo(context.Background())
	compare(t, ok, err, mockErr, result13)

	result14, ok, err := client.PullRequests.ListPullRequestCommits(context.Background(), owner, repo, number, page)
	compare(t, ok, err, mockErr, result14)

	result15, ok, err := client.Issues.GetIssueLabels(context.Background(), owner, "15423", page)
	compare(t, ok, err, mockErr, result15)

	result16, ok, err := client.PullRequests.GetPullRequestChangeFiles(context.Background(), owner, repo, number)
	compare(t, ok, err, mockErr, result16)

	result17, ok, err := client.PullRequests.ListPullRequestOperationLogs(context.Background(), owner, repo, number, "asc", page)
	compare(t, ok, err, mockErr, result17)

	result18, ok, err := client.PullRequests.MergePullRequest(context.Background(), owner, repo, number, nil)
	compare(t, ok, err, mockErr, result18)

	result19, ok, err := client.PullRequests.ListPullRequestComments(context.Background(), owner, repo, number, page, "pr_comment", "")
	compare(t, ok, err, mockErr, result19)

	ok, err = client.PullRequests.UpdatePullRequestComment(context.Background(), owner, repo, commentID, comment)
	compare(t, ok, err, mockErr, nil)

	ok, err = client.PullRequests.DeletePullRequestComment(context.Background(), owner, repo, commentID)
	compare(t, ok, err, mockErr, nil)

	result20, ok, err := client.PullRequests.GetLabelsOfPullRequest(context.Background(), owner, repo, number)
	compare(t, ok, err, mockErr, result20)

	result21, ok, err := client.Repository.GetRepoContentByPath(context.Background(), owner, repo, "1.txt", "main")
	compare(t, ok, err, mockErr, result21)

	result22, ok, err := client.Issues.CreateIssue(context.Background(), owner, nil)
	compare(t, ok, err, mockErr, result22)

	result23, ok, err := client.Issues.GetIssue(context.Background(), owner, repo, number)
	compare(t, ok, err, mockErr, result23)

	result24, ok, err := client.Org.ListOrgMember(context.Background(), owner, "all", page)
	compare(t, ok, err, mockErr, result24)

	ok, err = client.Org.AddOrgMember(context.Background(), owner, username, "admin", "1")
	compare(t, ok, err, mockErr, nil)

	ok, err = client.Org.RemoveOrgMember(context.Background(), owner, username)
	compare(t, ok, err, mockErr, nil)

	result25, ok, err := client.Repository.CreateOrgRepo(context.Background(), owner, &RepositoryRequest{})
	compare(t, ok, err, mockErr, result25)

	result26, ok, err := client.Repository.ListOrgRepo(context.Background(), owner, "private", page)
	compare(t, ok, err, mockErr, result26)

	ok, err = client.Repository.DeleteOrgRepo(context.Background(), owner, repo)
	compare(t, ok, err, mockErr, nil)

	result27, ok, err := client.Repository.GetRepoPermissionMode(context.Background(), owner, repo)
	compare(t, ok, err, mockErr, result27)

	ok, err = client.Repository.UpdateRepoPermissionMode(context.Background(), owner, repo, "6")
	compare(t, ok, err, mockErr, nil)

	mode := "a"
	ok, err = client.Repository.UpdateRepoPermissionMode(context.Background(), owner, repo, mode)
	_, err1 := strconv.Atoi(mode)
	compare(t, ok, err, err1, nil)

	result28, ok, err := client.Repository.GetRepoCustomRoles(context.Background(), owner, repo)
	compare(t, ok, err, mockErr, result28)

	result29, ok, err := client.Repository.GetRepoTrees(context.Background(), owner, repo, "124543", page, "", "")
	compare(t, ok, err, mockErr, result29)

	result30, ok, err := client.Repository.GetRepoFileList(context.Background(), owner, repo, "v0", "OWN")
	compare(t, ok, err, mockErr, result30)

	result31, ok, err := client.User.GetUserInfoByUsername(context.Background(), username)
	compare(t, ok, err, mockErr, result31)

	permission := "pull"
	ok, err = client.Repository.AddRepoMember(context.Background(), owner, repo, username, permission)
	compare(t, ok, err, mockErr, nil)

	ok, err = client.Repository.RemoveRepoMember(context.Background(), owner, repo, username)
	compare(t, ok, err, mockErr, nil)

	result32, ok, err := client.Repository.GetRepoAllBranch(context.Background(), owner, repo, page)
	compare(t, ok, err, mockErr, result32)

	branch1 := "release1"
	result33, ok, err := client.Repository.GetRepoBranch(context.Background(), owner, repo, branch)
	compare(t, ok, err, mockErr, result33)

	ok, err = client.Repository.CreateRepoBranch(context.Background(), owner, repo, branch, branch1)
	compare(t, ok, err, mockErr, nil)

	permission = "none"
	ok, err = client.Repository.CreateRepoBranchProtectedRule(context.Background(), owner, repo, branch, permission)
	compare(t, ok, err, mockErr, nil)

	ok, err = client.Repository.UpdateRepoBranchProtectedRule(context.Background(), owner, repo, branch, permission)
	compare(t, ok, err, mockErr, nil)

	ok, err = client.Repository.RemoveRepoBranchProtectedRule(context.Background(), owner, repo, branch)
	compare(t, ok, err, mockErr, nil)

	ok, err = client.Repository.UpdateRepoPullRequestSetting(context.Background(), owner, repo, nil)
	compare(t, ok, err, mockErr, nil)

	ok, err = client.Repository.UpdateRepoSetting(context.Background(), owner, repo)
	compare(t, ok, err, mockErr, nil)

	var time1 *Timestamp
	assert.True(t, time.Time{}.Equal(time1.ToTime()))

	result34, ok, err := client.Issues.ListIssueComments(context.Background(), owner, repo, number, page, "desc", "")
	compare(t, ok, err, mockErr, result34)
}

func compare(t *testing.T, ok bool, err, mock error, result any) {
	assert.Nil(t, result)
	assert.Equal(t, false, ok)
	assert.Equal(t, mock, err)
}
