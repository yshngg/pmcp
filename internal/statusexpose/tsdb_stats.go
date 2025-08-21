package statusexpose

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type TSDBStatsExposeParams struct{}

type TSDBStatsExposeResult = v1.TSDBResult

func (e *statusExposer) TSDBStatsExposeHandler(ctx context.Context, _ *mcp.ServerSession, _ *mcp.CallToolParamsFor[TSDBStatsExposeParams]) (*mcp.CallToolResultFor[TSDBStatsExposeResult], error) {
	result, err := e.API.TSDB(ctx)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[TSDBStatsExposeResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
