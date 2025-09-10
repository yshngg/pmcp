package statusexpose

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type WALReplayStatsExposeParams struct{}

type WALReplayStatsExposeResult = v1.WalReplayStatus

func (e *statusExposer) WALReplayStatsExposeHandler(ctx context.Context, request *mcp.CallToolRequest, input *WALReplayStatsExposeParams) (*mcp.CallToolResult, *WALReplayStatsExposeResult, error) {
	result, err := e.API.WalReplay(ctx)
	if err != nil {
		return nil, nil, err
	}
	return nil, &result, nil
}
