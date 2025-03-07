package protocol

// ListPromptsRequest 请求提示列表
type ListPromptsRequest struct {
	PaginatedRequest
}

// NewListPromptsRequest 创建新的提示列表请求
func NewListPromptsRequest(id RequestId, cursor *Cursor) *ListPromptsRequest {
	return &ListPromptsRequest{
		PaginatedRequest: *NewPaginatedRequest(id, "prompts/list", cursor),
	}
}

// Prompt 表示服务器提供的提示
type Prompt struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Arguments   []PromptArgument `json:"arguments,omitempty"`
}

// PromptArgument 提示参数
type PromptArgument struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

// ListPromptsResult 提示列表响应
type ListPromptsResult struct {
	Prompts []Prompt `json:"prompts"`
	PaginatedResult
}

// GetPromptRequest 获取特定提示
type GetPromptRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

// GetPromptResult 是 prompts/get 请求的响应
type GetPromptResult struct {
	Description string          `json:"description,omitempty"`
	Messages    []PromptMessage `json:"messages"`
}

// PromptMessage 定义提示中的消息
type PromptMessage struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"` // 支持 TextContent、ImageContent 等
}

// NewGetPromptRequest 创建新的获取提示请求
func NewGetPromptRequest(id RequestId, name string, args map[string]string) *GetPromptRequest {
	params := map[string]interface{}{
		"name": name,
	}
	if args != nil {
		params["arguments"] = args
	}
	return &GetPromptRequest{
		JSONRPC: JSONRPCVersion,
		ID:      id,
		Request: Request{
			Method: "prompts/get",
			Params: params,
		},
	}
}
