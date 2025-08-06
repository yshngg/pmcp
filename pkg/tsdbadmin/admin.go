package tsdbadmin

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/api"
)

type TSDBAdmin interface {
	SnapshotHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[SnapshotParams]) (*mcp.CallToolResultFor[SnapshotResult], error)
	DeleteSeriesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteSeriesParams]) (*mcp.CallToolResultFor[DeleteSeriesResult], error)
	CleanTombstonesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[CleanTombstonesParams]) (*mcp.CallToolResultFor[CleanTombstonesResult], error)
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
