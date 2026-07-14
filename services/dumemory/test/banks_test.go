// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package dumemory_test

import (
	"github.com/baidubce/bce-sdk-go/services/dumemory/api"
	"testing"
)

// TestBanksLifecycle exercises every Bank endpoint as an ordered sub-test
// flow: create -> get -> list -> stats -> get/update config -> consolidate
// -> delete. Sub-tests inherit the bank created in step 0 via t.Run order.
func TestBanksLifecycle(t *testing.T) {
	c := newClient(t)
	id := bankID()

	t.Run("CreateBank", func(t *testing.T) {
		req := dumemory.NewCreateBankRequest()
		out, err := c.CreateBank(ctx(t), id, *req)
		if err != nil {
			t.Fatalf("CreateBank: %v", err)
		}
		t.Logf("CreateBank => %+v", out)
	})

	t.Run("GetBank", func(t *testing.T) {
		out, err := c.GetBank(ctx(t), id)
		if err != nil {
			t.Fatalf("GetBank: %v", err)
		}
		t.Logf("GetBank => %+v", out)
	})

	t.Run("ListBanks", func(t *testing.T) {
		out, err := c.ListBanks(ctx(t))
		if err != nil {
			t.Fatalf("ListBanks: %v", err)
		}
		t.Logf("ListBanks => %+v", out)
	})

	t.Run("GetBankStats", func(t *testing.T) {
		out, err := c.GetBankStats(ctx(t), id)
		if err != nil {
			t.Fatalf("GetBankStats: %v", err)
		}
		t.Logf("GetBankStats => %+v", out)
	})

	t.Run("GetBankConfig", func(t *testing.T) {
		out, err := c.GetBankConfig(ctx(t), id)
		if err != nil {
			t.Fatalf("GetBankConfig: %v", err)
		}
		t.Logf("GetBankConfig => %+v", out)
	})

	t.Run("UpdateBankConfig", func(t *testing.T) {
		update := *dumemory.NewBankConfigUpdate(map[string]interface{}{})
		out, err := c.UpdateBankConfig(ctx(t), id, update)
		if err != nil {
			t.Fatalf("UpdateBankConfig: %v", err)
		}
		t.Logf("UpdateBankConfig => %+v", out)
	})

	t.Run("ConsolidateBank", func(t *testing.T) {
		out, err := c.ConsolidateBank(ctx(t), id, nil)
		if err != nil {
			t.Logf("ConsolidateBank (non-fatal): %v", err)
			return
		}
		t.Logf("ConsolidateBank => %+v", out)
	})

	t.Run("DeleteBank", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skip delete in -short")
		}
		out, err := c.DeleteBank(ctx(t), id)
		if err != nil {
			t.Fatalf("DeleteBank: %v", err)
		}
		t.Logf("DeleteBank => %+v", out)
	})
}
