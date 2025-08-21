package manage

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (m *manager) ReloadHandler(ctx context.Context, _ *mcp.ServerSession, _ *mcp.CallToolParamsFor[struct{}]) (*mcp.CallToolResultFor[ManagementResult], error) {
	result := ManagementResult{
		Success: true,
	}

	err := m.api.Reload(ctx)
	if err != nil {
		result.Success = false
		result.Message = err.Error()
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return &mcp.CallToolResultFor[ManagementResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
