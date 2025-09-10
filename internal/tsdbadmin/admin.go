package tsdbadmin

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/prometheus/api"
)

type TSDBAdmin interface {
	SnapshotHandler(ctx context.Context, request *mcp.CallToolRequest, input *SnapshotParams) (*mcp.CallToolResult, *SnapshotResult, error)
	DeleteSeriesHandler(ctx context.Context, request *mcp.CallToolRequest, input *DeleteSeriesParams) (*mcp.CallToolResult, *DeleteSeriesResult, error)
	CleanTombstonesHandler(ctx context.Context, request *mcp.CallToolRequest, input *CleanTombstonesParams) (*mcp.CallToolResult, *CleanTombstonesResult, error)
}

func NewTSDBAdmin(api api.PrometheusAPI) TSDBAdmin {
	return &tsdbAdmin{
		API: api,
	}
}

type tsdbAdmin struct {
	API api.PrometheusAPI
}

var _ TSDBAdmin = &tsdbAdmin{}
