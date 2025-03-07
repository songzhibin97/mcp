package jsonrpc

import (
	"fmt"

	"github.com/songzhibin97/mcp/protocol"
)

// Validate 验证 JSON-RPC 消息的格式
func Validate(msg protocol.JSONRPCMessage) error {
	switch m := msg.(type) {
	case *protocol.JSONRPCRequest:
		if m.JSONRPC != protocol.JSONRPCVersion {
			return fmt.Errorf("invalid JSON-RPC version: %s", m.JSONRPC)
		}
		if m.ID == "" {
			return fmt.Errorf("missing request ID")
		}
		if m.Method == "" {
			return fmt.Errorf("missing method")
		}
	case *protocol.JSONRPCNotification:
		if m.JSONRPC != protocol.JSONRPCVersion {
			return fmt.Errorf("invalid JSON-RPC version: %s", m.JSONRPC)
		}
		if m.Method == "" {
			return fmt.Errorf("missing method")
		}
	case *protocol.JSONRPCResponse:
		if m.JSONRPC != protocol.JSONRPCVersion {
			return fmt.Errorf("invalid JSON-RPC version: %s", m.JSONRPC)
		}
		if m.ID == "" {
			return fmt.Errorf("missing response ID")
		}
	case *protocol.JSONRPCError:
		if m.JSONRPC != protocol.JSONRPCVersion {
			return fmt.Errorf("invalid JSON-RPC version: %s", m.JSONRPC)
		}
		if m.ID == "" {
			return fmt.Errorf("missing error ID")
		}
		if m.Error.Code == 0 && m.Error.Message == "" {
			return fmt.Errorf("missing error code or message")
		}
	default:
		return fmt.Errorf("unknown message type")
	}
	return nil
}
