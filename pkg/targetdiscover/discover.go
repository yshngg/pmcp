package targetdiscover

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/api"
)

const TargetDiscoverEndpoint = "/targets"

type TargetDiscoverer interface {
	TargetDiscoverHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[TargetDiscoverParams]) (*mcp.CallToolResultFor[TargetDiscoverResult], error)
}

func NewTargetDiscoverer(api api.PrometheusAPI) TargetDiscoverer {
	return &targetDiscoverer{API: api}
}

type targetDiscoverer struct {
	API api.PrometheusAPI
}

var _ TargetDiscoverer = &targetDiscoverer{}
