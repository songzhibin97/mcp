package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

// SendProgress 发送进度通知
func (s *DefaultServer) SendProgress(ctx context.Context, token protocol.ProgressToken, progress, total int) error {
	notification := protocol.NewProgressNotification(token, progress, total)
	return s.transport.Send(ctx, notification)
}
