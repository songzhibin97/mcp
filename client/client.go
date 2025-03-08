package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/songzhibin97/mcp/internal/utils"
	"github.com/songzhibin97/mcp/protocol"
	"github.com/songzhibin97/mcp/transport"
)

// DefaultClient 是客户端的默认实现
type DefaultClient struct {
	transport       transport.Transport
	logger          *utils.Logger
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
		logger:          utils.NewLogger(),
		handlers:        make(map[string]func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)),
		pendingRequests: make(map[protocol.RequestId]chan response),
	}
	// 启动消息处理协程
	go c.handleMessages(context.Background())
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
		close(ch)
		delete(c.pendingRequests, id)
	}
	return c.transport.Close()
}

// handleMessages 处理所有接收到的消息，分发到通知或响应处理
func (c *DefaultClient) handleMessages(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.transport.Receive(ctx)
			if err != nil {
				c.logger.Error("Failed to receive message", "error", err)
				continue
			}
			switch msg := msg.(type) {
			case *protocol.JSONRPCResponse:
				c.handleResponse(ctx, msg)
			case *protocol.JSONRPCError:
				c.handleError(ctx, msg)
			case protocol.JSONRPCMessage:
				c.handleNotification(ctx, msg)
			default:
				c.logger.Warn("Received unexpected message type", "type", fmt.Sprintf("%T", msg))
			}
		}
	}
}

// handleResponse 处理响应
func (c *DefaultClient) handleResponse(ctx context.Context, resp *protocol.JSONRPCResponse) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ch, ok := c.pendingRequests[resp.ID]; ok {
		ch <- response{result: resp.Result, err: nil}
		delete(c.pendingRequests, resp.ID)
	} else {
		c.logger.Warn("Received response for unknown request ID", "id", resp.ID)
	}
}

// handleError 处理错误响应
func (c *DefaultClient) handleError(ctx context.Context, errResp *protocol.JSONRPCError) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ch, ok := c.pendingRequests[errResp.ID]; ok {
		ch <- response{result: nil, err: fmt.Errorf("server error: code=%d, message=%s", errResp.Error.Code, errResp.Error.Message)}
		delete(c.pendingRequests, errResp.ID)
	} else {
		c.logger.Warn("Received error for unknown request ID", "id", errResp.ID)
	}
}

// handleNotification 处理通知
func (c *DefaultClient) handleNotification(ctx context.Context, msg protocol.JSONRPCMessage) {
	switch n := msg.(type) {
	case *protocol.ProgressNotification:
		token, _ := n.Params["progressToken"].(string)
		progress, _ := n.Params["progress"].(float64)
		total, totalOk := n.Params["total"].(float64)
		if totalOk {
			c.logger.Info("Progress update", "token", token, "progress", progress, "total", total)
		} else {
			c.logger.Info("Progress update", "token", token, "progress", progress)
		}

	case *protocol.CancelledNotification:
		requestId, _ := n.Params["requestId"].(string)
		reason, reasonOk := n.Params["reason"].(string)
		if reasonOk {
			c.logger.Info("Request cancelled", "requestId", requestId, "reason", reason)
		} else {
			c.logger.Info("Request cancelled", "requestId", requestId)
		}
		c.mu.Lock()
		if ch, ok := c.pendingRequests[protocol.RequestId(requestId)]; ok {
			ch <- response{result: nil, err: fmt.Errorf("request cancelled by server: %s", reason)}
			delete(c.pendingRequests, protocol.RequestId(requestId))
		}
		c.mu.Unlock()

	case *protocol.LoggingMessageNotification:
		level, _ := n.Params["level"].(string)
		data, _ := n.Params["data"]
		loggerName, loggerOk := n.Params["logger"].(string)
		if loggerOk {
			c.logger.Log(ctx, level, "Server log", "logger", loggerName, "data", data)
		} else {
			c.logger.Log(ctx, level, "Server log", "data", data)
		}

	case *protocol.ResourceUpdatedNotification:
		uri, _ := n.Params["uri"].(string)
		c.logger.Info("Resource updated", "uri", uri)

	case *protocol.ResourceListChangedNotification:
		c.logger.Info("Resource list changed")

	case *protocol.PromptListChangedNotification:
		c.logger.Info("Prompt list changed")

	case *protocol.ToolListChangedNotification:
		c.logger.Info("Tool list changed")

	case *protocol.InitializedNotification:
		c.logger.Info("Server initialized")

	default:
		c.logger.Warn("Received unknown notification", "msg", fmt.Sprintf("%+v", msg))
	}
}
