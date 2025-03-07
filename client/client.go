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
	transport transport.Transport
	logger    *util.Logger
	handlers  map[string]func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)
	mu        sync.Mutex
}

// NewDefaultClient 创建新的默认客户端
func NewDefaultClient(tr transport.Transport) *DefaultClient {
	return &DefaultClient{
		transport: tr,
		logger:    util.NewLogger(),
		handlers:  make(map[string]func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)),
	}
}

// RegisterHandler 注册请求处理器
func (c *DefaultClient) RegisterHandler(method string, handler func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[method] = handler
}

// Close 关闭客户端
func (c *DefaultClient) Close() error {
	return c.transport.Close()
}
