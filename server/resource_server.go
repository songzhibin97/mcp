package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

func (s *DefaultServer) HandleListResources(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
	return &protocol.ListResourcesResult{
		Resources: []protocol.Resource{
			{URI: "file://example", Name: "Example Resource"},
		},
	}, nil
}
