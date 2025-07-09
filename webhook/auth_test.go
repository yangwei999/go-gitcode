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
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	payloadData  = "{\n  \"note\": \"/ibforuorg/community-test/pulls/2#note_30974945\" \n }"
	dummySignKey = "1234"
)

func TestGitCodeAuthenticationAuth(t *testing.T) {

	type args struct {
		r   *GitCodeAuthentication
		w   http.ResponseWriter
		req *http.Request
	}

	testCases := []struct {
		no  string
		in  args
		out string
		fn  func(i *args)
	}{
		{
			"case2",
			args{
				&GitCodeAuthentication{},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case2", nil)
					return req
				}(),
			},
			ErrorMessageEvent,
			nil,
		},
		{
			"case3",
			args{
				&GitCodeAuthentication{},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case3", nil)
					req.Header.Set(HeaderEventType, noteEvent)
					return req
				}(),
			},
			ErrorMessageEmptyToken,
			nil,
		},
		{
			"case4",
			args{
				&GitCodeAuthentication{SignKey: dummySignKey},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case4", func() io.Reader {
						b := &bytes.Buffer{}
						_, err := b.ReadFrom(bytes.NewBufferString(payloadData))
						assert.NoError(t, err)
						return b
					}())
					req.Header.Set(HeaderEventType, noteEvent)
					req.Header.Set(HeaderEventToken, "123451")
					return req
				}(),
			},
			ErrorMessageInvalidToken,
			nil,
		},
		{
			"case5",
			args{
				&GitCodeAuthentication{SignKey: dummySignKey},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case5", func() io.Reader {
						b := &bytes.Buffer{}
						return b
					}())
					req.Header.Set(HeaderEventType, noteEvent)
					req.Header.Set(HeaderEventToken, "sha256=36acf017ea0974457577506ef75268ac93ed6d61864ee994f438b63916ed1736")
					return req
				}(),
			},
			ErrorMessageInvalidAuthenticationObject,
			nil,
		},
		{
			"case6",
			args{
				&GitCodeAuthentication{SignKey: dummySignKey},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case6", func() io.Reader {
						var b io.Reader
						buf := &bytes.Buffer{}
						buf.Write([]byte(payloadData))
						b = buf
						return b
					}())
					req.Header.Set(HeaderEventType, noteEvent)
					req.Header.Set(HeaderEventToken, "sha256=f585860d0ca237e0550da0e166370b9c372e8aeb2e639b0ac9884cd52681c576")
					return req
				}(),
			},
			"",
			func(i *args) {
				assert.Equal(t, noteEvent, i.r.eventType)
				if i.r.Payload == nil {
					t.Error("payload should be non-nil")
				}
				assert.Equal(t, i.r.Payload.String(), payloadData)
				assert.Equal(t, "1234", i.r.SignKey)
			},
		},
		{
			"case7",
			args{
				&GitCodeAuthentication{SignKey: dummySignKey},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case7", func() io.Reader {
						var b io.Reader
						buf := &bytes.Buffer{}
						buf.Write([]byte(payloadData))
						b = buf
						return b
					}())
					req.Header.Set(HeaderEventType, noteEvent)
					req.Header.Set(HeaderEventToken, "sha256=36acf017ea0974457577506ef75268ac93ed6d61864ee994f438b63916ed1736")

					_, _ = io.Copy(io.Discard, req.Body)
					return req
				}(),
			},
			"",
			func(i *args) {
				assert.Equal(t, "", i.r.Payload.String())
			},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].no, func(t *testing.T) {
			if testCases[i].in.req.Body != nil && testCases[i].in.req.Body != http.NoBody {
				var payload bytes.Buffer
				_, err := io.Copy(&payload, testCases[i].in.req.Body)
				assert.Nil(t, err)
				testCases[i].in.r.Payload = &payload
			}
			got := testCases[i].in.r.Auth(testCases[i].in.w, testCases[i].in.req)
			if got != nil {
				assert.Equal(t, testCases[i].out, got.Error())
			}
			if testCases[i].fn != nil {
				testCases[i].fn(&testCases[i].in)
			}
		})
	}
}

func TestGitCodeAuthenticationGetEventType(t *testing.T) {

	type args struct {
		r GitCodeAuthentication
	}

	testCases := []struct {
		no  string
		in  args
		out string
		fn  func(i *args)
	}{
		{
			"case1",
			args{
				GitCodeAuthentication{},
			},
			"",
			nil,
		},
		{
			"case2",
			args{
				GitCodeAuthentication{
					eventType: "",
				},
			},
			"",
			nil,
		},
		{
			"case3",
			args{
				GitCodeAuthentication{
					eventType: "1234hfas",
				},
			},
			"1234hfas",
			nil,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].no, func(t *testing.T) {
			got := testCases[i].in.r.GetEventType()
			assert.Equal(t, testCases[i].out, got)
			if testCases[i].fn != nil {
				testCases[i].fn(&testCases[i].in)
			}
		})
	}
}

func TestGitCodeAuthenticationGetEventGUID(t *testing.T) {

	type args struct {
		r GitCodeAuthentication
	}

	testCases := []struct {
		no  string
		in  args
		out string
		fn  func(i *args)
	}{
		{
			"case1",
			args{
				GitCodeAuthentication{},
			},
			"",
			nil,
		},
		{
			"case2",
			args{
				GitCodeAuthentication{
					eventGUID: "",
				},
			},
			"",
			nil,
		},
		{
			"case3",
			args{
				GitCodeAuthentication{
					eventGUID: "1234hfas",
				},
			},
			"1234hfas",
			nil,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].no, func(t *testing.T) {
			got := testCases[i].in.r.GetEventGUID()
			assert.Equal(t, testCases[i].out, got)
			if testCases[i].fn != nil {
				testCases[i].fn(&testCases[i].in)
			}
		})
	}
}

func TestSignSuccess(t *testing.T) {
	type args struct {
		token   string
		signKey string
		payload *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with valid token",
			args: args{
				token:   "sha256=3938a65bf0a111e17a7dfe928ea0c73a38c5af80006a939a81c46b372e8f8815",
				signKey: "secret",
				payload: bytes.NewBufferString("{\n    \"content\": \"MTI0MTQxMjQxMjQ=\",\n    \"message\": \"fas\",\n    \"branch\": \"test1-patch-1\"\n}"),
			},
			want: true,
		},
		{
			name: "Test with invalid token prefix",
			args: args{
				token:   "md5=123456",
				signKey: "secret",
				payload: bytes.NewBufferString("test payload"),
			},
			want: false,
		},
		{
			name: "Test with empty token",
			args: args{
				token:   "",
				signKey: "secret",
				payload: bytes.NewBufferString("test payload"),
			},
			want: false,
		},
		{
			name: "Test with empty payload",
			args: args{
				token:   "sha256=123456",
				signKey: "secret",
				payload: bytes.NewBufferString(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := signSuccess(tt.args.token, tt.args.signKey, tt.args.payload); got != tt.want {
				t.Errorf("signSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}
