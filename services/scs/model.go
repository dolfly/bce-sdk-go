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

// model.go - definitions of the request arguments and results data structure model

package scs

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type Billing struct {
	PaymentTiming string       `json:"paymentTiming"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}

type Subnet struct {
	Name     string `json:"name,omitempty"`
	SubnetID string `json:"subnetId"`
	ZoneName string `json:"zoneName"`
	Cidr     string `json:"cidr,omitempty"`
	VpcID    string `json:"vpcId,omitempty"`
}

type CreateInstanceArgs struct {
	Billing       Billing `json:"billing"`
	PurchaseCount int     `json:"purchaseCount"`
	InstanceName  string  `json:"instanceName"`
	NodeType      string  `json:"nodeType"`
	ShardNum      int     `json:"shardNum"`
	ProxyNum      int     `json:"proxyNum"`
	ClusterType   string  `json:"clusterType"`
	// Deprecated: 该字段已废弃，请使用 ReplicationInfo 替代
	ReplicationNum  int           `json:"replicationNum"`
	ReplicationInfo []Replication `json:"replicationInfo"`
	Port            int           `json:"port"`
	Engine          int           `json:"engine,omitempty"`
	EngineVersion   string        `json:"engineVersion"`
	DiskFlavor      int           `json:"diskFlavor,omitempty"`
	DiskType        string        `json:"diskType,omitempty"`
	VpcID           string        `json:"vpcId"`
	// Deprecated: 该字段已废弃，请使用 ReplicationInfo 替代
	Subnets           []Subnet                  `json:"subnets,omitempty"`
	AutoRenewTimeUnit string                    `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int                       `json:"autoRenewTime,omitempty"`
	BgwGroupId        string                    `json:"bgwGroupId,omitempty"`
	ClientToken       string                    `json:"-"`
	ClientAuth        string                    `json:"clientAuth"`
	ResourceGroupId   string                    `json:"resourceGroupId"`
	ConfTpl           string                    `json:"confTpl,omitempty"`
	StoreType         int                       `json:"storeType"`
	EnableReadOnly    int                       `json:"enableReadOnly,omitempty"`
	Tags              []model.TagModel          `json:"tags"`
	BcmInstanceGroups []BcmInstanceGroupRequest `json:"bcmInstanceGroups"`
	AutoBackupConfig  string                    `json:"autoBackupConfig,omitempty"`
	DeployIdList      []string                  `json:"deployIdList,omitempty"`
}

type CreateInstanceResult struct {
	InstanceIds []string `json:"instanceIds"`
	OrderId     string   `json:"orderId"`
}

type InstanceModel struct {
	InstanceID         string           `json:"instanceId"`
	InstanceName       string           `json:"instanceName"`
	InstanceStatus     string           `json:"instanceStatus"`
	InstanceExpireTime string           `json:"instanceExpireTime"`
	ShardNum           int              `json:"shardNum"`
	ReplicationNum     int              `json:"replicationNum"`
	ClusterType        string           `json:"clusterType"`
	Engine             string           `json:"engine"`
	EngineVersion      string           `json:"engineVersion"`
	VnetIP             string           `json:"vnetIp"`
	Domain             string           `json:"domain"`
	Port               int              `json:"port"`
	InstanceCreateTime string           `json:"instanceCreateTime"`
	Capacity           float64          `json:"capacity"`
	UsedCapacity       float64          `json:"usedCapacity"`
	PaymentTiming      string           `json:"paymentTiming"`
	ZoneNames          []string         `json:"zoneNames"`
	Tags               []model.TagModel `json:"tags"`
	ResourceGroupId    string           `json:"resourceGroupId"`
	ResourceGroupName  string           `json:"resourceGroupName"`
	DeployIdList       []string         `json:"deployIdList"`
	OrderStatus        string           `json:"orderStatus"`
	NodeType           string           `json:"nodeType"`
	DiskFlavor         int              `json:"diskFlavor"`
	DiskType           string           `json:"diskType"`
	StoreType          int              `json:"storeType"`
	Eip                string           `json:"eip"`
	Entrance           string           `json:"entrance"`
}

type ListInstancesArgs struct {
	Marker      string
	MaxKeys     int
	InstanceIds string
	VnetIp      string
}

type ListInstancesResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Instances   []InstanceModel `json:"instances"`
}

type ResizeInstanceArgs struct {
	NodeType        string        `json:"nodeType"`
	ShardNum        int           `json:"shardNum"`
	IsDefer         bool          `json:"isDefer"`
	ClientToken     string        `json:"-"`
	DiskFlavor      int           `json:"diskFlavor"`
	DiskType        string        `json:"diskType"`
	ModifyMethod    string        `json:"modifyMethod"`
	ReplicationInfo []Replication `json:"replicationInfo"`
}

type ReplicationArgs struct {
	ResizeType      string        `json:"resizeType"`
	ReplicationInfo []Replication `json:"replicationInfo"`
	ClientToken     string        `json:"-"`
}

type Replication struct {
	AvailabilityZone string `json:"availabilityZone"`
	SubnetId         string `json:"subnetId"`
	IsMaster         int    `json:"isMaster"`
}

type BcmInstanceGroupRequest struct {
	GroupIds          string `json:"groupIds"`
	GroupResourceType string `json:"groupResourceType"`
}

type RestartInstanceArgs struct {
	IsDefer bool `json:"isDefer"`
}

