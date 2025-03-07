package handlers

import (
	"context"

	"github.com/songzhibin97/mcp/client"
	"github.com/songzhibin97/mcp/protocol"
)

func PromptHandler(c *client.DefaultClient) func(context.Context, *protocol.JSONRPCRequest) (interface{}, error) {
	return func(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
		if req.Method != "prompts/get" {
			return nil, nil
		}
		return protocol.GetPromptResult{
			Messages: []protocol.PromptMessage{
				{Role: "user", Content: protocol.TextContent{Type: "text", Text: "Hello"}},
			},
		}, nil
	}
}
