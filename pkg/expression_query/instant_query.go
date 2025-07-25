package expressionquery

import (
	"context"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

const InstantQueryEndpoint = "/query"

// URL query parameters:
// query=<string>: Prometheus expression query string.
// time=<rfc3339 | unix_timestamp>: Evaluation timestamp. Optional.
// timeout=<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag.
// limit=<number>: Maximum number of returned series. Doesn’t affect scalars or strings but truncates the number of series for matrices and vectors. Optional. 0 means disabled.
type InstantQueryArguments struct {
	Query   string        `jsonschema:"<string>: Prometheus expression query string."`
	Time    string        `jsonschema:"<rfc3339 | unix_timestamp>: Evaluation timestamp. Optional."`
	Timeout time.Duration `jsonschema:"<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag."`
	Limit   uint64        `jsonschema:"<number>: Maximum number of returned series. Doesn't affect scalars or strings but truncates the number of series for matrices and vectors. Optional. 0 means disabled."`
}

type InstantQueryResult struct{}

func (q *expressionQuerier) InstantQueryHandler(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[InstantQueryArguments]) (*mcp.CallToolResultFor[InstantQueryResult], error) {
	ts, err := time.Parse(time.RFC3339, params.Arguments.Time)
	if err != nil {
		return nil, err
	}
	q.Client.Query(
		ctx,
		params.Arguments.Query,
		ts,
		v1.WithTimeout(params.Arguments.Timeout),
		v1.WithLimit(params.Arguments.Limit),
	)
	return &mcp.CallToolResultFor[InstantQueryResult]{
		Content:           []mcp.Content{},
		StructuredContent: InstantQueryResult{},
	}, nil
}