type GetInstanceDetailResult struct {
	InstanceID              string                  `json:"instanceId"`
	InstanceName            string                  `json:"instanceName"`
	InstanceStatus          string                  `json:"instanceStatus"`
	ClusterType             string                  `json:"clusterType"`
	Engine                  string                  `json:"engine"`
	EngineVersion           string                  `json:"engineVersion"`
	VnetIP                  string                  `json:"vnetIp"`
	Domain                  string                  `json:"domain"`
	Port                    int                     `json:"port"`
	InstanceCreateTime      string                  `json:"instanceCreateTime"`
	InstanceExpireTime      string                  `json:"instanceExpireTime"`
	Capacity                float64                 `json:"capacity"`
	UsedCapacity            float64                 `json:"usedCapacity"`
	PaymentTiming           string                  `json:"paymentTiming"`
	VpcID                   string                  `json:"vpcId"`
	ZoneNames               []string                `json:"zoneNames"`
	Subnets                 []Subnet                `json:"subnets"`
	AutoRenew               string                  `json:"autoRenew"`
	Tags                    []model.TagModel        `json:"tags"`
	ShardNum                int                     `json:"shardNum"`
	ReplicationNum          int                     `json:"replicationNum"`
	NodeType                string                  `json:"nodeType"`
	DiskFlavor              int                     `json:"diskFlavor"`
	DiskType                string                  `json:"diskType"`
	StoreType               int                     `json:"storeType"`
	Eip                     string                  `json:"eip"`
	PublicDomain            string                  `json:"publicDomain"`
	EnableReadOnly          int                     `json:"enableReadOnly"`
	ReplicationInfo         []Replication           `json:"replicationInfo"`
	BnsGroup                string                  `json:"bnsGroup"`
	ResourceGroupId         string                  `json:"resourceGroupId"`
	ResourceGroupName       string                  `json:"resourceGroupName"`
	DeployIdList            []string                `json:"deployIdList"`
	OrderStatus             string                  `json:"orderStatus"`
	Entrance                string                  `json:"entrance"`
	EnableSlowLog           int                     `json:"enableSlowLog"`
	RedisList               []RedisNode             `json:"redisList"`
	ProxyList               []ProxyItem             `json:"proxyList"`
	CacheClusterInstances   []CacheClusterNode      `json:"cacheClusterInstances"`
	FullVersionInfo         InstanceFullVersionInfo `json:"fullVersionInfo"`
	MaintainTime            MaintainTime            `json:"maintainTime"`
	EnableHotkey            bool                    `json:"enableHotkey"`
	FeatureSwitches         FeatureSwitches         `json:"featureSwitches"`
	CrossAzNearest          string                  `json:"crossAzNearest"`
	EntranceList            []EntranceItem          `json:"entranceList"`
	SupportSentinelCommands bool                    `json:"supportSentinelCommands"`
}

type UpdateInstanceNameArgs struct {
	InstanceName string `json:"instanceName"`
	ClientToken  string `json:"-"`
}

type NodeType struct {
	InstanceFlavor          float64 `json:"instanceFlavor"`
	NodeType                string  `json:"nodeType"`
	CPUNum                  int     `json:"cpuNum"`
	NetworkThroughputInGbps float64 `json:"networkThroughputInGbps"`
	PeakQPS                 int     `json:"peakQps"`
	MaxConnections          int     `json:"maxConnections"`
	AllowedNodeNumList      []int   `json:"allowedNodeNumList"`
	MinDiskFlavor           int     `json:"minDiskFlavor"`
	MaxDiskFlavor           int     `json:"maxDiskFlavor"`
}

type GetNodeTypeListResult struct {
	ClusterNodeTypeList     []NodeType `json:"clusterNodeTypeList"`
	DefaultNodeTypeList     []NodeType `json:"defaultNodeTypeList"`
	HsdbNodeTypeList        []NodeType `json:"hsdbNodeTypeList"`
	PegaClusterNodeTypeList []NodeType `json:"pegaClusterNodeTypeList"`
}

type ListSubnetsArgs struct {
	VpcID    string `json:"vpcId"`
	ZoneName string `json:"zoneName"`
}

type ListSubnetsResult struct {
	SubnetOriginals []SubnetOriginal `json:"subnets"`
}

type SubnetOriginal struct {
	Name     string `json:"name"`
	SubnetID string `json:"subnetId"`
	ZoneName string `json:"zoneName"`
	Cidr     string `json:"cidr"`
	VpcID    string `json:"vpcId"`
}

type UpdateInstanceDomainNameArgs struct {
	Domain      string `json:"domain"`
	ClientToken string `json:"-"`
}

type GetZoneListResult struct {
	Zones []ZoneNames `json:"zones"`
}

type ZoneNames struct {
	ZoneNames []string `json:"zoneNames"`
}

type FlushInstanceArgs struct {
	Password       string `json:"password"`
	ClientToken    string `json:"-"`
	DbIndex        int    `json:"dbIndex"`
	IsFlushExpired bool   `json:"isFlushExpired"`
	IsDefer        bool   `json:"isDefer"`
}

type BindingTagArgs struct {
	ChangeTags []model.TagModel `json:"changeTags"`
}

type GetSecurityIpResult struct {
	SecurityIps []string `json:"securityIps"`
}

type SecurityIpArgs struct {
	SecurityIps []string `json:"securityIps"`
	ClientToken string   `json:"-"`
}

type WhiteListGroupResult struct {
	ClusterIPGroups []WhiteListGroup `json:"clusterIPGroups"`
}

type WhiteListGroup struct {
	GroupName string   `json:"groupName"`
	IPList    []string `json:"ipList"`
}

