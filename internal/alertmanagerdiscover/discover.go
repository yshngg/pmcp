package alertmanagerdiscover

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/prometheus-mcp-server/internal/prometheus/api"
)

type AlertmanagerDiscoverer interface {
	AlertmanagerDiscoverHandler(ctx context.Context, request *mcp.CallToolRequest, input *AlertmanagerDiscoverParams) (*mcp.CallToolResult, *AlertmanagerDiscoverResult, error)
}

// NewAlertmanagerDiscoverer returns an AlertmanagerDiscoverer backed by the provided PrometheusAPI.
func NewAlertmanagerDiscoverer(api api.PrometheusAPI) AlertmanagerDiscoverer {
	return &alertmanagerDiscoverer{API: api}
}

type alertmanagerDiscoverer struct {
	API api.PrometheusAPI
}

var _ AlertmanagerDiscoverer = &alertmanagerDiscoverer{}
