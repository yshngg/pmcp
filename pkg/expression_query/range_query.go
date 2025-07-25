package expressionquery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RangeQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.ReadResourceParams) (*mcp.ReadResourceResult, error) {

	return nil, nil
}

var _ mcp.ResourceHandler = RangeQueryHandler
