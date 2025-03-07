# MCP - 模型通信协议实现

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 项目概述

**MCP（Model Communication Protocol）** 是一个基于 JSON-RPC 的通信协议，旨在实现客户端与服务端之间的高效交互。本项目使用 Go 语言实现，提供完整的客户端和服务端功能，支持资源管理、提示获取、工具调用和消息采样等特性，适用于分布式系统和模型服务的通信需求。

## 功能特性

- **双向通信**：基于 JSON-RPC 协议，支持客户端请求和服务端通知。
- **传输支持**：
  - 模拟传输（`MockTransport`）用于开发和测试。
  - HTTP 和 WebSocket 传输支持真实场景部署。
- **核心功能**：
  - 初始化（`initialize`）
  - 资源列举（`resources/list`）
  - 提示获取（`prompts/get`）
  - 工具调用（`tools/call`）
  - 消息采样（`sampling/createMessage`）
  - 通知支持（进度通知、日志通知）
- **模块化设计**：协议定义、客户端、服务端和传输层分离，便于维护和扩展。

## 目录结构

```
mcp/
├── README-zh.md                        # 项目中文说明文档
├── README.md                           # 项目英文说明文档
├── client                              # 客户端实现
│   ├── client.go                       # 客户端核心逻辑
│   ├── handlers                        # 客户端请求处理器
│   │   ├── prompt_handler.go           # 提示相关处理器
│   │   ├── resource_handler.go         # 资源相关处理器
│   │   ├── sampling_handler.go         # 采样相关处理器
│   │   └── tool_handler.go             # 工具相关处理器
│   ├── interface.go                    # 客户端接口定义
│   ├── jsonrpc_client.go               # JSON-RPC 客户端实现
│   ├── notifications.go                # 通知处理逻辑
│   ├── prompt_client.go                # 提示相关客户端功能
│   ├── resource_client.go              # 资源相关客户端功能
│   ├── sampling_client.go              # 采样相关客户端功能
│   └── tool_client.go                  # 工具相关客户端功能
├── example                             # 示例代码
│   ├── client                          # 客户端示例
│   │   ├── full_client                 # 完整客户端示例
│   │   │   └── full_client.go          # 完整客户端实现
│   │   ├── initialize                  # 初始化示例
│   │   │   └── initialize.go           # 初始化客户端实现
│   │   ├── minimal_client              # 最小化客户端示例
│   │   │   └── minimal_client.go       # 最小化客户端实现
│   │   └── resources                   # 资源示例
│   │       └── resources.go            # 资源客户端实现
│   └── server                          # 服务端示例
│       ├── full_server                 # 完整服务端示例
│       │   └── full_server.go          # 完整服务端实现
│       ├── initialize                  # 初始化服务端示例
│       │   └── initialize.go           # 初始化服务端实现
│       ├── minimal_server              # 最小化服务端示例
│       │   └── minimal_server.go       # 最小化服务端实现
│       └── tools                       # 工具服务端示例
│           └── tools.go                # 工具服务端实现
├── go.mod                              # Go 模块定义文件
├── go.sum                              # Go 依赖校验文件
├── internal                            # 内部工具
│   ├── jsonrpc                         # JSON-RPC 相关工具
│   │   ├── codec.go                    # JSON-RPC 编解码
│   │   └── validator.go                # JSON-RPC 验证器
│   └── utils                           # 通用工具
│       ├── config.go                   # 配置管理
│       └── logger.go                   # 日志工具
├── protocol                            # 协议定义
│   ├── capabilities.go                 # 能力定义
│   ├── completion.go                   # 补全相关定义
│   ├── errors.go                       # 错误定义
│   ├── initialize.go                   # 初始化请求/响应
│   ├── jsonrpc.go                      # JSON-RPC 基础结构
│   ├── logging.go                      # 日志通知定义
│   ├── notifications.go                # 通知定义
│   ├── pagination.go                   # 分页支持
│   ├── prompts.go                      # 提示定义
│   ├── requests.go                     # 请求基础结构
│   ├── resources.go                    # 资源定义
│   ├── sampling.go                     # 采样定义
│   ├── tools.go                        # 工具定义
│   └── types.go                        # 通用类型
├── schema.ts                           # TypeScript 协议定义（可选）
├── server                              # 服务端实现
│   ├── handlers                        # 服务端请求处理器
│   │   ├── prompt_handler.go           # 提示相关处理器
│   │   ├── resource_handler.go         # 资源相关处理器
│   │   ├── sampling_handler.go         # 采样相关处理器
│   │   └── tool_handler.go             # 工具相关处理器
│   ├── interface.go                    # 服务端接口定义
│   ├── jsonrpc_server.go               # JSON-RPC 服务端实现
│   ├── logging_server.go               # 日志功能
│   ├── notifications.go                # 通知功能
│   ├── prompt_server.go                # 提示相关服务端功能
│   ├── resource_server.go              # 资源相关服务端功能
│   ├── sampling_server.go              # 采样相关服务端功能
│   ├── server.go                       # 服务端核心逻辑
│   └── tool_server.go                  # 工具相关服务端功能
└── transport                           # 传输层
├── http_transport.go               # HTTP 传输实现
├── interface.go                    # 传输接口定义
├── jsonrpc_transport.go            # JSON-RPC 传输实现
├── mock_transport.go               # 模拟传输实现
└── websocket_transport.go          # WebSocket 传输实现
```

