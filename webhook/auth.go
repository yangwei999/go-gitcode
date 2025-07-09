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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
)

type GitCodeAuthentication struct {
	Payload   *bytes.Buffer
	SignKey   string
	eventType string
	eventGUID string
}

func (a *GitCodeAuthentication) GetEventType() string {
	return a.eventType
}

func (a *GitCodeAuthentication) GetEventGUID() string {
	return a.eventGUID
}

const (
	HeaderEventType  = "X-GitCode-Event"
	HeaderEventGUID  = "X-GitCode-Delivery"
	HeaderEventToken = "X-GitCode-Signature-256"

	// error message constants
	ErrorMessageEvent                       = "Missing X-GitCode-Event Header"
	ErrorMessageEmptyToken                  = "Missing X-GitCode-Signature-256"
	ErrorMessageInvalidToken                = "Invalid X-GitCode-Signature-256"
	ErrorMessageInvalidAuthenticationObject = "Invalid Authentication"
)

func (a *GitCodeAuthentication) Auth(w http.ResponseWriter, r *http.Request) error {

	if a.eventType = r.Header.Get(HeaderEventType); a.eventType == "" {
		http.Error(w, ErrorMessageEvent, http.StatusBadRequest)
		return errors.New(ErrorMessageEvent)
	}

	token := r.Header.Get(HeaderEventToken)
	if token == "" {
		http.Error(w, ErrorMessageEmptyToken, http.StatusUnauthorized)
		return errors.New(ErrorMessageEmptyToken)
	}

	if a.Payload == nil || strings.TrimSpace(a.SignKey) == "" {
		return errors.New(ErrorMessageInvalidAuthenticationObject)
	}

	// Validate the payload with our HMAC secret.
	if !signSuccess(token, a.SignKey, a.Payload) {
		http.Error(w, ErrorMessageInvalidToken, http.StatusUnauthorized)
		return errors.New(ErrorMessageInvalidToken)
	}

	a.eventGUID = r.Header.Get(HeaderEventGUID)

	return nil
}

func signSuccess(token, signKey string, payload *bytes.Buffer) bool {
	if !strings.HasPrefix(token, "sha256=") {
		return false
	}

	mac := hmac.New(sha256.New, []byte(signKey))
	mac.Write(payload.Bytes())

	expected := hex.EncodeToString(mac.Sum(nil))
	return expected == token[7:]
}
