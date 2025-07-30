package expressionquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
)

type ExpressionQuerier interface {
	InstantQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[InstantQueryArguments]) (*mcp.CallToolResultFor[InstantQueryResult], error)
	RangeQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[RangeQueryArguments]) (*mcp.CallToolResultFor[RangeQueryResult], error)
}

func NewExpressionQuerier(cli client.PrometheusClient) ExpressionQuerier {
	return &expressionQuerier{Client: cli}
}

type expressionQuerier struct {
	Client client.PrometheusClient
}

var _ ExpressionQuerier = &expressionQuerier{}
