package server

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	util "github.com/songzhibin97/mcp/internal/utils"
	"github.com/songzhibin97/mcp/protocol"
	"github.com/songzhibin97/mcp/transport"
)

// DefaultServer 是服务器的默认实现
type DefaultServer struct {
	transport transport.Transport
	logger    *util.Logger
	handlers  map[string]func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)
	mu        sync.Mutex
}

// NewDefaultServer 创建新的默认服务器
func NewDefaultServer(tr transport.Transport) *DefaultServer {
	s := &DefaultServer{
		transport: tr,
		logger:    util.NewLogger(),
		handlers:  make(map[string]func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)),
	}
	s.RegisterHandler("initialize", s.HandleInitialize)
	s.RegisterHandler("resources/list", s.HandleListResources)
	s.RegisterHandler("prompts/get", s.HandleGetPrompt)
	go s.handleMessages(context.Background())
	return s
}

// RegisterHandler 注册请求处理器
func (s *DefaultServer) RegisterHandler(method string, handler func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[method] = handler
}

// Start 启动服务器，处理传入请求
func (s *DefaultServer) Start(ctx context.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := s.transport.Receive(ctx)
				if err != nil {
					s.logger.Error("Failed to receive message", "error", err)
					continue
				}
				s.handleMessage(ctx, msg)
			}
		}
	}()
	return nil
}

// Close 关闭服务器
func (s *DefaultServer) Close() error {
	return s.transport.Close()
}

// handleMessages 处理接收到的消息
func (s *DefaultServer) handleMessages(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := s.transport.Receive(ctx)
			if err != nil {
				s.logger.Error("Failed to receive message", "error", err)
				continue
			}
			switch req := msg.(type) {
			case *protocol.JSONRPCRequest:
				go s.handleRequest(ctx, req)
			default:
				s.logger.Warn("Received unexpected message type", "type", fmt.Sprintf("%T", msg))
			}
		}
	}
}

// handleRequest 处理请求
func (s *DefaultServer) handleRequest(ctx context.Context, req *protocol.JSONRPCRequest) {
	handler, ok := s.handlers[req.Method]
	if !ok {
		s.sendError(ctx, req.ID, protocol.MethodNotFound, fmt.Sprintf("method %s not found", req.Method))
		return
	}
	result, err := handler(ctx, req)
	if err != nil {
		s.sendError(ctx, req.ID, protocol.InternalError, err.Error())
		return
	}
	// 类型断言：将 interface{} 转换为 protocol.Result
	respResult, ok := result.(protocol.Result)
	if !ok {
		// 如果 result 不是 protocol.Result，尝试转换为可序列化的类型
		respResult = protocol.Result{}
		data, marshalErr := json.Marshal(result)
		if marshalErr != nil {
			s.sendError(ctx, req.ID, protocol.InternalError, fmt.Sprintf("failed to marshal result: %v", marshalErr))
			return
		}
		if unmarshalErr := json.Unmarshal(data, &respResult); unmarshalErr != nil {
			s.sendError(ctx, req.ID, protocol.InternalError, fmt.Sprintf("failed to unmarshal result: %v", unmarshalErr))
			return
		}
	}
	resp := &protocol.JSONRPCResponse{
		JSONRPC: protocol.JSONRPCVersion,
		ID:      req.ID,
		Result:  respResult,
	}
	if err := s.transport.Send(ctx, resp); err != nil {
		s.logger.Error("Failed to send response", "error", err)
	}
}

// handleMessage 处理接收到的消息
func (s *DefaultServer) handleMessage(ctx context.Context, msg interface{}) {
	var id protocol.RequestId
	var method string
	var handler func(context.Context, *protocol.JSONRPCRequest) (interface{}, error)

	switch req := msg.(type) {
	case *protocol.JSONRPCRequest:
		id = req.ID
		method = req.Method
		handler = s.handlers[method]
	case *protocol.InitializeRequest:
		id = req.ID
		method = req.Method
		handler = s.handlers[method]
	case *protocol.PingRequest:
		id = req.ID
		method = req.Method
		handler = s.handlers[method]
	case *protocol.ListResourcesRequest:
		id = req.ID
		method = req.Method
		handler = s.handlers[method]
	case *protocol.GetPromptRequest:
		id = req.ID
		method = req.Method
		handler = s.handlers[method]
	case *protocol.CallToolRequest:
		id = req.ID
		method = req.Method
		handler = s.handlers[method]
	default:
		s.logger.Warn("Received non-request message", "msg", msg)
		return
	}

	if handler == nil {
		s.sendError(ctx, id, protocol.MethodNotFound, "Method not found")
		return
	}

	result, err := handler(ctx, &protocol.JSONRPCRequest{ID: id, Request: protocol.Request{Method: method}})
	if err != nil {
		s.sendError(ctx, id, protocol.InternalError, err.Error())
		return
	}
	s.sendResponse(ctx, id, result)
}

// sendResponse 发送响应
func (s *DefaultServer) sendResponse(ctx context.Context, id protocol.RequestId, result interface{}) {
	resp := &protocol.JSONRPCResponse{
		JSONRPC: protocol.JSONRPCVersion,
		ID:      id,
		Result:  result.(protocol.Result),
	}
	if err := s.transport.Send(ctx, resp); err != nil {
		s.logger.Error("Failed to send response", "error", err)
	}
}

// sendError 发送错误响应
func (s *DefaultServer) sendError(ctx context.Context, id protocol.RequestId, code int, message string) {
	errResp := &protocol.JSONRPCError{
		JSONRPC: protocol.JSONRPCVersion,
		ID:      id,
		Error: struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Data    interface{} `json:"data,omitempty"`
		}{
			Code:    code,
			Message: message,
		},
	}
	if err := s.transport.Send(ctx, errResp); err != nil {
		s.logger.Error("Failed to send error response", "error", err)
	}
}
