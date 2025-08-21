package alertmanagerdiscover

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/prometheus/api"
)

type AlertmanagerDiscoverer interface {
	AlertmanagerDiscoverHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[AlertmanagerDiscoverParams]) (*mcp.CallToolResultFor[AlertmanagerDiscoverResult], error)
}

func NewAlertmanagerDiscoverer(api api.PrometheusAPI) AlertmanagerDiscoverer {
	return &alertmanagerDiscoverer{API: api}
}

type alertmanagerDiscoverer struct {
	API api.PrometheusAPI
}

var _ AlertmanagerDiscoverer = &alertmanagerDiscoverer{}
