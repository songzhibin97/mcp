package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/songzhibin97/mcp/protocol"

	"github.com/songzhibin97/mcp/server"
	"github.com/songzhibin97/mcp/server/handlers"
	"github.com/songzhibin97/mcp/transport"
)

func main() {
	// 使用模拟传输
	tr := transport.NewMockTransport()
	s := server.NewDefaultServer(tr)

	// 注册所有处理器
	s.RegisterHandler("initialize", s.HandleInitialize)
	s.RegisterHandler("resources/list", handlers.ResourceHandler(s))
	s.RegisterHandler("prompts/get", handlers.PromptHandler(s))
	s.RegisterHandler("tools/call", handlers.ToolHandler(s))
	s.RegisterHandler("sampling/createMessage", handlers.SamplingHandler(s))
	s.RegisterHandler("ping", s.HandlePing)

	// 创建可取消的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 启动服务器
	if err := s.Start(ctx); err != nil {
		log.Fatalf("Server start failed: %v", err)
	}

	fmt.Println("Full server started")

	// 模拟客户端发送请求
	go func() {
		time.Sleep(100 * time.Millisecond) // 等待服务器启动
		req := protocol.NewPingRequest("ping-1")
		tr.InjectMessage(req)
	}()

	// 发送日志和进度通知
	go func() {
		time.Sleep(1 * time.Second)
		s.SendLog(ctx, "info", "Server running")
		s.SendProgress(ctx, "task-1", 50, 100)
	}()

	// 等待上下文超时或取消
	<-ctx.Done()
	if err := s.Close(); err != nil {
		log.Fatalf("Close failed: %v", err)
	}
	fmt.Println("Server stopped")
}
