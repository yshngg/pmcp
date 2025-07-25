package expressionquery

import (
	"context"
	"fmt"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const InstantQueryPath = "/query"

func InstantQueryHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.ReadResourceParams) (*mcp.ReadResourceResult, error) {
	values, err := url.ParseQuery(params.URI)
	if err != nil {
		return nil, err
	}
	fmt.Printf("URI: %#v", values)
	return nil, nil
}

var _ mcp.ResourceHandler = InstantQueryHandler
