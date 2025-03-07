package protocol

// ProgressToken 用于关联进度通知与原始请求，可以是字符串或数字
type ProgressToken string

// Cursor 用于分页的Opaque Token
type Cursor string

// RequestId 是 JSON-RPC 请求的唯一标识符，可以是字符串或数字
type RequestId string

// TextContent 表示文本内容
type TextContent struct {
	Type string `json:"type"` // 固定为 "text"
	Text string `json:"text"`
}
