package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

func (s *DefaultServer) HandleGetPrompt(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
	name, _ := req.Params["name"].(string)
	return &protocol.GetPromptResult{
		Messages: []protocol.PromptMessage{
			{Role: "user", Content: protocol.TextContent{Type: "text", Text: "Hello " + name}},
		},
	}, nil
}
