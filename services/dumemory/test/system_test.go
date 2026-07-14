// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package dumemory_test

import "testing"

func TestHealth(t *testing.T) {
	c := newClient(t)
	out, err := c.Health(ctx(t))
	if err != nil {
		t.Fatalf("Health failed: %v", err)
	}
	t.Logf("Health => %v", out)
}

func TestVersion(t *testing.T) {
	c := newClient(t)
	out, err := c.Version(ctx(t))
	if err != nil {
		t.Fatalf("Version failed: %v", err)
	}
	t.Logf("Version => %v", out)
}
