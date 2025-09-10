package statusexpose

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/prometheus/api"
)

type StatusExposer interface {
	ConfigExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *ConfigExposeParams) (*mcp.CallToolResult, *ConfigExposeResult, error)
	FlagsExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *FlagsExposeParams) (*mcp.CallToolResult, *FlagsExposeResult, error)
	RuntimeInformationExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *RuntimeInformationExposeParams) (*mcp.CallToolResult, *RuntimeInformationExposeResult, error)
	BuildInformationExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *BuildInformationExposeParams) (*mcp.CallToolResult, *BuildInformationExposeResult, error)
	TSDBStatsExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *TSDBStatsExposeParams) (*mcp.CallToolResult, *TSDBStatsExposeResult, error)
	WALReplayStatsExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *WALReplayStatsExposeParams) (*mcp.CallToolResult, *WALReplayStatsExposeResult, error)
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