type WhiteListGroupArgs struct {
	GroupName     string   `json:"groupName"`
	NewGroupName  string   `json:"newGroupName,omitempty"`
	ClusterIPList []string `json:"clusterIpList"`
}

type ModifyPasswordArgs struct {
	Password    string `json:"password"`
	ClientToken string `json:"-"`
}

type GetParametersResult struct {
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Default      string `json:"default"`
	ForceRestart string `json:"forceRestart"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	Type         int    `json:"type"`
	Range        string `json:"range"`
	Module       int    `json:"module"`
	Desc         string `json:"desc"`
}

type ModifyParametersArgs struct {
	Parameter   InstanceParam `json:"parameter"`
	ClientToken string        `json:"-"`
}

type InstanceParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetBackupListResult struct {
	TotalCount int          `json:"totalCount"`
	Backups    []BackupInfo `json:"backups"`
}

type BackupInfo struct {
	BackupType  string         `json:"backupType"`
	Comment     string         `json:"comment"`
	StartTime   string         `json:"startTime"`
	Recoverable string         `json:"recoverable"`
	BatchID     string         `json:"batchId"`
	Records     []BackupRecord `json:"records"`
}

type BackupRecord struct {
	BackupRecordId string `json:"backupRecordId"`
	BackupStatus   string `json:"backupStatus"`
	Duration       int64  `json:"duration"`
	ObjectSize     int64  `json:"objectSize"`
	ShardName      string `json:"shardName"`
	StartTime      string `json:"startTime"`
	BackupType     string `json:"backupType"`
	Comment        string `json:"comment"`
}

type ModifyBackupPolicyArgs struct {
	BackupDays           string `json:"backupDays"`
	BackupTime           string `json:"backupTime"`
	ClientToken          string `json:"clientToken"`
	ExpireDay            int    `json:"expireDay"`
	IsEncrypt            string `json:"isEncrypt"`
	IsEnableBackupLog    string `json:"isEnableBackupLog"`
	HighFreqBackupEnable bool   `json:"highFreqBackupEnable"`
	HighFreqBackupConfig string `json:"highFreqBackupConfig"`
}

type BackupPolicyResult struct {
	BackupTime           string   `json:"backupTime"`
	BackupDays           string   `json:"backupDays"`
	ExpireDay            int      `json:"expireDay"`
	IsEncrypt            string   `json:"isEncrypt"`
	IsEnableBackupLog    string   `json:"isEnableBackupLog"`
	HighFreqBackupEnable bool     `json:"highFreqBackupEnable"`
	HighFreqBackupConfig string   `json:"highFreqBackupConfig"`
	HighFreqList         []string `json:"highFreqList"`
}

type BackupCommentArgs struct {
	Comment string `json:"comment"`
}

type AclUserListResult struct {
	Success bool          `json:"success"`
	Result  []AclUserItem `json:"result"`
}

type AclUserItem struct {
	UserName     string `json:"userName"`
	UpdateStatus int    `json:"updateStatus"`
	Extra        string `json:"extra"`
	UserType     int    `json:"userType"`
}

type AclUserCreateArgs struct {
	UserName   string `json:"userName"`
	ClientAuth string `json:"clientAuth"`
	Extra      string `json:"extra,omitempty"`
	UserType   int    `json:"userType"`
}

type AclUserDeleteArgs struct {
	UserName string `json:"userName"`
}

type AclUserPasswordArgs struct {
	UserName   string `json:"userName"`
	ClientAuth string `json:"clientAuth"`
}

type ToPrepayArgs struct {
	Duration    int      `json:"duration"`
	InstanceIds []string `json:"instanceIds"`
}

type ToPostpayArgs struct {
	InstanceIds []string `json:"instanceIds"`
}

type SwitchMasterSlaveShard struct {
	HashName   string `json:"hashName"`
	NodeShowID string `json:"nodeShowId"`
}

type SwitchMasterSlaveArgs struct {
	Shards []SwitchMasterSlaveShard `json:"shards"`
}

type MigrateAvailabilityZoneArgs struct {
	IsDefer         bool          `json:"isDefer"`
	ReplicationInfo []Replication `json:"replicationInfo"`
}

type ModifyEntranceArgs struct {
	IsDefer bool `json:"isDefer"`
}

type UpgradeVersionArgs struct {
	KernelVersion string `json:"kernelVersion,omitempty"`
	IsDefer       bool   `json:"isDefer,omitempty"`
}

type UpgradeProxyArgs struct {
	ProxyList   []string `json:"proxyList,omitempty"`
	UpgradeType string   `json:"upgradeType"`
	IsDefer     bool     `json:"isDefer,omitempty"`
}

type MemAutoScalingConfig struct {
	MemUsageUpperThreshold        int    `json:"memUsageUpperThreshold"`
	MemUsageDownThreshold         int    `json:"memUsageDownThreshold"`
	MaxNodeType                   string `json:"maxNodeType"`
	MinNodeType                   string `json:"minNodeType"`
	ObservationWindowSizeForUpper string `json:"observationWindowSizeForUpper"`
	ObservationWindowSizeForDown  string `json:"observationWindowSizeForDown"`
}

type AutoScalingConfigResult struct {
	MemSpec *MemAutoScalingConfig `json:"memSpec"`
}

type ToPrepayResult struct {
	OrderId string `json:"orderId"`
}

type BackupUsageResult struct {
	LogicalLogBackupBillingSizeBytes   int64  `json:"logicalLogBackupBillingSizeBytes"`
	SnapshotDataBackupSizeBytes        int64  `json:"snapshotDataBackupSizeBytes"`
	PhysicalDataBackupSizeBytes        int64  `json:"physicalDataBackupSizeBytes"`
	LogicalLogBackupSizeBytes          int64  `json:"logicalLogBackupSizeBytes"`
	LogicalDataBackupSizeBytes         int64  `json:"logicalDataBackupSizeBytes"`
	PhysicalDataBackupBillingSizeBytes int64  `json:"physicalDataBackupBillingSizeBytes"`
	LogicalDataBackupBillingSizeBytes  int64  `json:"logicalDataBackupBillingSizeBytes"`
	PhysicalLogBackupSizeBytes         int64  `json:"physicalLogBackupSizeBytes"`
	PhysicalLogBackupBillingSizeBytes  int64  `json:"physicalLogBackupBillingSizeBytes"`
	SnapshotDataBackupBillingSizeBytes int64  `json:"snapshotDataBackupBillingSizeBytes"`
	DataType                           string `json:"dataType"`
	ClusterID                          string `json:"clusterID"`
	AppID                              string `json:"appID"`
}

type ListVpcSecurityGroupsResult struct {
	Groups []SecurityGroup `json:"groups"`
}

type SecurityGroup struct {
	Name                 string `json:"name"`
	SecurityGroupID      string `json:"securityGroupId"`
	Description          string `json:"description"`
	TenantID             string `json:"tenantId"`
	AssociateNum         int    `json:"associateNum"`
	VpcID                string `json:"vpcId"`
	VpcShortID           string `json:"vpcShortId"`
	VpcName              string `json:"vpcName"`
	CreatedTime          string `json:"createdTime"`
	Version              int    `json:"version"`
	DefaultSecurityGroup int    `json:"defaultSecurityGroup"`
}

type SecurityGroupArgs struct {
	InstanceIds      []string `json:"instanceIds"`
	SecurityGroupIds []string `json:"securityGroupIds"`
}

type UnbindSecurityGroupArgs struct {
	InstanceId       string   `json:"instanceId"`
	SecurityGroupIds []string `json:"securityGroupIds"`
}

type ListSecurityGroupResult struct {
	Groups      []SecurityGroupDetail `json:"groups"`
	ActiveRules []SecurityGroupRule   `json:"activeRules"`
}

type SecurityGroupRule struct {
	PortRange           string `json:"portRange"`
	Protocol            string `json:"protocol"`
	RemoteGroupID       string `json:"remoteGroupId"`
	RemoteGroupName     string `json:"remoteGroupName"`
	RemoteIP            string `json:"remoteIP"`
	Ethertype           string `json:"ethertype"`
	TenantID            string `json:"tenantId"`
	Name                string `json:"name"`
	ID                  string `json:"id"`
	SecurityGroupRuleID string `json:"securityGroupRuleId"`
	SecurityGroupID     string `json:"securityGroupId"`
	SecurityGroupUuid   string `json:"securityGroupUuid"`
	Direction           string `json:"direction"`
}

type SecurityGroupDetail struct {
	SecurityGroupName   string              `json:"securityGroupName"`
	SecurityGroupID     string              `json:"securityGroupId"`
	SecurityGroupUuid   string              `json:"securityGroupUuid"`
	SecurityGroupRemark string              `json:"securityGroupRemark"`
	Inbound             []SecurityGroupRule `json:"inbound"`
	Outbound            []SecurityGroupRule `json:"outbound"`
	VpcName             string              `json:"vpcName"`
	VpcID               string              `json:"vpcId"`
	ProjectID           string              `json:"projectId"`
}

type RequestBuilder struct {
}

type GetPriceRequest struct {
	Engine         string `json:"engine,omitempty"`
	ShardNum       int    `json:"shardNum,omitempty"`
	Period         int    `json:"period,omitempty"`
	ChargeType     string `json:"chargeType,omitempty"`
	NodeType       string `json:"nodeType,omitempty"`
	ReplicationNum int    `json:"replicationNum,omitempty"`
	ClusterType    string `json:"clusterType,omitempty"`
}

type GetPriceResult struct {
	Price float64 `json:"price,omitempty"`
}

type Marker struct {
	Marker  string `json:"marker,omitempty"`
	MaxKeys int    `json:"maxKeys,omitempty"`
}

type ListResultWithMarker struct {
	IsTruncated bool   `json:"isTruncated"`
	Marker      string `json:"marker"`
	MaxKeys     int    `json:"maxKeys"`
	NextMarker  string `json:"nextMarker"`
}

type RecycleInstance struct {
	InstanceID         string           `json:"cacheClusterShowId"`
	InstanceName       string           `json:"instanceName"`
	InstanceStatus     string           `json:"instanceStatus"`
	IsolatedStatus     string           `json:"isolatedStatus"`
	ClusterType        string           `json:"clusterType"`
	Engine             string           `json:"engine"`
	EngineVersion      string           `json:"engineVersion"`
	VnetIP             string           `json:"vnetIp"`
	Domain             string           `json:"domain"`
	Port               string           `json:"port"`
	InstanceCreateTime string           `json:"instanceCreateTime"`
	Capacity           float64          `json:"capacity"`
	UsedCapacity       float64          `json:"usedCapacity"`
	PaymentTiming      string           `json:"paymentTiming"`
	ZoneNames          []string         `json:"zoneNames"`
	Tags               []model.TagModel `json:"tags"`
}

type RecyclerInstanceList struct {
	ListResultWithMarker
	Result []RecycleInstance `json:"result"`
}

type BatchInstanceIds struct {
	InstanceIds []string `json:"cacheClusterShowIds,omitempty"`
}

type RenewInstanceArgs struct {
	Duration    int      `json:"duration,omitempty"`
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type OrderIdResult struct {
	OrderId string `json:"orderId"`
}

type ListLogArgs struct {
	FileType  string `json:"fileType"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type ListLogResult struct {
	LogList []ShardLog `json:"logList"`
}
type LogItem struct {
	LogStartTime    string `json:"logStartTime"`
	LogEndTime      string `json:"logEndTime"`
	DownloadURL     string `json:"downloadUrl"`
	LogID           string `json:"logId"`
	LogSizeInBytes  int    `json:"logSizeInBytes"`
	DownloadExpires string `json:"downloadExpires"`
}
type ShardLog struct {
	ShardShowID string    `json:"shardShowId"`
	TotalNum    int       `json:"totalNum"`
	ShardID     int       `json:"shardId"`
	LogItem     []LogItem `json:"logItem"`
}

type GetLogArgs struct {
	ValidSeconds int `json:"validSeconds"`
}
type RedisNode struct {
	UUID              string `json:"uuid"`
	NodeShowID        string `json:"nodeShowId"`
	CacheInstanceType int    `json:"cacheInstanceType"`
	IsReadOnly        int    `json:"isReadOnly"`
	InGroup           int    `json:"inGroup"`
	AvailabilityZone  string `json:"availabilityZone"`
	SubnetID          string `json:"subnetId"`
	Status            int    `json:"status"`
	Weight            int    `json:"weight"`
	HashName          string `json:"hashName"`
	ShardID           int    `json:"shardId"`
	NodeID            int    `json:"nodeId"`
}

type ProxyItem struct {
	UUID             string `json:"uuid"`
	NodeShowID       string `json:"nodeShowId"`
	AvailabilityZone string `json:"availabilityZone"`
	NodeID           string `json:"nodeId"`
}

type CacheClusterNode struct {
	InstanceID string `json:"instanceId"`
	FlavorInGB string `json:"flavorInGB"`
	HashName   string `json:"hashName"`
	Domain     string `json:"domain"`
	CreateTime string `json:"createTime"`
	ShardID    string `json:"shardId"`
}

type InstanceFullVersionInfo struct {
	ProxyFullVersion             string `json:"proxyFullVersion"`
	RedisOrPegaFullVerison       string `json:"redisOrPegaFullVerison"`
	ProxyLatestFullVersion       string `json:"proxyLatestFullVersion"`
	RedisOrPegaLatestFullVersion string `json:"redisOrPegaLatestFullVersion"`
	IsProxyCanUpgrade            bool   `json:"isProxyCanUpgrade"`
	IsRedisOrPegaCanUpgrade      bool   `json:"isRedisOrPegaCanUpgrade"`
	IsPegaCanRestart             bool   `json:"isPegaCanRestart"`
}

type FeatureSwitches struct {
	GroupModify            bool `json:"groupModify"`
	CrossAzNearest         bool `json:"crossAzNearest"`
	ProxyUpgradeSupport    bool `json:"proxyUpgradeSupport"`
	RecoverInOriginSupport bool `json:"recoverInOriginSupport"`
	RecoverInNewSupport    bool `json:"recoverInNewSupport"`
	SupportSentinelSwitch  bool `json:"supportSentinelSwitch"`
	WhitelistGroupSupport  bool `json:"whitelistGroupSupport"`
	BandwidthModify        bool `json:"bandwidthModify"`
	ShadowBackupSupport    bool `json:"shadowBackupSupport"`
}

type EntranceItem struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
	Zone string `json:"zone"`
}

