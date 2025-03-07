package handlers

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
	"github.com/songzhibin97/mcp/server"
)

func ResourceHandler(s *server.DefaultServer) func(context.Context, *protocol.JSONRPCRequest) (interface{}, error) {
	return s.HandleListResources
}
