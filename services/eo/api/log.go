package api

import (
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
)

// LogEntry defined a struct for one offline log file information
type LogEntry struct {
	Domain       string `json:"domain"`
	Url          string `json:"url"`
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	LogTimeBegin string `json:"logTimeBegin"`
	LogTimeEnd   string `json:"logTimeEnd"`
}

// LogQueryData defined a struct for offline log query conditions, used as request body
type LogQueryData struct {
	Site       string   `json:"site"`
	StartTime  string   `json:"startTime,omitempty"`
	EndTime    string   `json:"endTime,omitempty"`
	DomainList []string `json:"domainList,omitempty"`
	PageNo     int      `json:"pageNo,omitempty"`
	PageSize   int      `json:"pageSize,omitempty"`
}

// LogQueryResult defined a struct for the offline log query result
type LogQueryResult struct {
	LogEntryList []LogEntry `json:"logEntryList"`
	TotalCount   string     `json:"totalCount"`
}

// GetOfflineLog - get the offline log download urls of one or multi domains under a site
// For details, please refer https://cloud.baidu.com/doc/GEO/s/hmmljdegd
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: the querying conditions
//
// RETURNS:
//   - *LogQueryResult: the offline log entries and total count
//   - error: nil if success otherwise the specific error
func GetOfflineLog(cli bce.Client, queryData *LogQueryData) (*LogQueryResult, error) {
	if queryData == nil || queryData.Site == "" {
		return nil, errors.New("site is required")
	}
	if queryData.PageNo < 0 {
		return nil, errors.New("invalid PageNo, it should be larger than or equal to 0")
	}
	if queryData.PageSize < 0 {
		return nil, errors.New("invalid PageSize, it should be larger than or equal to 0")
	}

	respObj := &LogQueryResult{}
	if err := httpRequest(cli, "POST", "/v2/geo/log", nil, queryData, respObj); err != nil {
		return nil, err
	}

	return respObj, nil
}
