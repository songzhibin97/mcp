package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

func (s *DefaultServer) HandleCallTool(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
	return &protocol.CallToolResult{
		Content: []interface{}{
			protocol.TextContent{ContentType: "text", Text: "Tool executed"},
		},
	}, nil
}
