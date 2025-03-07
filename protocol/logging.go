package protocol

// SetLevelRequest 设置日志级别
type SetLevelRequest struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      RequestId `json:"id"`
	Request
}

// NewSetLevelRequest 创建新的设置日志级别请求
func NewSetLevelRequest(id RequestId, level string) *SetLevelRequest {
	return &SetLevelRequest{
		JSONRPC: JSONRPCVersion,
		ID:      id,
		Request: Request{
			Method: "logging/setLevel",
			Params: map[string]interface{}{
				"level": level,
			},
		},
	}
}

// LoggingMessageNotification 日志通知
type LoggingMessageNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Notification
}

// NewLoggingMessageNotification 创建新的日志通知
func NewLoggingMessageNotification(level, msg string) *LoggingMessageNotification {
	return &LoggingMessageNotification{
		JSONRPC: JSONRPCVersion,
		Notification: Notification{
			Method: "notifications/message",
			Params: map[string]interface{}{
				"level": level,
				"data":  msg,
			},
		},
	}
}
