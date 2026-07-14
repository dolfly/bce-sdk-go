// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package dumemory_test

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilesRetain(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)

	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	if err := os.WriteFile(path, []byte("integration test file content"), 0o600); err != nil {
		t.Fatalf("write tmp: %v", err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("open tmp: %v", err)
	}
	defer f.Close()

	requestJSON := `{"items":[{"content":"integration test file upload"}]}`
	out, err := c.FilesRetain(ctx(t), bankID(), []*os.File{f}, requestJSON)
	if err != nil {
		t.Fatalf("FilesRetain: %v", err)
	}
	t.Logf("FilesRetain => %+v", out)
}
