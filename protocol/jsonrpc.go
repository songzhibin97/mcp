package protocol

import "encoding/json"

const (
	LatestProtocolVersion = "2024-11-05"
	JSONRPCVersion        = "2.0"
)

type JSONRPCMessage interface{}

type Request struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type Notification struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type Result map[string]interface{}

type JSONRPCRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

type JSONRPCNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Notification
}

type JSONRPCResponse struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Result  Result    `json:"result"`
}

func (r JSONRPCRequest) MarshalJSON() ([]byte, error) {
	type Alias JSONRPCRequest
	return json.Marshal(&struct {
		JSONRPC string `json:"jsonrpc"`
		ID      RequestId
		*Alias
	}{
		JSONRPC: JSONRPCVersion,
		ID:      r.ID,
		Alias:   (*Alias)(&r),
	})
}
