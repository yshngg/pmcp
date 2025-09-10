package alertquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// curl http://localhost:9090/api/v1/alerts
type AlertQueryParams struct{}

type AlertQueryResult = v1.AlertsResult

func (d *alertQuerier) AlertQueryHandler(ctx context.Context, request *mcp.CallToolRequest, input *AlertQueryParams) (*mcp.CallToolResult, *AlertQueryResult, error) {
	var (
		result = &AlertQueryResult{}
		err    error
	)
	if *result, err = d.API.Alerts(ctx); err != nil {
		return nil, nil, err
	}
	return nil, result, nil
}
