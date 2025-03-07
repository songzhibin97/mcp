package protocol

// PaginatedRequest 分页请求的通用结构
type PaginatedRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

// NewPaginatedRequest 创建新的分页请求
func NewPaginatedRequest(id RequestId, method string, cursor *Cursor) *PaginatedRequest {
	params := map[string]interface{}{}
	if cursor != nil {
		params["cursor"] = *cursor
	}
	return &PaginatedRequest{
		JSONRPC: JSONRPCVersion,
		ID:      id,
		Request: Request{
			Method: method,
			Params: params,
		},
	}
}

// PaginatedResult 分页结果的通用结构
type PaginatedResult struct {
	NextCursor *Cursor `json:"nextCursor,omitempty"`
}
