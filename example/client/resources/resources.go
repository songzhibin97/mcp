package main

import (
	"context"
	"fmt"
	"log"

	"github.com/songzhibin97/mcp/protocol"

	"github.com/songzhibin97/mcp/client"
	"github.com/songzhibin97/mcp/transport"
)

func main() {
	// 使用模拟传输
	tr := transport.NewMockTransport()
	c := client.NewDefaultClient(tr)

	ctx := context.Background()

	// 模拟服务器响应
	go func() {
		resp := &protocol.JSONRPCResponse{
			JSONRPC: protocol.JSONRPCVersion,
			ID:      "res-1", // 与 ListResources 请求的 ID 匹配
			Result: protocol.Result{
				"resources": []protocol.Resource{
					{URI: "file://test", Name: "Test Resource"},
				},
			},
		}
		tr.InjectMessage(resp)
	}()

	// 发送 ListResources 请求
	result, err := c.ListResources(ctx, nil)
	if err != nil {
		log.Fatalf("ListResources failed: %v", err)
	}
	for _, res := range result.Resources {
		fmt.Printf("Resource: %s (%s)\n", res.Name, res.URI)
	}

	// 关闭客户端
	if err := c.Close(); err != nil {
		log.Fatalf("Close failed: %v", err)
	}
}
