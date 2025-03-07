package handlers

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
	"github.com/songzhibin97/mcp/server"
)

func SamplingHandler(s *server.DefaultServer) func(context.Context, *protocol.JSONRPCRequest) (interface{}, error) {
	return func(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
		return nil, nil // 客户端处理采样请求，服务器不直接响应
	}
}
