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
			ID:      "ping-1",
			Result:  protocol.Result{"ok": true},
		}
		tr.InjectMessage(resp)
	}()

	// 发送 Ping 请求
	err := c.Ping(ctx)
	if err != nil {
		log.Fatalf("Ping failed: %v", err)
	}
	fmt.Println("Ping successful")

	// 关闭客户端
	if err := c.Close(); err != nil {
		log.Fatalf("Close failed: %v", err)
	}
}
