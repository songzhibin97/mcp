package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/songzhibin97/mcp/client"
	"github.com/songzhibin97/mcp/protocol"
	"github.com/songzhibin97/mcp/transport"
)

func main() {
	tr := transport.NewMockTransport()
	c := client.NewDefaultClient(tr)

	ctx := context.Background()

	// 模拟服务器响应多个请求
	go func() {
		time.Sleep(100 * time.Millisecond)
		responses := []interface{}{
			&protocol.JSONRPCResponse{
				JSONRPC: protocol.JSONRPCVersion,
				ID:      "1",
				Result: protocol.Result{
					"protocolVersion": protocol.LatestProtocolVersion,
					"capabilities":    protocol.ServerCapabilities{},
					"serverInfo":      protocol.Implementation{Name: "mock-server", Version: "1.0"},
				},
			},
			&protocol.JSONRPCResponse{
				JSONRPC: protocol.JSONRPCVersion,
				ID:      "res-1",
				Result: protocol.Result{
					"resources": []protocol.Resource{{URI: "file://test", Name: "Test Resource"}},
				},
			},
			&protocol.JSONRPCResponse{
				JSONRPC: protocol.JSONRPCVersion,
				ID:      "prompt-1",
				Result: protocol.Result{
					"messages": []protocol.PromptMessage{{Role: "user", Content: protocol.TextContent{Type: "text", Text: "Hello"}}},
				},
			},
		}
		for _, resp := range responses {
			tr.InjectMessage(resp)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(3)

	// 并发发送请求
	go func() {
		defer wg.Done()
		result, err := c.Initialize(ctx, protocol.LatestProtocolVersion, protocol.ClientCapabilities{}, protocol.Implementation{Name: "full-client", Version: "1.0"})
		if err != nil {
			log.Printf("Initialize failed: %v", err)
			return
		}
		fmt.Printf("Initialized: %s\n", result.ServerInfo.Name)
	}()

	go func() {
		defer wg.Done()
		result, err := c.ListResources(ctx, nil)
		if err != nil {
			log.Printf("ListResources failed: %v", err)
			return
		}
		fmt.Printf("Resources: %d\n", len(result.Resources))
	}()

	go func() {
		defer wg.Done()
		result, err := c.GetPrompt(ctx, "test-prompt", nil)
		if err != nil {
			log.Printf("GetPrompt failed: %v", err)
			return
		}
		fmt.Printf("Prompt Messages: %d\n", len(result.Messages))
	}()

	wg.Wait()

	if err := c.Close(); err != nil {
		log.Fatalf("Close failed: %v", err)
	}
}
