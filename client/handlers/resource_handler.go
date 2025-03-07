package handlers

import (
	"context"

	"github.com/songzhibin97/mcp/client"
	"github.com/songzhibin97/mcp/protocol"
)

func ResourceHandler(c *client.DefaultClient) func(context.Context, *protocol.JSONRPCRequest) (interface{}, error) {
	return func(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
		if req.Method != "resources/list" {
			return nil, nil // 不处理
		}
		// 假设客户端有本地资源列表
		return protocol.ListResourcesResult{
			Resources: []protocol.Resource{
				{URI: "file://test", Name: "Test Resource"},
			},
		}, nil
	}
}
