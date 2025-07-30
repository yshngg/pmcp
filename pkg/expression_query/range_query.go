package expressionquery

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
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
	Query   string        `json:"query" jsonschema:"<string>: Prometheus expression query string."`
	Start   string        `json:"start" jsonschema:"<rfc3339 | unix_timestamp>: Start timestamp, inclusive."`
	End     string        `json:"end" jsonschema:"<rfc3339 | unix_timestamp>: End timestamp, inclusive."`
	Step    time.Duration `json:"step" jsonschema:"<duration | float>: Query resolution step width in duration format or float number of seconds."`
	Timeout time.Duration `json:"timeout,omitzero" jsonschema:"<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag."`
	Limit   uint64        `json:"limit,omitzero" jsonschema:"<number>: Maximum number of returned series. Optional. 0 means disabled."`
}

type RangeQueryResult struct {
	Value    model.Value `json:"value"`
	Warnings v1.Warnings `json:"warnings,omitempty"`
}

func (q *expressionQuerier) RangeQueryHandler(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[RangeQueryArguments]) (*mcp.CallToolResultFor[RangeQueryResult], error) {
	var (
		start, end time.Time
		step       time.Duration
		err        error
	)
	if len(params.Arguments.Start) != 0 {
		start, err = time.Parse(time.RFC3339, params.Arguments.Start)
		if err != nil {
			return nil, err
		}
	}
	if len(params.Arguments.End) != 0 {
		end, err = time.Parse(time.RFC3339, params.Arguments.End)
		if err != nil {
			return nil, err
		}
	}
	if params.Arguments.Step == 0 {
		return nil, errors.New("step cannot be 0")
	}
	step = params.Arguments.Step * time.Second

	opts := make([]v1.Option, 0)
	if params.Arguments.Timeout != 0 {
		opts = append(opts, v1.WithTimeout(params.Arguments.Timeout))
	}
	if params.Arguments.Limit != 0 {
		opts = append(opts, v1.WithLimit(params.Arguments.Limit))
	}

	result := RangeQueryResult{}
	if result.Value, result.Warnings, err = q.Client.QueryRange(
		ctx,
		params.Arguments.Query,
		v1.Range{
			Start: start,
			End:   end,
			Step:  step,
		},
		opts...,
	); err != nil {
		return nil, err
	}
	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return &mcp.CallToolResultFor[RangeQueryResult]{
		Content: []mcp.Content{&mcp.TextContent{
			Text: string(content),
		}},
		StructuredContent: result,
	}, nil
}
