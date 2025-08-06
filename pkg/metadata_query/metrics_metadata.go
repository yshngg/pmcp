package metadataquery

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// URL query parameters:
// limit=<number>: Maximum number of metrics to return.
// limit_per_metric=<number>: Maximum number of metadata to return per metric.
// metric=<string>: A metric name to filter metadata for. All metric metadata is retrieved if left empty.
// The following example returns two metrics. Note that the metric http_requests_total has more than one object in the list. At least one target has a value for HELP that do not match with the rest.
// curl -G http://localhost:9090/api/v1/metadata?limit=2
type MetricsMetadataQueryParams struct {
	Limit  string `json:"limit,omitzero" jsonschema:"<number>: Maximum number of metrics to return."`
	Metric string `json:"metric,omitzero" jsonschema:"<string>: A metric name to filter metadata for. All metric metadata is retrieved if left empty."`
}

type MetricsMetadataQueryResult struct {
	Data map[string][]v1.Metadata `json:"data" jsonschema:"Data consists of an object where each key is a metric name and each value is a list of unique metadata objects, as exposed for that metric name across all targets."`
}

func (q *metadataQuerier) MetricsMetadataQueryHandler(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[MetricsMetadataQueryParams]) (*mcp.CallToolResultFor[MetricsMetadataQueryResult], error) {
	var (
		result MetricsMetadataQueryResult
		err    error
	)
	if result.Data, err = q.API.Metadata(
		ctx,
		params.Arguments.Metric,
		params.Arguments.Limit,
	); err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[MetricsMetadataQueryResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
