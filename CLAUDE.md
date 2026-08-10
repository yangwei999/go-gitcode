# go-gitcode

GitCode OpenAPI 的 Go SDK（文档：https://docs.gitcode.com/docs/openapi/guide/）。模块名 `github.com/opensourceways/go-gitcode`，Go 1.26.5。

## 目录结构

- `openapi/` — API 客户端及各资源模块（issues、pull_requests、repository、user、org、dashboard）。每个资源通常有三个文件：`<resource>.go`（方法实现）、`<resource>_model.go`（请求/响应结构体）、`<resource>_test.go`。
- `webhook/` — 处理入站 webhook：`accessor.go` 根据 `X-Gitcode-Event`/GUID 等 header 分发事件，另有 issues/PR/note/push 各类事件类型定义，以及 `auth.go` 做签名校验。
- `testdata/` — 测试用的 JSON fixture，分布在 `repos/`、`pr/`、`issues/`、`user/`、`webhook/` 目录下，通过 `testdata/test_data.go` 中的类型化路径常量引用（如 `testdata.IssuesUpdate`）。使用 `testdata.ReadTestData(t, testdata.X, &target)` 加载。

## 客户端架构（openapi/service.go、client.go）

- `APIClient` 持有一个 `service{api *APIClient}`，通过类型转换（`(*IssuesService)(&c.common)`）复用给所有 `*Service` 类型，避免每个 service 单独分配堆内存。新增资源分组时遵循同样写法。
- `NewAPIClientWithAuthorization(token []byte)` 构造客户端，内部使用自定义 `roundTripperFunc` 注入 `Authorization: Bearer`、`User-Agent`、`Accept` 等 header。
- 请求构建统一走 `newRequest(api, method, urlStr, body, handlers...)`。`RequestHandler.t` 决定编码方式：`Json`（默认）、`Form`、或 `Query`（此时 body 传 `*url.Values`）。
- `APIClient.Do` 在 `resp.StatusCode > http.StatusPreconditionRequired` 时最多重试 3 次，每次间隔 1 秒，最终将 JSON 解码进 `receiver`（必须是指针或 nil）。
- 成功判定辅助函数：`successGetData`（200）、`successCreated`（200/201）、`successModified`（200/202/204）——每个 service 方法都返回 `(*T, bool, error)`，其中 bool 来自这几个函数之一。

## 新增接口时的约定

- 每个 API 调用对应一个方法，方法上写文档注释并链接到 GitCode API 文档页，例如 `// api Docs: https://docs.gitcode.com/docs/apis/...`。
- 方法签名模式：`func (s *XService) Verb(ctx context.Context, ...ids, body *XRequest) (*X, bool, error)`。列表接口返回 `[]*X`；无响应体的接口返回 `(bool, error)`。
- 用 `fmt.Sprintf("repos/%s/...", owner, ...)` 拼 URL（不要加前导斜杠——会基于已经以 `/api/v5/` 结尾的 `baseURL` 解析）。
- 分页 GET 接口用 `pagedQuery(page, perPage)` / `pagedLimitQuery(page)` / `buildQuery` 构造查询参数，并传入 `RequestHandler{t: Query}`。
- `newRequest` 失败时按现有方法的写法立即返回 `(nil, false, err)`（或对应的列表/bool 形式）。

## 测试约定

- 测试文件位于 `openapi/*_test.go`，使用 `openapi/service_test.go` 中的 `mockServer(t)`（启动一个挂载在 `/api/v5/` 下的 `httptest.NewServer`，返回 `client, mux, serverURL`）和 `mockResponse(t, mux, urlStr, body)`。共享的 fixture 常量：`owner="111"`、`repo="222"`、`number="1"`、`page="1"`、`branch="master"`、`pusher`、`merger`。
- 期望的响应体优先从 `testdata/` 下的 JSON fixture 加载，而不是内联构造结构体，例如 `testdata.ReadTestData(t, testdata.IssuesUpdate, &want)`。
- `openapi/common_test.go`（`TestNewRequestError`）使用 `github.com/agiledragon/gomonkey/v2` mock `newRequest`，断言所有 service 方法都能正确传播错误，统一用 `compare(t, ok, err, mockErr, result)` 辅助函数校验。新增方法时记得在这里补一条对应调用，覆盖错误路径。
- 在仓库根目录用 `go test ./...` 运行测试；目前仓库里没有 Makefile 或 CI 配置文件。

## License 头

每个 `.go` 文件开头都要有 Apache 2.0 的版权注释（`Copyright (c) Huawei Technologies Co., Ltd. <year>...`）。新建文件时严格按此格式，年份用当前年份。
