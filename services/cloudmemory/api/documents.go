// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package cloudmemory

import (
	"context"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ListDocumentsOptions are optional filters for ListDocuments.
type ListDocumentsOptions struct {
	Q         string
	Tags      []string
	TagsMatch string
	Limit     int32
	Offset    int32
}

// ListDocuments lists documents via GET /documents.
func (c *Client) ListDocuments(ctx context.Context, bankID string, opts ListDocumentsOptions) (*hindsight.ListDocumentsResponse, error) {
	call := c.hindsight.DocumentsAPI.ListDocuments(ctx, bankID)
	if opts.Q != "" {
		call = call.Q(opts.Q)
	}
	if len(opts.Tags) > 0 {
		call = call.Tags(opts.Tags)
	}
	if opts.TagsMatch != "" {
		call = call.TagsMatch(opts.TagsMatch)
	}
	if opts.Limit != 0 {
		call = call.Limit(opts.Limit)
	}
	if opts.Offset != 0 {
		call = call.Offset(opts.Offset)
	}
	out, _, err := call.Execute()
	return out, err
}

// GetDocument fetches a document via GET /documents/{documentId}.
func (c *Client) GetDocument(ctx context.Context, bankID, documentID string) (*hindsight.DocumentResponse, error) {
	out, _, err := c.hindsight.DocumentsAPI.GetDocument(ctx, bankID, documentID).Execute()
	return out, err
}

// UpdateDocument patches document tags via PATCH /documents/{documentId}.
func (c *Client) UpdateDocument(ctx context.Context, bankID, documentID string, req hindsight.UpdateDocumentRequest) (*hindsight.UpdateDocumentResponse, error) {
	out, _, err := c.hindsight.DocumentsAPI.UpdateDocument(ctx, bankID, documentID).UpdateDocumentRequest(req).Execute()
	return out, err
}

// DeleteDocument deletes a document via DELETE /documents/{documentId}.
func (c *Client) DeleteDocument(ctx context.Context, bankID, documentID string) (*hindsight.DeleteDocumentResponse, error) {
	out, _, err := c.hindsight.DocumentsAPI.DeleteDocument(ctx, bankID, documentID).Execute()
	return out, err
}

// ListDocumentChunks lists chunks via GET /documents/{documentId}/chunks.
func (c *Client) ListDocumentChunks(ctx context.Context, bankID, documentID string, limit, offset int32) (*hindsight.ListChunksResponse, error) {
	call := c.hindsight.DocumentsAPI.ListDocumentChunks(ctx, bankID, documentID)
	if limit != 0 {
		call = call.Limit(limit)
	}
	if offset != 0 {
		call = call.Offset(offset)
	}
	out, _, err := call.Execute()
	return out, err
}
