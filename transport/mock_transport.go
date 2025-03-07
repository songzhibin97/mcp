package transport

import (
	"context"
	"encoding/json"
	"fmt"

	util "github.com/songzhibin97/mcp/internal/utils"
)

// MockTransport 实现用于测试的模拟传输
type MockTransport struct {
	sendChan    chan interface{}
	receiveChan chan interface{}
	logger      *util.Logger
	closed      bool
}

// NewMockTransport 创建新的 MockTransport
func NewMockTransport() *MockTransport {
	return &MockTransport{
		sendChan:    make(chan interface{}, 10),
		receiveChan: make(chan interface{}, 10),
		logger:      util.NewLogger(),
	}
}

func (t *MockTransport) Send(ctx context.Context, msg interface{}) error {
	if t.closed {
		return fmt.Errorf("mock transport closed")
	}
	data, _ := json.Marshal(msg)
	t.logger.Debug("Mock sent", "msg", string(data))
	select {
	case t.sendChan <- msg:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *MockTransport) Receive(ctx context.Context) (interface{}, error) {
	if t.closed {
		return nil, fmt.Errorf("mock transport closed")
	}
	select {
	case msg := <-t.receiveChan:
		data, _ := json.Marshal(msg)
		t.logger.Debug("Mock received", "msg", string(data))
		return msg, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// Close 关闭传输
func (t *MockTransport) Close() error {
	t.closed = true
	close(t.sendChan)
	close(t.receiveChan)
	return nil
}

// InjectMessage 注入消息以模拟接收
func (t *MockTransport) InjectMessage(msg interface{}) {
	t.receiveChan <- msg
}
