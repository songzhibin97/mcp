package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

func (s *DefaultServer) HandleGetPrompt(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
	return protocol.GetPromptResult{
		Messages: []protocol.PromptMessage{
			{Role: "user", Content: protocol.TextContent{ContentType: "text", Text: "Hello"}},
		},
	}, nil
}
