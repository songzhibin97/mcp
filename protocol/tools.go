package protocol

// ListToolsRequest 请求工具列表
type ListToolsRequest struct {
	PaginatedRequest
}

// NewListToolsRequest 创建新的工具列表请求
func NewListToolsRequest(id RequestId, cursor *Cursor) *ListToolsRequest {
	return &ListToolsRequest{
		PaginatedRequest: *NewPaginatedRequest(id, "tools/list", cursor),
	}
}

// Tool 定义工具
type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	InputSchema map[string]interface{} `json:"inputSchema"`
}

// ListToolsResult 工具列表响应
type ListToolsResult struct {
	Tools []Tool `json:"tools"`
	PaginatedResult
}

// CallToolRequest 调用工具
type CallToolRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

// CallToolResult 是 tools/call 请求的响应
type CallToolResult struct {
	Content []interface{} `json:"content"` // 支持 TextContent、ImageContent 等
	IsError bool          `json:"isError,omitempty"`
}

// NewCallToolRequest 创建新的工具调用请求
func NewCallToolRequest(id RequestId, name string, args map[string]interface{}) *CallToolRequest {
	params := map[string]interface{}{
		"name": name,
	}
	if args != nil {
		params["arguments"] = args
	}
	return &CallToolRequest{
		JSONRPC: JSONRPCVersion,
		ID:      id,
		Request: Request{
			Method: "tools/call",
			Params: params,
		},
	}
}
