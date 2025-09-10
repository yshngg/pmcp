package statusexpose

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type ConfigExposeParams struct{}

type ConfigExposeResult = v1.ConfigResult

func (e *statusExposer) ConfigExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *ConfigExposeParams) (*mcp.CallToolResult, *ConfigExposeResult, error) {
	result, err := e.API.Config(ctx)
	if err != nil {
		return nil, nil, err
	}
	return nil, &result, nil
}
