// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package cloudmemory

import "context"

// Health checks server liveness via GET /health.
// Returns the JSON-decoded body (interface{} per upstream).
func (c *Client) Health(ctx context.Context) (interface{}, error) {
	out, _, err := c.hindsight.MonitoringAPI.HealthEndpointHealthGet(ctx).Execute()
	return out, err
}

// Version returns the server version via GET /version (requires auth).
func (c *Client) Version(ctx context.Context) (interface{}, error) {
	out, _, err := c.hindsight.MonitoringAPI.GetVersion(ctx).Execute()
	if err != nil {
		return nil, err
	}
	return out, nil
}
