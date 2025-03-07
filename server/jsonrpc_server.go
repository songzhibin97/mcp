package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

func (s *DefaultServer) HandleInitialize(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
	// 假设简单实现
	return &protocol.InitializeResult{
		ProtocolVersion: protocol.LatestProtocolVersion,
		Capabilities:    protocol.ServerCapabilities{},
		ServerInfo:      protocol.Implementation{Name: "mcp-server", Version: "1.0"},
	}, nil
}

func (s *DefaultServer) HandlePing(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
	return protocol.Result{"ok": true}, nil
}
