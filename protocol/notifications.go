package protocol

// CancelledNotification 取消通知
type CancelledNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Notification
}

// NewCancelledNotification 创建新的取消通知
func NewCancelledNotification(requestId RequestId, reason string) *CancelledNotification {
	params := map[string]interface{}{
		"requestId": requestId,
	}
	if reason != "" {
		params["reason"] = reason
	}
	return &CancelledNotification{
		JSONRPC: JSONRPCVersion,
		Notification: Notification{
			Method: "notifications/cancelled",
			Params: params,
		},
	}
}

// ProgressNotification 进度通知
type ProgressNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Notification
}

// NewProgressNotification 创建新的进度通知
func NewProgressNotification(token ProgressToken, progress, total int) *ProgressNotification {
	params := map[string]interface{}{
		"progressToken": token,
		"progress":      progress,
	}
	if total > 0 {
		params["total"] = total
	}
	return &ProgressNotification{
		JSONRPC: JSONRPCVersion,
		Notification: Notification{
			Method: "notifications/progress",
			Params: params,
		},
	}
}
