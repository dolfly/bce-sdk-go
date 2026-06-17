/*
 * Copyright 2020 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// scs.go - the SCS for Redis APIs definition supported by the redis service
package scs

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	KEY_MARKER                = "marker"
	KEY_MAX_KEYS              = "maxKeys"
	INSTANCE_URL_V1           = bce.URI_PREFIX + "v1" + "/instance"
	INSTANCE_URL_V2           = bce.URI_PREFIX + "v2" + "/instance"
	URI_PREFIX_V2             = bce.URI_PREFIX + "v2"
	URI_PREFIX_V1             = bce.URI_PREFIX + "v1"
	REQUEST_SECURITYGROUP_URL = "/security"
	REQUEST_RECYCLER_URL      = "/recycler"
)

func (c *Client) request(method, url string, result, body interface{}) (interface{}, error) {
	var err error
	if result != nil {
		err = bce.NewRequestBuilder(c).
			WithMethod(method).
			WithURL(url).
			WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
			WithBody(body).
			WithResult(result).
			Do()
	} else {
		err = bce.NewRequestBuilder(c).
			WithMethod(method).
			WithURL(url).
			WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
			WithBody(body).
			Do()
	}
	// fmt.Println(Json(result))
	return result, err
}

func getInstanceUrlWithId(instanceId string) string {
	return INSTANCE_URL_V1 + "/" + instanceId
}

func getWhiteListUrlWithInstanceId(instanceId string) string {
	return getInstanceUrlWithId(instanceId) + "/whitelist"
}

// List Security Group By Vpc URL
func getSecurityGroupWithVpcIdUrl(vpcId string) string {
	return URI_PREFIX_V2 + REQUEST_SECURITYGROUP_URL + "/vpc/" + vpcId
}

// List Security Group By Instance URL
func getSecurityGroupWithInstanceIdUrl(instanceId string) string {
	return INSTANCE_URL_V1 + "/" + instanceId + "/securityGroup"
}

// Bind Security Group To Instance URL
func getBindSecurityGroupWithUrl() string {
	return URI_PREFIX_V2 + REQUEST_SECURITYGROUP_URL + "/bind"
}

// UnBind Security Group To Instance URL
func getUnBindSecurityGroupWithUrl() string {
	return URI_PREFIX_V2 + REQUEST_SECURITYGROUP_URL + "/unbind"
}

// Batch Replace Security Group URL
func getReplaceSecurityGroupWithUrl() string {
	return URI_PREFIX_V2 + REQUEST_SECURITYGROUP_URL + "/update"
}

// Recycler URL
func getRecyclerUrl() string {
	return URI_PREFIX_V2 + REQUEST_RECYCLER_URL + "/list"
}

// Recycler Recover URL
func getRecyclerRecoverUrl() string {
	return URI_PREFIX_V2 + REQUEST_RECYCLER_URL + "/recover"
}

// Recycler Recover URL
func getRecyclerDeleteUrl() string {
	return URI_PREFIX_V2 + REQUEST_RECYCLER_URL + "/delete"
}

// Renew Instance URL
func getRenewUrl() string {
	return INSTANCE_URL_V2 + "/renew"
}

func getBackupUrlWithInstanceId(instanceId string) string {
	return INSTANCE_URL_V1 + "/" + instanceId + "/backup"
}

func getLogsUrlWithInstanceId(instanceId string) string {
	return INSTANCE_URL_V1 + "/" + instanceId + "/log"
}

func getLogsUrlWithLogId(instanceId, logId string) string {
	return INSTANCE_URL_V1 + "/" + instanceId + "/log/" + logId
}

func getGroupUrl() string {
	return "/v2/group"
}
func getTemplateUrl() string {
	return "/v2/template"
}

func getSyncGroupUrl() string {
	return URI_PREFIX_V1 + "/syncGroup"
}

func Json(v interface{}) string {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		panic("convert to json faild")
	}
	return string(jsonStr)
}

// Convert marker to request params
func getMarkerParams(marker *Marker) map[string]string {
	if marker == nil {
		marker = &Marker{Marker: "-1"}
	}
	params := make(map[string]string, 2)
	params[KEY_MARKER] = marker.Marker
	if marker.MaxKeys > 0 {
		params[KEY_MAX_KEYS] = strconv.Itoa(marker.MaxKeys)
	}
	return params
}

// Convert struct to request params
func getQueryParams(val interface{}) (map[string]string, error) {
	var params map[string]string
	if val != nil {
		err := json.Unmarshal([]byte(Json(val)), &params)
		if err != nil {
			return nil, err
		}
	}
	return params, nil
}

// CreateInstance - create an instance with specified parameters
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - reqBody: the request body to create instance
//
// RETURNS:
//   - *CreateInstanceResult: result of the instance ids newly created
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(args *CreateInstanceArgs) (*CreateInstanceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create scs argments")
	}
	if len(args.ClientAuth) != 0 {
		cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.ClientAuth)
		if err != nil {
			return nil, err
		}
		args.ClientAuth = cryptedPass
	}
	result := &CreateInstanceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListInstances - list all instances with the specified parameters
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to list instances
//
// RETURNS:
//   - *ListInstanceResult: result of the instance list
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstances(args *ListInstancesArgs) (*ListInstancesResult, error) {
	if args == nil {
		args = &ListInstancesArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListInstancesResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V2).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParamFilter("instanceIds", args.InstanceIds).
		WithQueryParamFilter("vnetIp", args.VnetIp).
		WithResult(result).
		Do()

	return result, err
}

// GetInstanceDetail - get details of the specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//
// RETURNS:
//   - *GetInstanceDetailResult: result of the instance details
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(instanceId string) (*GetInstanceDetailResult, error) {
	result := &GetInstanceDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V2 + "/" + instanceId).
		WithResult(result).
		Do()

	return result, err
}

// ResizeInstance - resize a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance to be resized
//   - reqBody: the request body to resize instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeInstance(instanceId string, args *ResizeInstanceArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/change").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// GetCreatePrice - get create price
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - reqBody: the request body to get price
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GetCreatePriceResult: result of the create price
func (c *Client) GetCreatePrice(args *CreatePriceArgs) (*CreatePriceResult, error) {
	result := &CreatePriceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL("/v2/price").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// GetResizePrice - get resize price
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance to be resized
//   - reqBody: the request body to get resize price
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GetResizePriceResult: result of the resize price
func (c *Client) GetResizePrice(instanceId string, args *ResizePriceArgs) (*ResizePriceResult, error) {
	result := &ResizePriceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2+"/"+instanceId+"/price").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// AddReplication - add replications
//
// PARAMS:
//   - instanceId: id of the instance to be resized
//   - args: replicationInfo
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AddReplication(instanceId string, args *ReplicationArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2+"/"+instanceId+"/resizeReplication").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteReplication - delete replications
//
// PARAMS:
//   - instanceId: id of the instance to be resized
//   - args: replicationInfo
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteReplication(instanceId string, args *ReplicationArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(INSTANCE_URL_V2+"/"+instanceId+"/resizeReplication").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// RestartInstance - restart a specified instance
//
// PARAMS:
//   - instanceId: id of the instance to be resized
//   - args: specify restart immediately or postpone restart to time window
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RestartInstance(instanceId string, args *RestartInstanceArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceUrlWithId(instanceId)+"/restart").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteInstance - delete a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance to be deleted
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(instanceId string, clientToken string) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(INSTANCE_URL_V1+"/"+instanceId).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// UpdateInstanceName - update name of a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance to be deleted
//   - args: the arguments to Update instanceName
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(instanceId string, args *UpdateInstanceNameArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/rename").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// GetNodeTypeList - list all nodetype
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance to be deleted
//   - args: the arguments to Update instanceName
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GetNodeTypeList() (*GetNodeTypeListResult, error) {
	getNodeTypeListResult := &GetNodeTypeListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL("/v2/nodetypes").
		WithResult(getNodeTypeListResult).
		Do()

	return getNodeTypeListResult, err
}

// ListsSubnet - list all Subnets
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to list all subnets, not necessary
//
// RETURNS:
//   - *ListSubnetsResult: result of the subnet list
//   - error: nil if success otherwise the specific error
func (c *Client) ListSubnets(args *ListSubnetsArgs) (*ListSubnetsResult, error) {
	if args == nil {
		args = &ListSubnetsArgs{}
	}

	result := &ListSubnetsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL("/v1/subnet").
		WithQueryParamFilter("vpcId", args.VpcID).
		WithQueryParamFilter("zoneName", args.ZoneName).
		WithResult(result).
		Do()

	return result, err
}

// UpdateInstanceDomainName - update name of a specified instance domain
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to update domainName
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceDomainName(instanceId string, args *UpdateInstanceDomainNameArgs) error {

	if args == nil || args.Domain == "" {
		return fmt.Errorf("unset Domain")
	}
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/renameDomain").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// GetZoneList - list all zone
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//
// RETURNS:
//   - *GetZoneListResult: result of the zone list
//   - error: nil if success otherwise the specific error
func (c *Client) GetZoneList() (*GetZoneListResult, error) {
	result := &GetZoneListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL("/v1/zone").
		WithResult(result).
		Do()

	return result, err
}

// FlushInstance - flush a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to flush instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) FlushInstance(instanceId string, args *FlushInstanceArgs) error {

	cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.Password)
	if err != nil {
		return err
	}
	args.Password = cryptedPass

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/flush").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// BindingTags - bind tags to a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to bind tags to instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) BindingTag(instanceId string, args *BindingTagArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1 + "/" + instanceId + "/bindTag").
		WithBody(args).
		Do()
}

// UnBindingTags - unbind tags to a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to unbind tags to instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UnBindingTag(instanceId string, args *BindingTagArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1 + "/" + instanceId + "/unBindTag").
		WithBody(args).
		Do()
}

// SetAsMaster - set instance as master
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SetAsMaster(instanceId string) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V2 + "/" + instanceId + "/setAsMaster").
		Do()
}

// SetAsSlave - set instance as master
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to set instance as slave
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SetAsSlave(instanceId string, args *SetAsSlaveArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithBody(args).
		WithURL(INSTANCE_URL_V2 + "/" + instanceId + "/setAsSlave").
		Do()
}

// GetSecurityIp - list all securityIps
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//
// RETURNS:
//   - *ListSecurityIp: result of the security IP list
//   - error: nil if success otherwise the specific error
func (c *Client) GetSecurityIp(instanceId string) (*GetSecurityIpResult, error) {

	result := &GetSecurityIpResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V1 + "/" + instanceId + "/securityIp").
		WithResult(result).
		Do()

	return result, err
}

// AddSecurityIp - add securityIp to access a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to add securityIp
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AddSecurityIp(instanceId string, args *SecurityIpArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/securityIp").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DeleteSecurityIp - delete securityIp to access a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to delete securityIp
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteSecurityIp(instanceId string, args *SecurityIpArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/securityIp").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// GetWhiteListGroup - get white list groups of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - groupName: name of the white list group
//
// RETURNS:
//   - *WhiteListGroupResult: result of the white list groups
//   - error: nil if success otherwise the specific error
func (c *Client) GetWhiteListGroup(instanceId, groupName string) (*WhiteListGroupResult, error) {
	result := &WhiteListGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getWhiteListUrlWithInstanceId(instanceId)).
		WithQueryParamFilter("groupName", groupName).
		WithResult(result).
		Do()

	return result, err
}

// CreateWhiteListGroup - create a white list group of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to create white list group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateWhiteListGroup(instanceId string, args *WhiteListGroupArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getWhiteListUrlWithInstanceId(instanceId)).
		WithBody(args).
		Do()
}

// ModifyWhiteListGroup - modify a white list group of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to modify white list group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyWhiteListGroup(instanceId string, args *WhiteListGroupArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getWhiteListUrlWithInstanceId(instanceId)).
		WithBody(args).
		Do()
}

// DeleteWhiteListGroup - delete a white list group of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - groupName: name of the white list group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteWhiteListGroup(instanceId, groupName string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getWhiteListUrlWithInstanceId(instanceId)).
		WithQueryParamFilter("groupName", groupName).
		Do()
}

// ModifyPassword - modify the password of a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to Modify Password
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyPassword(instanceId string, args *ModifyPasswordArgs) error {

	cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.Password)
	if err != nil {
		return err
	}
	args.Password = cryptedPass

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/modifyPassword").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// RenameDomain - rename domain
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to rename domain
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RenameDomain(instanceId string, args *RenameDomainArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V2+"/"+instanceId+"/renameDomain").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// SwapDomain - swap domain
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to swap domain
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SwapDomain(instanceId string, args *SwapDomainArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2+"/"+instanceId+"/swapDomain").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// GetParameters - query the configuration parameters and running parameters of redis instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//
// RETURNS:
//   - *GetParameterResult: result of the parameters
//   - error: nil if success otherwise the specific error
func (c *Client) GetParameters(instanceId string) (*GetParametersResult, error) {

	result := &GetParametersResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V1 + "/" + instanceId + "/parameter").
		WithResult(result).
		Do()

	return result, err
}

// ModifyParameters - modify the parameters of a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to modify parameters
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyParameters(instanceId string, args *ModifyParametersArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/parameter").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// GetBackupList - get backup list of the instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//
// RETURNS:
//   - *GetBackupListResult: result of the backup list
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupList(instanceId string) (*GetBackupListResult, error) {

	result := &GetBackupListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackupUrlWithInstanceId(instanceId)).
		WithResult(result).
		Do()

	return result, err
}

// GetBackupPolicy - get backup policy of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//
// RETURNS:
//   - *BackupPolicyResult: result of the backup policy
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupPolicy(instanceId string) (*BackupPolicyResult, error) {
	result := &BackupPolicyResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackupUrlWithInstanceId(instanceId) + "/policy").
		WithResult(result).
		Do()

	return result, err
}

// CreateBackup - create a manual backup of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to create manual backup
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateBackup(instanceId string, args *BackupCommentArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBackupUrlWithInstanceId(instanceId)).
		WithBody(args).
		Do()
}

// DeleteBackup - delete a manual backup of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - batchId: id of the manual backup batch
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteBackup(instanceId, batchId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getBackupUrlWithInstanceId(instanceId) + "/" + batchId).
		Do()
}

// ModifyBackupComment - modify comment of a backup
//
// PARAMS:
//   - instanceId: id of the instance
//   - batchId: id of the backup batch
//   - args: the arguments to modify backup comment
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupComment(instanceId, batchId string, args *BackupCommentArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBackupUrlWithInstanceId(instanceId) + "/" + batchId + "/comment").
		WithBody(args).
		Do()
}

// GetBackupUsage - get backup usage of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//
// RETURNS:
//   - *BackupUsageResult: result of backup usage
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupUsage(instanceId string) (*BackupUsageResult, error) {
	result := &BackupUsageResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackupUrlWithInstanceId(instanceId) + "/usage").
		WithResult(result).
		Do()

	return result, err
}

// GetBackupDetail - get backup detail of the instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - backupRecordId: the backup record id
//
// RETURNS:
//   - *GetBackupDetailResult: result of the backup detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupDetail(instanceId, backupRecordId string) (*GetBackupDetailResult, error) {

	result := &GetBackupDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V1 + "/" + instanceId + "/backup/" + backupRecordId + "/url").
		WithResult(result).
		Do()

	return result, err
}

// ModifyBackupPolicy - modify the BackupPolicy of a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to Modify BackupPolicy
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupPolicy(instanceId string, args *ModifyBackupPolicyArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/modifyBackupPolicy").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// ListSecurityGroupByVpcId - list security groups by vpc id
//
// PARAMS:
//   - vpcId: id of vpc
//
// RETURNS:
//   - *[]SecurityGroup:security groups of vpc
//   - error: nil if success otherwise the specific error
func (c *Client) ListSecurityGroupByVpcId(vpcId string) (*ListVpcSecurityGroupsResult, error) {
	if len(vpcId) < 1 {
		return nil, fmt.Errorf("unset vpcId")
	}
	result := &ListVpcSecurityGroupsResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSecurityGroupWithVpcIdUrl(vpcId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// ListSecurityGroupByInstanceId - list security groups by instance id
//
// PARAMS:
//   - instanceId: id of instance
//
// RETURNS:
//   - *ListSecurityGroupResult: list secrity groups result of instance
//   - error: nil if success otherwise the specific error
func (c *Client) ListSecurityGroupByInstanceId(instanceId string) (*ListSecurityGroupResult, error) {
	if len(instanceId) < 1 {
		return nil, fmt.Errorf("unset instanceId")
	}
	result := &ListSecurityGroupResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSecurityGroupWithInstanceIdUrl(instanceId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// BindSecurityGroups - bind SecurityGroup to instances
//
// PARAMS:
//   - args: http request body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) BindSecurityGroups(args *SecurityGroupArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.InstanceIds) < 1 {
		return fmt.Errorf("unset instanceIds")
	}
	if len(args.SecurityGroupIds) < 1 {
		return fmt.Errorf("unset securityGroupIds")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindSecurityGroupWithUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// UnBindSecurityGroups - unbind SecurityGroup to instances
//
// PARAMS:
//   - args: http request body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UnBindSecurityGroups(args *UnbindSecurityGroupArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.InstanceId) < 1 {
		return fmt.Errorf("unset instanceId")
	}
	if len(args.SecurityGroupIds) < 1 {
		return fmt.Errorf("unset securityGroupIds")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnBindSecurityGroupWithUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// ReplaceSecurityGroups - replace SecurityGroup to instances
//
// PARAMS:
//   - args: http request body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ReplaceSecurityGroups(args *SecurityGroupArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.InstanceIds) < 1 {
		return fmt.Errorf("unset instanceIds")
	}
	if len(args.SecurityGroupIds) < 1 {
		return fmt.Errorf("unset securityGroupIds")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getReplaceSecurityGroupWithUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// ListRecycleInstances - list all instances in recycler with marker
//
// PARAMS:
//   - marker: marker page
//
// RETURNS:
//   - *RecyclerInstanceList: the result of instances in recycler
//   - error: nil if success otherwise the specific error
func (c *Client) ListRecycleInstances(marker *Marker) (*RecyclerInstanceList, error) {
	result := &RecyclerInstanceList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithQueryParams(getMarkerParams(marker)).
		WithURL(getRecyclerUrl()).
		WithResult(result).
		Do()

	return result, err
}

// RecoverRecyclerInstances - batch recover instances that in recycler
//
// PARAMS:
//   - instanceIds: instanceId list to recover
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RecoverRecyclerInstances(instanceIds []string) error {
	if instanceIds == nil || len(instanceIds) < 1 {
		return fmt.Errorf("unset instanceIds")
	}
	if len(instanceIds) > 10 {
		return fmt.Errorf("the instanceIds length max value is 10")
	}

	args := &BatchInstanceIds{
		InstanceIds: instanceIds,
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRecyclerRecoverUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// DeleteRecyclerInstances - batch delete instances that in recycler
//
// PARAMS:
//   - instanceIds: instanceId list to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteRecyclerInstances(instanceIds []string) error {
	if instanceIds == nil || len(instanceIds) < 1 {
		return fmt.Errorf("unset instanceIds")
	}
	if len(instanceIds) > 10 {
		return fmt.Errorf("the instanceIds length max value is 10")
	}

	args := &BatchInstanceIds{
		InstanceIds: instanceIds,
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRecyclerDeleteUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// RenewInstances - batch renew instances
//
// PARAMS:
//   - args: renew instanceIds and duration
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RenewInstances(args *RenewInstanceArgs) (*OrderIdResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}
	if args.InstanceIds == nil || len(args.InstanceIds) < 1 {
		return nil, fmt.Errorf("unset instanceIds")
	}
	if len(args.InstanceIds) > 10 {
		return nil, fmt.Errorf("the instanceIds length max value is 10")
	}
	result := &OrderIdResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRenewUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// ListLogByInstanceId - list error or slow logs of instance
//
// PARAMS:
//   - instanceId: id of instance
//
// RETURNS:
//   - *[]Log:logs of instance
//   - error: nil if success otherwise the specific error
func (c *Client) ListLogByInstanceId(instanceId string, args *ListLogArgs) (*ListLogResult, error) {
	if len(instanceId) < 1 {
		return nil, fmt.Errorf("unset instanceId")
	}
	if args == nil {
		return nil, fmt.Errorf("unset list log args")
	}
	params, err2 := getQueryParams(args)
	if err2 != nil {
		return nil, err2
	}
	result := &ListLogResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getLogsUrlWithInstanceId(instanceId)).
		WithQueryParams(params).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// GetLogById - get log's detail of instance
//
// PARAMS:
//   - instanceId: id of instance
//
// RETURNS:
//   - *Log:log's detail of instance
//   - error: nil if success otherwise the specific error
func (c *Client) GetLogById(instanceId, logId string, args *GetLogArgs) (*LogItem, error) {
	if len(instanceId) < 1 {
		return nil, fmt.Errorf("unset instanceId")
	}
	if len(logId) < 1 {
		return nil, fmt.Errorf("unset logId")
	}
	if args == nil {
		return nil, fmt.Errorf("unset get log args")
	}

	result := &LogItem{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getLogsUrlWithLogId(instanceId, logId)).
		WithQueryParam("validSeconds", strconv.Itoa(args.ValidSeconds)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// GetMaintainTime - get maintainTime of instance
//
// PARAMS:
//   - instanceId: id of instance
//
// RETURNS:
//   - *GetMaintainTimeResult:maintainTime of instance
//   - error: nil if success otherwise the specific error
func (c *Client) GetMaintainTime(instanceId string) (*GetMaintainTimeResult, error) {
	if len(instanceId) < 1 {
		return nil, fmt.Errorf("unset instanceId")
	}

	result := &GetMaintainTimeResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceUrlWithId(instanceId)+"/timeWindow").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// ModifyMaintainTime - modify MaintainTime of instance
//
// PARAMS:
//   - args: new maintainTime
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyMaintainTime(instanceId string, args *MaintainTime) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(instanceId) < 1 {
		return fmt.Errorf("unset instanceIds")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceUrlWithId(instanceId)+"/timeWindow").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	if err != nil {
		return err
	}
	return nil
}

// GroupPreCheck - group preCheck
//
// PARAMS:
//   - args: the argumetns to group preCheck
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GroupPreCheckResult: the result of group preCheck
func (c *Client) GroupPreCheck(args *GroupPreCheckArgs) (*GroupPreCheckResult, error) {
	result := &GroupPreCheckResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupUrl()+"/check").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// CreateGroup - create group
//
// PARAMS:
//   - args: the argumetns to create group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *CreateGroupResult: the result of create group
func (c *Client) CreateGroup(args *CreateGroupArgs) (*CreateGroupResult, error) {
	result := &CreateGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupUrl()+"/create").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// GetGroupList - get group list
//
// PARAMS:
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GroupListResult: the result of group list
func (c *Client) GetGroupList(args *GetGroupListArgs) (*GroupListResult, error) {
	result := &GroupListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupUrl()+"/list").
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// GetGroupDetail - get group detail
//
// PARAMS:
//   - groupId: the group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GroupDetailResult: the result of group detail
func (c *Client) GetGroupDetail(groupId string) (*GroupDetailResult, error) {
	result := &GroupDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGroupUrl()+"/"+groupId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// DeleteGroup - delete group
//
// PARAMS:
//   - groupId: the group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteGroup(groupId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getGroupUrl()+"/"+groupId+"/release").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// GroupAddFollower - add follower to a group
//
// PARAMS:
//   - groupId: the group id
//   - args: the arguments to add follower
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GroupAddFollower(groupId string, args *FollowerInfo) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupUrl()+"/"+groupId+"/join").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// GroupRemoveFollower - remove follower to a group
//
// PARAMS:
//   - groupId: the group id
//   - instanceId: the instance id which to remove
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GroupRemoveFollower(groupId, instanceId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupUrl()+"/"+groupId+"/quit/"+instanceId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// SetAsLeader - set instance as leader
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SetAsLeader(groupId, instanceId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupUrl() + "/" + groupId + "/setAsLeader/" + instanceId).
		Do()
}

// UpdateGroupName - update group name
//
// PARAMS:
//   - groupId: the group id
//   - args: the arguments to update group name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateGroupName(groupId string, args *GroupNameArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGroupUrl()+"/"+groupId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// GroupForbidWrite - forbid write permission
//
// PARAMS:
//   - groupId: the group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GroupForbidWrite(groupId string, args *ForbidWriteArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGroupUrl()+"/"+groupId+"/forbidWrite").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// GroupSetQps - set group qps
// PARAMS:
//   - groupId: the group id
//   - args: the arguments to set group qps
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GroupSetQps(groupId string, args *GroupSetQpsArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGroupUrl()+"/"+groupId+"/qps").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// GroupSyncStatus - get group sync status
//
// PARAMS:
//   - groupId: the group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GroupSyncStatusResult: the result of  group sync status
func (c *Client) GroupSyncStatus(groupId string) (*GroupSyncStatusResult, error) {
	result := &GroupSyncStatusResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGroupUrl()+"/"+groupId+"/syncStatus").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// GroupWhiteList - get group white list
//
// PARAMS:
//   - groupId: the group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GroupWhiteListResult: the result of  group sync status
func (c *Client) GroupWhiteList(groupId string) (*GroupWhiteList, error) {
	result := &GroupWhiteList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGroupUrl()+"/"+groupId+"/white_list").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// GroupWhiteListAdd - add group white list
//
// PARAMS:
//   - groupId: the group id
//   - args: the arguments to add  group white list
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GroupWhiteListAdd(groupId string, args *GroupWhiteList) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGroupUrl()+"/"+groupId+"/white_list/add").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// GroupWhiteListDelete - delete group white list
//
// PARAMS:
//   - groupId: the group id
//   - args: the arguments to delete  group white list
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GroupWhiteListDelete(groupId string, args *GroupWhiteList) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGroupUrl()+"/"+groupId+"/white_list/delete").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// GroupStaleReadable - set group follower stale readable
//
// PARAMS:
//   - groupId: the group id
//   - args: the arguments to set group follower stale readable
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GroupStaleReadable(groupId string, args *StaleReadableArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGroupUrl()+"/"+groupId+"/stale_readable").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// CreateParamsTemplate - create params template
//
// PARAMS:
//   - args: the arguments to create params template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *CreateParamsTemplateResult: the result of  create params template
func (c *Client) CreateParamsTemplate(args *CreateTemplateArgs) (*CreateParamsTemplateResult, error) {
	result := &CreateParamsTemplateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTemplateUrl()+"/create").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// GetParamsTemplateList - get params template list
//
// PARAMS:
//   - marker: pagination marker
//   - maxkeys  : max keys
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *ParamsTemplateListResult: the result of get params template list
func (c *Client) GetParamsTemplateList(marker *Marker) (*ParamsTemplateListResult, error) {
	result := &ParamsTemplateListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getTemplateUrl()+"/list").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParams(getMarkerParams(marker)).
		WithResult(result).
		Do()
	return result, err
}

// GetParamsTemplateDetail - get params template detail
//
// PARAMS:
//   - templateId: the template id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *ParamsTemplateDetailResult: the result of get params template detail
func (c *Client) GetParamsTemplateDetail(templateId string) (*ResultItem, error) {
	result := &ResultItem{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getTemplateUrl()+"/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// DeleteParamsTemplate - delete params template
//
// PARAMS:
//   - templateId: the template id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteParamsTemplate(templateId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getTemplateUrl()+"/delete/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// RenameParamsTemplate - rename params template
//
// PARAMS:
//   - templateId: the template id
//   - args : new template name args
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RenameParamsTemplate(templateId string, args *RenameTemplateArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getTemplateUrl()+"/rename/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// ApplyParamsTemplate - apply params template
//
// PARAMS:
//   - templateId: the template id
//   - args : the args to apply params template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ApplyParamsTemplate(templateId string, args *ApplyTemplateArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTemplateUrl()+"/apply/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// TemplateAddParams - add params to template
//
// PARAMS:
//   - templateId: the template id
//   - args : the args to add params template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) TemplateAddParams(templateId string, args *AddParamsArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTemplateUrl()+"/addParams/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// TemplateModifyParams - modify params to template
//
// PARAMS:
//   - templateId: the template id
//   - args : the args to modify params template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) TemplateModifyParams(templateId string, args *ModifyParamsArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTemplateUrl()+"/modifyParams/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// TemplateDeleteParams - delete params to template
//
// PARAMS:
//   - templateId: the template id
//   - args : the args to delete params template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) TemplateDeleteParams(templateId string, args *DeleteParamsArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTemplateUrl()+"/deleteParams/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// GetSystemTemplate - get system template
//
// PARAMS:
//   - args : the args to get system params template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *SystemTemptaleResult: the result of get system template
func (c *Client) GetSystemTemplate(args *GetSystemTemplateArgs) (*SystemTemplateResult, error) {
	result := &SystemTemplateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getTemplateUrl()+"/system").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithQueryParamFilter("engine", args.Engine).
		WithQueryParamFilter("engineVersion", args.EngineVersion).
		WithQueryParamFilter("clusterType", args.ClusterType).
		WithResult(result).
		Do()
	return result, err
}

// GetApplyRecords - get template apply records
//
// PARAMS:
//   - args : the args to get  template apply records
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GetApplyRecordsResult: the result of get template apply records
func (c *Client) GetApplyRecords(templateId string, marker *Marker) (*GetApplyRecordsResult, error) {
	result := &GetApplyRecordsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getTemplateUrl()+"/record/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParams(getMarkerParams(marker)).
		WithResult(result).
		Do()
	return result, err
}

// SyncGroupPreCheck - pre-check for creating a sync group
//
// PARAMS:
//   - args: the arguments to pre-check sync group
//
// RETURNS:
//   - *GroupPreCheckResult: result of pre-check
//   - error: nil if success otherwise the specific error
func (c *Client) SyncGroupPreCheck(args *SyncGroupCheckRequest) (*SyncGroupPreCheckResult, error) {
	result := &SyncGroupPreCheckResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSyncGroupUrl()+"/check").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// CreateSyncGroup - create a sync group
//
// PARAMS:
//   - args: the arguments to create sync group
//
// RETURNS:
//   - *SyncGroupCreateResult: result of create sync group
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSyncGroup(args *SyncGroupCreateRequest) (*SyncGroupCreateResult, error) {
	result := &SyncGroupCreateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSyncGroupUrl()+"/create").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// ListSyncGroup - list sync groups
//
// PARAMS:
//   - page: page number, default 1
//   - pageSize: page size, default 10
//
// RETURNS:
//   - *GroupListResult: result of sync group list
//   - error: nil if success otherwise the specific error
func (c *Client) ListSyncGroup(page, pageSize int) (*SyncGroupListResult, error) {
	result := &SyncGroupListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSyncGroupUrl()+"/list").
		WithQueryParamFilter("page", strconv.Itoa(page)).
		WithQueryParamFilter("pageSize", strconv.Itoa(pageSize)).
		WithResult(result).
		Do()
	return result, err
}

// GetSyncGroupDetail - get sync group detail
//
// PARAMS:
//   - groupId: the sync group id
//
// RETURNS:
//   - *SyncGroupDetailResult: result of sync group detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSyncGroupDetail(groupId string) (*SyncGroupDetailResult, error) {
	result := &SyncGroupDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSyncGroupUrl() + "/" + groupId).
		WithResult(result).
		Do()
	return result, err
}

// DeleteSyncGroup - delete a sync group
//
// PARAMS:
//   - groupId: the sync group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteSyncGroup(groupId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getSyncGroupUrl() + "/" + groupId).
		Do()
}

// RemoveSyncGroupCluster - remove a cluster from sync group
//
// PARAMS:
//   - groupId: the sync group id
//   - args: the arguments of the cluster member to remove
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RemoveSyncGroupCluster(groupId string, args *SyncGroupMember) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSyncGroupUrl() + "/" + groupId + "/removeCluster").
		WithBody(args).
		Do()
}

// AddSyncGroupCluster - add a cluster to sync group
//
// PARAMS:
//   - groupId: the sync group id
//   - args: the arguments of the cluster member to add
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AddSyncGroupCluster(groupId string, args *SyncGroupMember) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSyncGroupUrl() + "/" + groupId + "/addCluster").
		WithBody(args).
		Do()
}

// ModifySyncGroupName - modify sync group name
//
// PARAMS:
//   - groupId: the sync group id
//   - args: the arguments to modify sync group name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifySyncGroupName(groupId string, args *SyncGroupRenameArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSyncGroupUrl() + "/" + groupId).
		WithBody(args).
		Do()
}

// GetSyncGroupStatus - get sync group sync status
//
// PARAMS:
//   - groupId: the sync group id
//
// RETURNS:
//   - *SyncGroupSyncStatusResult: result of sync group status
//   - error: nil if success otherwise the specific error
func (c *Client) GetSyncGroupStatus(groupId string) (*SyncGroupSyncStatusResult, error) {
	result := &SyncGroupSyncStatusResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSyncGroupUrl() + "/" + groupId + "/syncStatus").
		WithResult(result).
		Do()
	return result, err
}

// GetSyncGroupDelayInfo - get sync group delay info
//
// PARAMS:
//   - groupId: the sync group id
//
// RETURNS:
//   - *SyncGroupDelayInfoResult: result of sync group delay info
//   - error: nil if success otherwise the specific error
func (c *Client) GetSyncGroupDelayInfo(groupId string) (*SyncGroupDelayInfoResult, error) {
	result := &SyncGroupDelayInfoResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSyncGroupUrl() + "/" + groupId + "/delayInfo").
		WithResult(result).
		Do()
	return result, err
}

// ModifySyncGroupBnsGroup - modify sync group bns group
//
// PARAMS:
//   - groupId: the sync group id
//   - args: the arguments to modify sync group bns group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifySyncGroupBnsGroup(groupId string, args *SyncGroupBnsGroupArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSyncGroupUrl() + "/" + groupId + "/modifyBnsGroup").
		WithBody(args).
		Do()
}

// ListAclUsers - list ACL users of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//
// RETURNS:
//   - *AclUserListResult: result of the ACL user list
//   - error: nil if success otherwise the specific error
func (c *Client) ListAclUsers(instanceId string) (*AclUserListResult, error) {
	result := &AclUserListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V2 + "/" + instanceId + "/aclUserActions").
		WithResult(result).
		Do()
	return result, err
}

// CreateAclUser - create an ACL user of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to create ACL user
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAclUser(instanceId string, args *AclUserCreateArgs) error {
	if args == nil {
		return fmt.Errorf("please set create acl user argments")
	}
	if len(args.ClientAuth) != 0 {
		cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.ClientAuth)
		if err != nil {
			return err
		}
		args.ClientAuth = cryptedPass
	}
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2 + "/" + instanceId + "/aclUserActions").
		WithBody(args).
		Do()
}

// DeleteAclUser - delete an ACL user of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to delete ACL user
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAclUser(instanceId string, args *AclUserDeleteArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2 + "/" + instanceId + "/aclUserActions/delete").
		WithBody(args).
		Do()
}

// SetAclUserAuthority - set authority of an ACL user of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to set ACL user authority
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SetAclUserAuthority(instanceId string, args *AclUserCreateArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2 + "/" + instanceId + "/aclUserActions/authority").
		WithBody(args).
		Do()
}

// ModifyAclUserPassword - modify password of an ACL user of a specified instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to modify ACL user password
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyAclUserPassword(instanceId string, args *AclUserPasswordArgs) error {
	if args == nil {
		return fmt.Errorf("please set modify acl user password argments")
	}
	if len(args.ClientAuth) != 0 {
		cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.ClientAuth)
		if err != nil {
			return err
		}
		args.ClientAuth = cryptedPass
	}
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2 + "/" + instanceId + "/aclUserActions/modifyPasswd").
		WithBody(args).
		Do()
}

// ToPrepay - convert postpaid instances to prepaid
//
// PARAMS:
//   - args: the arguments to convert postpaid instances to prepaid
//
// RETURNS:
//   - *ToPrepayResult: result containing the prepaid order id
//   - error: nil if success otherwise the specific error
func (c *Client) ToPrepay(args *ToPrepayArgs) (*ToPrepayResult, error) {
	result := &ToPrepayResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V1+"/toPrepay").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// ToPostpay - convert prepaid instances to postpaid
//
// PARAMS:
//   - args: the arguments to convert prepaid instances to postpaid
//
// RETURNS:
//   - *ToPrepayResult: result containing order ids joined by commas
//   - error: nil if success otherwise the specific error
func (c *Client) ToPostpay(args *ToPostpayArgs) (*ToPrepayResult, error) {
	result := &ToPrepayResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V1+"/toPostpay").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// CancelToPostpay - cancel prepaid to postpaid conversion
//
// PARAMS:
//   - args: the arguments to cancel prepaid to postpaid conversion
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CancelToPostpay(args *ToPostpayArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V1+"/cancelToPostpay").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// SwitchMasterSlave - switch master and slave of specified shards
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments containing shards to switch
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SwitchMasterSlave(instanceId string, args *SwitchMasterSlaveArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/switchMasterSlave").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// MigrateAvailabilityZone - migrate the instance to another availability zone / subnet
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments containing the full replication info
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) MigrateAvailabilityZone(instanceId string, args *MigrateAvailabilityZoneArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/azoneMigration").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ModifyEntrance - modify the internal access IP of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments controlling whether to defer the change to maintain window
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyEntrance(instanceId string, args *ModifyEntranceArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/azoneMigration/modifyEntrance").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpgradeVersion - upgrade the engine major version of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments containing target kernel version and execution time
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpgradeVersion(instanceId string, args *UpgradeVersionArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/upgradeVersion").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpgradeProxy - upgrade or restart the proxy of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments containing proxy list, upgrade type and execution time
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpgradeProxy(instanceId string, args *UpgradeProxyArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/upgradeProxy").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// GetAutoScalingConfig - get memory auto scaling config of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//
// RETURNS:
//   - *AutoScalingConfigResult: result of memory auto scaling config
//   - error: nil if success otherwise the specific error
func (c *Client) GetAutoScalingConfig(instanceId string) (*AutoScalingConfigResult, error) {
	result := &AutoScalingConfigResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V1 + "/" + instanceId + "/autoScalingConfig").
		WithResult(result).
		Do()
	return result, err
}

// SetAutoScalingConfig - set memory auto scaling config of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments containing memory auto scaling config
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SetAutoScalingConfig(instanceId string, args *AutoScalingConfigResult) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/autoScalingConfig").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteAutoScalingConfig - delete memory auto scaling config of the instance
//
// PARAMS:
//   - instanceId: id of the instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAutoScalingConfig(instanceId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V1+"/"+instanceId+"/deleteAutoScalingConfig").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
}
