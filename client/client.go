package client

import (
	"context"
	"sync"

	util "github.com/songzhibin97/mcp/internal/utils"
	"github.com/songzhibin97/mcp/protocol"
	"github.com/songzhibin97/mcp/transport"
)

// DefaultClient 是客户端的默认实现
type DefaultClient struct {
	transport       transport.Transport
	logger          *util.Logger
	handlers        map[string]func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)
	pendingRequests map[protocol.RequestId]chan response
	mu              sync.Mutex
}

// response 用于封装响应结果或错误
type response struct {
	result interface{}
	err    error
}

// NewDefaultClient 创建新的默认客户端
func NewDefaultClient(tr transport.Transport) *DefaultClient {
	c := &DefaultClient{
		transport:       tr,
		logger:          util.NewLogger(),
		handlers:        make(map[string]func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)),
		pendingRequests: make(map[protocol.RequestId]chan response),
	}
	// 启动响应处理协程
	go c.handleResponse(context.Background())
	return c
}

// RegisterHandler 注册请求处理器
func (c *DefaultClient) RegisterHandler(method string, handler func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[method] = handler
}

// Close 关闭客户端
func (c *DefaultClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	for id, ch := range c.pendingRequests {
		close(ch) // 关闭所有未完成的响应通道
		delete(c.pendingRequests, id)
	}
	return c.transport.Close()
}
