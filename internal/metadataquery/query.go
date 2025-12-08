package metadataquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/prometheus-mcp-server/internal/prometheus/api"
)

type MetadataQuerier interface {
	SeriesHandler(ctx context.Context, request *mcp.CallToolRequest, input *SeriesArguments) (*mcp.CallToolResult, *SeriesResult, error)
	LabelNamesHandler(ctx context.Context, request *mcp.CallToolRequest, input *LabelNamesArguments) (*mcp.CallToolResult, *LabelNamesResult, error)
	LabelValuesHandler(ctx context.Context, request *mcp.CallToolRequest, input *LabelValuesArguments) (*mcp.CallToolResult, *LabelValuesResult, error)

	TargetMetadataQueryHandler(ctx context.Context, request *mcp.CallToolRequest, input *TargetMetadataQueryParams) (*mcp.CallToolResult, *TargetMetadataQueryResult, error)
	MetricsMetadataQueryHandler(ctx context.Context, request *mcp.CallToolRequest, input *MetricsMetadataQueryParams) (*mcp.CallToolResult, *MetricsMetadataQueryResult, error)
}

// NewMetadataQuerier returns a MetadataQuerier implementation that uses the provided PrometheusAPI.
// The returned implementation is a concrete metadataQuerier with its API field set to the given api.
func NewMetadataQuerier(api api.PrometheusAPI) MetadataQuerier {
	return &metadataQuerier{API: api}
}

type metadataQuerier struct {
	API api.PrometheusAPI
}

var _ MetadataQuerier = &metadataQuerier{}
