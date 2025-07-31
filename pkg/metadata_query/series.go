package metadataquery

import (
	"context"
	"encoding/json"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/yshngg/pmcp/pkg/utils"
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

func (q *metadataQuerier) SeriesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[SeriesArguments]) (*mcp.CallToolResultFor[SeriesResult], error) {
	var (
		start, end time.Time
		err        error
	)
	if start, err = utils.ParseTime(params.Arguments.Start); err != nil {
		return nil, err
	}
	if end, err = utils.ParseTime(params.Arguments.End); err != nil {
		return nil, err
	}

	opts := make([]v1.Option, 0)
	if params.Arguments.Limit != 0 {
		opts = append(opts, v1.WithLimit(params.Arguments.Limit))
	}

	result := SeriesResult{}
	result.LabelSets, result.Warnings, err = q.Client.Series(
		ctx,
		params.Arguments.Match,
		start,
		end,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return &mcp.CallToolResultFor[SeriesResult]{
		Content: []mcp.Content{&mcp.TextContent{
			Text: string(content),
		}},
		StructuredContent: result,
	}, nil
}
