package statusexpose

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type TSDBStatsExposeParams struct{}

type TSDBStatsExposeResult = v1.TSDBResult

func (e *statusExposer) TSDBStatsExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *TSDBStatsExposeParams) (*mcp.CallToolResult, *TSDBStatsExposeResult, error) {
	result, err := e.API.TSDB(ctx)
	if err != nil {
		return nil, nil, err
	}
	return nil, &result, nil
}
