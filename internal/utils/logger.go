package utils

import (
	"context"
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
			Level: slog.LevelDebug, // 默认记录所有级别
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
	l.logger.Log(nil, slog.LevelInfo+1, msg, args...)
}

// Warn 记录警告信息
func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error 记录错误信息
func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// Log 根据指定级别动态记录日志
func (l *Logger) Log(ctx context.Context, level string, msg string, args ...any) {
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
	case "critical":
		slogLevel = slog.LevelError + 1
	case "alert":
		slogLevel = slog.LevelError + 2
	case "emergency":
		slogLevel = slog.LevelError + 3
	default:
		slogLevel = slog.LevelInfo
	}
	l.logger.Log(ctx, slogLevel, msg, args...)
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
	case "critical":
		slogLevel = slog.LevelError + 1
	case "alert":
		slogLevel = slog.LevelError + 2
	case "emergency":
		slogLevel = slog.LevelError + 3
	default:
		slogLevel = slog.LevelInfo
	}
	l.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLevel,
	}))
}
