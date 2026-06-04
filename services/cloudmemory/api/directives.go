// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package cloudmemory

import (
	"context"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ListDirectivesOptions are optional filters for ListDirectives.
type ListDirectivesOptions struct {
	Tags       []string
	TagsMatch  string
	ActiveOnly bool
	Limit      int32
	Offset     int32
}

// ListDirectives lists directives via GET /banks/{bankId}/directives.
func (c *Client) ListDirectives(ctx context.Context, bankID string, opts ListDirectivesOptions) (*hindsight.DirectiveListResponse, error) {
	call := c.hindsight.DirectivesAPI.ListDirectives(ctx, bankID)
	if len(opts.Tags) > 0 {
		call = call.Tags(opts.Tags)
	}
	if opts.TagsMatch != "" {
		call = call.TagsMatch(opts.TagsMatch)
	}
	if opts.ActiveOnly {
		call = call.ActiveOnly(true)
	}
	if opts.Limit != 0 {
		call = call.Limit(opts.Limit)
	}
	if opts.Offset != 0 {
		call = call.Offset(opts.Offset)
	}
	out, _, err := call.Execute()
	return out, err
}

// CreateDirective creates a directive via POST /banks/{bankId}/directives.
func (c *Client) CreateDirective(ctx context.Context, bankID string, req hindsight.CreateDirectiveRequest) (*hindsight.DirectiveResponse, error) {
	out, _, err := c.hindsight.DirectivesAPI.CreateDirective(ctx, bankID).CreateDirectiveRequest(req).Execute()
	return out, err
}

// GetDirective fetches a directive via GET /banks/{bankId}/directives/{directiveId}.
func (c *Client) GetDirective(ctx context.Context, bankID, directiveID string) (*hindsight.DirectiveResponse, error) {
	out, _, err := c.hindsight.DirectivesAPI.GetDirective(ctx, bankID, directiveID).Execute()
	return out, err
}

// UpdateDirective patches a directive via PATCH /banks/{bankId}/directives/{directiveId}.
func (c *Client) UpdateDirective(ctx context.Context, bankID, directiveID string, req hindsight.UpdateDirectiveRequest) (*hindsight.DirectiveResponse, error) {
	out, _, err := c.hindsight.DirectivesAPI.UpdateDirective(ctx, bankID, directiveID).UpdateDirectiveRequest(req).Execute()
	return out, err
}

// DeleteDirective deletes a directive via DELETE /banks/{bankId}/directives/{directiveId}.
func (c *Client) DeleteDirective(ctx context.Context, bankID, directiveID string) (interface{}, error) {
	out, _, err := c.hindsight.DirectivesAPI.DeleteDirective(ctx, bankID, directiveID).Execute()
	return out, err
}