type GetMaintainTimeResult struct {
	CacheClusterShowId string       `json:"cacheClusterShowId"`
	MaintainTime       MaintainTime `json:"maintainTime"`
}
type MaintainTime struct {
	StartTime string `json:"startTime"`
	Duration  int    `json:"duration"`
	Period    []int  `json:"period"`
}

type CreatePriceArgs struct {
	Engine         int    `json:"engine,omitempty"`
	ClusterType    string `json:"clusterType,omitempty"`
	NodeType       string `json:"nodeType,omitempty"`
	ShardNum       int    `json:"shardNum,omitempty"`
	ReplicationNum int    `json:"replicationNum,omitempty"`
	InstanceNum    int    `json:"instanceNum,omitempty"`
	DiskType       string `json:"diskType,omitempty"`
	DiskFlavor     int    `json:"diskFlavor,omitempty"`
	ChargeType     string `json:"chargeType,omitempty"`
	Period         int    `json:"period,omitempty"`
}
type CreatePriceResult struct {
	Price        float64 `json:"price,omitempty"`
	CatalogPrice float64 `json:"catalogPrice,omitempty"`
}

type ResizePriceArgs struct {
	NodeType       string `json:"nodeType"`
	ShardNum       int    `json:"shardNum,omitempty"`
	ReplicationNum int    `json:"replicationNum,omitempty"`
	DiskFlavor     int    `json:"diskFlavor,omitempty"`
	ChargeType     string `json:"chargeType,omitempty"`
	Period         int    `json:"period,omitempty"`
	ChangeType     string `json:"changeType,omitempty"`
}
type ResizePriceResult struct {
	Price float64 `json:"price,omitempty"`
}

