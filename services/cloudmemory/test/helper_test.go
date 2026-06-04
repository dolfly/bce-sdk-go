// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package cloudmemory_test

import (
	"context"
	"net"
	"os"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/services/cloudmemory/api"
)

const (
	defaultEndpoint = "http://127.0.0.1:8888"
	defaultBankID   = "bce-sdk-it-bank"
)

// envOr returns os.Getenv(key) or fallback when unset/empty.
func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// endpoint returns the cloudmemory endpoint under test.
func endpoint() string { return envOr("CLOUDMEMORY_ENDPOINT", defaultEndpoint) }

// apiKey returns the bearer token (may be empty for local servers).
func apiKey() string { return os.Getenv("CLOUDMEMORY_API_KEY") }

// bankID returns the bank id used by the integration tests.
func bankID() string { return envOr("CLOUDMEMORY_BANK_ID", defaultBankID) }

// newClient builds a cloudmemory.Client and skips the test if the server
// is not reachable.
func newClient(t *testing.T) *cloudmemory.Client {
	t.Helper()
	addr := "127.0.0.1:8888"
	if v := os.Getenv("CLOUDMEMORY_ADDR"); v != "" {
		addr = v
	}
	conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
	if err != nil {
		t.Skipf("cloudmemory server not reachable at %s: %v", addr, err)
	}
	_ = conn.Close()
	return cloudmemory.NewWithTimeout(endpoint(), apiKey(), 2*time.Minute)
}

// ctx returns a 2min context for a single test.
func ctx(t *testing.T) context.Context {
	t.Helper()
	c, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	t.Cleanup(cancel)
	return c
}
