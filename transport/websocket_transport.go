package transport

import (
	"context"
	"fmt"
	"sync"

	util "github.com/songzhibin97/mcp/internal/utils"
	"github.com/songzhibin97/mcp/protocol"

	"github.com/gorilla/websocket"
)

// WebSocketTransport 实现基于 WebSocket 的传输
type WebSocketTransport struct {
	conn   *websocket.Conn
	base   *JSONRPCTransport
	logger *util.Logger
	mu     sync.Mutex
}

// NewWebSocketTransport 创建新的 WebSocketTransport
func NewWebSocketTransport(url string) (*WebSocketTransport, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("websocket dial failed: %w", err)
	}
	return &WebSocketTransport{
		conn:   conn,
		base:   NewJSONRPCTransport(),
		logger: util.NewLogger(),
	}, nil
}

// Send 实现 Transport 接口的发送方法
func (t *WebSocketTransport) Send(ctx context.Context, msg interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.base.SendMsg(ctx, msg.(protocol.JSONRPCMessage), func(data []byte) error {
		return t.conn.WriteMessage(websocket.TextMessage, data)
	})
}

// Receive 实现 Transport 接口的接收方法
func (t *WebSocketTransport) Receive(ctx context.Context) (interface{}, error) {
	return t.base.ReceiveMsg(ctx, func() ([]byte, error) {
		_, data, err := t.conn.ReadMessage()
		return data, err
	})
}

// Close 关闭传输
func (t *WebSocketTransport) Close() error {
	return t.conn.Close()
}
