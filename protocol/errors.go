package protocol

import "encoding/json"

const (
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603
)

type JSONRPCError struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Error   struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	} `json:"error"`
}

func (e JSONRPCError) MarshalJSON() ([]byte, error) {
	type Alias JSONRPCError
	return json.Marshal(&struct {
		JSONRPC string `json:"jsonrpc"`
		ID      RequestId
		Error   struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Data    interface{} `json:"data,omitempty"`
		}
	}{
		JSONRPC: JSONRPCVersion,
		ID:      e.ID,
		Error:   e.Error,
	})
}
