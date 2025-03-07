package client

import (
	"context"
	"fmt"

	"github.com/songzhibin97/mcp/protocol"
)

func (c *DefaultClient) GetPrompt(ctx context.Context, name string, args map[string]string) (*protocol.GetPromptResult, error) {
	req := protocol.NewGetPromptRequest("prompt-1", name, args)
	resp, err := c.sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	result := &protocol.GetPromptResult{}
	if err := mapToStruct(resp, result); err != nil {
		return nil, fmt.Errorf("decode result failed: %w", err)
	}
	return result, nil
}
