package tsdbadmin

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type CleanTombstonesParams struct{}

type CleanTombstonesResult struct {
	Success bool   `json:"success" jsonschema:"Indicate the result of the management operation, true means success, false means failure"`
	Message string `json:"message,omitempty" jsonschema:"Explanation message when the operation fails."`
}

func (a *tsdbAdmin) CleanTombstonesHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[CleanTombstonesParams]) (*mcp.CallToolResultFor[CleanTombstonesResult], error) {
	result := CleanTombstonesResult{Success: true}
	if err := a.API.CleanTombstones(ctx); err != nil {
		result.Success = false
		result.Message = err.Error()
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[CleanTombstonesResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
