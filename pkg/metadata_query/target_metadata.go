package metadataquery

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// URL query parameters:
// match_target=<label_selectors>: Label selectors that match targets by their label sets. All targets are selected if left empty.
// metric=<string>: A metric name to retrieve metadata for. All metric metadata is retrieved if left empty.
// limit=<number>: Maximum number of targets to match.
// The following example returns all metadata entries for the go_goroutines metric from the first two targets with label job="prometheus".
//
//	curl -G http://localhost:9091/api/v1/targets/metadata \
//	    --data-urlencode 'metric=go_goroutines' \
//	    --data-urlencode 'match_target={job="prometheus"}' \
//	    --data-urlencode 'limit=2'
type TargetMetadataQueryParams struct {
	MatchTarget string `json:"matchTarget,omitzero" jsonschema:"<label_selectors>: Label selectors that match targets by their label sets. All targets are selected if left empty."`
	Metric      string `json:"metrics,omitzero" jsonschema:"<string>: A metric name to retrieve metadata for. All metric metadata is retrieved if left empty."`
	Limit       string `json:"limit,omitzero" jsonschema:"<number>: Maximum number of targets to match."`
}

type TargetMetadataQueryResult struct {
	Data []v1.MetricMetadata `json:"data"`
}

func (d *metadataQuerier) TargetMetadataQueryHandler(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[TargetMetadataQueryParams]) (*mcp.CallToolResultFor[TargetMetadataQueryResult], error) {
	var (
		result TargetMetadataQueryResult
		err    error
	)
	if result.Data, err = d.API.TargetsMetadata(
		ctx,
		params.Arguments.MatchTarget,
		params.Arguments.Metric,
		params.Arguments.Limit,
	); err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[TargetMetadataQueryResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
