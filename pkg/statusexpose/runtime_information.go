package statusexpose

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type RuntimeInformationExposeParams struct{}

type RuntimeInformationExposeResult = v1.RuntimeinfoResult

func (e *statusExposer) RuntimeInformationExposeHandler(ctx context.Context, _ *mcp.ServerSession, _ *mcp.CallToolParamsFor[RuntimeInformationExposeParams]) (*mcp.CallToolResultFor[RuntimeInformationExposeResult], error) {
	result, err := e.API.Runtimeinfo(ctx)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[RuntimeInformationExposeResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
