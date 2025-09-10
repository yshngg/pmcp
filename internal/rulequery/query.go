package rulequery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/prometheus/api"
)

type RuleQuerier interface {
	RuleQueryHandler(ctx context.Context, request *mcp.CallToolRequest, input *RuleQueryArguments) (*mcp.CallToolResult, *RuleQueryResult, error)
}

// NewRuleQuerier returns a RuleQuerier implementation that uses the provided PrometheusAPI to execute rule queries.
func NewRuleQuerier(api api.PrometheusAPI) RuleQuerier {
	return &ruleQuerier{API: api}
}

type ruleQuerier struct {
	API api.PrometheusAPI
}

var _ RuleQuerier = &ruleQuerier{}
