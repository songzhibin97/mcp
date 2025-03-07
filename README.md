# MCP - Model Communication Protocol Implementation

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## Project Overview

**MCP (Model Communication Protocol)** is a communication protocol based on JSON-RPC, designed for efficient interaction between clients and servers. This project is implemented in Go and provides full client and server functionalities, supporting resource management, prompt retrieval, tool invocation, and message sampling. It is suitable for communication in distributed systems and model services.

## Features

- **Bidirectional Communication**: Based on the JSON-RPC protocol, supporting both client requests and server notifications.
- **Transport Support**:
  - Mock Transport (`MockTransport`) for development and testing.
  - HTTP and WebSocket transport for real-world deployments.
- **Core Functions**:
  - Initialization (`initialize`)
  - Resource Listing (`resources/list`)
  - Prompt Retrieval (`prompts/get`)
  - Tool Invocation (`tools/call`)
  - Message Sampling (`sampling/createMessage`)
  - Notifications (Progress and Log Notifications)
- **Modular Design**: Separation of protocol definition, client, server, and transport layers for maintainability and extensibility.

## Directory Structure

```
mcp/
├── README-zh.md                        # Project documentation in Chinese
├── README.md                           # Project documentation in English
├── client                              # Client implementation
│   ├── client.go                       # Core client logic
│   ├── handlers                        # Client request handlers
│   │   ├── prompt_handler.go           # Prompt-related handler
│   │   ├── resource_handler.go         # Resource-related handler
│   │   ├── sampling_handler.go         # Sampling-related handler
│   │   └── tool_handler.go             # Tool-related handler
│   ├── interface.go                    # Client interface definition
│   ├── jsonrpc_client.go               # JSON-RPC client implementation
│   ├── notifications.go                # Notification handling logic
│   ├── prompt_client.go                # Prompt-related client functions
│   ├── resource_client.go              # Resource-related client functions
│   ├── sampling_client.go              # Sampling-related client functions
│   └── tool_client.go                  # Tool-related client functions
├── example                             # Example code
│   ├── client                          # Client examples
│   │   ├── full_client                 # Full client example
│   │   │   └── full_client.go          # Full client implementation
│   │   ├── initialize                  # Initialization example
│   │   │   └── initialize.go           # Initialization client implementation
│   │   ├── minimal_client              # Minimal client example
│   │   │   └── minimal_client.go       # Minimal client implementation
│   │   └── resources                   # Resource example
│   │       └── resources.go            # Resource client implementation
│   └── server                          # Server examples
│       ├── full_server                 # Full server example
│       │   └── full_server.go          # Full server implementation
│       ├── initialize                  # Initialization server example
│       │   └── initialize.go           # Initialization server implementation
│       ├── minimal_server              # Minimal server example
│       │   └── minimal_server.go       # Minimal server implementation
│       └── tools                       # Tool server example
│           └── tools.go                # Tool server implementation
├── go.mod                              # Go module definition file
├── go.sum                              # Go dependency checksum file
├── internal                            # Internal utilities
│   ├── jsonrpc                         # JSON-RPC utilities
│   │   ├── codec.go                    # JSON-RPC codec
│   │   └── validator.go                # JSON-RPC validator
│   └── utils                           # General utilities
│       ├── config.go                   # Configuration management
│       └── logger.go                   # Logging utility
├── protocol                            # Protocol definitions
│   ├── capabilities.go                 # Capability definitions
│   ├── completion.go                   # Completion-related definitions
│   ├── errors.go                       # Error definitions
│   ├── initialize.go                   # Initialization request/response
│   ├── jsonrpc.go                      # JSON-RPC base structures
│   ├── logging.go                      # Log notification definitions
│   ├── notifications.go                # Notification definitions
│   ├── pagination.go                    # Pagination support
│   ├── prompts.go                      # Prompt definitions
│   ├── requests.go                     # Request base structures
│   ├── resources.go                    # Resource definitions
│   ├── sampling.go                     # Sampling definitions
│   ├── tools.go                        # Tool definitions
│   └── types.go                        # General types
├── schema.ts                           # TypeScript protocol definition (optional)
├── server                              # Server implementation
│   ├── handlers                        # Server request handlers
│   │   ├── prompt_handler.go           # Prompt-related handler
│   │   ├── resource_handler.go         # Resource-related handler
│   │   ├── sampling_handler.go         # Sampling-related handler
│   │   └── tool_handler.go             # Tool-related handler
│   ├── interface.go                    # Server interface definition
│   ├── jsonrpc_server.go               # JSON-RPC server implementation
│   ├── logging_server.go               # Logging functionality
│   ├── notifications.go                # Notification functionality
│   ├── prompt_server.go                # Prompt-related server functions
│   ├── resource_server.go              # Resource-related server functions
│   ├── sampling_server.go              # Sampling-related server functions
│   ├── server.go                       # Core server logic
│   └── tool_server.go                  # Tool-related server functions
└── transport                           # Transport layer
    ├── http_transport.go               # HTTP transport implementation
    ├── interface.go                    # Transport interface definition
    ├── jsonrpc_transport.go            # JSON-RPC transport implementation
    ├── mock_transport.go               # Mock transport implementation
    └── websocket_transport.go          # WebSocket transport implementation
```

## Installation Guide

### Prerequisites
- Go version: 1.21 or later

### Installation Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/songzhibin97/mcp.git
   cd mcp
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage Guide

### Running Examples

#### Minimal Client
Run a simple Ping request:
```bash
go run example/client/minimal_client/minimal_client.go
```
Expected output:
```
Ping successful
```

#### Minimal Server
Start a server responding to Ping requests:
```bash
go run example/server/minimal_server/minimal_server.go
```
Expected output:
```
Server started
[After 3 seconds]
Server stopped
```

## Contribution Guide

1. Fork this repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add new feature: your-feature"
   ```
4. Push the branch:
   ```bash
   git push origin feature/your-feature
   ```
5. Submit a Pull Request.

## License

This project is licensed under the MIT License.

## Acknowledgements

- Developed using Go.
- References JSON-RPC standards and distributed system protocol design.

