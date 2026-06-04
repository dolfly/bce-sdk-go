// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package cloudmemory

import (
	"context"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ListEntities lists entities via GET /entities.
func (c *Client) ListEntities(ctx context.Context, bankID string, limit, offset int32) (*hindsight.EntityListResponse, error) {
	call := c.hindsight.EntitiesAPI.ListEntities(ctx, bankID)
	if limit != 0 {
		call = call.Limit(limit)
	}
	if offset != 0 {
		call = call.Offset(offset)
	}
	out, _, err := call.Execute()
	return out, err
}

// EntityGraph fetches the entity graph via GET /entities/graph.
func (c *Client) EntityGraph(ctx context.Context, bankID string, limit, minCount int32) (*hindsight.EntityGraphResponse, error) {
	call := c.hindsight.EntitiesAPI.GetEntityGraph(ctx, bankID)
	if limit != 0 {
		call = call.Limit(limit)
	}
	if minCount != 0 {
		call = call.MinCount(minCount)
	}
	out, _, err := call.Execute()
	return out, err
}
