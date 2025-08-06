package alertquery

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// curl http://localhost:9090/api/v1/alerts
type AlertQueryParams struct{}

type AlertQueryResult = v1.AlertsResult

func (d *alertQuerier) AlertQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[AlertQueryParams]) (*mcp.CallToolResultFor[AlertQueryResult], error) {
	var (
		result AlertQueryResult
		err    error
	)
	if result, err = d.API.Alerts(ctx); err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[AlertQueryResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
