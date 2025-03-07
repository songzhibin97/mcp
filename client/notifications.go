package client

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

// HandleNotifications 处理通知（需在客户端运行循环中调用）
func (c *DefaultClient) HandleNotifications(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := c.transport.Receive(ctx)
				if err != nil {
					c.logger.Error("Failed to receive notification", "error", err)
					continue
				}
				switch n := msg.(type) {
				case *protocol.ProgressNotification:
					c.logger.Info("Progress update", "token", n.Params["progressToken"], "progress", n.Params["progress"])
				}
			}
		}
	}()
}
