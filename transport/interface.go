package transport

import "context"

// Transport 定义了传输层的基本接口
type Transport interface {
	// Send 发送消息到远程端
	Send(ctx context.Context, msg interface{}) error
	// Receive 接收来自远程端的消息
	Receive(ctx context.Context) (interface{}, error)
	// Close 关闭传输连接
	Close() error
}
