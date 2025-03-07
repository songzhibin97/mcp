package client

import (
	"context"
	"fmt"

	"github.com/songzhibin97/mcp/protocol"
)

func (c *DefaultClient) CreateMessage(ctx context.Context, req *protocol.CreateMessageRequest) (*protocol.CreateMessageResult, error) {
	resp, err := c.sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	result := &protocol.CreateMessageResult{}
	if err := mapToStruct(resp, result); err != nil {
		return nil, fmt.Errorf("decode result failed: %w", err)
	}
	return result, nil
}
