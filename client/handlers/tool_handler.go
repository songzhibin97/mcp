package handlers

import (
	"context"

	"github.com/songzhibin97/mcp/client"
	"github.com/songzhibin97/mcp/protocol"
)

func ToolHandler(c *client.DefaultClient) func(context.Context, *protocol.JSONRPCRequest) (interface{}, error) {
	return func(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
		if req.Method != "tools/call" {
			return nil, nil
		}
		return protocol.CallToolResult{
			Content: []interface{}{
				protocol.TextContent{Type: "text", Text: "Tool result"},
			},
		}, nil
	}
}
