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
	"net/http"
)

type GitCodeAccessor struct {
	Issues  *IssueEvent
	PR      *PullRequestEvent
	Note    *NoteEvent
	Push    *PushEvent
	Payload *bytes.Buffer
}

const (
	pushEvent        = "Push Hook"
	issueEvent       = "Issue Hook"
	pullRequestEvent = "Merge Request Hook"
	noteEvent        = "Note Hook"
)

func (a *GitCodeAccessor) GetAccessor(r *http.Request) (any, *string, *string) {

	if a.Payload == nil {
		return nil, nil, nil
	}

	eventGUID := r.Header.Get(HeaderEventGUID)
	eventType := r.Header.Get(HeaderEventType)

	switch eventType {
	case issueEvent:
		a.Issues = new(IssueEvent)
		_ = json.Unmarshal(a.Payload.Bytes(), a.Issues)
		return a.Issues, &eventType, &eventGUID
	case pullRequestEvent:
		a.PR = new(PullRequestEvent)
		_ = json.Unmarshal(a.Payload.Bytes(), a.PR)
		return a.PR, &eventType, &eventGUID
	case noteEvent:
		a.Note = new(NoteEvent)
		_ = json.Unmarshal(a.Payload.Bytes(), a.Note)
		return a.Note, &eventType, &eventGUID
	case pushEvent:
		a.Push = new(PushEvent)
		_ = json.Unmarshal(a.Payload.Bytes(), a.Push)
		return a.Push, &eventType, &eventGUID
	default:
		// do nothing
	}

	return nil, &eventType, &eventGUID
}
