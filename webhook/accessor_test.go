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
package webhook

import (
	"bytes"
	"encoding/json"
	"github.com/opensourceways/go-gitcode/testdata"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"reflect"
	"testing"
)

const (
	htmlUrl = "https://gitcode.com/ibforuorg/test1/issues/4"
)

func TestGetAccessor(t *testing.T) {

	createIssue(t)
	createPR(t)
	notePR(t)
	noteIssue(t)
	pushCode(t)

	buf := &bytes.Buffer{}
	buf.Write([]byte("kjhygadsskhj"))
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/11", buf)
	req.Header.Set(HeaderEventType, "Dummy Hook")
	req.Header.Set(HeaderEventGUID, "fasgasd")

	a := new(GitCodeAccessor)
	getPayload(t, a, req)
	got1, got2, got3 := a.GetAccessor(req)
	assert.Equal(t, nil, got1)
	assert.Equal(t, "Dummy Hook", *got2)
	assert.Equal(t, "fasgasd", *got3)

	a1 := new(GitCodeAccessor)
	got1, got2, got3 = a1.GetAccessor(req)
	assert.Nil(t, got1)
	assert.Nil(t, got2)
	assert.Nil(t, got3)
}

func getPayload(t *testing.T, a *GitCodeAccessor, r *http.Request) {
	var payload bytes.Buffer
	_, err := io.Copy(&payload, r.Body)
	assert.Nil(t, err)
	a.Payload = &payload
}

func createIssue(t *testing.T) {
	want := GitCodeAccessor{Issues: new(IssueEvent)}
	data := testdata.ReadTestData(t, testdata.IssuesCreate, want.Issues)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/0", buf)
	req.Header.Set(HeaderEventType, issueEvent)
	req.Header.Set(HeaderEventGUID, "1231321")

	a := new(GitCodeAccessor)
	getPayload(t, a, req)
	got1, got2, got3 := a.GetAccessor(req)
	d1, _ := json.Marshal(want.Issues)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, issueEvent, *got2)
	assert.Equal(t, "1231321", *got3)

	issue, _ := got1.(*IssueEvent)
	assert.Equal(t, "open", *issue.GetAction())
	assert.Equal(t, (*string)(nil), issue.GetActionDetail())
	assert.Equal(t, "opened", *issue.GetState())
	assert.Equal(t, "ibforuorg", *issue.GetOrg())
	assert.Equal(t, "test1", *issue.GetRepo())
	assert.Equal(t, htmlUrl, *issue.GetHtmlURL())
	assert.Equal(t, "4", *issue.GetNumber())
	assert.Equal(t, "515443", *issue.GetID())
	assert.Equal(t, "*****", *issue.GetAuthor())
	assert.Equal(t, "2024-10-26T10:28:03+08:00", *issue.GetCreateTime())
	assert.Equal(t, "2024-10-26T10:28:03+08:00", *issue.GetUpdateTime())
	assert.Equal(t, "public", *issue.GetRepoVisibility())

	issue = new(IssueEvent)
	rt := reflect.TypeOf(issue)
	n := rt.NumMethod()
	for i := 0; i < n; i++ {
		rm := rt.Method(i)
		ret := rm.Func.Call([]reflect.Value{reflect.ValueOf(issue)})
		assert.Equal(t, (*string)(nil), ret[0].Interface())
	}
}

func pushCode(t *testing.T) {
	want := GitCodeAccessor{Push: new(PushEvent)}
	data := testdata.ReadTestData(t, testdata.PushCode, want.Push)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/1", buf)
	req.Header.Set(HeaderEventType, "Push Hook")
	req.Header.Set(HeaderEventGUID, "fasgasd")

	a := new(GitCodeAccessor)
	getPayload(t, a, req)
	got1, got2, got3 := a.GetAccessor(req)
	d1, _ := json.Marshal(want.Push)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, "Push Hook", *got2)
	assert.Equal(t, "fasgasd", *got3)

	pr, _ := got1.(*PushEvent)
	assert.Equal(t, (*string)(nil), pr.GetAction())
	assert.Equal(t, (*string)(nil), pr.GetState())
	assert.Equal(t, "ibforuorg", *pr.GetOrg())
	assert.Equal(t, "org-repo-role-member-manage", *pr.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/org-repo-role-member-manage", *pr.GetHtmlURL())
	assert.Equal(t, "dev", *pr.GetBase())
	assert.Equal(t, "ibforu", *pr.GetAuthor())
	assert.Nil(t, pr.GetRepoVisibility())

	pr = new(PushEvent)
	rt := reflect.TypeOf(pr)
	n := rt.NumMethod()
	for i := 0; i < n; i++ {
		rm := rt.Method(i)
		ret := rm.Func.Call([]reflect.Value{reflect.ValueOf(pr)})
		assert.Equal(t, (*string)(nil), ret[0].Interface())
	}
}

