# CloudMemory服务

# 概述

本文档主要介绍 CloudMemory（云端记忆）GO SDK 的使用。CloudMemory 是基于向量数据库 VectorDB 提供的长期记忆服务，可用于为大模型应用提供持久化的记忆读写、心智模型、指令、文档与实体关系等能力。SDK 在开源 [Hindsight Go Client](https://github.com/vectorize-io/hindsight)（MIT License）之上做了一层简化封装，对外暴露与 [接口速查表](https://cloud.baidu.com/doc/VDB/s/cmpl4uayz) 完全对齐的 API。

> 注意：CloudMemory 的鉴权方式不同于其它 BCE 服务，**不使用 AK/SK 签名**，而是通过 HTTP Bearer Token（API Key）进行认证。

# 初始化

## 确认 Endpoint

CloudMemory 的服务端点由用户在控制台或自建部署中获取，通常形如：

- 公有云端点：`https://cloud.memory.<region>.baidubce.com/api`
- 自建/本地部署：`http://127.0.0.1:8888`

API 支持 HTTP 和 HTTPS 两种调用方式。生产环境建议使用 HTTPS。

## 获取 API Key

通过 [百度智能云控制台](https://console.bce.baidu.com) 创建并获取 CloudMemory 的 API Key。所有请求会通过下面的请求头传递：

```
Authorization: Bearer <your-api-key>
```

## 新建 CloudMemory Client

CloudMemory Client 是与服务交互的入口。子包路径为 `services/cloudmemory/api`，包名为 `cloudmemory`。

### 使用 API Key 新建 Client

```go
import (
    "context"

    "github.com/baidubce/bce-sdk-go/services/cloudmemory/api"
)

func main() {
    endpoint, apiKey := "<your-endpoint>", "<your-api-key>"

    client := cloudmemory.New(endpoint, apiKey)
    _ = client
}
```

### 自定义超时

如果需要为底层 HTTP Client 设置每次请求的超时时间，使用 `NewWithTimeout`：

```go
import (
    "time"

    "github.com/baidubce/bce-sdk-go/services/cloudmemory/api"
)

client := cloudmemory.NewWithTimeout(endpoint, apiKey, 60*time.Second)
```

传入 `0` 表示禁用超时（不推荐）。

### 复用已有的 hindsight.APIClient

若已存在一个配置好的 `*hindsight.APIClient`（例如自定义了 Transport 或拦截器），可以直接复用：

```go
import (
    "github.com/baidubce/bce-sdk-go/services/cloudmemory/api"
    hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

raw := hindsight.NewAPIClientWithToken(endpoint, apiKey)
// raw.GetConfig().HTTPClient = customHTTPClient // 可选：自定义底层 http.Client
client := cloudmemory.NewFromAPIClient(raw)
```

### 访问底层 hindsight 客户端

封装层只暴露速查表中的端点。如需调用未被封装的高级特性（builder 选项、原始响应、流式 API 等），通过 `Underlying()` 拿到底层 `*hindsight.APIClient`：

```go
api := client.Underlying().BanksAPI.ListBanks(context.Background())
out, httpResp, err := api.Execute()
_ = httpResp // 完整的 *http.Response，含 status、header
_, _ = out, err
```

## 引用类型

请求体类型来自 hindsight 包，使用以下别名引入：

```go
import hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
```

下文中的 `hindsight.NewXxxRequest(...)` 等构造器即来自该包。

# 系统检查

## 健康检查

`GET /health`，无需认证即可访问。

```go
resp, err := client.Health(context.Background())
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 版本信息

`GET /version`，需要认证。

```go
resp, err := client.Version(context.Background())
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

# Bank 管理

Bank 是隔离记忆数据的逻辑单元，类似一个"记忆库"。绝大多数读写接口都按 Bank 维度进行隔离。

## 列出 Bank

`GET /banks`

```go
resp, err := client.ListBanks(context.Background())
if err != nil {
    panic(err)
}
fmt.Println(resp.Items)
```

## 创建 / 更新 Bank

`POST /banks`，按 `bankID` 幂等创建。

```go
import hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"

req := hindsight.NewCreateBankRequest()
// 可选：req.SetName("user-bank") 等

resp, err := client.CreateBank(context.Background(), "your-bank-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 获取 Bank 详情

`GET /banks/{bankId}`

```go
resp, err := client.GetBank(context.Background(), "your-bank-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 删除 Bank

`DELETE /banks/{bankId}`，会**同时删除 Bank 下所有数据**，请谨慎使用。

```go
resp, err := client.DeleteBank(context.Background(), "your-bank-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 获取 / 更新 Bank 配置

`GET /banks/{bankId}/config` 与 `PATCH /banks/{bankId}/config`。`UpdateBankConfig` 接收一个 `BankConfigUpdate`，其中 `Updates` 字段是 `map[string]interface{}`，**必须为非空字典**（即使没有要更新的键也要传 `map[string]interface{}{}`）。

```go
// 获取
cfg, err := client.GetBankConfig(context.Background(), "your-bank-id")
if err != nil {
    panic(err)
}
fmt.Println(cfg)

// 更新
update := *hindsight.NewBankConfigUpdate(map[string]interface{}{
    "auto_consolidate": true,
})
out, err := client.UpdateBankConfig(context.Background(), "your-bank-id", update)
if err != nil {
    panic(err)
}
fmt.Println(out)
```

## 获取统计信息

`GET /banks/{bankId}/stats`

```go
stats, err := client.GetBankStats(context.Background(), "your-bank-id")
if err != nil {
    panic(err)
}
fmt.Println(stats)
```

## 触发记忆整合

`POST /banks/{bankId}/consolidate`。`req` 可传 `nil` 使用服务端默认参数。

```go
resp, err := client.ConsolidateBank(context.Background(), "your-bank-id", nil)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

# 记忆读写

## 同步写入记忆

`POST /memories/retain`，请求中 `items` 至少包含一项。

```go
items := []hindsight.MemoryItem{
    *hindsight.NewMemoryItem("用户喜欢简洁的回答"),
    *hindsight.NewMemoryItem("用户使用 Go 1.21"),
}
req := hindsight.NewRetainRequest(items)

resp, err := client.Retain(context.Background(), "your-bank-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 异步写入记忆

`POST /memories/retain_async`，立即返回，记忆在后台处理。

```go
resp, err := client.RetainAsync(context.Background(), "your-bank-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp.OperationId) // 可用于后续查询
```

## 上传文件并写入记忆

`POST /files/retain`，支持批量上传文件并把内容写入记忆。`requestJSON` 是必填的多部分表单字段（即便仅写入文件也需要传，至少为 `{}` 或 `{"items":[...]}`）。

```go
import "os"

f, err := os.Open("/path/to/notes.txt")
if err != nil {
    panic(err)
}
defer f.Close()

requestJSON := `{"items":[{"content":"file upload"}]}`
resp, err := client.FilesRetain(context.Background(), "your-bank-id", []*os.File{f}, requestJSON)
if err != nil {
    panic(err)
}
fmt.Println(resp.OperationIds)
```

## 检索相关记忆

`POST /recall`

```go
req := hindsight.NewRecallRequest("用户的编程语言偏好是什么？")
resp, err := client.Recall(context.Background(), "your-bank-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 基于记忆生成综合回答

`POST /reflect`，由服务端调用 LLM 综合现有记忆给出答案。**该接口耗时较长**（一般在数十秒级别），建议把 Client 的超时和上下文超时都设置为 2 分钟以上。

```go
req := hindsight.NewReflectRequest("总结一下你对用户的了解")
resp, err := client.Reflect(context.Background(), "your-bank-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 列出记忆

`GET /list`，支持类型、关键词、整合状态、分页等过滤。

```go
opts := cloudmemory.ListMemoriesOptions{
    Type:               "",  // "raw" / "consolidated" 等
    Q:                  "",  // 关键词
    ConsolidationState: "",  // "pending" / "done" 等
    Limit:              20,
    Offset:             0,
}
resp, err := client.ListMemories(context.Background(), "your-bank-id", opts)
if err != nil {
    panic(err)
}
fmt.Println(resp.Items)
```

# 文档管理

> `ListDocuments` 返回的 `Items` 字段是 `[]map[string]interface{}`（弱类型），文档 ID 通常以 `id` 键访问，需自行做 `string` 类型断言。

## 列出文档

`GET /documents`

```go
opts := cloudmemory.ListDocumentsOptions{
    Q:         "",
    Tags:      []string{"work"},
    TagsMatch: "any",  // any / all
    Limit:     10,
    Offset:    0,
}
resp, err := client.ListDocuments(context.Background(), "your-bank-id", opts)
if err != nil {
    panic(err)
}
for _, item := range resp.Items {
    fmt.Println(item["id"], item["title"])
}
```

## 获取文档详情

`GET /documents/{documentId}`

```go
resp, err := client.GetDocument(context.Background(), "your-bank-id", "your-document-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 更新文档标签

`PATCH /documents/{documentId}`。请求体**至少包含一个待更新字段**，否则服务端会返回 422。

```go
req := hindsight.NewUpdateDocumentRequest()
req.SetTags([]string{"important", "reviewed"})

resp, err := client.UpdateDocument(context.Background(), "your-bank-id", "your-document-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 删除文档

`DELETE /documents/{documentId}`，会一并删除该文档关联的记忆。

```go
resp, err := client.DeleteDocument(context.Background(), "your-bank-id", "your-document-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 列出文档分块

`GET /documents/{documentId}/chunks`

```go
resp, err := client.ListDocumentChunks(context.Background(), "your-bank-id", "your-document-id", 20, 0)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

# 实体关系

## 列出实体

`GET /entities`

```go
resp, err := client.ListEntities(context.Background(), "your-bank-id", 50, 0)
if err != nil {
    panic(err)
}
fmt.Println(resp.Items)
```

## 获取实体关系图

`GET /entities/graph`，可指定返回的实体数量上限和最小出现次数。

```go
resp, err := client.EntityGraph(context.Background(), "your-bank-id", 100, 2)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

# 后台操作

异步任务（`RetainAsync`、`RefreshMentalModel`、`FilesRetain` 等）会返回 `operationId`，可在后台操作接口中查询和取消。

## 列出后台操作

`GET /operations/{bankId}`

```go
opts := cloudmemory.ListOperationsOptions{
    Status:         "",   // queued / running / done / failed
    Type:           "",
    Limit:          50,
    Offset:         0,
    ExcludeParents: false,
}
resp, err := client.ListOperations(context.Background(), "your-bank-id", opts)
if err != nil {
    panic(err)
}
for _, op := range resp.Operations {
    fmt.Println(op.Id, op.Status)
}
```

## 取消后台操作

按 `operationId` 取消单个操作。

```go
resp, err := client.CancelOperation(context.Background(), "your-bank-id", "your-operation-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

# 心智模型

心智模型（Mental Model）是一种基于记忆的可命名、可刷新的"认知摘要"，对应 `/banks/{bankId}/mental-models` 系列接口。

## 列出心智模型

`GET /banks/{bankId}/mental-models`

```go
opts := cloudmemory.ListMentalModelsOptions{
    Tags:      nil,
    TagsMatch: "any",
    Detail:    "",  // 详情级别，由服务端约定
    Limit:     20,
    Offset:    0,
}
resp, err := client.ListMentalModels(context.Background(), "your-bank-id", opts)
if err != nil {
    panic(err)
}
fmt.Println(resp.Items)
```

## 创建心智模型

`POST /banks/{bankId}/mental-models`，需要传入 `name` 和 `sourceQuery`。

```go
req := hindsight.NewCreateMentalModelRequest("user-profile", "what do you remember about the user?")
resp, err := client.CreateMentalModel(context.Background(), "your-bank-id", *req)
if err != nil {
    panic(err)
}
modelID := resp.GetMentalModelId()  // 注意：返回值字段是 NullableString，需用 GetMentalModelId
fmt.Println(modelID, resp.OperationId)
```

## 获取心智模型详情

`GET /banks/{bankId}/mental-models/{modelId}`

```go
resp, err := client.GetMentalModel(context.Background(), "your-bank-id", "your-model-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 更新心智模型

`PATCH /banks/{bankId}/mental-models/{modelId}`。请求体**至少包含一个待更新字段**，否则服务端返回 404/422。

```go
req := hindsight.NewUpdateMentalModelRequest()
req.SetSourceQuery("what does the user prefer?")

resp, err := client.UpdateMentalModel(context.Background(), "your-bank-id", "your-model-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 删除心智模型

`DELETE /banks/{bankId}/mental-models/{modelId}`

```go
resp, err := client.DeleteMentalModel(context.Background(), "your-bank-id", "your-model-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 刷新心智模型

`POST /banks/{bankId}/mental-models/{modelId}/refresh`，异步触发服务端重新生成模型内容，返回 `operationId`。

```go
resp, err := client.RefreshMentalModel(context.Background(), "your-bank-id", "your-model-id")
if err != nil {
    panic(err)
}
fmt.Println(resp.OperationId, resp.Status)
```

# 指令

指令（Directive）用于注入到下游 LLM 的 system prompt 中，对应 `/banks/{bankId}/directives` 系列接口。

## 列出指令

`GET /banks/{bankId}/directives`

```go
opts := cloudmemory.ListDirectivesOptions{
    Tags:       nil,
    TagsMatch:  "any",
    ActiveOnly: true,
    Limit:      20,
    Offset:     0,
}
resp, err := client.ListDirectives(context.Background(), "your-bank-id", opts)
if err != nil {
    panic(err)
}
fmt.Println(resp.Items)
```

## 创建指令

`POST /banks/{bankId}/directives`，必填 `name` 与 `content`。

```go
req := hindsight.NewCreateDirectiveRequest("be-concise", "Always answer concisely and avoid filler.")
resp, err := client.CreateDirective(context.Background(), "your-bank-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 获取指令详情

`GET /banks/{bankId}/directives/{directiveId}`

```go
resp, err := client.GetDirective(context.Background(), "your-bank-id", "your-directive-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 更新指令

`PATCH /banks/{bankId}/directives/{directiveId}`

```go
req := hindsight.NewUpdateDirectiveRequest()
req.SetContent("Be concise but complete.")

resp, err := client.UpdateDirective(context.Background(), "your-bank-id", "your-directive-id", *req)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

## 删除指令

`DELETE /banks/{bankId}/directives/{directiveId}`

```go
resp, err := client.DeleteDirective(context.Background(), "your-bank-id", "your-directive-id")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

# 接口与速查表对照

下表给出 SDK 方法与 [接口速查表](https://cloud.baidu.com/doc/VDB/s/cmpl4uayz) 中 REST 端点的一一对应关系（共 35 个端点）：

| 分组 | HTTP 方法 | 路径 | SDK 方法 |
| --- | --- | --- | --- |
| 系统检查 | GET | `/health` | `Health` |
|  | GET | `/version` | `Version` |
| Bank 管理 | GET | `/banks` | `ListBanks` |
|  | POST | `/banks` | `CreateBank` |
|  | GET | `/banks/{bankId}` | `GetBank` |
|  | DELETE | `/banks/{bankId}` | `DeleteBank` |
|  | GET | `/banks/{bankId}/config` | `GetBankConfig` |
|  | PATCH | `/banks/{bankId}/config` | `UpdateBankConfig` |
|  | GET | `/banks/{bankId}/stats` | `GetBankStats` |
|  | POST | `/banks/{bankId}/consolidate` | `ConsolidateBank` |
| 记忆读写 | POST | `/memories/retain` | `Retain` |
|  | POST | `/memories/retain_async` | `RetainAsync` |
|  | POST | `/files/retain` | `FilesRetain` |
|  | POST | `/recall` | `Recall` |
|  | POST | `/reflect` | `Reflect` |
|  | GET | `/list` | `ListMemories` |
| 文档管理 | GET | `/documents` | `ListDocuments` |
|  | GET | `/documents/{documentId}` | `GetDocument` |
|  | PATCH | `/documents/{documentId}` | `UpdateDocument` |
|  | DELETE | `/documents/{documentId}` | `DeleteDocument` |
|  | GET | `/documents/{documentId}/chunks` | `ListDocumentChunks` |
| 实体关系 | GET | `/entities` | `ListEntities` |
|  | GET | `/entities/graph` | `EntityGraph` |
| 后台操作 | GET | `/operations/{bankId}` | `ListOperations` |
|  | DELETE | `/operations/{bankId}` | `CancelOperation` |
| 心智模型 | GET | `/banks/{bankId}/mental-models` | `ListMentalModels` |
|  | POST | `/banks/{bankId}/mental-models` | `CreateMentalModel` |
|  | GET | `/banks/{bankId}/mental-models/{modelId}` | `GetMentalModel` |
|  | PATCH | `/banks/{bankId}/mental-models/{modelId}` | `UpdateMentalModel` |
|  | DELETE | `/banks/{bankId}/mental-models/{modelId}` | `DeleteMentalModel` |
|  | POST | `/banks/{bankId}/mental-models/{modelId}/refresh` | `RefreshMentalModel` |
| 指令 | GET | `/banks/{bankId}/directives` | `ListDirectives` |
|  | POST | `/banks/{bankId}/directives` | `CreateDirective` |
|  | GET | `/banks/{bankId}/directives/{directiveId}` | `GetDirective` |
|  | PATCH | `/banks/{bankId}/directives/{directiveId}` | `UpdateDirective` |
|  | DELETE | `/banks/{bankId}/directives/{directiveId}` | `DeleteDirective` |

# 错误处理

所有 SDK 方法的最后一个返回值都是 `error`。当服务端返回非 2xx 状态码时，错误信息中会带有上游响应体。常见错误：

- `401 Unauthorized` — API Key 无效或缺失。
- `404 Not Found` — Bank / 文档 / 模型 / 指令 ID 不存在；或 PATCH 请求体为空导致路由不匹配。
- `422 Unprocessable Entity` — 请求体校验失败，例如 `BankConfigUpdate.Updates` 为 `nil`、或 `UpdateDocumentRequest` 没有任何字段。
- `context deadline exceeded` — 请求超时。`Reflect`、`RefreshMentalModel` 等接口由服务端调用 LLM，建议把 client 与 ctx 超时都放宽到 2 分钟以上。

# 开源协议

CloudMemory Go SDK 在开源项目 [Hindsight](https://github.com/vectorize-io/hindsight)（MIT License, Copyright 2025 Vectorize AI, Inc.）基础上做封装。原始 LICENSE 文件保留在 `services/cloudmemory/api/LICENSE` 中，每个封装源文件顶部均带有 `Based on Hindsight Go Client - MIT License` 声明。
