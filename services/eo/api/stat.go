package api

import (
	"encoding/json"
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
)

// StatFilterItem defined a filter condition for stat query
type StatFilterItem struct {
	Key       string   `json:"key"`       // "site" | "host"
	Value     []string `json:"value"`     // values for the key
	Operation string   `json:"operation"` // "equal" | "notequal"
}

// StatTimeQueryData - request body for type 1: no group, showType=time
type StatTimeQueryData struct {
	Metric    []string         `json:"metrics"`
	StartTime string           `json:"startTime,omitempty"`
	EndTime   string           `json:"endTime,omitempty"`
	Filter    []StatFilterItem `json:"filter,omitempty"`
}

// StatPeakQueryData - request body for type 2: no group, showType=peak
type StatPeakQueryData struct {
	Metric    []string         `json:"metrics"`
	StartTime string           `json:"startTime,omitempty"`
	EndTime   string           `json:"endTime,omitempty"`
	Filter    []StatFilterItem `json:"filter,omitempty"`
}

// StatTimeByGroupQueryData - request body for type 3: showType=time, group required
type StatTimeByGroupQueryData struct {
	Metric    []string         `json:"metrics"`
	StartTime string           `json:"startTime,omitempty"`
	EndTime   string           `json:"endTime,omitempty"`
	Filter    []StatFilterItem `json:"filter,omitempty"`
	Group     []string         `json:"group"` // required
}

// StatSumByGroupQueryData - request body for type 4: showType=sum, group required
type StatSumByGroupQueryData struct {
	Metric    []string         `json:"metrics"`
	StartTime string           `json:"startTime,omitempty"`
	EndTime   string           `json:"endTime,omitempty"`
	Filter    []StatFilterItem `json:"filter,omitempty"`
	Group     []string         `json:"group"` // required
}

// StatLimit defined a limit option for top queries
type StatLimit struct {
	PageSize int `json:"pageSize"`
}

// StatTopByGroupQueryData - request body for type 5: showType=top, group required
type StatTopByGroupQueryData struct {
	Metric    []string         `json:"metrics"`
	StartTime string           `json:"startTime,omitempty"`
	EndTime   string           `json:"endTime,omitempty"`
	Filter    []StatFilterItem `json:"filter,omitempty"`
	Group     []string         `json:"group"`           // required
	Limit     *StatLimit       `json:"limit,omitempty"` // optional, top result count limit
}

// StatDataPoint defined a struct for one metric data point
//   - Timestamp: Int per the API doc, but the server returns null when ShowType is peak/sum/top, so use *int64
//   - Value: the server may return either a number (e.g. 306) or a string, so use json.Number for transparent decoding;
//     callers can convert via Value.Int64() / Value.Float64() / Value.String() as needed
//   - GroupParams: List<GroupParams> per the API doc; the server returns an empty array `[]` when no group is set,
//     and an object like `{"code":"2xx"}` when group is set, so use json.RawMessage for transparent decoding,
//     and call GroupParamsMap() to extract the group key-value pairs
type StatDataPoint struct {
	Timestamp   *int64          `json:"timestamp"`
	Value       json.Number     `json:"value"`
	GroupParams json.RawMessage `json:"groupParams"`
}

// GroupParamsMap parses GroupParams into a map of group key to value.
// When GroupParams is empty or an empty array, returns an empty map.
func (p StatDataPoint) GroupParamsMap() (map[string]string, error) {
	if len(p.GroupParams) == 0 {
		return map[string]string{}, nil
	}
	// the server returns an empty array `[]` when no group is set
	var arr []interface{}
	if err := json.Unmarshal(p.GroupParams, &arr); err == nil {
		return map[string]string{}, nil
	}
	m := map[string]string{}
	if err := json.Unmarshal(p.GroupParams, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatResult defined the response of stat query, the key is metric name (e.g. "sum_bps", "pv")
type StatResult map[string][]StatDataPoint

// sendStatRequest is the internal helper that performs the actual HTTP POST to /v2/geo/stat
func sendStatRequest(cli bce.Client, body interface{}) (StatResult, error) {
	respObj := StatResult{}
	if err := httpRequest(cli, "POST", "/v2/geo/stat", nil, body, &respObj); err != nil {
		return nil, err
	}
	return respObj, nil
}

// GetStatTime - type 1: no group, showType=time
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: the querying conditions, `Metric` is required
//
// RETURNS:
//   - StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func GetStatTime(cli bce.Client, queryData *StatTimeQueryData) (StatResult, error) {
	if queryData == nil || len(queryData.Metric) == 0 {
		return nil, errors.New("metric is required")
	}
	return sendStatRequest(cli, struct {
		StatTimeQueryData
		ShowType string `json:"showType"`
	}{
		StatTimeQueryData: *queryData,
		ShowType:          "time",
	})
}

// GetStatPeak - type 2: no group, showType=peak
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: the querying conditions, `Metric` is required
//
// RETURNS:
//   - StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func GetStatPeak(cli bce.Client, queryData *StatPeakQueryData) (StatResult, error) {
	if queryData == nil || len(queryData.Metric) == 0 {
		return nil, errors.New("metric is required")
	}
	return sendStatRequest(cli, struct {
		StatPeakQueryData
		ShowType string `json:"showType"`
	}{
		StatPeakQueryData: *queryData,
		ShowType:          "peak",
	})
}

// GetStatTimeByGroup - type 3: showType=time, group required
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func GetStatTimeByGroup(cli bce.Client, queryData *StatTimeByGroupQueryData) (StatResult, error) {
	if queryData == nil || len(queryData.Metric) == 0 {
		return nil, errors.New("metric is required")
	}
	if len(queryData.Group) == 0 {
		return nil, errors.New("group is required")
	}
	return sendStatRequest(cli, struct {
		StatTimeByGroupQueryData
		ShowType string `json:"showType"`
	}{
		StatTimeByGroupQueryData: *queryData,
		ShowType:                 "time",
	})
}

// GetStatSumByGroup - type 4: showType=sum, group required
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func GetStatSumByGroup(cli bce.Client, queryData *StatSumByGroupQueryData) (StatResult, error) {
	if queryData == nil || len(queryData.Metric) == 0 {
		return nil, errors.New("metric is required")
	}
	if len(queryData.Group) == 0 {
		return nil, errors.New("group is required")
	}
	return sendStatRequest(cli, struct {
		StatSumByGroupQueryData
		ShowType string `json:"showType"`
	}{
		StatSumByGroupQueryData: *queryData,
		ShowType:                "sum",
	})
}

// GetStatTopByGroup - type 5: showType=top, group required
// Note: only `sum_flow`, `upstream_flow`, `download_flow`, `pv` metrics are supported by the server for top queries.
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func GetStatTopByGroup(cli bce.Client, queryData *StatTopByGroupQueryData) (StatResult, error) {
	if queryData == nil || len(queryData.Metric) == 0 {
		return nil, errors.New("metric is required")
	}
	if len(queryData.Group) == 0 {
		return nil, errors.New("group is required")
	}
	return sendStatRequest(cli, struct {
		StatTopByGroupQueryData
		ShowType string `json:"showType"`
	}{
		StatTopByGroupQueryData: *queryData,
		ShowType:                "top",
	})
}