type SetAsSlaveArgs struct {
	MasterDomain string `json:"masterDomain"`
	MasterPort   int    `json:"masterPort"`
}

type RenameDomainArgs struct {
	Domain      string `json:"domain"`
	ClientToken string `json:"clientToken,omitempty"`
}

type SwapDomainArgs struct {
	SourceInstanceId string `json:"sourceInstanceId"`
	TargetInstanceId string `json:"targetInstanceId"`
	ClientToken      string `json:"clientToken,omitempty"`
}

type GetBackupDetailResult struct {
	Url           string `json:"url"`
	UrlExpiration int    `json:"urlExpiration"`
}

type GroupPreCheckArgs struct {
	Leader    GroupLeader     `json:"leader"`
	Followers []GroupFollower `json:"followers"`
}
type GroupLeader struct {
	LeaderRegion string `json:"leaderRegion"`
	LeaderId     string `json:"leaderId"`
}
type GroupFollower struct {
	FollowerId     string `json:"followerId"`
	FollowerRegion string `json:"followerRegion"`
}
type GroupPreCheckResult struct {
	LeaderResult      GroupLeaderResult       `json:"leaderResult"`
	FollowerResult    []GroupFollowerResult   `json:"followerResult"`
	ConnectionResults []GroupConnectionResult `json:"connectionResults"`
}
type GroupLeaderResult struct {
	Version         bool `json:"version"`
	ClusterStatus   bool `json:"clusterStatus"`
	ReplicationNum  bool `json:"replicationNum"`
	Flavor          bool `json:"flavor"`
	Joined          bool `json:"joined"`
	NoPasswd        bool `json:"noPasswd"`
	NoSecurityGroup bool `json:"noSecurityGroup"`
	IsHitX1         bool `json:"isHitX1"`
}
type GroupFollowerResult struct {
	FollowerId      string `json:"followerId"`
	NoData          bool   `json:"noData"`
	Version         bool   `json:"version"`
	EngineVersion   bool   `json:"engineVersion"`
	ClusterStatus   bool   `json:"clusterStatus"`
	ShardNum        bool   `json:"shardNum"`
	ReplicationNum  bool   `json:"replicationNum"`
	Flavor          bool   `json:"flavor"`
	Joined          bool   `json:"joined"`
	NoPasswd        bool   `json:"noPasswd"`
	NoSecurityGroup bool   `json:"noSecurityGroup"`
	IsHitX1         bool   `json:"isHitX1"`
}
type GroupConnectionResult struct {
	SourceId    string `json:"sourceId"`
	SourceRole  string `json:"sourceRole"`
	TargetId    string `json:"targetId"`
	TargetRole  string `json:"targetRole"`
	Connectable bool   `json:"connectable"`
}

