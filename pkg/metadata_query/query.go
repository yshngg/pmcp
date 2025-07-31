package metadataquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
)

type MetadataQuerier interface {
	SeriesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[SeriesArguments]) (*mcp.CallToolResultFor[SeriesResult], error)
	LabelNamesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[LabelNamesArguments]) (*mcp.CallToolResultFor[LabelNamesResult], error)
	LabelValuesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[LabelValuesArguments]) (*mcp.CallToolResultFor[LabelValuesResult], error)
}

func NewMetadataQuerier(cli client.PrometheusClient) MetadataQuerier {
	return &metadataQuerier{Client: cli}
}

type metadataQuerier struct {
	Client client.PrometheusClient
}

var _ MetadataQuerier = &metadataQuerier{}
