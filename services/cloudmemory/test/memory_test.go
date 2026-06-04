// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package cloudmemory_test

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/cloudmemory/api"
	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ensureBank creates the integration bank if it does not exist.
func ensureBank(t *testing.T, c *cloudmemory.Client) {
	t.Helper()
	if _, err := c.GetBank(ctx(t), bankID()); err == nil {
		return
	}
	req := hindsight.NewCreateBankRequest()
	if _, err := c.CreateBank(ctx(t), bankID(), *req); err != nil {
		t.Fatalf("ensureBank: %v", err)
	}
}

func TestRetain(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	items := []hindsight.MemoryItem{*hindsight.NewMemoryItem("hello from bce-sdk integration test")}
	req := hindsight.NewRetainRequest(items)
	out, err := c.Retain(ctx(t), bankID(), *req)
	if err != nil {
		t.Fatalf("Retain: %v", err)
	}
	t.Logf("Retain => %+v", out)
}

func TestRetainAsync(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	items := []hindsight.MemoryItem{*hindsight.NewMemoryItem("hello async")}
	req := hindsight.NewRetainRequest(items)
	out, err := c.RetainAsync(ctx(t), bankID(), *req)
	if err != nil {
		t.Fatalf("RetainAsync: %v", err)
	}
	t.Logf("RetainAsync => %+v", out)
}

func TestRecall(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	req := hindsight.NewRecallRequest("hello")
	out, err := c.Recall(ctx(t), bankID(), *req)
	if err != nil {
		t.Fatalf("Recall: %v", err)
	}
	t.Logf("Recall => %+v", out)
}

func TestReflect(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	req := hindsight.NewReflectRequest("what do you remember?")
	out, err := c.Reflect(ctx(t), bankID(), *req)
	if err != nil {
		t.Fatalf("Reflect: %v", err)
	}
	t.Logf("Reflect => %+v", out)
}

func TestListMemories(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	out, err := c.ListMemories(ctx(t), bankID(), cloudmemory.ListMemoriesOptions{Limit: 10})
	if err != nil {
		t.Fatalf("ListMemories: %v", err)
	}
	t.Logf("ListMemories => %+v", out)
}
