// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package cloudmemory

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// VersionInfo is a forward/backward compatible representation of GET /version.
// Only Version is guaranteed; Features carries all feature flags returned by
// the server as a generic map so new flags do not break decoding.
type VersionInfo struct {
	APIVersion string                 `json:"api_version"`
	Features   map[string]interface{} `json:"features"`
}

// Health checks server liveness via GET /health.
// Returns the JSON-decoded body (interface{} per upstream).
func (c *Client) Health(ctx context.Context) (interface{}, error) {
	out, _, err := c.hindsight.MonitoringAPI.HealthEndpointHealthGet(ctx).Execute()
	return out, err
}

// Version returns the server version via GET /version (requires auth).
//
// This implementation decodes into a permissive map so it works against any
// server version regardless of which feature flags it returns. The bundled
// hindsight client uses strict decoding (DisallowUnknownFields + required
// field validation) which fails when client/server schemas drift.
func (c *Client) Version(ctx context.Context) (*VersionInfo, error) {
	cfg := c.hindsight.GetConfig()

	base, err := cfg.ServerURLWithContext(ctx, "MonitoringAPIService.GetVersion")
	if err != nil {
		return nil, err
	}
	url := strings.TrimRight(base, "/") + "/version"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	for k, v := range cfg.DefaultHeader {
		req.Header.Set(k, v)
	}
	if cfg.UserAgent != "" {
		req.Header.Set("User-Agent", cfg.UserAgent)
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("version: unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var out VersionInfo
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("version: decode response: %w", err)
	}
	return &out, nil
}
