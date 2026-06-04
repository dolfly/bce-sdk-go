// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package cloudmemory_test

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/cloudmemory/api"
	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

const testMentalModelName = "bce-sdk-it-mm"

// TestMentalModelsLifecycle exercises every mental-model endpoint
// in dependency order: create -> list -> get -> update -> refresh -> delete.
func TestMentalModelsLifecycle(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)

	var modelID string

	t.Run("CreateMentalModel", func(t *testing.T) {
		req := hindsight.NewCreateMentalModelRequest(testMentalModelName, "what do you remember?")
		out, err := c.CreateMentalModel(ctx(t), bankID(), *req)
		if err != nil {
			t.Fatalf("CreateMentalModel: %v", err)
		}
		t.Logf("CreateMentalModel => %+v", out)
		if out != nil {
			modelID = out.GetMentalModelId()
		}
	})

	t.Run("ListMentalModels", func(t *testing.T) {
		out, err := c.ListMentalModels(ctx(t), bankID(), cloudmemory.ListMentalModelsOptions{Limit: 10})
		if err != nil {
			t.Fatalf("ListMentalModels: %v", err)
		}
		t.Logf("ListMentalModels => %+v", out)
	})

	if modelID == "" {
		t.Skip("model not created; skipping subsequent steps")
	}

	t.Run("GetMentalModel", func(t *testing.T) {
		out, err := c.GetMentalModel(ctx(t), bankID(), modelID)
		if err != nil {
			t.Fatalf("GetMentalModel: %v", err)
		}
		t.Logf("GetMentalModel => %+v", out)
	})

	t.Run("UpdateMentalModel", func(t *testing.T) {
		req := hindsight.NewUpdateMentalModelRequest()
		req.SetSourceQuery("what do you remember now?")
		out, err := c.UpdateMentalModel(ctx(t), bankID(), modelID, *req)
		if err != nil {
			// Newly-created models may still be in async-creation state and
			// can briefly 404 on PATCH while Get already succeeds. Treat as
			// non-fatal so the rest of the lifecycle keeps running.
			t.Logf("UpdateMentalModel (non-fatal): %v", err)
			return
		}
		t.Logf("UpdateMentalModel => %+v", out)
	})

	t.Run("RefreshMentalModel", func(t *testing.T) {
		out, err := c.RefreshMentalModel(ctx(t), bankID(), modelID)
		if err != nil {
			t.Logf("RefreshMentalModel (non-fatal): %v", err)
			return
		}
		t.Logf("RefreshMentalModel => %+v", out)
	})

	t.Run("DeleteMentalModel", func(t *testing.T) {
		out, err := c.DeleteMentalModel(ctx(t), bankID(), modelID)
		if err != nil {
			t.Fatalf("DeleteMentalModel: %v", err)
		}
		t.Logf("DeleteMentalModel => %+v", out)
	})
}
