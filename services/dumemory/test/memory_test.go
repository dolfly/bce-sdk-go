// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package dumemory_test

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/dumemory/api"
)

// ensureBank creates the integration bank if it does not exist.
func ensureBank(t *testing.T, c *dumemory.Client) {
	t.Helper()
	if _, err := c.GetBank(ctx(t), bankID()); err == nil {
		return
	}
	req := dumemory.NewCreateBankRequest()
	if _, err := c.CreateBank(ctx(t), bankID(), *req); err != nil {
		t.Fatalf("ensureBank: %v", err)
	}
}

func TestRetain(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	items := []dumemory.MemoryItem{*dumemory.NewMemoryItem("hello from bce-sdk integration test")}
	req := dumemory.NewRetainRequest(items)
	out, err := c.Retain(ctx(t), bankID(), *req)
	if err != nil {
		t.Fatalf("Retain: %v", err)
	}
	t.Logf("Retain => %+v", out)
}

func TestRetainAsync(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	items := []dumemory.MemoryItem{*dumemory.NewMemoryItem("hello async")}
	req := dumemory.NewRetainRequest(items)
	out, err := c.RetainAsync(ctx(t), bankID(), *req)
	if err != nil {
		t.Fatalf("RetainAsync: %v", err)
	}
	t.Logf("RetainAsync => %+v", out)
}

func TestRecall(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	req := dumemory.NewRecallRequest("hello")
	out, err := c.Recall(ctx(t), bankID(), *req)
	if err != nil {
		t.Fatalf("Recall: %v", err)
	}
	t.Logf("Recall => %+v", out)
}

func TestReflect(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	req := dumemory.NewReflectRequest("what do you remember?")
	out, err := c.Reflect(ctx(t), bankID(), *req)
	if err != nil {
		t.Fatalf("Reflect: %v", err)
	}
	t.Logf("Reflect => %+v", out)
}

func TestListMemories(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	out, err := c.ListMemories(ctx(t), bankID(), dumemory.ListMemoriesOptions{Limit: 10})
	if err != nil {
		t.Fatalf("ListMemories: %v", err)
	}
	t.Logf("ListMemories => %+v", out)
}
