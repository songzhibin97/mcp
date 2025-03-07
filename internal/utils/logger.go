package util

import (
	"log/slog"
	"os"
)

// Logger 是一个基于 slog 的结构化日志工具
type Logger struct {
	logger *slog.Logger
}

// NewLogger 创建一个新的 Logger 实例，默认输出到标准输出
func NewLogger() *Logger {
	return &Logger{
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug, // 默认记录所有级别，后续可通过配置调整
		})),
	}
}

// Debug 记录调试信息
func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info 记录常规信息
func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Notice 记录通知信息
func (l *Logger) Notice(msg string, args ...any) {
	l.logger.Log(nil, slog.LevelInfo+1, msg, args...) // Notice 介于 Info 和 Warn 之间
}

// Warn 记录警告信息
func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error 记录错误信息
func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// SetLevel 设置日志级别
func (l *Logger) SetLevel(level string) {
	var slogLevel slog.Level
	switch level {
	case "debug":
		slogLevel = slog.LevelDebug
	case "info":
		slogLevel = slog.LevelInfo
	case "notice":
		slogLevel = slog.LevelInfo + 1
	case "warning":
		slogLevel = slog.LevelWarn
	case "error":
		slogLevel = slog.LevelError
	case "critical", "alert", "emergency": // 映射到 slog 的更高级别
		slogLevel = slog.LevelError + 4
	default:
		slogLevel = slog.LevelInfo // 默认级别
	}
	l.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLevel,
	}))
}
