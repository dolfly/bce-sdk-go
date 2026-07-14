package api

import (
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
)

// CacheTtl defined cache rule for a site
type CacheTtl struct {
	Value          string `json:"value"`
	Weight         int    `json:"weight"`
	OverrideOrigin bool   `json:"override_origin"`
	Ttl            int    `json:"ttl"`
	Type           string `json:"type"`
}

// CacheCodeTtl defined status-code cache rule for a site.
type CacheCodeTtl struct {
	Value          string `json:"value"`
	Weight         int    `json:"weight"`
	OverrideOrigin bool   `json:"overrideOrigin"`
	Ttl            int    `json:"ttl"`
	Type           string `json:"type"`
}

// CacheKey defined the query string strategy for cache key.
type CacheKey struct {
	Query       *bool     `json:"query,omitempty"`
	IncludeArgs *[]string `json:"include_args,omitempty"`
	ExcludeArgs *[]string `json:"exclude_args,omitempty"`
}

// HSTS defined the HSTS strategy for a site.
type HSTS struct {
	MaxAge            *int  `json:"maxAge,omitempty"`
	IncludeSubDomains *bool `json:"includeSubDomains,omitempty"`
	Preload           *bool `json:"preload,omitempty"`
}

// HTTP3 defined the HTTP3 strategy for a site.
type HTTP3 struct {
	Enable *bool `json:"enable,omitempty"`
}

// HttpHeader defined HTTP header rule.
type HttpHeader struct {
	Type   string `json:"type"`
	Header string `json:"header"`
	Value  string `json:"value"`
	Action string `json:"action"`
}

// PageRule defined a rule-engine entry of a site.
type PageRule struct {
	Name   string      `json:"name"`
	Status string      `json:"status"`
	Rules  [][]Rule    `json:"rules"`
	Config *RuleConfig `json:"config,omitempty"`
}

// Rule defined a single match condition inside PageRule.Rules.
// Outer slice means OR; inner slice means AND.
type Rule struct {
	MatchFrom  string   `json:"matchFrom"`
	Operator   string   `json:"operator"`
	MatchKey   string   `json:"matchKey,omitempty"`
	Values     []string `json:"values"`
	IgnoreCase *bool    `json:"ignoreCase,omitempty"`
}

// RuleCacheKey defined the query string strategy for cache key in rule scope.
// It has an extra `ignore_case` field compared to the site-scope CacheKey.
type RuleCacheKey struct {
	Query       *bool     `json:"query,omitempty"`
	IncludeArgs *[]string `json:"include_args,omitempty"`
	ExcludeArgs *[]string `json:"exclude_args,omitempty"`
	IgnoreCase  *bool     `json:"ignore_case,omitempty"`
}

// RefreshRevalidate defined the refresh-revalidate strategy in rule scope.
type RefreshRevalidate struct {
	Enabled *string `json:"enabled,omitempty"`
}

// WebSocket defined the WebSocket strategy in rule scope.
type WebSocket struct {
	Enabled *string `json:"enabled,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`
}

// OriginTimeout defined the origin timeout strategy in rule scope.
type OriginTimeout struct {
	LoadTimeout    *int `json:"loadTimeout,omitempty"`
	ConnectTimeout *int `json:"connectTimeout,omitempty"`
}

// OriginRedirectOptions defined the origin redirect strategy in rule scope.
type OriginRedirectOptions struct {
	EnableRedirectFollow   *string `json:"enableRedirectFollow,omitempty"`
	MaxRedirectFollowCount *int    `json:"maxRedirectFollowCount,omitempty"`
}

// OriginOptions defined origin range/part-size strategy in rule scope.
type OriginOptions struct {
	Range    *string `json:"range,omitempty"`
	PartSize *int    `json:"partSize,omitempty"`
}

// RealIp defined the real-IP strategy in rule scope.
type RealIp struct {
	Enabled *string `json:"enabled,omitempty"`
	Name    *string `json:"name,omitempty"`
}

// ErrorPage defined a single custom error-page entry in rule scope.
type ErrorPage struct {
	Code *string `json:"code,omitempty"`
	Url  *string `json:"url,omitempty"`
}

// TrafficLimit defined the traffic-limit strategy in rule scope.
type TrafficLimit struct {
	Enable         *string `json:"enable,omitempty"`
	LimitRate      *int    `json:"limitRate,omitempty"`
	LimitStartHour *int    `json:"limitStartHour,omitempty"`
	LimitEndHour   *int    `json:"limitEndHour,omitempty"`
	LimitRateUnit  *string `json:"limitRateUnit,omitempty"`
}

// AntiHotLink defined the anti-hot-link strategy in rule scope.
type AntiHotLink struct {
	AntiType        *string `json:"antiType,omitempty"`
	SecretKey       *string `json:"secretKey,omitempty"`
	NewsecretKey    *string `json:"newsecretKey,omitempty"`
	Timeout         *int    `json:"timeout,omitempty"`
	TimestampFormat *string `json:"timestampFormat,omitempty"`
	AuthArg         *string `json:"authArg,omitempty"`
}

// OriginArg defined the origin-args strategy in rule scope.
type OriginArg struct {
	Ignore *string   `json:"ignore,omitempty"`
	Args   *[]string `json:"args,omitempty"`
}

