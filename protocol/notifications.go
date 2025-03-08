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

// ResourceUpdatedNotification 资源更新通知
type ResourceUpdatedNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Notification
}

// NewResourceUpdatedNotification 创建新的资源更新通知
func NewResourceUpdatedNotification(uri string) *ResourceUpdatedNotification {
	return &ResourceUpdatedNotification{
		JSONRPC: JSONRPCVersion,
		Notification: Notification{
			Method: "notifications/resources/updated",
			Params: map[string]interface{}{
				"uri": uri,
			},
		},
	}
}

// ResourceListChangedNotification 资源列表变更通知
type ResourceListChangedNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Notification
}

// NewResourceListChangedNotification 创建新的资源列表变更通知
func NewResourceListChangedNotification() *ResourceListChangedNotification {
	return &ResourceListChangedNotification{
		JSONRPC: JSONRPCVersion,
		Notification: Notification{
			Method: "notifications/resources/list_changed",
		},
	}
}

// PromptListChangedNotification 提示列表变更通知
type PromptListChangedNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Notification
}

// NewPromptListChangedNotification 创建新的提示列表变更通知
func NewPromptListChangedNotification() *PromptListChangedNotification {
	return &PromptListChangedNotification{
		JSONRPC: JSONRPCVersion,
		Notification: Notification{
			Method: "notifications/prompts/list_changed",
		},
	}
}

// ToolListChangedNotification 工具列表变更通知
type ToolListChangedNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Notification
}

// NewToolListChangedNotification 创建新的工具列表变更通知
func NewToolListChangedNotification() *ToolListChangedNotification {
	return &ToolListChangedNotification{
		JSONRPC: JSONRPCVersion,
		Notification: Notification{
			Method: "notifications/tools/list_changed",
		},
	}
}
