package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/songzhibin97/mcp/protocol"

	"github.com/songzhibin97/mcp/server"
	"github.com/songzhibin97/mcp/transport"
)

func main() {
	// 使用模拟传输
	tr := transport.NewMockTransport()
	s := server.NewDefaultServer(tr)

	// 注册 Initialize 处理器
	s.RegisterHandler("initialize", s.HandleInitialize)

	// 创建可取消的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 启动服务器
	if err := s.Start(ctx); err != nil {
		log.Fatalf("Server start failed: %v", err)
	}

	fmt.Println("Server started for initialization")

	// 模拟客户端发送 Initialize 请求（可选）
	go func() {
		time.Sleep(100 * time.Millisecond) // 等待服务器启动
		req := protocol.NewInitializeRequest("1", protocol.LatestProtocolVersion, protocol.ClientCapabilities{}, protocol.Implementation{Name: "test-client", Version: "1.0"})
		tr.InjectMessage(req)
	}()

	// 等待上下文超时或取消
	<-ctx.Done()
	if err := s.Close(); err != nil {
		log.Fatalf("Close failed: %v", err)
	}
	fmt.Println("Server stopped")
}
