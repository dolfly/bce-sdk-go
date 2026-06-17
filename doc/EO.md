EO服务

# 概述

本文档主要介绍EO GO SDK的使用。在使用本文档前，您需要先了解EO的一些基本知识，并已开通了EO服务。若您还不了解CDN，可以参考[产品介绍](https://cloud.baidu.com/doc/GEO/s/lmj18vwxu)和[快速入门](https://cloud.baidu.com/doc/GEO/s/ymocqbtaz)。

# 初始化

## 确认Endpoint

目前使用EO服务时，EO的 Endpoint 统一使用`https://geo.baidubce.com`，这也是默认值。

## 获取密钥

要使用百度云EO，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问EO做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 使用AK/SK新建EO Client

通过AK/SK方式访问EO，用户可以参考如下代码新建一个EO Client：

```go
ak := "your_access_key_id"
sk := "your_secret_key_id"
endpoint := "geo.baidubce.com"

cli, err := eo.NewClient(ak, sk, endpoint)
```

在上面代码中，变量`ak`对应控制台中的“Access Key ID”，变量`sk`对应控制台中的“Access Key Secret”，获取方式请参考《 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。变量`endpoint`必须为`https://geo.baidubce.com`，也是默认值，为空表示使用默认值，设置为其他则SDK无法工作。

在下面的示例中，会频繁使用到GetDefaultClient函数，它的定义为：

```go
func GetDefaultClient() *eo.Client {
	ak := "your_access_key_id"
	sk := "your_secret_key_id"
	endpoint := "https://geo.baidubce.com"

	// ignore error in test, but you should handle error in dev
	client, _ := eo.NewClient(ak, sk, endpoint)
	return client
}
```

## 缓存管理接口

### 刷新缓存/查询刷新状态 Purge/GetPurgedStatus

> 缓存清除方式有URL刷新、目录刷新。URL刷新是以文件或一个资源为单位进行缓存刷新。目录刷新是以目录为单位，将目录下的所有文件进行缓存清除。提交刷新任务时需要指定站点（site）。

```go
// 刷除
cli := GetDefaultClient()
purgedId, err := cli.Purge("your_site.com", []api.PurgeTask{
	{
		Url:  "http://your_site.com/path/to/purge/1.data",
		Type: "file",
	},
	{
		Url:  "http://your_site.com/path/to/purge/html/",
		Type: "directory",
	},
})
fmt.Printf("purgedId:%+v\n", purgedId)
fmt.Printf("err:%+v\n", err)

// 方式一：根据站点和任务ID查询刷除状态
purgedStatus, err := cli.GetPurgedStatus(&api.PurgeStatusQueryData{
	Site: "your_site.com",
	Id:   string(purgedId),
})
fmt.Printf("purgedStatus:%+v\n", purgedStatus)
fmt.Printf("err:%+v\n", err)

// 方式二：根据站点和时间范围查询刷除状态（可选按刷新类型过滤）
purgedStatus, err := cli.GetPurgedStatus(&api.PurgeStatusQueryData{
	Site:      "your_site.com",
	StartTime: "2026-05-01T00:00:00Z",
	EndTime:   "2026-05-31T23:59:59Z",
	Type:      "file", // 可选，按刷新类型过滤，可选值为 "file" 或 "directory"
})
fmt.Printf("purgedStatus:%+v\n", purgedStatus)
fmt.Printf("err:%+v\n", err)
```

接口更多细节可以参考缓存管理文档：https://cloud.baidu.com/doc/GEO/s/4mhsrv9ry  、 https://cloud.baidu.com/doc/GEO/s/mmhssw91q

### 预热资源/查询预热状态 Prefetch/GetPrefetchStatus

> URL预热是以文件为单位进行资源预热。

```go
// 预热
cli := GetDefaultClient()    
prefetchId, err := cli.Prefetch("your_site.com", []api.PrefetchTask{
	{
		Url: "http://your_site.com/path/to/prefetch/1.data",
	},
	{
		Url: "http://your_site.com/path/to/prefetch/2.data",
	},
})
fmt.Printf("prefetchId:%+v\n", prefetchId)
fmt.Printf("err:%+v\n", err)

// 方式一：根据站点和任务ID查询预热状态
prefetchStatus, err := cli.GetPrefetchStatus(&api.PrefetchStatusQueryData{
	Site: "your_site.com",
	Id:   string(prefetchId),
})
fmt.Printf("prefetchStatus:%+v\n", prefetchStatus)
fmt.Printf("err:%+v\n", err)

// 方式二：根据站点和时间范围查询预热状态
prefetchStatus, err := cli.GetPrefetchStatus(&api.PrefetchStatusQueryData{
	Site:      "your_site.com",
	StartTime: "2026-05-01T00:00:00Z",
	EndTime:   "2026-05-31T23:59:59Z",
})
fmt.Printf("prefetchStatus:%+v\n", prefetchStatus)
fmt.Printf("err:%+v\n", err)
```

接口更多细节可以参考缓存管理文档：https://cloud.baidu.com/doc/GEO/s/5mhsuituv 、 https://cloud.baidu.com/doc/GEO/s/Bmhsv5i9u

## 离线日志接口

### 获取离线日志下载地址 GetOfflineLog

> 获取用户某个站点下单个域名或多个域名某一指定时间段内的日志下载地址。日志的保存时间为 180 天。

```go
cli := GetDefaultClient()
logResult, err := cli.GetOfflineLog(&api.LogQueryData{
	Site:       "your_site.com",
	StartTime:  "2026-05-01T00:00:00Z",
	EndTime:    "2026-05-31T23:59:59Z",
	DomainList: []string{"your.domain.com", "your.domain.com"},
	PageNo:     1,
	PageSize:   20,
})
fmt.Printf("logResult:%+v\n", logResult)
fmt.Printf("err:%+v\n", err)
```

`api.LogQueryData`字段说明：

| 字段       | 类型     | 是否必选 | 说明                                                                  |
| ---------- | -------- | -------- | --------------------------------------------------------------------- |
| Site       | String   | 是       | 站点名称。                                                            |
| StartTime  | String   | 否       | 查询时间范围起始值，UTC 时间。默认为 `EndTime` 前推 8 小时。 |
| EndTime    | String   | 否       | 查询时间范围结束值，UTC 时间。默认为当前时间。          |
| DomainList | []String | 否       | 查询的域名列表。                                                      |
| PageNo     | int      | 否       | 分页编号，默认值为 1。                                                |
| PageSize   | int      | 否       | 每页返回日志数目，默认值为 20。                                       |

`logResult`是`*api.LogQueryResult`类型的对象，详细说明如下：

| 字段          | 类型        | 说明                 |
| ------------- | ----------- | -------------------- |
| LogEntryList  | []LogEntry  | 离线日志列表。       |
| TotalCount    | String      | 离线日志列表总数。   |

`LogEntry`类型说明：

| 字段          | 类型   | 说明                                  |
| ------------- | ------ | ------------------------------------- |
| Domain        | String | 域名。                                |
| Url           | String | 可下载离线日志的 URL。                |
| Name          | String | 离线日志文件名称。                    |
| Size          | int64  | 离线日志文件大小，单位为 B。          |
| LogTimeBegin  | String | 文件中日志开始时间，UTC 时间。        |
| LogTimeEnd    | String | 文件中日志结束时间，UTC 时间。        |

## 统计接口

> 查询用户站点维度或域名维度的统计指标信息，支持总带宽 / 上下行带宽、总流量 / 上下行流量、PV 等指标。
> 按接口规范，请求被划分为 5 种类型，分别对应不同的 `(group, showType)` 组合，SDK 提供 5 个独立方法，每个方法对应一种类型，调用方无需关心 `showType`：

| 类型 | 方法                  | 是否带 group | 内置 showType | 说明              |
| ---- | --------------------- | ------------ | ------------- | ----------------- |
| 1    | `GetStatTime`         | 否           | `time`        | 时间打点          |
| 2    | `GetStatPeak`         | 否           | `peak`        | 峰值              |
| 3    | `GetStatTimeByGroup`  | 是           | `time`        | 时间打点带分组    |
| 4    | `GetStatSumByGroup`   | 是           | `sum`         | 聚合带分组        |
| 5    | `GetStatTopByGroup`   | 是           | `top`         | TOP 带分组        |

### 类型 1：GetStatTime — 时间打点（无 group）

```go
cli := GetDefaultClient()
result, err := cli.GetStatTime(&api.StatTimeQueryData{
	Metric:    []string{"sum_bps"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T09:06:29Z",
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
})
b, _ := json.Marshal(result)
fmt.Printf("result:%s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 类型 2：GetStatPeak — 峰值（无 group）

```go
cli := GetDefaultClient()
result, err := cli.GetStatPeak(&api.StatPeakQueryData{
	Metric:    []string{"sum_bps", "upstream_bps", "download_bps"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T09:06:29Z",
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
})
b, _ := json.Marshal(result)
fmt.Printf("result:%s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 类型 3：GetStatTimeByGroup — 时间打点带 group

```go
cli := GetDefaultClient()
result, err := cli.GetStatTimeByGroup(&api.StatTimeByGroupQueryData{
	Metric:    []string{"pv"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T11:05:47Z",
	Group:     []string{"code"},
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
})
b, _ := json.Marshal(result)
fmt.Printf("result:%s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 类型 4：GetStatSumByGroup — 聚合带 group

```go
cli := GetDefaultClient()
result, err := cli.GetStatSumByGroup(&api.StatSumByGroupQueryData{
	Metric:    []string{"pv"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T11:05:47Z",
	Group:     []string{"code"},
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
})

b, _ := json.Marshal(result)
fmt.Printf("stat sum by group: %s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 类型 5：GetStatTopByGroup — TOP 带 group

```go
cli := GetDefaultClient()
result, err := cli.GetStatTopByGroup(&api.StatTopByGroupQueryData{
	Metric:    []string{"pv"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T11:05:47Z",
	Group:     []string{"host"},
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
	Limit: &api.StatLimit{PageSize: 100}, // 可选，限制 TOP 返回条数
})
b, _ := json.Marshal(result)
fmt.Printf("result: %s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 请求字段说明

5 种请求类型共享如下基础字段：

| 字段       | 类型              | 是否必选 | 说明                                                                                                            |
| ---------- | ----------------- | -------- |---------------------------------------------------------------------------------------------------------------|
| Metric     | []String          | 是       | 指标类型，可选值：`sum_bps` / `upstream_bps` / `download_bps` / `sum_flow` / `upstream_flow` / `download_flow` / `pv`。 |
| StartTime  | String            | 否       | 查询时间范围起始值，UTC 时间。默认 `EndTime` 前推 24 小时。最长可查近 31 天。                                                            |
| EndTime    | String            | 否       | 查询时间范围结束值，UTC 时间。默认当前时间。                                                                                      |
| Filter     | []StatFilterItem  | 否       | 过滤条件列表。                                                                                                       |

类型 3、4、5 在以上字段基础上额外要求 `Group` 字段：

| 字段   | 类型     | 是否必选 | 说明                                  |
| ------ | -------- | -------- | ------------------------------------- |
| Group  | []String | 是       | 聚合字段，例如 `code`、`host`。       |

类型 5（TOP 查询）在以上字段基础上还可选 `Limit` 字段：

| 字段   | 类型         | 是否必选 | 说明                                   |
| ------ | ------------ | -------- | -------------------------------------- |
| Limit  | *StatLimit   | 否       | TOP 返回条数限制，例如 `{PageSize:100}`。 |

`api.StatFilterItem` 字段说明：

| 字段       | 类型     | 是否必选 | 说明                                          |
| ---------- | -------- | -------- | --------------------------------------------- |
| Key        | String   | 是       | 过滤的 key，支持 `site`、`host`。             |
| Value      | []String | 是       | key 对应的值，支持多个。                      |
| Operation  | String   | 是       | 操作类型，支持 `equal`、`notequal`。          |

### 返回值说明

`result`是`api.StatResult`类型（即 `map[string][]api.StatDataPoint`），key 为请求中的 metric 名称（例如 `sum_bps`、`pv`），value 是该指标的数据点列表。

`api.StatDataPoint` 字段说明：

| 字段        | 类型             | 说明                                                                                  |
| ----------- | ---------------- | ------------------------------------------------------------------------------------- |
| Timestamp   | *int64           | 时间戳。类型 1 / 3（`time`）时为有效值；类型 2 / 4 / 5（`peak` / `sum` / `top`）时为 `nil`。 |
| Value       | json.Number      | 指标值。服务端可能返回数字或字符串，统一用 `json.Number` 承载，可通过 `Value.Int64()` / `Value.Float64()` / `Value.String()` 取值。 |
| GroupParams | json.RawMessage  | 聚合参数原始 JSON。无 group（类型 1 / 2）时为空数组 `[]`；带 group（类型 3 / 4 / 5）时为对象（如 `{"code":"2xx"}`）。 |


更多详细说明可以参考API文档：https://cloud.baidu.com/doc/GEO/s/4mp0unbx3

# 站点配置接口

EO 站点的全局配置功能共用同一接口：

- `PUT /v2/geo/site/{site}/config`：设置配置
- `GET /v2/geo/site/{site}/config`：查询配置（返回站点全部配置项）

`api.SiteConfig` 是统一的配置容器，所有字段均为指针类型 + `omitempty`：未赋值的字段不会出现在请求 body 中，未来新增配置项时按相同方式扩展即可。如需将某个配置显式设置为空（例如 `cacheTtl: []` 表示遵循源站-默认缓存策略），传一个非 nil 的空值（如 `&[]api.CacheTtl{}`）。

相关配置接口说明可参考对应配置的API文档：https://cloud.baidu.com/doc/GEO/s/Pmiigxbf0

## 设置节点缓存配置

```go
cli := GetDefaultClient()
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	CacheTtl: &[]api.CacheTtl{
		{
			Value:          "/",
			Weight:         100,
			OverrideOrigin: true,
			Ttl:            2592000,
			Type:           "path",
		},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.CacheTtl` 字段说明：

| 字段           | 类型   | 是否必选 | 说明                                                             |
| -------------- | ------ | -------- |----------------------------------------------------------------|
| Type           | String | 是       | 其合法值为“path”。表示缓存目录的路径。                                         |
| Value          | String | 是       | 其合法值为“/”。表示根目录。                                                |
| Weight         | Int    | 是       | 权重，合法值 `100`。                                                  |
| OverrideOrigin | Bool   | 是       | 表示缓存是否遵循源站。值为 `true` 时，表示不遵循源站，按照该条配置规则缓存。值为 `false` 时，表示遵循源站。 |
| Ttl            | Int    | 是       | 缓存时间，单位秒；`0` 表示不缓存。                                            |

## 设置查询字符串配置

```go
cli := GetDefaultClient()
query := false
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	CacheKey: &api.CacheKey{
		Query:       &query,
		IncludeArgs: &[]string{"test1"},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.CacheKey` 字段说明（`IncludeArgs` 与 `ExcludeArgs` 不可同时设置；二者仅在 `Query=false` 时生效）：

| 字段        | 类型      | 是否必选 | 说明                                    |
| ----------- | --------- | -------- |---------------------------------------|
| Query       | *Bool     | 是       | `true` 保留全部参数参与缓存；`false` 忽略全部参数参与缓存。 |
| IncludeArgs | *[]String | 否       | 保留指定参数参与缓存（仅 `Query=false` 时有效）。      |
| ExcludeArgs | *[]String | 否       | 忽略指定参数参与缓存（仅 `Query=false` 时有效）。      |

## 设置离线模式配置

```go
cli := GetDefaultClient()
offlineMode := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	OfflineMode: &offlineMode,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`OfflineMode` 字段说明：

| 字段        | 类型    | 是否必选 | 说明                                            |
| ----------- | ------- |------| ----------------------------------------------- |
| OfflineMode | *String | 是    | `ON` 开启离线模式；`OFF` 关闭离线模式。         |

## 设置强制 HTTPS 配置

```go
cli := GetDefaultClient()
httpToHttpsEnabled := "ON"
httpToHttpsCode := "302"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	HttpToHttpsEnabled: &httpToHttpsEnabled,
	HttpToHttpsCode:    &httpToHttpsCode,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`HttpToHttps*` 字段说明：

| 字段               | 类型    | 是否必选 | 说明                                                              |
| ------------------ | ------- |------| ----------------------------------------------------------------- |
| HttpToHttpsEnabled | *String | 是    | `ON` 开启强制 HTTPS 重定向；`OFF` 关闭。                          |
| HttpToHttpsCode    | *String | 否    | 重定向状态码，`301` 或 `302`；`HttpToHttpsEnabled=OFF` 时无效。   |

## 设置 HSTS 配置

```go
cli := GetDefaultClient()
maxAge := -1
includeSubDomains := false
preload := false
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Hsts: &api.HSTS{
		MaxAge:            &maxAge,
		IncludeSubDomains: &includeSubDomains,
		Preload:           &preload,
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.HSTS` 字段说明：

| 字段              | 类型  | 是否必选 | 说明                                             |
| ----------------- | ----- | -------- |------------------------------------------------|
| MaxAge            | *Int  | 是       | 配置保存时间，单位为天, 用户输入值为 0 ~ 730 或者 -1，为 -1 时表示关闭该配置项。 |
| IncludeSubDomains | *Bool | 是       | 是否包含子域名。                                       |
| Preload           | *Bool | 是       | 是否支持预加载。                                       |

## 设置 HTTP2 配置

```go
cli := GetDefaultClient()
http2Disable := "OFF"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Http2Disable: &http2Disable,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Http2Disable` 字段说明：

| 字段         | 类型    | 是否必选 | 说明                                          |
| ------------ | ------- |------| --------------------------------------------- |
| Http2Disable | *String | 是    | `OFF` 开启 HTTP2；`ON` 关闭 HTTP2。           |

## 设置 HTTP3 配置

```go
cli := GetDefaultClient()
http3Enable := true
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Http3: &api.HTTP3{
		Enable: &http3Enable,
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.HTTP3` 字段说明：

| 字段   | 类型  | 是否必选 | 说明                                |
| ------ | ----- | -------- |-----------------------------------|
| Enable | *Bool | 是       | `true` 开启 HTTP3；`false` 关闭 HTTP3。 |

## 设置最大上传大小配置

```go
cli := GetDefaultClient()
clientMaxBodySize := "500m"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	ClientMaxBodySize: &clientMaxBodySize,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`ClientMaxBodySize` 字段说明：

| 字段              | 类型    | 是否必选 | 说明                                                                                          |
| ----------------- | ------- |------| --------------------------------------------------------------------------------------------- |
| ClientMaxBodySize | *String | 是    | 单位支持 `b`、`k`、`m`（忽略大小写，`b` 可省略）。最大 `500m`。示例：`100`、`100k`、`100M`。 |

## 设置页面压缩配置

```go
cli := GetDefaultClient()
compress := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Compress:            &compress,
	CompressMethodArray: &[]string{"gzip", "br"},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Compress*` 字段说明：

| 字段                | 类型      | 是否必选 | 说明                          |
| ------------------- | --------- |------|-----------------------------|
| Compress            | *String   | 是    | `ON` 开启页面压缩；`OFF` 关闭页面压缩。   |
| CompressMethodArray | *[]String | 是    | 压缩方式，合法值：`gzip`、`br`，可同时启用。 |

## 设置智能加速配置

```go
cli := GetDefaultClient()
isa := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Isa: &isa,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Isa` 字段说明：

| 字段 | 类型    | 是否必选 | 说明                        |
| ---- | ------- |------|---------------------------|
| Isa  | *String | 是    | `ON` 开启智能加速；`OFF` 关闭智能加速。 |

## 设置自定义 HTTP 头配置

> 注意：本接口为全量更新，每次设置需带上希望保留的全部规则，否则原有配置会被覆盖。

```go
cli := GetDefaultClient()
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	HttpHeader: &[]api.HttpHeader{
		{Type: "response", Header: "Cache-Control", Value: "test", Action: "add"},
		{Type: "origin", Header: "Expires", Value: "test", Action: "add"},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.HttpHeader` 字段说明：

| 字段   | 类型   | 是否必选 | 说明                                                                                       |
| ------ | ------ | -------- | ------------------------------------------------------------------------------------------ |
| Type   | String | 是       | `origin` 表示回源生效；`response` 表示给用户响应时生效。                                   |
| Header | String | 是       | header 为 http 头字段，一般为 HTTP 的标准 Header，其长度限制为 128。也可以是用户自定义的 Header。                                                     |
| Value  | String | 是       | 指定 header 的值，其长度限制为 1000。可以是常量，也可以是变量；删除 HTTP 头时可以传空字符串 `""`。<br>**变量约束**：以 `$` 开始的子串必须符合 `${x}` 模式，合法变量为：<br>• `${uri}` 客户端请求的 URL 路径部分（不含查询参数）<br>• `${host}` 客户端请求 host 头部值<br>• `${scheme}` 客户端请求协议（`http` 或 `https`）<br>• `${request_uri}` 客户端请求路径和参数（含查询参数）<br>• `${jvip}` 节点 IP<br>• `${remote_addr}` 客户端 IP（存在代理时不准确）<br>• `${request_id}` 请求的唯一标识符<br>**典型非法值**：<br>• 变量不符合限制，如 `X-REQ-${url}`<br>• 含 `$` 但不符合 `${x}` 模式，如 `X-REQ-$uri`<br>注意：value 不支持 `$` 纯字符透传，如 `X-$` 非法。 |
| Action | String | 是       | 设置 HTTP 头合法值为 `add`，删除 HTTP 头合法值为 `remove`。                                                                |

注：最多设置 20 条 HTTP 回源请求头规则，最多设置 20 条 HTTP 节点响应头规则；不支持删除以 `ohc`、`baidu` 开头的回源请求头。

## 设置状态码缓存配置

```go
cli := GetDefaultClient()
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	CacheCodeTtl: &[]api.CacheTtlCode{
		{Value: "404", Weight: 100, OverrideOrigin: true, Ttl: 10, Type: "code"},
		{Value: "400", Weight: 100, OverrideOrigin: true, Ttl: 10, Type: "code"},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.CacheTtlCode` 字段说明：

| 字段           | 类型   | 是否必选 | 说明                                                                                  |
| -------------- | ------ | -------- | ------------------------------------------------------------------------------------- |
| Type           | String | 是       | 合法值 `code`，表示异常状态码缓存。                                                   |
| Value          | String | 是       | 4xx：`400`/`401`/`403`/`404`/`405`/`407`/`414`/`451`；<br/>5xx：`500`/`501`/`502`/`503`/`504`/`509`/`514`。 |
| Ttl            | Int    | 是       | 缓存时间，单位秒，取值范围 `0~315360000`。                                            |
| OverrideOrigin | Bool   | 是       | 合法值 `true`。                                                                       |
| Weight         | Int    | 是       | 权重，合法值 `100`。                                                                  |

注：最多 15 条状态码缓存规则，且不可重复。

## 设置 gRPC 回源配置

```go
cli := GetDefaultClient()
grpcOrigin := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	GrpcOrigin: &grpcOrigin,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`GrpcOrigin` 字段说明：

| 字段       | 类型    | 是否必选 | 说明                                  |
| ---------- | ------- |------| ------------------------------------- |
| GrpcOrigin | *String | 是    | `ON` 开启 gRPC 回源；`OFF` 关闭 gRPC 回源。     |

## 设置 HTTP2 回源配置

```go
cli := GetDefaultClient()
http2Origin := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Http2Origin: &http2Origin,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Http2Origin` 字段说明：

| 字段        | 类型    | 是否必选 | 说明                                       |
| ----------- | ------- |------| ------------------------------------------ |
| Http2Origin | *String | 是    | `ON` 开启 HTTP2 回源；`OFF` 关闭 HTTP2 回源。         |

## 查询站点配置

```go
cli := GetDefaultClient()
cfg, err := cli.GetSiteConfig("your_site.com")
if err != nil {
    fmt.Printf("err:%+v\n", err)
    return
}

// 只查询某一个配置项
if cfg.CacheTtl != nil {
    cacheTtlData, _ := json.Marshal(cfg.CacheTtl)
    fmt.Printf("cacheTtl: %s\n", cacheTtlData)
}

// 查询由多个字段组成的配置项（如「页面压缩」由 Compress + CompressMethodArray 共同组成）
if cfg.Compress != nil || cfg.CompressMethodArray != nil {
    compressData, _ := json.Marshal(struct {
        Compress            *string   `json:"compress,omitempty"`
        CompressMethodArray *[]string `json:"compressMethodArray,omitempty"`
    }{cfg.Compress, cfg.CompressMethodArray})
    fmt.Printf("compress: %s\n", compressData)
}
// 同理「强制 HTTPS」由 HttpToHttpsEnabled 和 HttpToHttpsCode 组成，可按相同方式聚合查询。
// 注意：当 HttpToHttpsEnabled = "OFF" 时，服务端不返回 HttpToHttpsCode
if cfg.HttpToHttpsEnabled != nil {
    httptohttpsData, _ := json.Marshal(struct {
        HttpToHttpsEnabled *string `json:"httpToHttpsEnabled,omitempty"`
        HttpToHttpsCode    *string `json:"httpToHttpsCode,omitempty"`
    }{cfg.HttpToHttpsEnabled, cfg.HttpToHttpsCode})
    fmt.Printf("httpToHttps: %s\n", httptohttpsData)
}

// 查询全部配置
allData, _ := json.Marshal(cfg)
fmt.Printf("SiteConfig: %s\n", allData)

// cfg.CacheTtl 为该站点当前的节点缓存配置；
// cfg.CacheKey 为该站点当前的查询字符串配置;
// cfg.OfflineMode 为该站点当前的离线模式配置;
// cfg.HttpToHttpsEnabled 和 cfg.HttpToHttpsCode 为该站点当前的强制 HTTPS 配置;
// cfg.Hsts 为该站点当前的 HSTS 配置;
// cfg.Http2Disable 为该站点当前的 HTTP2 配置;
// cfg.Http3 为该站点当前的 HTTP3 配置;
// cfg.ClientMaxBodySize 为该站点当前的最大上传大小配置;
// cfg.Compress 和 cfg.CompressMethodArray 为该站点当前的页面压缩配置;
// cfg.Isa 为该站点当前的智能加速配置;
// cfg.HttpHeader 为该站点当前的自定义 HTTP 头配置;
// cfg.CacheCodeTtl 为该站点当前的状态码缓存配置;
// cfg.GrpcOrigin 为该站点当前的 gRPC 回源配置;
// cfg.Http2Origin 为该站点当前的 HTTP2 回源配置;
```


