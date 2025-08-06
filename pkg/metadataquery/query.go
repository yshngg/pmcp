package metadataquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/api"
)

type MetadataQuerier interface {
	SeriesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[SeriesArguments]) (*mcp.CallToolResultFor[SeriesResult], error)
	LabelNamesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[LabelNamesArguments]) (*mcp.CallToolResultFor[LabelNamesResult], error)
	LabelValuesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[LabelValuesArguments]) (*mcp.CallToolResultFor[LabelValuesResult], error)

	TargetMetadataQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[TargetMetadataQueryParams]) (*mcp.CallToolResultFor[TargetMetadataQueryResult], error)
	MetricsMetadataQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[MetricsMetadataQueryParams]) (*mcp.CallToolResultFor[MetricsMetadataQueryResult], error)
}

func NewMetadataQuerier(api api.PrometheusAPI) MetadataQuerier {
	return &metadataQuerier{API: api}
}

type metadataQuerier struct {
	API api.PrometheusAPI
}

var _ MetadataQuerier = &metadataQuerier{}
