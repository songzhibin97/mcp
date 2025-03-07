package handlers

import (
	"context"

	"github.com/songzhibin97/mcp/client"
	"github.com/songzhibin97/mcp/protocol"
)

func SamplingHandler(c *client.DefaultClient) func(context.Context, *protocol.JSONRPCRequest) (interface{}, error) {
	return func(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
		if req.Method != "sampling/createMessage" {
			return nil, nil
		}
		return protocol.CreateMessageResult{
			SamplingMessage: protocol.SamplingMessage{Role: "assistant", Content: "Sample response"},
			Model:           "mock-model",
		}, nil
	}
}
