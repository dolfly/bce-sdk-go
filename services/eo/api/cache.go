package api

import (
	"errors"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util"
)

type PurgedId string
type PrefetchId string

// PurgeStatusQueryData defined a struct for the query conditions about the purge tasks' progress
type PurgeStatusQueryData struct {
	Site      string
	Id        string
	StartTime string
	EndTime   string
	Type      string
	Marker    string
}

// PrefetchStatusQueryData defined a struct for the query conditions about the prefetch tasks' progress
type PrefetchStatusQueryData struct {
	Site      string
	Id        string
	StartTime string
	EndTime   string
	Marker    string
}

// PurgeTask defined a struct for purge task
type PurgeTask struct {
	Url  string `json:"url"`
	Type string `json:"type,omitempty"`
}

// PrefetchTask defined a struct for prefetch task
type PrefetchTask struct {
	Url string `json:"url"`
}

// PurgeRecord defined a struct for purged task information
type PurgeRecord struct {
	Status     string    `json:"status"`
	Task       PurgeTask `json:"task"`
	CreatedAt  string    `json:"createdAt"`
	FinishedAt string    `json:"finishedAt,omitempty"`
	Progress   int64     `json:"progress"`
	Operator   string    `json:"operator"`
}

// PrefetchRecord defined a struct for prefetch task information
type PrefetchRecord struct {
	Status     string       `json:"status"`
	Task       PrefetchTask `json:"task"`
	CreatedAt  string       `json:"createdAt"`
	StartedAt  string       `json:"startedAt,omitempty"`
	FinishedAt string       `json:"finishedAt,omitempty"`
	Progress   int64        `json:"progress"`
	Operator   string       `json:"operator"`
	Reason     string       `json:"reason,omitempty"`
	Id         string       `json:"id"`
}

// PurgeRecords defined a struct for multi operating purged task records
type PurgeRecords struct {
	Details     []PurgeRecord `json:"details"`
	IsTruncated bool          `json:"isTruncated"`
	NextMarker  string        `json:"nextMarker,omitempty"`
}

// PrefetchRecords defined a struct for multi operating prefetch task records
type PrefetchRecords struct {
	Details     []PrefetchRecord `json:"details"`
	IsTruncated bool             `json:"isTruncated"`
	NextMarker  string           `json:"nextMarker,omitempty"`
}

// Purge - tells the EO system to purge the specified files
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mhsrv9ry
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site that the purge tasks belong to
//   - tasks: the tasks about purging the files from the EO nodes
//
// RETURNS:
//   - PurgedId: an ID representing a purged task, using it to search the task progress
//   - error: nil if success otherwise the specific error
func Purge(cli bce.Client, site string, tasks []PurgeTask) (PurgedId, error) {
	if site == "" {
		return "", errors.New("site is required")
	}

	respObj := &struct {
		Id string `json:"id"`
	}{}

	err := httpRequest(cli, "POST", "/v2/geo/cache/purge", nil, &struct {
		Tasks []PurgeTask `json:"tasks"`
		Site  string      `json:"site"`
	}{
		Tasks: tasks,
		Site:  site,
	}, respObj)
	if err != nil {
		return "", err
	}

	return PurgedId(respObj.Id), nil
}

// GetPurgedStatus - get the purged progress
// For details, please refer https://cloud.baidu.com/doc/GEO/s/mmhssw91q
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: querying conditions, it contains the site, the time interval, the task ID, the type and the marker
//
// RETURNS:
//   - *PurgeRecords: the details about the purged
//   - error: nil if success otherwise the specific error
func GetPurgedStatus(cli bce.Client, queryData *PurgeStatusQueryData) (*PurgeRecords, error) {
	if queryData == nil || queryData.Site == "" {
		return nil, errors.New("site is required")
	}

	params := map[string]string{
		"site": queryData.Site,
	}

	// when querying by task id, no time and marker parameters
	if queryData.Id != "" {
		params["id"] = queryData.Id
	} else {
		if err := getTimeParams(params, queryData.StartTime, queryData.EndTime); err != nil {
			return nil, err
		}
		if queryData.Marker != "" {
			params["marker"] = queryData.Marker
		}
	}

	if queryData.Type != "" {
		params["type"] = queryData.Type
	}

	respObj := &PurgeRecords{}
	err := httpRequest(cli, "GET", "/v2/geo/cache/purge", params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// Prefetch - tells the EO system to prefetch the specified files
// For details, please refer https://cloud.baidu.com/doc/GEO/s/5mhsuituv
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site that the prefetch tasks belong to
//   - tasks: the tasks about prefetch the files from the EO nodes
//
// RETURNS:
//   - PrefetchId: an ID representing a prefetch task, using it to search the task progress
//   - error: nil if success otherwise the specific error
func Prefetch(cli bce.Client, site string, tasks []PrefetchTask) (PrefetchId, error) {
	if site == "" {
		return "", errors.New("site is required")
	}

	respObj := &struct {
		Id string `json:"id"`
	}{}

	err := httpRequest(cli, "POST", "/v2/geo/cache/prefetch", nil, &struct {
		Tasks []PrefetchTask `json:"tasks"`
		Site  string         `json:"site"`
	}{
		Tasks: tasks,
		Site:  site,
	}, respObj)
	if err != nil {
		return "", err
	}

	return PrefetchId(respObj.Id), nil
}

// GetPrefetchStatus - get the prefetch progress
// For details, please refer https://cloud.baidu.com/doc/GEO/s/Bmhsv5i9u
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: querying conditions, it contains the site, the time interval, the task ID and the marker
//
// RETURNS:
//   - *PrefetchRecords: the details about the prefetch
//   - error: nil if success otherwise the specific error
func GetPrefetchStatus(cli bce.Client, queryData *PrefetchStatusQueryData) (*PrefetchRecords, error) {
	if queryData == nil || queryData.Site == "" {
		return nil, errors.New("site is required")
	}

	params := map[string]string{
		"site": queryData.Site,
	}

	// when querying by task id, no time and marker parameters
	if queryData.Id != "" {
		params["id"] = queryData.Id
	} else {
		if err := getTimeParams(params, queryData.StartTime, queryData.EndTime); err != nil {
			return nil, err
		}
		if queryData.Marker != "" {
			params["marker"] = queryData.Marker
		}
	}

	respObj := &PrefetchRecords{}
	err := httpRequest(cli, "GET", "/v2/geo/cache/prefetch", params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

func getTimeParams(params map[string]string, startTime, endTime string) error {
	// get "endTime"
	endTs := int64(0)
	if endTime == "" {
		// default current time
		endTs = time.Now().Unix()
		params["endTime"] = util.FormatISO8601Date(endTs)
	} else {
		t, err := util.ParseISO8601Date(endTime)
		if err != nil {
			return err
		}
		endTs = t.Unix()
		params["endTime"] = endTime
	}

	// get "startTime", the default "startTime" is one day later than the "endTime"
	startTs := int64(0)
	if startTime == "" {
		startTs = endTs - 24*60*60
		params["startTime"] = util.FormatISO8601Date(startTs)
	} else {
		t, err := util.ParseISO8601Date(startTime)
		if err != nil {
			return err
		}
		startTs = t.Unix()
		params["startTime"] = startTime
	}

	// the "startTime" should be less than the "endTime"
	if startTs > endTs {
		return errors.New("error time range, the startTime should be less than the endTime")
	}

	return nil
}
