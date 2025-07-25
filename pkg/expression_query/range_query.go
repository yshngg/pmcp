package expressionquery

import (
	"context"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

const RangeQueryEndpoint = "/query_range"

// URL query parameters:
// query=<string>: Prometheus expression query string.
// start=<rfc3339 | unix_timestamp>: Start timestamp, inclusive.
// end=<rfc3339 | unix_timestamp>: End timestamp, inclusive.
// step=<duration | float>: Query resolution step width in duration format or float number of seconds.
// timeout=<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag.
// limit=<number>: Maximum number of returned series. Optional. 0 means disabled.
type RangeQueryArguments struct {
	Query   string        `jsonschema:"<string>: Prometheus expression query string."`
	Start   string        `jsonschema:"<rfc3339 | unix_timestamp>: Start timestamp, inclusive."`
	End     string        `jsonschema:"<rfc3339 | unix_timestamp>: End timestamp, inclusive."`
	Step    time.Duration `jsonschema:"<duration | float>: Query resolution step width in duration format or float number of seconds."`
	Timeout time.Duration `jsonschema:"<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag."`
	Limit   uint64        `jsonschema:"<number>: Maximum number of returned series. Optional. 0 means disabled."`
}

type RangeQueryResult struct{}

func (q *expressionQuerier) RangeQueryHandler(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[RangeQueryArguments]) (*mcp.CallToolResultFor[RangeQueryResult], error) {
	start, err := time.Parse(time.RFC3339, params.Arguments.Start)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(time.RFC3339, params.Arguments.End)
	if err != nil {
		return nil, err
	}

	q.Client.QueryRange(
		ctx,
		params.Arguments.Query,
		v1.Range{
			Start: start,
			End:   end,
			Step:  params.Arguments.Step,
		},
		v1.WithTimeout(params.Arguments.Timeout),
		v1.WithLimit(params.Arguments.Limit),
	)
	return &mcp.CallToolResultFor[RangeQueryResult]{
		Content:           []mcp.Content{},
		StructuredContent: RangeQueryResult{},
	}, nil
}
