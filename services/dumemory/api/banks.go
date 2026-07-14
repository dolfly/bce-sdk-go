// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package dumemory

import (
	"context"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ListBanks lists banks via GET /banks.
func (c *Client) ListBanks(ctx context.Context) (*hindsight.BankListResponse, error) {
	out, _, err := c.hindsight.BanksAPI.ListBanks(ctx).Execute()
	return out, err
}

// CreateBank creates (or updates) a bank via POST /banks.
// Per the upstream OpenAPI, bank creation is keyed by bankId; req carries
// the optional metadata such as name/description/config.
func (c *Client) CreateBank(ctx context.Context, bankID string, req hindsight.CreateBankRequest) (*hindsight.BankProfileResponse, error) {
	out, _, err := c.hindsight.BanksAPI.CreateOrUpdateBank(ctx, bankID).CreateBankRequest(req).Execute()
	return out, err
}

// GetBank fetches a bank profile via GET /banks/{bankId}.
func (c *Client) GetBank(ctx context.Context, bankID string) (*hindsight.BankProfileResponse, error) {
	out, _, err := c.hindsight.BanksAPI.GetBankProfile(ctx, bankID).Execute()
	return out, err
}

// DeleteBank deletes a bank via DELETE /banks/{bankId}.
func (c *Client) DeleteBank(ctx context.Context, bankID string) (*hindsight.DeleteResponse, error) {
	out, _, err := c.hindsight.BanksAPI.DeleteBank(ctx, bankID).Execute()
	return out, err
}

// GetBankConfig fetches a bank's config via GET /banks/{bankId}/config.
func (c *Client) GetBankConfig(ctx context.Context, bankID string) (*hindsight.BankConfigResponse, error) {
	out, _, err := c.hindsight.BanksAPI.GetBankConfig(ctx, bankID).Execute()
	return out, err
}

// UpdateBankConfig patches a bank's config via PATCH /banks/{bankId}/config.
func (c *Client) UpdateBankConfig(ctx context.Context, bankID string, update hindsight.BankConfigUpdate) (*hindsight.BankConfigResponse, error) {
	out, _, err := c.hindsight.BanksAPI.UpdateBankConfig(ctx, bankID).BankConfigUpdate(update).Execute()
	return out, err
}

// GetBankStats fetches stats via GET /banks/{bankId}/stats.
func (c *Client) GetBankStats(ctx context.Context, bankID string) (*hindsight.BankStatsResponse, error) {
	out, _, err := c.hindsight.BanksAPI.GetAgentStats(ctx, bankID).Execute()
	return out, err
}

// ConsolidateBank manually triggers consolidation via POST /banks/{bankId}/consolidate.
func (c *Client) ConsolidateBank(ctx context.Context, bankID string, req *hindsight.ConsolidationRequest) (*hindsight.ConsolidationResponse, error) {
	call := c.hindsight.BanksAPI.TriggerConsolidation(ctx, bankID)
	if req != nil {
		call = call.ConsolidationRequest(*req)
	}
	out, _, err := call.Execute()
	return out, err
}