type CreateGroupArgs struct {
	Leader CreateGroupLeader `json:"leader"`
}
type CreateGroupLeader struct {
	GroupName    string `json:"groupName"`
	LeaderRegion string `json:"leaderRegion"`
	LeaderId     string `json:"leaderId"`
}
type CreateGroupResult struct {
	GroupId string `json:"groupId"`
}
type GroupListResult struct {
	TotalCount int           `json:"totalCount"`
	PageNo     int           `json:"pageNo"`
	PageSize   int           `json:"pageSize"`
	Result     []GroupResult `json:"result"`
}
type GroupResult struct {
	GroupId         string `json:"groupId"`
	GroupName       string `json:"groupName"`
	GroupStatus     string `json:"groupStatus"`
	ClusterNum      int    `json:"clusterNum"`
	GroupCreateTime string `json:"groupCreateTime"`
	ForbidWrite     int    `json:"forbidWrite"`
	GroupType       string `json:"groupType"`
	LeaderName      string `json:"leaderName"`
	LeaderShowId    string `json:"leaderShowId"`
	LeaderRegion    string `json:"leaderRegion"`
}
type GetGroupListArgs struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type GroupDetailResult struct {
	GroupId         string              `json:"groupId"`
	GroupName       string              `json:"groupName"`
	GroupStatus     string              `json:"groupStatus"`
	ClusterNum      int                 `json:"clusterNum"`
	GroupCreateTime string              `json:"groupCreateTime"`
	ForbidWrite     int                 `json:"forbidWrite"`
	GroupType       string              `json:"groupType"`
	Leader          GroupLeaderInfo     `json:"leader"`
	Followers       []GroupFollowerInfo `json:"followers"`
}
type GroupLeaderInfo struct {
	ClusterName       string  `json:"clusterName"`
	ClusterShowId     string  `json:"clusterShowId"`
	Region            string  `json:"region"`
	Status            string  `json:"status"`
	TotalCapacityInGB float64 `json:"totalCapacityInGB"`
	UsedCapacityInGB  int     `json:"usedCapacityInGB"`
	ShardNum          int     `json:"shardNum"`
	Flavor            float64 `json:"flavor"`
	QpsWrite          int64   `json:"qpsWrite"`
	QpsRead           int64   `json:"qpsRead"`
	StableReadable    bool    `json:"stableReadable"`
	ForbidWrite       int     `json:"forbidWrite"`
	AvailabilityZone  string  `json:"availabilityZone"`
	ExpiredTime       string  `json:"expiredTime"`
}
type GroupFollowerInfo struct {
	ClusterName       string  `json:"clusterName"`
	ClusterShowId     string  `json:"clusterShowId"`
	Region            string  `json:"region"`
	Status            string  `json:"status"`
	TotalCapacityInGB float64 `json:"totalCapacityInGB"`
	UsedCapacityInGB  int     `json:"usedCapacityInGB"`
	ShardNum          int     `json:"shardNum"`
	Flavor            float64 `json:"flavor"`
	QpsWrite          int64   `json:"qpsWrite"`
	QpsRead           int64   `json:"qpsRead"`
	StableReadable    bool    `json:"stableReadable"`
	ForbidWrite       int     `json:"forbidWrite"`
	AvailabilityZone  string  `json:"availabilityZone"`
	ExpiredTime       string  `json:"expiredTime"`
}

