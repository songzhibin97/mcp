package main

import (
	"context"
	"fmt"
	"log"

	"github.com/songzhibin97/mcp/client"
	"github.com/songzhibin97/mcp/protocol"
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
			ID:      "1", // 与 Initialize 请求的 ID 匹配
			Result: protocol.Result{
				"protocolVersion": protocol.LatestProtocolVersion,
				"capabilities":    protocol.ServerCapabilities{},
				"serverInfo":      protocol.Implementation{Name: "mock-server", Version: "1.0"},
			},
		}
		tr.InjectMessage(resp)
	}()

	// 发送 Initialize 请求
	result, err := c.Initialize(ctx, protocol.LatestProtocolVersion, protocol.ClientCapabilities{}, protocol.Implementation{Name: "test-client", Version: "1.0"})
	if err != nil {
		log.Fatalf("Initialize failed: %v", err)
	}
	fmt.Printf("Initialized: Server Version=%s\n", result.ServerInfo.Version)

	// 关闭客户端
	if err := c.Close(); err != nil {
		log.Fatalf("Close failed: %v", err)
	}
}
