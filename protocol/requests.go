package protocol

// PingRequest Ping 请求
type PingRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

// NewPingRequest 创建新的 Ping 请求
func NewPingRequest(id RequestId) *PingRequest {
	return &PingRequest{
		JSONRPC: JSONRPCVersion,
		ID:      id,
		Request: Request{
			Method: "ping",
		},
	}
}
