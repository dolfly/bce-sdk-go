// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package dumemory_test

import "testing"

func TestListEntities(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	out, err := c.ListEntities(ctx(t), bankID(), 10, 0)
	if err != nil {
		t.Fatalf("ListEntities: %v", err)
	}
	t.Logf("ListEntities => %+v", out)
}

func TestEntityGraph(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	out, err := c.EntityGraph(ctx(t), bankID(), 10, 0)
	if err != nil {
		t.Fatalf("EntityGraph: %v", err)
	}
	t.Logf("EntityGraph => %+v", out)
}
