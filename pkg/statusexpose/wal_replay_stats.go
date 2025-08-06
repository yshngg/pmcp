package statusexpose

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type WALReplayStatsExposeParams struct{}

type WALReplayStatsExposeResult = v1.WalReplayStatus

func (e *statusExposer) WALReplayStatsExposeHandler(ctx context.Context, _ *mcp.ServerSession, _ *mcp.CallToolParamsFor[WALReplayStatsExposeParams]) (*mcp.CallToolResultFor[WALReplayStatsExposeResult], error) {
	result, err := e.API.WalReplay(ctx)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[WALReplayStatsExposeResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
