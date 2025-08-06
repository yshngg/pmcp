package alertquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/api"
)

type AlertQuerier interface {
	AlertQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[AlertQueryParams]) (*mcp.CallToolResultFor[AlertQueryResult], error)
}

func NewAlertQueryer(api api.PrometheusAPI) AlertQuerier {
	return &alertQuerier{API: api}
}

type alertQuerier struct {
	API api.PrometheusAPI
}

var _ AlertQuerier = &alertQuerier{}
