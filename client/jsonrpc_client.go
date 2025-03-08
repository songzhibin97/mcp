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
	// 从请求中提取 ID
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

	// 注册请求
	c.mu.Lock()
	c.pendingRequests[id] = respChan
	c.mu.Unlock()

	// 发送请求
	if err := c.transport.Send(ctx, req); err != nil {
		c.mu.Lock()
		delete(c.pendingRequests, id)
		c.mu.Unlock()
		close(respChan)
		return nil, fmt.Errorf("send request failed: %w", err)
	}

	// 等待响应或上下文超时
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

// handleResponse 持续处理来自 transport 的响应
func (c *DefaultClient) handleResponse(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			resp, err := c.transport.Receive(ctx)
			if err != nil {
				c.logger.Error("Receive failed", "error", err)
				continue
			}
			switch r := resp.(type) {
			case *protocol.JSONRPCResponse:
				c.mu.Lock()
				if ch, ok := c.pendingRequests[r.ID]; ok {
					ch <- response{result: r.Result, err: nil}
					delete(c.pendingRequests, r.ID)
				} else {
					c.logger.Warn("Received response for unknown request ID", "id", r.ID)
				}
				c.mu.Unlock()
			case *protocol.JSONRPCError:
				c.mu.Lock()
				if ch, ok := c.pendingRequests[r.ID]; ok {
					ch <- response{result: nil, err: fmt.Errorf("server error: code=%d, message=%s", r.Error.Code, r.Error.Message)}
					delete(c.pendingRequests, r.ID)
				} else {
					c.logger.Warn("Received error for unknown request ID", "id", r.ID)
				}
				c.mu.Unlock()
			default:
				c.logger.Warn("Received unexpected message type", "type", fmt.Sprintf("%T", resp))
			}
		}
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
