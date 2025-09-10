package metadataquery

import (
	"context"
	"log/slog"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/yshngg/pmcp/internal/utils"
)

const LabelNamesEndpoint = "/labels"

// URL query parameters:
// start=<rfc3339 | unix_timestamp>: Start timestamp. Optional.
// end=<rfc3339 | unix_timestamp>: End timestamp. Optional.
// match[]=<series_selector>: Repeated series selector argument that selects the series from which to read the label names. Optional.
// limit=<number>: Maximum number of returned series. Optional. 0 means disabled.
// Here is an example:
// curl 'localhost:9090/api/v1/labels'
type LabelNamesArguments struct {
	Start string   `json:"start,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: Start timestamp. Optional."`
	End   string   `json:"end,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: End timestamp. Optional."`
	Match []string `json:"match[],omitzero" jsonschema:"<series_selector>: Repeated series selector argument that selects the series from which to read the label names. Optional."`
	Limit uint64   `json:"limit,omitzero" jsonschema:"<number>: Maximum number of returned series. Optional. 0 means disabled."`
}

type LabelNamesResult struct {
	// TODO: Replace []string with model.LabelNames (see https://github.com/prometheus/client_golang/pull/1850).
	LabelNames []string    `json:"names" jsonschema:"Names is a list of string label names."`
	Warnings   v1.Warnings `json:"warnings,omitempty"`
}

func (q *metadataQuerier) LabelNamesHandler(ctx context.Context, request *mcp.CallToolRequest, input *LabelNamesArguments) (*mcp.CallToolResult, *LabelNamesResult, error) {
	var (
		start, end time.Time
		err        error
	)
	if start, err = utils.ParseTime(input.Start); err != nil {
		slog.Warn("parse start time", "err", err)
	}
	if end, err = utils.ParseTime(input.End); err != nil {
		slog.Warn("parse end time", "err", err)
	}

	opts := make([]v1.Option, 0)
	if input.Limit != 0 {
		opts = append(opts, v1.WithLimit(input.Limit))
	}

	result := &LabelNamesResult{}
	if result.LabelNames, result.Warnings, err = q.API.LabelNames(
		ctx,
		input.Match,
		start,
		end,
		opts...,
	); err != nil {
		return nil, nil, err
	}
	return nil, result, nil
}
