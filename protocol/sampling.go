package protocol

// CreateMessageRequest 服务器请求采样
type CreateMessageRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

// NewCreateMessageRequest 创建新的采样请求
func NewCreateMessageRequest(id RequestId, messages []SamplingMessage, maxTokens int) *CreateMessageRequest {
	return &CreateMessageRequest{
		JSONRPC: JSONRPCVersion,
		ID:      id,
		Request: Request{
			Method: "sampling/createMessage",
			Params: map[string]interface{}{
				"messages":  messages,
				"maxTokens": maxTokens,
			},
		},
	}
}

// CreateMessageResult 是 sampling/createMessage 请求的响应
type CreateMessageResult struct {
	SamplingMessage
	Model      string `json:"model"`
	StopReason string `json:"stopReason,omitempty"`
}

// SamplingMessage LLM 消息
type SamplingMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"` // 简化，假设文本内容
}
