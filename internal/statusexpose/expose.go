package statusexpose

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/prometheus/api"
)

type StatusExposer interface {
	ConfigExposeHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[ConfigExposeParams]) (*mcp.CallToolResultFor[ConfigExposeResult], error)
	FlagsExposeHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[FlagsExposeParams]) (*mcp.CallToolResultFor[FlagsExposeResult], error)
	RuntimeInformationExposeHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[RuntimeInformationExposeParams]) (*mcp.CallToolResultFor[RuntimeInformationExposeResult], error)
	BuildInformationExposeHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[BuildInformationExposeParams]) (*mcp.CallToolResultFor[BuildInformationExposeResult], error)
	TSDBStatsExposeHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[TSDBStatsExposeParams]) (*mcp.CallToolResultFor[TSDBStatsExposeResult], error)
	WALReplayStatsExposeHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[WALReplayStatsExposeParams]) (*mcp.CallToolResultFor[WALReplayStatsExposeResult], error)
}

func NewStatusExposer(api api.PrometheusAPI) StatusExposer {
	return &statusExposer{
		API: api,
	}
}

type statusExposer struct {
	API api.PrometheusAPI
}

var _ StatusExposer = &statusExposer{}
