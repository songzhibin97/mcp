package jsonrpc

import (
	"encoding/json"
	"fmt"

	"github.com/songzhibin97/mcp/protocol"
)

// Encode 将 JSON-RPC 消息编码为字节数组
func Encode(msg protocol.JSONRPCMessage) ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode JSON-RPC message: %w", err)
	}
	return data, nil
}

// Decode 将字节数组解码为 JSON-RPC 消息
// Decode 将字节数组解码为 JSON-RPC 消息
func Decode(data []byte) (protocol.JSONRPCMessage, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to decode JSON-RPC message: %w", err)
	}

	// 检查 jsonrpc 字段
	if version, ok := raw["jsonrpc"].(string); !ok || version != protocol.JSONRPCVersion {
		return nil, fmt.Errorf("invalid JSON-RPC version: %v", raw["jsonrpc"])
	}

	// 根据是否有 id 和 error 字段判断消息类型
	if _, hasID := raw["id"]; hasID {
		if _, hasError := raw["error"]; hasError {
			var errMsg protocol.JSONRPCError
			if err := json.Unmarshal(data, &errMsg); err != nil {
				return nil, fmt.Errorf("failed to decode JSONRPCError: %w", err)
			}
			return &errMsg, nil
		}
		if _, hasResult := raw["result"]; hasResult {
			var resp protocol.JSONRPCResponse
			if err := json.Unmarshal(data, &resp); err != nil {
				return nil, fmt.Errorf("failed to decode JSONRPCResponse: %w", err)
			}
			return &resp, nil
		}
		var req protocol.JSONRPCRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return nil, fmt.Errorf("failed to decode JSONRPCRequest: %w", err)
		}
		return &req, nil
	}

	// 处理通知类型
	if method, ok := raw["method"].(string); ok {
		switch method {
		case "notifications/progress":
			var notif protocol.ProgressNotification
			if err := json.Unmarshal(data, &notif); err != nil {
				return nil, fmt.Errorf("failed to decode ProgressNotification: %w", err)
			}
			return &notif, nil
		case "notifications/cancelled":
			var notif protocol.CancelledNotification
			if err := json.Unmarshal(data, &notif); err != nil {
				return nil, fmt.Errorf("failed to decode CancelledNotification: %w", err)
			}
			return &notif, nil
		case "notifications/message":
			var notif protocol.LoggingMessageNotification
			if err := json.Unmarshal(data, &notif); err != nil {
				return nil, fmt.Errorf("failed to decode LoggingMessageNotification: %w", err)
			}
			return &notif, nil
		case "notifications/resources/updated":
			var notif protocol.ResourceUpdatedNotification
			if err := json.Unmarshal(data, &notif); err != nil {
				return nil, fmt.Errorf("failed to decode ResourceUpdatedNotification: %w", err)
			}
			return &notif, nil
		case "notifications/resources/list_changed":
			var notif protocol.ResourceListChangedNotification
			if err := json.Unmarshal(data, &notif); err != nil {
				return nil, fmt.Errorf("failed to decode ResourceListChangedNotification: %w", err)
			}
			return &notif, nil
		case "notifications/prompts/list_changed":
			var notif protocol.PromptListChangedNotification
			if err := json.Unmarshal(data, &notif); err != nil {
				return nil, fmt.Errorf("failed to decode PromptListChangedNotification: %w", err)
			}
			return &notif, nil
		case "notifications/tools/list_changed":
			var notif protocol.ToolListChangedNotification
			if err := json.Unmarshal(data, &notif); err != nil {
				return nil, fmt.Errorf("failed to decode ToolListChangedNotification: %w", err)
			}
			return &notif, nil
		case "notifications/initialized":
			var notif protocol.InitializedNotification
			if err := json.Unmarshal(data, &notif); err != nil {
				return nil, fmt.Errorf("failed to decode InitializedNotification: %w", err)
			}
			return &notif, nil
		}
	}

	var notif protocol.JSONRPCNotification
	if err := json.Unmarshal(data, &notif); err != nil {
		return nil, fmt.Errorf("failed to decode JSONRPCNotification: %w", err)
	}
	return &notif, nil
}
