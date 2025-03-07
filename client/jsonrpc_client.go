package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/songzhibin97/mcp/protocol"
)

// requestResponse 用于存储请求和响应的映射
type requestResponse struct {
	response chan interface{}
	err      chan error
}

func (c *DefaultClient) sendRequest(ctx context.Context, req protocol.JSONRPCMessage) (interface{}, error) {
	if err := c.transport.Send(ctx, req); err != nil {
		return nil, fmt.Errorf("send request failed: %w", err)
	}

	resp, err := c.transport.Receive(ctx)
	if err != nil {
		return nil, fmt.Errorf("receive response failed: %w", err)
	}

	switch r := resp.(type) {
	case *protocol.JSONRPCResponse:
		return r.Result, nil
	case *protocol.JSONRPCError:
		return nil, fmt.Errorf("server error: code=%d, message=%s", r.Error.Code, r.Error.Message)
	default:
		return nil, fmt.Errorf("unexpected response type")
	}
}

// Initialize 实现初始化
func (c *DefaultClient) Initialize(ctx context.Context, protocolVersion string, capabilities protocol.ClientCapabilities, clientInfo protocol.Implementation) (*protocol.InitializeResult, error) {
	req := protocol.NewInitializeRequest("1", protocolVersion, capabilities, clientInfo)
	resp, err := c.sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	result := &protocol.InitializeResult{}
	if err := mapToStruct(resp, result); err != nil {
		return nil, fmt.Errorf("decode result failed: %w", err)
	}
	return result, nil
}

// Ping 实现 Ping
func (c *DefaultClient) Ping(ctx context.Context) error {
	req := protocol.NewPingRequest("ping-1")
	_, err := c.sendRequest(ctx, req)
	return err
}

// mapToStruct 将 map 转换为结构体
func mapToStruct(src interface{}, dst interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dst)
}
