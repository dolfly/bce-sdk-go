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
purgedStatus, err = cli.GetPurgedStatus(&api.PurgeStatusQueryData{
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
prefetchStatus, err = cli.GetPrefetchStatus(&api.PrefetchStatusQueryData{
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
| StartTime  | String   | 否       | 查询时间范围起始值，UTC 时间，ISO8601 格式。默认为 `EndTime` 前推 8 小时。 |
| EndTime    | String   | 否       | 查询时间范围结束值，UTC 时间，ISO8601 格式。默认为当前时间。          |
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
fmt.Printf("result:%+v\n", result)
fmt.Printf("err:%+v\n", err)
```

### 类型 2：GetStatPeak — 峰值（无 group）

```go
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
```

### 类型 3：GetStatTimeByGroup — 时间打点带 group

```go
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
```

### 类型 4：GetStatSumByGroup — 聚合带 group

```go
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

// 解析按 code 聚合的 groupParams
for _, item := range result["pv"] {
	groupMap, _ := item.GroupParamsMap()
	fmt.Printf("code=%s, value=%s\n", groupMap["code"], item.Value)
}
```

### 类型 5：GetStatTopByGroup — TOP 带 group

```go
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
```

### 请求字段说明

5 种请求类型共享如下基础字段：

| 字段       | 类型              | 是否必选 | 说明                                                                                |
| ---------- | ----------------- | -------- | ----------------------------------------------------------------------------------- |
| Metric     | []String          | 是       | 指标类型，可选值：`sum_bps` / `upstream_bps` / `download_bps` / `sum_flow` / `upstream_flow` / `download_flow` / `pv`。 |
| StartTime  | String            | 否       | 查询时间范围起始值，UTC ISO8601。默认 `EndTime` 前推 24 小时。最长可查近 31 天。     |
| EndTime    | String            | 否       | 查询时间范围结束值，UTC ISO8601。默认当前时间。                                     |
| Filter     | []StatFilterItem  | 否       | 过滤条件列表。                                                                      |

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

由于 `GroupParams` 在不同场景下形态不同，`StatDataPoint` 提供了 `GroupParamsMap()` 辅助方法将其解析为 `map[string]string`：

```go
groupMap, err := item.GroupParamsMap()
// 无 group 时返回空 map；有 group 时返回 {"code":"2xx"} 这类映射
```

更多详细说明可以参考API文档：https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
