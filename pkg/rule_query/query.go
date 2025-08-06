package rulequery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/api"
)

type RuleQuerier interface {
	RuleQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[struct{}]) (*mcp.CallToolResultFor[RuleQueryResult], error)
}

func NewRuleQuerier(api api.PrometheusAPI) RuleQuerier {
	return &ruleQuerier{API: api}
}

type ruleQuerier struct {
	API api.PrometheusAPI
}

var _ RuleQuerier = &ruleQuerier{}
