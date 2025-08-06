package statusexpose

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type FlagsExposeParams struct{}

type FlagsExposeResult = v1.FlagsResult

func (e *statusExposer) FlagsExposeHandler(ctx context.Context, _ *mcp.ServerSession, _ *mcp.CallToolParamsFor[FlagsExposeParams]) (*mcp.CallToolResultFor[FlagsExposeResult], error) {
	result, err := e.API.Flags(ctx)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[FlagsExposeResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
