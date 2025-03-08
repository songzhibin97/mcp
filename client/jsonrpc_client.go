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

// sendRequest 发送请求并异步等待响应
func (c *DefaultClient) sendRequest(ctx context.Context, req interface{}) (interface{}, error) {
	var id protocol.RequestId
	switch r := req.(type) {
	case *protocol.JSONRPCRequest:
		id = r.ID
	case *protocol.InitializeRequest:
		id = r.ID
	case *protocol.ListResourcesRequest:
		id = r.ID
	case *protocol.GetPromptRequest:
		id = r.ID
	case *protocol.CallToolRequest:
		id = r.ID
	case *protocol.PingRequest:
		id = r.ID
	case *protocol.CreateMessageRequest:
		id = r.ID
	default:
		return nil, fmt.Errorf("unsupported request type: %T", req)
	}

	respChan := make(chan response, 1)

	c.mu.Lock()
	c.pendingRequests[id] = respChan
	c.mu.Unlock()

	if err := c.transport.Send(ctx, req); err != nil {
		c.mu.Lock()
		delete(c.pendingRequests, id)
		c.mu.Unlock()
		close(respChan)
		return nil, fmt.Errorf("send request failed: %w", err)
	}

	select {
	case resp := <-respChan:
		close(respChan)
		return resp.result, resp.err
	case <-ctx.Done():
		c.mu.Lock()
		delete(c.pendingRequests, id)
		c.mu.Unlock()
		close(respChan)
		return nil, ctx.Err()
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
