package alertmanagerdiscover

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type AlertmanagerDiscoverParams struct{}

type AlertmanagerDiscoverResult = v1.AlertManagersResult

func (d *alertmanagerDiscoverer) AlertmanagerDiscoverHandler(ctx context.Context, _ *mcp.ServerSession, _ *mcp.CallToolParamsFor[AlertmanagerDiscoverParams]) (*mcp.CallToolResultFor[AlertmanagerDiscoverResult], error) {
	var (
		result AlertmanagerDiscoverResult
		err    error
	)

	if result, err = d.API.AlertManagers(ctx); err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[AlertmanagerDiscoverResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
