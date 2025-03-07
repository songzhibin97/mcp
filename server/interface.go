package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

// Server 定义服务器接口
type Server interface {
	HandleInitialize(ctx context.Context, req *protocol.InitializeRequest) (*protocol.InitializeResult, error)
	HandleListResources(ctx context.Context, req *protocol.ListResourcesRequest) (*protocol.ListResourcesResult, error)
	HandleGetPrompt(ctx context.Context, req *protocol.GetPromptRequest) (*protocol.GetPromptResult, error)
	HandleCallTool(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error)
	HandlePing(ctx context.Context, req *protocol.PingRequest) (interface{}, error)
	Start(ctx context.Context) error
	Close() error
}
