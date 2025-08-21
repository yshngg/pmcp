package statusexpose

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type ConfigExposeParams struct{}

type ConfigExposeResult = v1.ConfigResult

func (e *statusExposer) ConfigExposeHandler(ctx context.Context, _ *mcp.ServerSession, _ *mcp.CallToolParamsFor[ConfigExposeParams]) (*mcp.CallToolResultFor[ConfigExposeResult], error) {
	result, err := e.API.Config(ctx)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[ConfigExposeResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