type SyncGroupBaseInfo struct {
	SyncGroupShowId     string `json:"syncGroupShowId"`
	SyncGroupName       string `json:"syncGroupName"`
	Status              string `json:"status"`
	ClusterNum          int    `json:"clusterNum"`
	AvailabilityZone    string `json:"availabilityZone"`
	NodeType            string `json:"nodeType"`
	NetConn             string `json:"netConn"`
	ConfilctResolution  string `json:"confilctResolution"`
	SyncGroupCreateTime string `json:"syncGroupCreateTime"`
	SameSpec            bool   `json:"sameSpec"`
	SameShardNum        bool   `json:"sameShardNum"`
	Engine              string `json:"engine"`
	BnsGroup            string `json:"bnsGroup"`
}

type SyncGroupMember struct {
	MemberId string `json:"memberId"`
	Region   string `json:"region"`
}

type SyncGroupCheckRequest struct {
	SyncGroupShowId string            `json:"syncGroupShowId"`
	Members         []SyncGroupMember `json:"members"`
}

type SyncGroupPreCheckResult struct {
	CheckResult       []SyncGroupCheckItem      `json:"checkResult"`
	ConnectionResults []SyncGroupConnectionItem `json:"connectionResults"`
}

type SyncGroupCheckItem struct {
	MemberId        string `json:"memberId"`
	NoData          bool   `json:"noData"`
	Version         bool   `json:"version"`
	EngineVersion   bool   `json:"engineVersion"`
	ClusterStatus   bool   `json:"clusterStatus"`
	ShardNum        bool   `json:"shardNum"`
	ReplicationNum  bool   `json:"replicationNum"`
	Flavor          bool   `json:"flavor"`
	NotJoined       bool   `json:"notJoined"`
	NoSecurityGroup bool   `json:"noSecurityGroup"`
	IsHitX1         bool   `json:"isHitX1"`
	IsAppendOnlyOn  bool   `json:"isAppendOnlyOn"`
	SamePasswd      bool   `json:"samePasswd"`
}

type SyncGroupConnectionItem struct {
	SourceId    string `json:"sourceId"`
	TargetId    string `json:"targetId"`
	Connectable bool   `json:"connectable"`
}

type SyncGroupCreateRequest struct {
	SyncGroupName string            `json:"syncGroupName"`
	BnsName       string            `json:"bnsName"`
	Members       []SyncGroupMember `json:"members"`
}

type SyncGroupCreateResult struct {
	SyncGroupShowId string `json:"syncGroupShowId"`
}

type SyncGroupListResult struct {
	TotalCount int             `json:"totalCount"`
	PageNo     int             `json:"pageNo"`
	PageSize   int             `json:"pageSize"`
	Result     []SyncGroupItem `json:"result"`
}

type SyncGroupItem struct {
	SyncGroupBaseInfo
	Cluster []SyncGroupItemCluster `json:"cluster"`
}

type SyncGroupItemCluster struct {
	ClusterShowId string `json:"clusterShowId"`
	Region        string `json:"region"`
}

type SyncGroupDetailResult struct {
	SyncGroupBaseInfo
	Cluster []CacheSyncGroupInstance `json:"cluster"`
}

type CacheSyncGroupInstance struct {
	ClusterId         int64          `json:"clusterId"`
	ClusterName       string         `json:"clusterName"`
	ClusterShowId     string         `json:"clusterShowId"`
	Region            string         `json:"region"`
	ClusterStatus     string         `json:"clusterStatus"`
	ClusterEngine     string         `json:"clusterEngine"`
	CreateTime        string         `json:"createTime"`
	TotalCapacityInGb float64        `json:"totalCapacityInGb"`
	UsedCapacityInGb  float64        `json:"usedCapacityInGb"`
	ExpiredTime       string         `json:"expiredTime"`
	ShardList         []string       `json:"shardList"`
	SyncFlow          []SyncFlowItem `json:"syncFlow"`
}

type SyncFlowItem struct {
	TargetBLBIp         string `json:"targetBLBIp"`
	TargetBLBPort       int    `json:"targetBLBPort"`
	TargetClusterShowId string `json:"targetClusterShowId"`
}

type SyncGroupSyncStatusResult struct {
	SyncGroupShowId string                        `json:"syncGroupShowId"`
	SyncStatus      []SyncGroupInstanceSyncStatus `json:"syncStatus"`
}

type SyncGroupInstanceSyncStatus struct {
	MemberId string `json:"memberId"`
	Status   string `json:"status"`
}

type SyncGroupDelayInfoResult struct {
	DelayInfo []SyncGroupDelayInfoItem `json:"delayInfo"`
}

type SyncGroupDelayInfoItem struct {
	SourceCluster string `json:"sourceCluster"`
	DestCluster   string `json:"destCluster"`
	DelayResult   int64  `json:"delayResult"`
	TimeResult    int64  `json:"timeResult"`
}

type SyncGroupRenameArgs struct {
	GroupName string `json:"groupName"`
}

type SyncGroupBnsGroupArgs struct {
	BnsGroup string `json:"bnsGroup"`
}

type FollowerInfo struct {
	FollowerId     string `json:"followerId"`
	FollowerRegion string `json:"followerRegion"`
	SyncMaster     string `json:"syncMaster"`
}

type GroupNameArgs struct {
	GroupName string `json:"groupName"`
}

type ForbidWriteArgs struct {
	ForbidWriteFlag bool `json:"forbidWriteFlag"`
}

