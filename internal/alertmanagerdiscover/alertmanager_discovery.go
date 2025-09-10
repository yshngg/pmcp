package alertmanagerdiscover

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type AlertmanagerDiscoverParams struct{}

type AlertmanagerDiscoverResult = v1.AlertManagersResult

func (d *alertmanagerDiscoverer) AlertmanagerDiscoverHandler(ctx context.Context, request *mcp.CallToolRequest, input *AlertmanagerDiscoverParams) (*mcp.CallToolResult, *AlertmanagerDiscoverResult, error) {
	result, err := d.API.AlertManagers(ctx)
	if err != nil {
		return nil, nil, err
	}
	return nil, &result, nil
}
