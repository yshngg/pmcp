package manage

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (m *manager) HealthCheckHandler(ctx context.Context, request *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, *ManagementResult, error) {
	result := &ManagementResult{
		Success: true,
	}

	err := m.api.HealthCheck(ctx)
	if err != nil {
		result.Success = false
		result.Message = err.Error()
	}
	return nil, result, nil
}
