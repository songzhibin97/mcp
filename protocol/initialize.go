package protocol

import "encoding/json"

// InitializeRequest 客户端发送的初始化请求
type InitializeRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

// NewInitializeRequest 创建新的初始化请求
func NewInitializeRequest(id RequestId, protocolVersion string, capabilities ClientCapabilities, clientInfo Implementation) *InitializeRequest {
	return &InitializeRequest{
		JSONRPC: JSONRPCVersion,
		ID:      id,
		Request: Request{
			Method: "initialize",
			Params: map[string]interface{}{
				"protocolVersion": protocolVersion,
				"capabilities":    capabilities,
				"clientInfo":      clientInfo,
			},
		},
	}
}

// InitializeResult 服务器响应的初始化结果
type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      Implementation     `json:"serverInfo"`
	Instructions    string             `json:"instructions,omitempty"`
}

// InitializedNotification 客户端发送的初始化完成通知
type InitializedNotification struct {
	JSONRPC      string `json:"jsonrpc"`
	Notification struct {
		Method string `json:"method"`
	} `json:"method"`
}

// NewInitializedNotification 创建新的初始化完成通知
func NewInitializedNotification() *InitializedNotification {
	return &InitializedNotification{
		JSONRPC: JSONRPCVersion,
		Notification: struct {
			Method string `json:"method"`
		}{
			Method: "notifications/initialized",
		},
	}
}

// MarshalJSON 自定义序列化
func (n InitializedNotification) MarshalJSON() ([]byte, error) {
	type Alias InitializedNotification
	return json.Marshal(&struct {
		JSONRPC string `json:"jsonrpc"`
		*Alias
	}{
		JSONRPC: JSONRPCVersion,
		Alias:   (*Alias)(&n),
	})
}
