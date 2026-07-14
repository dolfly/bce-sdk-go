// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package dumemory

import (
	"context"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ListMentalModelsOptions are optional filters for ListMentalModels.
type ListMentalModelsOptions struct {
	Tags      []string
	TagsMatch string
	Detail    string
	Limit     int32
	Offset    int32
}

// ListMentalModels lists models via GET /banks/{bankId}/mental-models.
func (c *Client) ListMentalModels(ctx context.Context, bankID string, opts ListMentalModelsOptions) (*hindsight.MentalModelListResponse, error) {
	call := c.hindsight.MentalModelsAPI.ListMentalModels(ctx, bankID)
	if len(opts.Tags) > 0 {
		call = call.Tags(opts.Tags)
	}
	if opts.TagsMatch != "" {
		call = call.TagsMatch(opts.TagsMatch)
	}
	if opts.Detail != "" {
		call = call.Detail(opts.Detail)
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

// CreateMentalModel creates a model via POST /banks/{bankId}/mental-models.
func (c *Client) CreateMentalModel(ctx context.Context, bankID string, req hindsight.CreateMentalModelRequest) (*hindsight.CreateMentalModelResponse, error) {
	out, _, err := c.hindsight.MentalModelsAPI.CreateMentalModel(ctx, bankID).CreateMentalModelRequest(req).Execute()
	return out, err
}

// GetMentalModel fetches a model via GET /banks/{bankId}/mental-models/{modelId}.
func (c *Client) GetMentalModel(ctx context.Context, bankID, modelID string) (*hindsight.MentalModelResponse, error) {
	out, _, err := c.hindsight.MentalModelsAPI.GetMentalModel(ctx, bankID, modelID).Execute()
	return out, err
}

// UpdateMentalModel patches a model via PATCH /banks/{bankId}/mental-models/{modelId}.
func (c *Client) UpdateMentalModel(ctx context.Context, bankID, modelID string, req hindsight.UpdateMentalModelRequest) (*hindsight.MentalModelResponse, error) {
	out, _, err := c.hindsight.MentalModelsAPI.UpdateMentalModel(ctx, bankID, modelID).UpdateMentalModelRequest(req).Execute()
	return out, err
}

// DeleteMentalModel deletes a model via DELETE /banks/{bankId}/mental-models/{modelId}.
func (c *Client) DeleteMentalModel(ctx context.Context, bankID, modelID string) (interface{}, error) {
	out, _, err := c.hindsight.MentalModelsAPI.DeleteMentalModel(ctx, bankID, modelID).Execute()
	return out, err
}

// RefreshMentalModel triggers refresh via POST /banks/{bankId}/mental-models/{modelId}/refresh.
func (c *Client) RefreshMentalModel(ctx context.Context, bankID, modelID string) (*hindsight.AsyncOperationSubmitResponse, error) {
	out, _, err := c.hindsight.MentalModelsAPI.RefreshMentalModel(ctx, bankID, modelID).Execute()
	return out, err
}
