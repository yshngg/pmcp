package targetdiscover

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/prometheus/api"
)

const TargetDiscoverEndpoint = "/targets"

type TargetDiscoverer interface {
	TargetDiscoverHandler(ctx context.Context, request *mcp.CallToolRequest, input *TargetDiscoverParams) (*mcp.CallToolResult, *TargetDiscoverResult, error)
}

func NewTargetDiscoverer(api api.PrometheusAPI) TargetDiscoverer {
	return &targetDiscoverer{API: api}
}

type targetDiscoverer struct {
	API api.PrometheusAPI
}

var _ TargetDiscoverer = &targetDiscoverer{}
