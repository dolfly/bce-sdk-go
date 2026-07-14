// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package dumemory

import (
	"context"
	"os"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// FilesRetain uploads files and writes them as memories via POST /files/retain.
// requestJSON is an optional JSON-encoded RetainRequest accompanying the
// upload (pass "" to omit).
func (c *Client) FilesRetain(ctx context.Context, bankID string, files []*os.File, requestJSON string) (*hindsight.FileRetainResponse, error) {
	call := c.hindsight.FilesAPI.FileRetain(ctx, bankID).Files(files)
	if requestJSON != "" {
		call = call.Request(requestJSON)
	}
	out, _, err := call.Execute()
	return out, err
}