type GroupSetQpsArgs struct {
	ClusterShowId string `json:"clusterShowId"`
	QpsWrite      int    `json:"qpsWrite"`
	QpsRead       int    `json:"qpsRead"`
}

type GroupSyncStatusResult struct {
	Followers []FollowerSyncInfo `json:"followers"`
}

type FollowerSyncInfo struct {
	ClusterShowId string `json:"clusterShowId"`
	SyncStatus    string `json:"syncStatus"`
	MaxOffset     int    `json:"maxOffset"`
	Lag           int    `json:"lag"`
}

type GroupWhiteList struct {
	WhiteLists []string `json:"whiteLists"`
}

type StaleReadableArgs struct {
	FollowerId    string `json:"followerId"`
	StaleReadable bool   `json:"staleReadable"`
}

type CreateTemplateArgs struct {
	EngineVersion string          `json:"engineVersion"`
	TemplateType  int             `json:"templateType"`
	ClusterType   string          `json:"clusterType"`
	Engine        string          `json:"engine"`
	Name          string          `json:"name"`
	Comment       string          `json:"comment"`
	Parameters    []ParameterItem `json:"parameters"`
}
type ParameterItem struct {
	ConfName   string `json:"confName"`
	ConfModule int    `json:"confModule"`
	ConfValue  string `json:"confValue"`
	ConfType   int    `json:"confType"`
}
type CreateParamsTemplateResult struct {
	TemplateId     int    `json:"templateId"`
	TemplateShowId string `json:"templateShowId"`
}

type ParamsTemplateListResult struct {
	Marker      string       `json:"marker"`
	MaxKeys     int          `json:"maxKeys"`
	NextMarker  string       `json:"nextMarker"`
	IsTruncated bool         `json:"isTruncated"`
	Result      []ResultItem `json:"result"`
}
type ResultItem struct {
	EngineVersion  string      `json:"engineVersion"`
	TemplateType   int         `json:"templateType"`
	ClusterType    string      `json:"clusterType"`
	NeedReboot     int         `json:"needReboot"`
	TemplateShowId string      `json:"templateShowId"`
	UpdateTime     string      `json:"updateTime"`
	TemplateId     int         `json:"templateId"`
	ParameterNum   int         `json:"parameterNum"`
	TemplateName   string      `json:"templateName"`
	Engine         string      `json:"engine"`
	CreateTime     string      `json:"createTime"`
	Comment        string      `json:"comment"`
	Parameters     []ParamItem `json:"parameters"`
}

type ParamItem struct {
	ConfName         string `json:"confName"`
	ConfModule       int    `json:"confModule"`
	ConfCacheVersion int    `json:"confCacheVersion"`
	ConfValue        string `json:"confValue"`
	NeedReboot       int    `json:"needReboot"`
	ConfRedisVersion string `json:"confRedisVersion"`
	ConfDefault      string `json:"confDefault"`
	ConfType         int    `json:"confType"`
	ConfRange        string `json:"confRange"`
	ConfDesc         string `json:"confDesc"`
	ConfUserVisible  int    `json:"confUserVisible"`
}

type RenameTemplateArgs struct {
	Name string `json:"name"`
}

type ApplyTemplateArgs struct {
	RebootType             int                  `json:"rebootType"`
	Extra                  string               `json:"extra"`
	CacheClusterShowIdItem []CacheClusterShowId `json:"cacheClusterShowId"`
	Parameters             []ParameterItem      `json:"parameters"`
}
type CacheClusterShowId struct {
	CacheClusterShowId string `json:"cacheClusterShowId"`
	Region             string `json:"region"`
}

type AddParamsArgs struct {
	Parameters []ParameterItem `json:"parameters"`
}

type ModifyParamsArgs struct {
	Parameters []ParameterItem `json:"parameters"`
}

type DeleteParamsArgs struct {
	Parameters []string `json:"parameters"`
}

type GetSystemTemplateArgs struct {
	Engine        string `json:"engine"`
	EngineVersion string `json:"engineVersion"`
	ClusterType   string `json:"clusterType"`
}

type SystemTemplateResult struct {
	Success bool             `json:"success"`
	Result  []SystemTemplate `json:"result"`
}

type SystemTemplate struct {
	ConfName         string `json:"confName"`
	ConfDefault      string `json:"confDefault"`
	ConfValue        string `json:"confValue"`
	ConfType         int    `json:"confType"`
	ConfRange        string `json:"confRange"`
	ConfModule       int    `json:"confModule"`
	ConfDesc         string `json:"confDesc"`
	NeedReboot       int    `json:"needReboot"`
	ConfRedisVersion string `json:"confRedisVersion"`
	ConfCacheVersion int    `json:"confCacheVersion"`
}

type GetApplyRecordsResult struct {
	Marker      string        `json:"marker"`
	MaxKeys     int           `json:"maxKeys"`
	NextMarker  string        `json:"nextMarker"`
	IsTruncated bool          `json:"isTruncated"`
	Result      []ApplyRecord `json:"result"`
}

type ApplyRecord struct {
	CacheClusterShowId string `json:"cacheClusterShowId"`
	CacheClusterName   string `json:"cacheClusterName"`
	AvailabilityZone   string `json:"availabilityZone"`
	Version            int    `json:"version"`
	Status             string `json:"status"`
	Engine             string `json:"engine"`
	EngineVersion      string `json:"engineVersion"`
	ClusterType        string `json:"clusterType"`
	CreateTime         string `json:"createTime"`
	ApplyTime          string `json:"applyTime"`
}