## 安装指南

### 前提条件
- Go 语言版本：1.21 或更高

### 安装步骤
1. 克隆仓库：
   ```bash
   git clone https://github.com/songzhibin97/mcp.git
   cd mcp
   ```
2. 安装依赖：
   ```bash
   go mod tidy
   ```

## 使用说明

### 示例运行

#### 最小化客户端
运行一个简单的 Ping 请求：
```bash
go run example/client/minimal_client/minimal_client.go
```
预期输出：
```
Ping successful
```

#### 最小化服务端
启动一个响应 Ping 的服务端：
```bash
go run example/server/minimal_server/minimal_server.go
```
预期输出：
```
Server started
[3秒后]
Server stopped
```

#### 完整客户端
展示所有客户端功能：
```bash
go run example/client/full_client/full_client.go
```
预期输出：
```
Initialized: mock-server
Resources: 1
Prompt Messages: 1
Tool Content: 1
```

#### 完整服务端
运行支持所有功能的服务端：
```bash
go run example/server/full_server/full_server.go
```
预期输出：
```
Full server started
[处理请求和通知的日志]
Server stopped
```

## 自定义开发

### 创建客户端
```go
package main

import (
    "context"
    "github.com/songzhibin97/mcp/client"
    "github.com/songzhibin97/mcp/protocol"
    "github.com/songzhibin97/mcp/transport"
)

func main() {
    tr := transport.NewMockTransport()
    c := client.NewDefaultClient(tr)
    ctx := context.Background()
    result, err := c.Initialize(ctx, "2024-11-05", protocol.ClientCapabilities{}, protocol.Implementation{Name: "my-client"})
    if err != nil {
        panic(err)
    }
    println("Server:", result.ServerInfo.Name)
}
```

### 创建服务端
```go
package main

import (
    "context"
    "github.com/songzhibin97/mcp/server"
    "github.com/songzhibin97/mcp/transport"
)

func main() {
    tr := transport.NewMockTransport()
    s := server.NewDefaultServer(tr)
    s.RegisterHandler("ping", s.HandlePing)
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    if err := s.Start(ctx); err != nil {
        panic(err)
    }
}
```

## 贡献指南

1. Fork 本仓库。
2. 创建功能分支：
   ```bash
   git checkout -b feature/your-feature
   ```
3. 提交更改：
   ```bash
   git commit -m "添加新功能：your-feature"
   ```
4. 推送分支：
   ```bash
   git push origin feature/your-feature
   ```
5. 提交 Pull Request。

## 许可证

本项目采用 MIT 许可证。

## 致谢

- 使用 Go 开发。
- 参考 JSON-RPC 标准和分布式系统协议设计。

