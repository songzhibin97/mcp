package transport

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	util "github.com/songzhibin97/mcp/internal/utils"
	"github.com/songzhibin97/mcp/protocol"
)

// HTTPTransport 实现基于 HTTP 的传输
type HTTPTransport struct {
	client *http.Client
	url    string
	base   *JSONRPCTransport
	logger *util.Logger
}

// NewHTTPTransport 创建新的 HTTPTransport
func NewHTTPTransport(url string) *HTTPTransport {
	return &HTTPTransport{
		client: &http.Client{},
		url:    url,
		base:   NewJSONRPCTransport(),
		logger: util.NewLogger(),
	}
}

// Send 实现 Transport 接口的发送方法
func (t *HTTPTransport) Send(ctx context.Context, msg interface{}) error {
	return t.base.SendMsg(ctx, msg.(protocol.JSONRPCMessage), func(data []byte) error {
		req, err := http.NewRequestWithContext(ctx, "POST", t.url, bytes.NewReader(data))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := t.client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status: %d", resp.StatusCode)
		}
		return nil
	})
}

// Receive 实现 Transport 接口的接收方法（HTTP 客户端通常不直接接收）
func (t *HTTPTransport) Receive(ctx context.Context) (interface{}, error) {
	return nil, fmt.Errorf("HTTP transport does not support direct receiving")
}

// Close 关闭传输
func (t *HTTPTransport) Close() error {
	t.client.CloseIdleConnections()
	return nil
}
