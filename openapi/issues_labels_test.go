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
	"encoding/json"
	"fmt"
	"github.com/opensourceways/go-gitcode/testdata"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIssuesLabels(t *testing.T) {
	client, mux, _ := mockServer(t)

	// ListRepoIssueLabels CreateRepoIssueLabel
	listOrCreateLabel(t, client, mux)
	// UpdateRepoIssueLabel DeleteRepoIssueLabel
	updateOrDeleteLabel(t, client, mux)
	// AddLabelsToIssue
	addIssueLabel(t, client, mux)
	// RemoveLabelsFromIssue
	removeIssueLabel(t, client, mux)
	// GetIssueLabels
	getIssueLabels(t, client, mux)
}

func listOrCreateLabel(t *testing.T, client *APIClient, mux *http.ServeMux) {

	var allLabels []*Label
	_ = testdata.ReadTestData(t, testdata.IssuesListLabels, &allLabels)
	var addLabel Label
	_ = testdata.ReadTestData(t, testdata.IssuesCreateLabel, &addLabel)
	urlStr := fmt.Sprintf("/repos/%s/%s/labels", owner, repo)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		var err error
		if r.Method == http.MethodGet {
			err = json.NewEncoder(w).Encode(allLabels)
		} else if r.Method == http.MethodPost {
			err = json.NewEncoder(w).Encode(addLabel)
		}
		if err != nil {
			t.Errorf("IssuesLabels mock response data error: %v", err)
		}
	})
	urlStr = fmt.Sprintf("/repos/%s/%s/labels", owner, "333")
	errMsg := "{\n    \"error_code\": 500,\n    \"error_code_name\": \"FAIL\",\n    \"error_message\": \"系统错误\",\n    \"trace_id\": \"d0834ebae0074f5cab66ef8b8fc529d5\"\n}"
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		// a word wrap \n write to response
		http.Error(w, errMsg, http.StatusInternalServerError)
	})

	got, ok, err := client.Issues.ListRepoIssueLabels(context.Background(), owner, repo)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range allLabels {
		assert.Equal(t, *allLabels[i], *got[i])
	}

	got, ok, err = client.Issues.ListRepoIssueLabels(context.Background(), owner, "333")
	assert.False(t, ok)
	assert.Nil(t, got)
	assert.Equal(t, errMsg+"\n", err.Error())

	got1, ok, err := client.Issues.CreateRepoIssueLabel(context.Background(), owner, repo, &Label{Name: "gfa", Color: "#000"})
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, addLabel, *got1)
}

func updateOrDeleteLabel(t *testing.T, client *APIClient, mux *http.ServeMux) {

	var want Label
	_ = testdata.ReadTestData(t, testdata.IssuesCreateLabel, &want)
	oldLabelName := "gfa"
	urlStr := fmt.Sprintf("/repos/%s/%s/labels/%s", owner, repo, oldLabelName)
	mockResponse(t, mux, urlStr, want)

	got, ok, err := client.Issues.UpdateRepoIssueLabel(context.Background(), owner, repo, oldLabelName, "gf3", "#001")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, want, *got)

	ok, err = client.Issues.DeleteRepoIssueLabel(context.Background(), owner, repo, oldLabelName)
	assert.Nil(t, err)
	assert.True(t, ok)

	data := testdata.ReadTestData(t, testdata.IssuesDeleteLabelFailed, nil)
	labelName := "canNotDeleteLabel"
	urlStr = fmt.Sprintf("/repos/%s/%s/labels/%s", owner, repo, labelName)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(data)
	})

	ok, err = client.Issues.DeleteRepoIssueLabel(context.Background(), owner, repo, labelName)
	assert.False(t, ok)
	assert.Equal(t, string(data), err.Error())
}

func addIssueLabel(t *testing.T, client *APIClient, mux *http.ServeMux) {
	var want []*Label
	_ = testdata.ReadTestData(t, testdata.IssuesAddLabels, &want)
	labelName := "add-issue-label"
	urlStr := fmt.Sprintf("/repos/%s/%s/issues/%s/labels", owner, repo, number)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(want)
		if err != nil {
			t.Errorf("Issues.AddLabelsToIssue mock response data error: %v", err)
		}
	})

	got, ok, err := client.Issues.AddLabelsToIssue(context.Background(), owner, repo, number, []string{labelName})
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range want {
		assert.Equal(t, *want[i], *got[i])
	}

}

func removeIssueLabel(t *testing.T, client *APIClient, mux *http.ServeMux) {
	labelName := "remove-issue-label"
	urlStr := fmt.Sprintf("/repos/%s/%s/issues/%s/labels/%s", owner, repo, number, labelName)
	mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNoContent)
	})

	ok, err := client.Issues.RemoveLabelsFromIssue(context.Background(), owner, repo, number, labelName)
	assert.Nil(t, err)
	assert.True(t, ok)

}

func getIssueLabels(t *testing.T, client *APIClient, mux *http.ServeMux) {
	var labels []*Label
	_ = testdata.ReadTestData(t, testdata.IssuesHavingLabels, &labels)
	issueID := "5342"
	urlStr := fmt.Sprintf("/enterprises/%s/issues/%s/labels", owner, issueID)
	mockResponse(t, mux, urlStr, labels)

	result, ok, err := client.Issues.GetIssueLabels(context.Background(), owner, issueID, page)
	assert.Nil(t, err)
	assert.True(t, ok)
	for i := range labels {
		assert.Equal(t, *labels[i], *result[i])
	}

}
