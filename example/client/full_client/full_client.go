package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
		// 等待客户端发送请求
		time.Sleep(100 * time.Millisecond)

		// 响应 Initialize
		initResp := &protocol.JSONRPCResponse{
			JSONRPC: protocol.JSONRPCVersion,
			ID:      "1",
			Result: protocol.Result{
				"protocolVersion": protocol.LatestProtocolVersion,
				"capabilities":    protocol.ServerCapabilities{},
				"serverInfo":      protocol.Implementation{Name: "mock-server", Version: "1.0"},
			},
		}
		tr.InjectMessage(initResp)

		// 响应 ListResources
		resResp := &protocol.JSONRPCResponse{
			JSONRPC: protocol.JSONRPCVersion,
			ID:      "res-1",
			Result: protocol.Result{
				"resources": []protocol.Resource{
					{URI: "file://test", Name: "Test Resource"},
				},
			},
		}
		tr.InjectMessage(resResp)

		// 响应 GetPrompt
		promptResp := &protocol.JSONRPCResponse{
			JSONRPC: protocol.JSONRPCVersion,
			ID:      "prompt-1",
			Result: protocol.Result{
				"messages": []protocol.PromptMessage{
					{Role: "user", Content: protocol.TextContent{Type: "text", Text: "Hello"}},
				},
			},
		}
		tr.InjectMessage(promptResp)

		// 响应 CallTool
		toolResp := &protocol.JSONRPCResponse{
			JSONRPC: protocol.JSONRPCVersion,
			ID:      "tool-1",
			Result: protocol.Result{
				"content": []interface{}{
					protocol.TextContent{Type: "text", Text: "Tool result"},
				},
			},
		}
		tr.InjectMessage(toolResp)
	}()

	// 初始化
	initResult, err := c.Initialize(ctx, protocol.LatestProtocolVersion, protocol.ClientCapabilities{}, protocol.Implementation{Name: "full-client", Version: "1.0"})
	if err != nil {
		log.Fatalf("Initialize failed: %v", err)
	}
	fmt.Printf("Initialized: %s\n", initResult.ServerInfo.Name)

	// 列出资源
	resResult, err := c.ListResources(ctx, nil)
	if err != nil {
		log.Fatalf("ListResources failed: %v", err)
	}
	fmt.Printf("Resources: %d\n", len(resResult.Resources))

	// 获取提示
	promptResult, err := c.GetPrompt(ctx, "test-prompt", nil)
	if err != nil {
		log.Fatalf("GetPrompt failed: %v", err)
	}
	fmt.Printf("Prompt Messages: %d\n", len(promptResult.Messages))

	// 调用工具
	toolResult, err := c.CallTool(ctx, "test-tool", nil)
	if err != nil {
		log.Fatalf("CallTool failed: %v", err)
	}
	fmt.Printf("Tool Content: %d\n", len(toolResult.Content))

	// 处理通知
	c.HandleNotifications(ctx)

	// 等待通知处理（简单模拟）
	time.Sleep(1 * time.Second)

	if err := c.Close(); err != nil {
		log.Fatalf("Close failed: %v", err)
	}
}
