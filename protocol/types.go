package protocol

import "fmt"

// ProgressToken 用于关联进度通知与原始请求，可以是字符串或数字
type ProgressToken string

// Cursor 用于分页的Opaque Token
type Cursor string

// RequestId 是 JSON-RPC 请求的唯一标识符，可以是字符串或数字
type RequestId string

// Content 接口定义了所有内容类型的通用方法
type Content interface {
	Type() string
	String() string
}

// TextContent 表示文本内容
type TextContent struct {
	ContentType string `json:"type"` // 重命名为 ContentType，避免与方法冲突
	Text        string `json:"text"`
}

func (t TextContent) Type() string {
	return t.ContentType
}

func (t TextContent) String() string {
	return t.Text
}

// ImageContent 表示图片内容
type ImageContent struct {
	ContentType string `json:"type"` // 同上
	Data        string `json:"data"`
	MimeType    string `json:"mimeType"`
}

func (t ImageContent) Type() string {
	return t.ContentType
}

func (t ImageContent) String() string {
	return fmt.Sprintf("Image (%s): %s", t.MimeType, t.Data)
}
