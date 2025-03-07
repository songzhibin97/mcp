package protocol

// ListResourcesRequest 请求资源列表
type ListResourcesRequest struct {
	PaginatedRequest
}

// NewListResourcesRequest 创建新的资源列表请求
func NewListResourcesRequest(id RequestId, cursor *Cursor) *ListResourcesRequest {
	return &ListResourcesRequest{
		PaginatedRequest: *NewPaginatedRequest(id, "resources/list", cursor),
	}
}

// Resource 表示服务器可读取的资源
type Resource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
	Size        int    `json:"size,omitempty"`
}

// ListResourcesResult 资源列表响应
type ListResourcesResult struct {
	Resources []Resource `json:"resources"`
	PaginatedResult
}
