package eo

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/eo/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "geo.baidubce.com"
)

// Client of EO service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the EO service client with default configuration
// Use `cli.Config.xxx` to access the config or change it to non-default value
func NewClient(ak, sk, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 && len(sk) == 0 { // to support public-read-write request
		credentials, err = nil, nil
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
		if err != nil {
			return nil, err
		}
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

// SendCustomRequest - send a HTTP request, and response data or error, it use the default times for retrying
//
// PARAMS:
//   - method: the HTTP requested method, e.g. "GET", "POST", "PUT" ...
//   - urlPath: a path component, consisting of a sequence of path segments separated by a slash ( / ).
//   - params: the query params, which will be append to the query path, and separate by "&"
//     e.g. http://www.baidu.com?query_param1=value1&query_param2=value2
//   - reqHeaders: the request http headers
//   - bodyObj: the HTTP requested body content transferred to a goland object
//   - respObj: the HTTP response content transferred to a goland object
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (cli *Client) SendCustomRequest(method string, urlPath string, params, reqHeaders map[string]string, bodyObj interface{}, respObj interface{}) error {
	return api.SendCustomRequest(cli, method, urlPath, params, reqHeaders, bodyObj, respObj)
}

// Purge - tells the EO system to purge the specified files
// For more details, please refer https://cloud.baidu.com/doc/GEO/s/4mhsrv9ry
//
// PARAMS:
//   - site: the site that the purge tasks belong to
//   - tasks: the tasks about purging the files from the EO nodes
//
// RETURNS:
//   - PurgedId: an ID representing a purged task, using it to search the task progress
//   - error: nil if success otherwise the specific error
func (cli *Client) Purge(site string, tasks []api.PurgeTask) (api.PurgedId, error) {
	return api.Purge(cli, site, tasks)
}

// GetPurgedStatus - get the purged progress
// For more details, please refer https://cloud.baidu.com/doc/GEO/s/mmhssw91q
//
// PARAMS:
//   - queryData: querying conditions, it contains the site, the time interval, the task ID, the type and the marker
//
// RETURNS:
//   - *PurgeRecords: the details about the purged
//   - error: nil if success otherwise the specific error
func (cli *Client) GetPurgedStatus(queryData *api.PurgeStatusQueryData) (*api.PurgeRecords, error) {
	return api.GetPurgedStatus(cli, queryData)
}

// Prefetch - tells the EO system to prefetch the specified files
// For more details, please refer https://cloud.baidu.com/doc/GEO/s/5mhsuituv
//
// PARAMS:
//   - site: the site that the prefetch tasks belong to
//   - tasks: the tasks about prefetch the files from the EO nodes
//
// RETURNS:
//   - PrefetchId: an ID representing a prefetch task, using it to search the task progress
//   - error: nil if success otherwise the specific error
func (cli *Client) Prefetch(site string, tasks []api.PrefetchTask) (api.PrefetchId, error) {
	return api.Prefetch(cli, site, tasks)
}

// GetPrefetchStatus - get the prefetch progress
// For more details, please refer https://cloud.baidu.com/doc/GEO/s/Bmhsv5i9u
//
// PARAMS:
//   - queryData: querying conditions, it contains the site, the time interval, the task ID and the marker
//
// RETURNS:
//   - *PrefetchRecords: the details about the prefetch
//   - error: nil if success otherwise the specific error
func (cli *Client) GetPrefetchStatus(queryData *api.PrefetchStatusQueryData) (*api.PrefetchRecords, error) {
	return api.GetPrefetchStatus(cli, queryData)
}

// GetOfflineLog - get the offline log download urls of one or multi domains under a site
//
// PARAMS:
//   - queryData: querying conditions, it contains the site, the time interval, the domain list and pagination
//
// RETURNS:
//   - *LogQueryResult: the offline log entries and total count
//   - error: nil if success otherwise the specific error
func (cli *Client) GetOfflineLog(queryData *api.LogQueryData) (*api.LogQueryResult, error) {
	return api.GetOfflineLog(cli, queryData)
}

// GetStatTime - type 1: showType=time , no group
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` is required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatTime(queryData *api.StatTimeQueryData) (api.StatResult, error) {
	return api.GetStatTime(cli, queryData)
}

// GetStatPeak - type 2: showType=peak, no group
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` is required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatPeak(queryData *api.StatPeakQueryData) (api.StatResult, error) {
	return api.GetStatPeak(cli, queryData)
}

// GetStatTimeByGroup - type 3: showType=time, group required
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatTimeByGroup(queryData *api.StatTimeByGroupQueryData) (api.StatResult, error) {
	return api.GetStatTimeByGroup(cli, queryData)
}

// GetStatSumByGroup - type 4: showType=sum, group required
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatSumByGroup(queryData *api.StatSumByGroupQueryData) (api.StatResult, error) {
	return api.GetStatSumByGroup(cli, queryData)
}

// GetStatTopByGroup - type 5: showType=top, group required
// Note: only `sum_flow`, `upstream_flow`, `download_flow`, `pv` metrics are supported by the server for top queries.
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatTopByGroup(queryData *api.StatTopByGroupQueryData) (api.StatResult, error) {
	return api.GetStatTopByGroup(cli, queryData)
}
