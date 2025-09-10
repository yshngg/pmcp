package manage

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/prometheus/api"
)

type ManagementResult struct {
	Success bool   `json:"success" jsonschema:"Indicate the result of the management operation, true means success, false means failure"`
	Message string `json:"message,omitempty" jsonschema:"Explanation message when the operation fails."`
}

type Manager interface {
	HealthCheckHandler(ctx context.Context, request *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, *ManagementResult, error)
	ReadinessCheckHandler(ctx context.Context, request *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, *ManagementResult, error)
	ReloadHandler(ctx context.Context, request *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, *ManagementResult, error)
	QuitHandler(ctx context.Context, request *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, *ManagementResult, error)
}

type manager struct {
	api api.PrometheusAPI
}

func NewManager(api api.PrometheusAPI) Manager {
	return &manager{api: api}
}

var _ Manager = &manager{}
