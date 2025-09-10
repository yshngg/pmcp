package metadataquery

import (
	"context"
	"log/slog"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/yshngg/pmcp/internal/utils"
)

const SeriesEndpoint = "/series"

// URL query parameters:
// match[]=<series_selector>: Repeated series selector argument that selects the series to return. At least one match[] argument must be provided.
// start=<rfc3339 | unix_timestamp>: Start timestamp.
// end=<rfc3339 | unix_timestamp>: End timestamp.
// limit=<number>: Maximum number of returned series. Optional. 0 means disabled.
// The following example returns all series that match either of the selectors up or process_start_time_seconds{job="prometheus"}:
// curl -g 'http://localhost:9090/api/v1/series?' --data-urlencode 'match[]=up' --data-urlencode 'match[]=process_start_time_seconds{job="prometheus"}'
type SeriesArguments struct {
	Match []string `json:"match[]" jsonschema:"<series_selector>: Repeated series selector argument that selects the series to return. At least one match[] argument must be provided."`
	Start string   `json:"start,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: Start timestamp."`
	End   string   `json:"end,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: End timestamp."`
	Limit uint64   `json:"limit,omitzero" jsonschema:"<number>: Maximum number of returned series. Optional. 0 means disabled."`
}

type SeriesResult struct {
	LabelSets []model.LabelSet `json:"labelsets" jsonschema:"LabelSets consists of a list of objects that contain the label name/value pairs which identify each series."`
	Warnings  v1.Warnings      `json:"warnings,omitempty"`
}

func (q *metadataQuerier) SeriesHandler(ctx context.Context, request *mcp.CallToolRequest, input *SeriesArguments) (*mcp.CallToolResult, *SeriesResult, error) {
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

	result := &SeriesResult{}
	result.LabelSets, result.Warnings, err = q.API.Series(
		ctx,
		input.Match,
		start,
		end,
		opts...,
	)
	if err != nil {
		return nil, nil, err
	}
	return nil, result, nil
}
