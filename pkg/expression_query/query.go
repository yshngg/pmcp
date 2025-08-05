package expressionquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/api"
)

type ExpressionQuerier interface {
	InstantQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[InstantQueryArguments]) (*mcp.CallToolResultFor[InstantQueryResult], error)
	RangeQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[RangeQueryArguments]) (*mcp.CallToolResultFor[RangeQueryResult], error)
}

func NewExpressionQuerier(api api.PrometheusAPI) ExpressionQuerier {
	return &expressionQuerier{API: api}
}

type expressionQuerier struct {
	API api.PrometheusAPI
}

var _ ExpressionQuerier = &expressionQuerier{}
