package server

import (
	"context"

	"github.com/songzhibin97/mcp/protocol"
)

func (s *DefaultServer) HandleListResources(ctx context.Context, req *protocol.JSONRPCRequest) (interface{}, error) {
	var cursor protocol.Cursor
	if c, ok := req.Params["cursor"]; ok && c != nil {
		cursor = protocol.Cursor(c.(string))
	}

	// 模拟资源数据
	allResources := []protocol.Resource{
		{URI: "file://test1", Name: "Test Resource 1"},
		{URI: "file://test2", Name: "Test Resource 2"},
		{URI: "file://test3", Name: "Test Resource 3"},
		{URI: "file://test4", Name: "Test Resource 4"},
	}

	const pageSize = 2
	startIdx := 0
	if cursor != "" {
		for i, r := range allResources {
			if r.URI == string(cursor) {
				startIdx = i
				break
			}
		}
	}

	endIdx := startIdx + pageSize
	if endIdx > len(allResources) {
		endIdx = len(allResources)
	}

	result := protocol.ListResourcesResult{
		Resources: allResources[startIdx:endIdx],
	}
	if endIdx < len(allResources) {
		next := protocol.Cursor(allResources[endIdx].URI)
		result.NextCursor = &next
	}

	return result, nil
}
