package protocol

// CompleteRequest 自动完成请求
type CompleteRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

// NewCompleteRequest 创建新的自动完成请求
func NewCompleteRequest(id RequestId, ref interface{}, argName, argValue string) *CompleteRequest {
	return &CompleteRequest{
		JSONRPC: JSONRPCVersion,
		ID:      id,
		Request: Request{
			Method: "completion/complete",
			Params: map[string]interface{}{
				"ref":      ref,
				"argument": map[string]string{"name": argName, "value": argValue},
			},
		},
	}
}

// CompleteResult 自动完成结果
type CompleteResult struct {
	Completion struct {
		Values  []string `json:"values"`
		Total   int      `json:"total,omitempty"`
		HasMore bool     `json:"hasMore,omitempty"`
	} `json:"completion"`
}
