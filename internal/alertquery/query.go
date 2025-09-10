package alertquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/prometheus/api"
)

type AlertQuerier interface {
	AlertQueryHandler(ctx context.Context, request *mcp.CallToolRequest, input *AlertQueryParams) (*mcp.CallToolResult, *AlertQueryResult, error)
}

func NewAlertQuerier(api api.PrometheusAPI) AlertQuerier {
	return &alertQuerier{API: api}
}

type alertQuerier struct {
	API api.PrometheusAPI
}

var _ AlertQuerier = &alertQuerier{}
