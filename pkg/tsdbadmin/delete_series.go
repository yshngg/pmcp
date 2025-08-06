package tsdbadmin

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/utils"
)

// URL query parameters:
// match[]=<series_selector>: Repeated label matcher argument that selects the series to delete. At least one match[] argument must be provided.
// start=<rfc3339 | unix_timestamp>: Start timestamp. Optional and defaults to minimum possible time.
// end=<rfc3339 | unix_timestamp>: End timestamp. Optional and defaults to maximum possible time.
// Example:
//
//	curl -X POST \
//		 -g 'http://localhost:9090/api/v1/admin/tsdb/delete_series?match[]=up&match[]=process_start_time_seconds{job="prometheus"}'
type DeleteSeriesParams struct {
	Match []string `json:"match[]" jsonschema:"<series_selector>: Repeated label matcher argument that selects the series to delete. At least one match[] argument must be provided."`
	Start string   `json:"start,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: Start timestamp. Optional and defaults to minimum possible time."`
	End   string   `json:"end,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: End timestamp. Optional and defaults to maximum possible time."`
}

type DeleteSeriesResult struct {
	Success bool   `json:"success" jsonschema:"Indicate the result of the management operation, true means success, false means failure"`
	Message string `json:"message,omitempty" jsonschema:"Explanation message when the operation fails."`
}

func (a *tsdbAdmin) DeleteSeriesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteSeriesParams]) (*mcp.CallToolResultFor[DeleteSeriesResult], error) {
	var (
		start, end time.Time
		err        error
	)
	if start, err = utils.ParseTime(params.Arguments.Start); err != nil {
		slog.Warn("parse start time", "err", err)
	}
	if end, err = utils.ParseTime(params.Arguments.End); err != nil {
		slog.Warn("parse end time", "err", err)
	}

	result := DeleteSeriesResult{Success: true}
	if err = a.API.DeleteSeries(ctx, params.Arguments.Match, start, end); err != nil {
		result.Success = false
		result.Message = err.Error()
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[DeleteSeriesResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
