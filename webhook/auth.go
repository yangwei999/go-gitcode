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
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	errorNilToken    = errors.New("token should be non-nil/non-empty")
	errorNilResponse = errors.New("http response should be non-nil")
	errorNilRequest  = errors.New("http request should be non-nil")
)

type GitCodeAuthentication struct {
	payload   *bytes.Buffer
	eventType string
	eventGUID string
	signKey   string
}

func (a *GitCodeAuthentication) SetSignKey(token []byte) error {
	if len(token) == 0 {
		return errorNilToken
	}
	a.signKey = string(token)
	return nil
}

func (a *GitCodeAuthentication) GetPayload() *bytes.Buffer {
	return a.payload
}

func (a *GitCodeAuthentication) GetEventType() string {
	return a.eventType
}

func (a *GitCodeAuthentication) GetEventGUID() string {
	return a.eventGUID
}

const (
	headerEventType      = "X-GitCode-Event"
	headerEventGUID      = "X-GitCode-Delivery"
	headerEventToken     = "X-GitCode-Signature-256"
	headerUserAgent      = "User-Agent"
	headerUserAgentValue = "git-gitcode-hook"

	headerContentTypeName      = "Content-Type"
	headerContentTypeJsonValue = "application/json"

	httpStatusCodeIncorrectErrorFormat = "http response status code can not be setting to %d"

	// error message constants
	bodyReadErrorMessage           = "400 Bad Request: Failed to read request body"
	headerContentTypeErrorMessage  = "400 Bad Request: Hook only accepts content-type: application/json"
	headerEventErrorMessage        = "400 Bad Request: Missing X-GitCode-Event Header"
	headerUserAgentErrorMessage    = "400 Bad Request: Invalid User-Agent Header"
	headerEmptyTokenErrorMessage   = "401 Unauthorized: Missing X-GitCode-Token"
	headerInvalidTokenErrorMessage = "403 Forbidden: Invalid X-GitCode-Token"
)

func (a *GitCodeAuthentication) Auth(w http.ResponseWriter, r *http.Request) (error, bool) {
	if r == nil {
		return errorNilRequest, false
	}

	if r.Header.Get(headerUserAgent) != headerUserAgentValue {
		return errors.New(headerUserAgentErrorMessage), true
	}

	var err error
	if a.payload, err = ReadPayload(w, r); err != nil {
		return err, false
	}

	// Header checks: It must be a POST with an event type and a signature.
	if r.Method != http.MethodPost {
		return handleErr(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)), false
	}

	if v := r.Header.Get(headerContentTypeName); !strings.HasPrefix(v, headerContentTypeJsonValue) {
		return handleErr(w, http.StatusBadRequest, headerContentTypeErrorMessage), false
	}

	if a.eventType = r.Header.Get(headerEventType); a.eventType == "" {
		return handleErr(w, http.StatusBadRequest, headerEventErrorMessage), false
	}

	token := r.Header.Get(headerEventToken)
	if token == "" {
		return handleErr(w, http.StatusUnauthorized, headerEmptyTokenErrorMessage), false
	}

	// Validate the payload with our HMAC secret.
	if !signSuccess(token, a.signKey, a.payload) {
		return handleErr(w, http.StatusUnauthorized, headerInvalidTokenErrorMessage), false
	}

	a.eventGUID = r.Header.Get(headerEventGUID)

	return nil, false
}

func ReadPayload(w http.ResponseWriter, r *http.Request) (*bytes.Buffer, error) {
	if r.Body == nil {
		return nil, nil
	}

	defer func() {
		_ = r.Body.Close()
	}()
	var payload bytes.Buffer
	if r.Body != http.NoBody {
		if _, err := io.Copy(&payload, r.Body); err != nil {
			http.Error(w, bodyReadErrorMessage, http.StatusBadRequest)
			return nil, err
		}
	}
	return &payload, nil
}

func handleErr(w http.ResponseWriter, errCode int, errMsg string) error {
	if errCode < http.StatusBadRequest ||
		(errCode > http.StatusUnavailableForLegalReasons && errCode < http.StatusInternalServerError) ||
		errCode > http.StatusNetworkAuthenticationRequired {
		return fmt.Errorf(httpStatusCodeIncorrectErrorFormat, errCode)
	}
	if w == nil {
		return errorNilResponse
	}

	http.Error(w, errMsg, errCode)
	return errors.New(errMsg)
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
