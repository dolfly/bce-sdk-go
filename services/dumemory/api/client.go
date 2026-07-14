// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package dumemory

import (
	"time"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// Client is the entry point of the Baidu Cloud Memory SDK. It embeds a
// *hindsight.APIClient and exposes a simplified, endpoint-table aligned
// surface. Use Underlying to access advanced builder-style options.
type Client struct {
	hindsight *hindsight.APIClient
}

// New creates a Client targeting baseURL and authenticated with apiKey.
// baseURL must include scheme, e.g. "https://cloud.memory.bj.baidubce.com/api" or
// "http://127.0.0.1:8888".
func New(baseURL, apiKey string) *Client {
	return &Client{
		hindsight: hindsight.NewAPIClientWithToken(baseURL, apiKey),
	}
}

// NewWithTimeout is like New but configures a per-request timeout.
// Pass 0 to disable the timeout.
func NewWithTimeout(baseURL, apiKey string, timeout time.Duration) *Client {
	return &Client{
		hindsight: hindsight.NewAPIClientWithTimeout(baseURL, apiKey, timeout),
	}
}

// NewFromAPIClient wraps an existing *hindsight.APIClient (advanced use).
func NewFromAPIClient(c *hindsight.APIClient) *Client {
	return &Client{hindsight: c}
}

// Underlying returns the wrapped *hindsight.APIClient. Use this when you
// need access to features that are not surfaced by this facade.
func (c *Client) Underlying() *hindsight.APIClient { return c.hindsight }
