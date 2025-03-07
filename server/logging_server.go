package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

// SendLog 发送日志通知
func (s *DefaultServer) SendLog(ctx context.Context, level, msg string) error {
	notification := protocol.NewLoggingMessageNotification(level, msg)
	return s.transport.Send(ctx, notification)
}
