package statusexpose

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type BuildInformationExposeParams struct{}

type BuildInformationExposeResult = v1.BuildinfoResult

func (e *statusExposer) BuildInformationExposeHandler(ctx context.Context, _ *mcp.ServerSession, _ *mcp.CallToolParamsFor[BuildInformationExposeParams]) (*mcp.CallToolResultFor[BuildInformationExposeResult], error) {
	result, err := e.API.Buildinfo(ctx)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[BuildInformationExposeResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
