package client

import (
	"context"
	"fmt"

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
					token, _ := n.Params["progressToken"].(string)
					progress, _ := n.Params["progress"].(float64) // JSON 数字默认解码为 float64
					total, totalOk := n.Params["total"].(float64)
					if totalOk {
						c.logger.Info("Progress update", "token", token, "progress", progress, "total", total)
					} else {
						c.logger.Info("Progress update", "token", token, "progress", progress)
					}

				case *protocol.CancelledNotification:
					requestId, _ := n.Params["requestId"].(string)
					reason, reasonOk := n.Params["reason"].(string)
					if reasonOk {
						c.logger.Info("Request cancelled", "requestId", requestId, "reason", reason)
					} else {
						c.logger.Info("Request cancelled", "requestId", requestId)
					}
					// 可选：清理 pendingRequests 中对应的请求
					c.mu.Lock()
					if ch, ok := c.pendingRequests[protocol.RequestId(requestId)]; ok {
						ch <- response{result: nil, err: fmt.Errorf("request cancelled by server: %s", reason)}
						delete(c.pendingRequests, protocol.RequestId(requestId))
					}
					c.mu.Unlock()

				case *protocol.LoggingMessageNotification:
					level, _ := n.Params["level"].(string)
					data, _ := n.Params["data"]
					loggerName, loggerOk := n.Params["logger"].(string)
					if loggerOk {
						c.logger.Log(ctx, level, "Server log", "logger", loggerName, "data", data)
					} else {
						c.logger.Log(ctx, level, "Server log", "data", data)
					}

				case *protocol.ResourceUpdatedNotification:
					uri, _ := n.Params["uri"].(string)
					c.logger.Info("Resource updated", "uri", uri)
					// 可选：客户端可以重新读取该资源

				case *protocol.ResourceListChangedNotification:
					c.logger.Info("Resource list changed")
					// 可选：客户端可以重新调用 ListResources

				case *protocol.PromptListChangedNotification:
					c.logger.Info("Prompt list changed")
					// 可选：客户端可以重新调用 ListPrompts

				case *protocol.ToolListChangedNotification:
					c.logger.Info("Tool list changed")
					// 可选：客户端可以重新调用 ListTools

				case *protocol.InitializedNotification:
					c.logger.Info("Server initialized")

				default:
					c.logger.Warn("Received unknown notification type", "type", fmt.Sprintf("%T", msg))
				}
			}
		}
	}()
}
