package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/songzhibin97/mcp/protocol"
)

// SendCreateMessage 发起采样请求
func (s *DefaultServer) SendCreateMessage(ctx context.Context) (*protocol.CreateMessageResult, error) {
	req := protocol.NewCreateMessageRequest("sample-1", []protocol.SamplingMessage{
		{Role: "user", Content: "Hello"},
	}, 100)
	if err := s.transport.Send(ctx, req); err != nil {
		return nil, fmt.Errorf("send failed: %w", err)
	}
	resp, err := s.transport.Receive(ctx)
	if err != nil {
		return nil, fmt.Errorf("receive failed: %w", err)
	}

	// 类型断言提取响应
	jsonResp, ok := resp.(*protocol.JSONRPCResponse)
	if !ok {
		if errResp, ok := resp.(*protocol.JSONRPCError); ok {
			return nil, fmt.Errorf("server error: code=%d, message=%s", errResp.Error.Code, errResp.Error.Message)
		}
		return nil, fmt.Errorf("unexpected response type: %T", resp)
	}

	// 将 Result 转换为 CreateMessageResult
	result := &protocol.CreateMessageResult{}
	data, err := json.Marshal(jsonResp.Result)
	if err != nil {
		return nil, fmt.Errorf("marshal result failed: %w", err)
	}
	if err := json.Unmarshal(data, result); err != nil {
		return nil, fmt.Errorf("unmarshal result failed: %w", err)
	}
	return result, nil
}
