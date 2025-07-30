package expressionquery

import (
	"context"
	"encoding/json"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

const InstantQueryEndpoint = "/query"

// URL query parameters:
// query=<string>: Prometheus expression query string.
// time=<rfc3339 | unix_timestamp>: Evaluation timestamp. Optional.
// timeout=<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag.
// limit=<number>: Maximum number of returned series. Doesn’t affect scalars or strings but truncates the number of series for matrices and vectors. Optional. 0 means disabled.
// The following example evaluates the expression up at the time 2015-07-01T20:10:51.781Z:
// curl 'http://localhost:9090/api/v1/query?query=up&time=2015-07-01T20:10:51.781Z'
type InstantQueryArguments struct {
	Query   string        `json:"query" jsonschema:"<string>: Prometheus expression query string."`
	Time    string        `json:"time,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: Evaluation timestamp. Optional."`
	Timeout time.Duration `json:"timeout,omitzero" jsonschema:"<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag."`
	Limit   uint64        `json:"limit,omitzero" jsonschema:"<number>: Maximum number of returned series. Doesn't affect scalars or strings but truncates the number of series for matrices and vectors. Optional. 0 means disabled."`
}

type InstantQueryResult struct {
	Value    model.Value `json:"value" jsonschema:"<value> refers to the query result data, which has varying formats depending on the resultType. See the [expression query result formats](https://prometheus.io/docs/prometheus/latest/querying/api/#expression-query-result-formats)."`
	Warnings v1.Warnings `json:"warnings,omitempty"`
}

func (q *expressionQuerier) InstantQueryHandler(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[InstantQueryArguments]) (*mcp.CallToolResultFor[InstantQueryResult], error) {
	var (
		ts  time.Time
		err error
	)
	if len(params.Arguments.Time) != 0 {
		ts, err = time.Parse(time.RFC3339, params.Arguments.Time)
		if err != nil {
			return nil, err
		}
	}

	opts := make([]v1.Option, 0)
	if params.Arguments.Timeout != 0 {
		opts = append(opts, v1.WithTimeout(params.Arguments.Timeout))
	}
	if params.Arguments.Limit != 0 {
		opts = append(opts, v1.WithLimit(params.Arguments.Limit))
	}

	result := InstantQueryResult{}
	if result.Value, result.Warnings, err = q.Client.Query(
		ctx,
		params.Arguments.Query,
		ts,
		opts...,
	); err != nil {
		return nil, err
	}
	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return &mcp.CallToolResultFor[InstantQueryResult]{
		Content: []mcp.Content{&mcp.TextContent{
			Text: string(content),
		}},
		StructuredContent: result,
	}, nil
}
