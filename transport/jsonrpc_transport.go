package transport

import (
	"context"
	"fmt"

	"github.com/songzhibin97/mcp/internal/jsonrpc"
	util "github.com/songzhibin97/mcp/internal/utils"
	"github.com/songzhibin97/mcp/protocol"
)

// JSONRPCTransport 是基于 JSON-RPC 的传输基类
type JSONRPCTransport struct {
	logger *util.Logger
}

// NewJSONRPCTransport 创建新的 JSONRPCTransport
func NewJSONRPCTransport() *JSONRPCTransport {
	return &JSONRPCTransport{
		logger: util.NewLogger(),
	}
}

// SendMsg 发送 JSON-RPC 消息（需由具体传输实现调用）
func (t *JSONRPCTransport) SendMsg(ctx context.Context, msg protocol.JSONRPCMessage, sender func([]byte) error) error {
	if err := jsonrpc.Validate(msg); err != nil {
		return fmt.Errorf("invalid message: %w", err)
	}
	data, err := jsonrpc.Encode(msg)
	if err != nil {
		return fmt.Errorf("encode failed: %w", err)
	}
	t.logger.Debug("Sending message", "data", string(data))
	return sender(data)
}

// ReceiveMsg 接收 JSON-RPC 消息（需由具体传输实现调用）
func (t *JSONRPCTransport) ReceiveMsg(ctx context.Context, receiver func() ([]byte, error)) (protocol.JSONRPCMessage, error) {
	data, err := receiver()
	if err != nil {
		return nil, fmt.Errorf("receive failed: %w", err)
	}
	t.logger.Debug("Received message", "data", string(data))
	msg, err := jsonrpc.Decode(data)
	if err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}
	if err := jsonrpc.Validate(msg); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return msg, nil
}
