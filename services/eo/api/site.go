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

// CacheTtlCode defined status-code cache rule for a site.
type CacheTtlCode struct {
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
	CacheCodeTtl        *[]CacheTtlCode `json:"cacheCodeTtl,omitempty"`
	GrpcOrigin          *string         `json:"grpcOrigin,omitempty"`
	Http2Origin         *string         `json:"http2Origin,omitempty"`
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
