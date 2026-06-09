// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package cloudmemory_test

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/cloudmemory/api"
)

// TestDirectivesLifecycle exercises every directive endpoint
// in dependency order: create -> list -> get -> update -> delete.
func TestDirectivesLifecycle(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)

	var directiveID string

	t.Run("CreateDirective", func(t *testing.T) {
		req := cloudmemory.NewCreateDirectiveRequest("bce-sdk-it-directive", "always be concise")
		out, err := c.CreateDirective(ctx(t), bankID(), *req)
		if err != nil {
			t.Fatalf("CreateDirective: %v", err)
		}
		t.Logf("CreateDirective => %+v", out)
		if out != nil {
			directiveID = out.Id
		}
	})

	t.Run("ListDirectives", func(t *testing.T) {
		out, err := c.ListDirectives(ctx(t), bankID(), cloudmemory.ListDirectivesOptions{Limit: 10})
		if err != nil {
			t.Fatalf("ListDirectives: %v", err)
		}
		t.Logf("ListDirectives => %+v", out)
	})

	if directiveID == "" {
		t.Skip("directive not created; skipping subsequent steps")
	}

	t.Run("GetDirective", func(t *testing.T) {
		out, err := c.GetDirective(ctx(t), bankID(), directiveID)
		if err != nil {
			t.Fatalf("GetDirective: %v", err)
		}
		t.Logf("GetDirective => %+v", out)
	})

	t.Run("UpdateDirective", func(t *testing.T) {
		req := cloudmemory.NewUpdateDirectiveRequest()
		out, err := c.UpdateDirective(ctx(t), bankID(), directiveID, *req)
		if err != nil {
			t.Fatalf("UpdateDirective: %v", err)
		}
		t.Logf("UpdateDirective => %+v", out)
	})

	t.Run("DeleteDirective", func(t *testing.T) {
		out, err := c.DeleteDirective(ctx(t), bankID(), directiveID)
		if err != nil {
			t.Fatalf("DeleteDirective: %v", err)
		}
		t.Logf("DeleteDirective => %+v", out)
	})
}
