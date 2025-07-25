package expressionquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
)

type ExpressionQuerier interface {
	InstantQueryHandler(context.Context, *mcp.ServerSession, *mcp.CallToolParamsFor[InstantQueryArguments]) (*mcp.CallToolResultFor[InstantQueryResult], error)
	RangeQueryHandler(context.Context, *mcp.ServerSession, *mcp.CallToolParamsFor[RangeQueryArguments]) (*mcp.CallToolResultFor[RangeQueryResult], error)
}

func NewExpressionQuerier(cli client.PrometheusClient) ExpressionQuerier {
	return &expressionQuerier{Client: cli}
}

type expressionQuerier struct {
	Client client.PrometheusClient
}

var _ ExpressionQuerier = &expressionQuerier{}
