package client

import (
	"context"
	"fmt"

	"github.com/songzhibin97/mcp/protocol"
)

func (c *DefaultClient) ListResources(ctx context.Context, cursor *protocol.Cursor) (*protocol.ListResourcesResult, error) {
	req := protocol.NewListResourcesRequest("res-1", cursor)
	resp, err := c.sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	result := &protocol.ListResourcesResult{}
	if err := mapToStruct(resp, result); err != nil {
		return nil, fmt.Errorf("decode result failed: %w", err)
	}
	return result, nil
}