// UrlRules defined a single URL-rewrite/redirect rule in rule scope.
type UrlRules struct {
	Scheme  *string `json:"scheme,omitempty"`
	Host    *string `json:"host,omitempty"`
	SrcPath *string `json:"srcPath,omitempty"`
	DstPath *string `json:"dstPath,omitempty"`
	Query   *string `json:"query,omitempty"`
	Status  *string `json:"status,omitempty"`
}

// RuleConfig defined the per-rule configurations that can be applied when a PageRule matches.
type RuleConfig struct {
	CacheTtl              *[]CacheTtl            `json:"cacheTtl,omitempty"`
	CacheKey              *RuleCacheKey          `json:"cacheKey,omitempty"`
	OfflineMode           *string                `json:"offlineMode,omitempty"`
	RefreshRevalidate     *RefreshRevalidate     `json:"refreshRevalidate,omitempty"`
	HttpToHttpsEnabled    *string                `json:"httpToHttpsEnabled,omitempty"`
	HttpToHttpsCode       *string                `json:"httpToHttpsCode,omitempty"`
	Hsts                  *HSTS                  `json:"hsts,omitempty"`
	Isa                   *string                `json:"isa,omitempty"`
	Http2Disable          *string                `json:"http2Disable,omitempty"`
	Http3                 *HTTP3                 `json:"http3,omitempty"`
	WebSocket             *WebSocket             `json:"webSocket,omitempty"`
	Http2Origin           *string                `json:"http2Origin,omitempty"`
	ClientMaxBodySize     *string                `json:"clientMaxBodySize,omitempty"`
	Compress              *string                `json:"compress,omitempty"`
	CompressMethodArray   *[]string              `json:"compressMethodArray,omitempty"`
	HttpHeader            *[]HttpHeader          `json:"httpHeader,omitempty"`
	OriginTimeout         *OriginTimeout         `json:"originTimeout,omitempty"`
	OriginRedirectOptions *OriginRedirectOptions `json:"originRedirectOptions,omitempty"`
	OriginOptions         *OriginOptions         `json:"originOptions,omitempty"`
	RealIp                *RealIp                `json:"realIp,omitempty"`
	ErrorPage             *[]ErrorPage           `json:"errorPage,omitempty"`
	TrafficLimit          *TrafficLimit          `json:"trafficLimit,omitempty"`
	AntiHotLink           *AntiHotLink           `json:"antiHotLink,omitempty"`
	OriginArg             *OriginArg             `json:"originArg,omitempty"`
	UrlRules              *[]UrlRules            `json:"urlRules,omitempty"`
}

// SiteConfig defined a unified container for site configurations.
// All fields use pointer types so that callers can distinguish "not set" (nil, will be omitted)
// from "explicitly set to empty" (non-nil empty value, will be sent as `[]` / `{}`).
// New configuration items should also be added as pointer fields with `omitempty`.
type SiteConfig struct {
	CacheTtl            *[]CacheTtl     `json:"cacheTtl,omitempty"`
	CacheKey            *CacheKey       `json:"cacheKey,omitempty"`
	OfflineMode         *string         `json:"offlineMode,omitempty"`
	HttpToHttpsEnabled  *string         `json:"httpToHttpsEnabled,omitempty"`
	HttpToHttpsCode     *string         `json:"httpToHttpsCode,omitempty"`
	Hsts                *HSTS           `json:"hsts,omitempty"`
	Http2Disable        *string         `json:"http2Disable,omitempty"`
	Http3               *HTTP3          `json:"http3,omitempty"`
	ClientMaxBodySize   *string         `json:"clientMaxBodySize,omitempty"`
	Compress            *string         `json:"compress,omitempty"`
	CompressMethodArray *[]string       `json:"compressMethodArray,omitempty"`
	Isa                 *string         `json:"isa,omitempty"`
	HttpHeader          *[]HttpHeader   `json:"httpHeader,omitempty"`
	CacheCodeTtl        *[]CacheCodeTtl `json:"cacheCodeTtl,omitempty"`
	GrpcOrigin          *string         `json:"grpcOrigin,omitempty"`
	Http2Origin         *string         `json:"http2Origin,omitempty"`
	PageRules           *[]PageRule     `json:"pageRules,omitempty"`
}

// SiteConfigUpdateResult defined the response of SetSiteConfig
type SiteConfigUpdateResult struct {
	Status string `json:"status"`
}

// SetSiteConfig - set the site-level configurations
// For details, please refer to https://cloud.baidu.com/doc/GEO/s/vmiia4s0j
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site to be configured
//   - config: the configurations to be set;
//
// RETURNS:
//   - *SiteConfigUpdateResult: the update status returned by the server
//   - error: nil if success otherwise the specific error
func SetSiteConfig(cli bce.Client, site string, config *SiteConfig) (*SiteConfigUpdateResult, error) {
	if site == "" {
		return nil, errors.New("site is required")
	}
	if config == nil {
		return nil, errors.New("config is required")
	}

	respObj := &SiteConfigUpdateResult{}
	err := httpRequest(cli, "PUT", "/v2/geo/site/"+site+"/config", nil, config, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// GetSiteConfig - get the site configurations
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site to be queried
//
// RETURNS:
//   - *SiteConfig: the configurations of the site
//   - error: nil if success otherwise the specific error
func GetSiteConfig(cli bce.Client, site string) (*SiteConfig, error) {
	if site == "" {
		return nil, errors.New("site is required")
	}

	respObj := &SiteConfig{}
	err := httpRequest(cli, "GET", "/v2/geo/site/"+site+"/config", nil, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}
