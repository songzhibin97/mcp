package client

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

// Client 定义客户端接口
type Client interface {
	Initialize(ctx context.Context, protocolVersion string, capabilities protocol.ClientCapabilities, clientInfo protocol.Implementation) (*protocol.InitializeResult, error)
	ListResources(ctx context.Context, cursor *protocol.Cursor) (*protocol.ListResourcesResult, error)
	GetPrompt(ctx context.Context, name string, args map[string]string) (*protocol.GetPromptResult, error)
	CallTool(ctx context.Context, name string, args map[string]interface{}) (*protocol.CallToolResult, error)
	CreateMessage(ctx context.Context, req *protocol.CreateMessageRequest) (*protocol.CreateMessageResult, error)
	Ping(ctx context.Context) error
	Close() error
}
