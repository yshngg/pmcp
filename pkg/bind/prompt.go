package bind

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (b *binder) addPrompts(server *mcp.Server) error {
	server.AddPrompt(&mcp.Prompt{
		Name: "All Available Metrics",
	}, func(ctx context.Context, ss *mcp.ServerSession, gpp *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
		return &mcp.GetPromptResult{
			Description: "What is the current value of a metric or calculation?",
			Messages: []*mcp.PromptMessage{
				{
					Content: &mcp.TextContent{
						Text: "List All Available Metrics is Equivalent to List Values of __name__ Label",
					},
				},
			},
		}, nil
	})
	return nil
}
