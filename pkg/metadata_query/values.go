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

const LabelValuesEndpoint = "/label/<label_name>/values"

// URL query parameters:
// start=<rfc3339 | unix_timestamp>: Start timestamp. Optional.
// end=<rfc3339 | unix_timestamp>: End timestamp. Optional.
// match[]=<series_selector>: Repeated series selector argument that selects the series from which to read the label values. Optional.
// limit=<number>: Maximum number of returned series. Optional. 0 means disabled.
// This example queries for all label values for the http_status_code label:
// curl http://localhost:9090/api/v1/label/http_status_code/values
// This example queries for all label values for the http.status_code label:
// curl http://localhost:9090/api/v1/label/U__http_2e_status_code/values
type LabelValuesArguments struct {
	Label string   `json:"label,omitzero" jsonschema:"<string>: Label names can optionally be encoded using the Values Escaping method, and is necessary if a name includes the / character. To encode a name in this way: 1. Prepend the label with U__. 2. Letters, numbers, and colons appear as-is. 3. Convert single underscores to double underscores. 4. For all other characters, use the UTF-8 codepoint as a hex integer, surrounded by underscores. So becomes _20_ and a . becomes _2e_."`
	Start string   `json:"start,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: Start timestamp. Optional."`
	End   string   `json:"end,omitzero" jsonschema:"<rfc3339 | unix_timestamp>: End timestamp. Optional."`
	Match []string `json:"match[],omitzero" jsonschema:"<series_selector>: Repeated series selector argument that selects the series from which to read the label values. Optional."`
	Limit uint64   `json:"limit,omitzero" jsonschema:"<number>: Maximum number of returned series. Optional. 0 means disabled."`
}

type LabelValuesResult struct {
	LabelValues model.LabelValues `json:"labelvalues" jsonschema:"LabelSets consists of a list of objects that contain the label name/value pairs which identify each series."`
	Warnings    v1.Warnings       `json:"warnings,omitempty"`
}

func (q *metadataQuerier) LabelValuesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[LabelValuesArguments]) (*mcp.CallToolResultFor[LabelValuesResult], error) {
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

	result := LabelValuesResult{}
	result.LabelValues, result.Warnings, err = q.Client.LabelValues(
		ctx,
		params.Arguments.Label,
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
	return &mcp.CallToolResultFor[LabelValuesResult]{
		Content: []mcp.Content{&mcp.TextContent{
			Text: string(content),
		}},
		StructuredContent: result,
	}, nil
}