func createPR(t *testing.T) {
	want := GitCodeAccessor{PR: new(PullRequestEvent)}
	data := testdata.ReadTestData(t, testdata.PrCreate, want.PR)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/2", buf)
	req.Header.Set(HeaderEventType, pullRequestEvent)
	req.Header.Set(HeaderEventGUID, "fasgasd")

	a := new(GitCodeAccessor)
	getPayload(t, a, req)
	got1, got2, got3 := a.GetAccessor(req)
	d1, _ := json.Marshal(want.PR)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, pullRequestEvent, *got2)
	assert.Equal(t, "fasgasd", *got3)

	pr, _ := got1.(*PullRequestEvent)
	assert.Equal(t, "open", *pr.GetAction())
	assert.Equal(t, "", *pr.GetActionDetail())
	assert.Equal(t, "opened", *pr.GetState())
	assert.Equal(t, "ibforuorg", *pr.GetOrg())
	assert.Equal(t, "test1", *pr.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1/merge_requests/4", *pr.GetHtmlURL())
	assert.Equal(t, "main", *pr.GetBase())
	assert.Equal(t, "ibforuorg/test1/24124124124", *pr.GetHead())
	assert.Equal(t, "4", *pr.GetNumber())
	assert.Equal(t, "190370", *pr.GetID())
	assert.Equal(t, "****", *pr.GetAuthor())
	assert.Equal(t, "2024-10-26T10:32:40+08:00", *pr.GetCreateTime())
	assert.Equal(t, "2024-10-26T10:32:41+08:00", *pr.GetUpdateTime())
	assert.Equal(t, "private", *pr.GetRepoVisibility())

	pr = new(PullRequestEvent)
	rt := reflect.TypeOf(pr)
	n := rt.NumMethod()
	for i := 0; i < n; i++ {
		rm := rt.Method(i)
		ret := rm.Func.Call([]reflect.Value{reflect.ValueOf(pr)})
		assert.Equal(t, (*string)(nil), ret[0].Interface())
	}
}

func notePR(t *testing.T) {
	want := GitCodeAccessor{Note: new(NoteEvent)}
	data := testdata.ReadTestData(t, testdata.PrNote, want.Note)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/21", buf)
	req.Header.Set(HeaderEventType, noteEvent)
	req.Header.Set(HeaderEventGUID, "651234123")

	a := new(GitCodeAccessor)
	getPayload(t, a, req)
	got1, got2, got3 := a.GetAccessor(req)
	d1, _ := json.Marshal(want.Note)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, noteEvent, *got2)
	assert.Equal(t, "651234123", *got3)

	note, _ := got1.(*NoteEvent)
	assert.Equal(t, "open", *note.GetAction())
	assert.Equal(t, (*string)(nil), note.GetActionDetail())
	assert.Equal(t, "opened", *note.GetState())
	assert.Equal(t, "ibforuorg", *note.GetOrg())
	assert.Equal(t, "test1", *note.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1/merge_requests/4#note_71e9657489bcddbed4c0a9d2b1e29eb7c8ab26c3", *note.GetHtmlURL())
	assert.Equal(t, "main", *note.GetBase())
	assert.Equal(t, "ibforuorg/test1/24124124124", *note.GetHead())
	assert.Equal(t, "4", *note.GetNumber())
	assert.Equal(t, "190370", *note.GetID())
	assert.Equal(t, "****", *note.GetAuthor())
	assert.Equal(t, "71e9657489bcddbed4c0a9d2b1e29eb7c8ab26c3", *note.GetCommentID())
	assert.Equal(t, "MergeRequest", *note.GetCommentKind())
	assert.Equal(t, "/lgtm\n/approve", *note.GetComment())
	assert.Equal(t, "****", *note.GetCommenter())
	assert.Equal(t, "2024-10-26T11:44:15+08:00", *note.GetCreateTime())
	assert.Equal(t, "2024-10-26T11:44:15+08:00", *note.GetUpdateTime())
	assert.Equal(t, "public", *note.GetRepoVisibility())
}

func noteIssue(t *testing.T) {
	want := GitCodeAccessor{Note: new(NoteEvent)}
	data := testdata.ReadTestData(t, testdata.IssuesNote, want.Note)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/26", buf)
	req.Header.Set(HeaderEventType, noteEvent)
	req.Header.Set(HeaderEventGUID, "151231321")

	a := new(GitCodeAccessor)
	getPayload(t, a, req)
	got1, got2, got3 := a.GetAccessor(req)
	d1, _ := json.Marshal(want.Note)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, noteEvent, *got2)
	assert.Equal(t, "151231321", *got3)

	note, _ := got1.(*NoteEvent)
	assert.Equal(t, (*string)(nil), note.GetAction())
	assert.Equal(t, (*string)(nil), note.GetActionDetail())
	assert.Equal(t, "opened", *note.GetState())
	assert.Equal(t, "ibforuorg", *note.GetOrg())
	assert.Equal(t, "test1", *note.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1/issues/4#note_d3ab73b290d6fcd8800177e2d34545c755af3af1", *note.GetHtmlURL())
	assert.Equal(t, (*string)(nil), note.GetBase())
	assert.Equal(t, (*string)(nil), note.GetHead())
	assert.Equal(t, "4", *note.GetNumber())
	assert.Equal(t, "515443", *note.GetID())
	assert.Equal(t, "****", *note.GetAuthor())
	assert.Equal(t, "d3ab73b290d6fcd8800177e2d34545c755af3af1", *note.GetCommentID())
	assert.Equal(t, "Issue", *note.GetCommentKind())
	assert.Equal(t, "oiugbfaijub", *note.GetComment())
	assert.Equal(t, "****", *note.GetCommenter())
	assert.Equal(t, "2024-10-26T11:42:05+08:00", *note.GetCreateTime())
	assert.Equal(t, "2024-10-26T11:42:05+08:00", *note.GetUpdateTime())
	assert.Equal(t, "private", *note.GetRepoVisibility())

	note = new(NoteEvent)
	rt := reflect.TypeOf(note)
	n := rt.NumMethod()
	for i := 0; i < n; i++ {
		rm := rt.Method(i)
		ret := rm.Func.Call([]reflect.Value{reflect.ValueOf(note)})
		assert.Equal(t, (*string)(nil), ret[0].Interface())
	}
}
