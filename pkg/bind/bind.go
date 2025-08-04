package bind

import (
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
)

type Binder interface {
	Bind() error
}

// NewBinder returns a Binder that binds components to the given MCP server using the provided Prometheus client.
func NewBinder(server *mcp.Server, promCli client.PrometheusClient) Binder {
	return &binder{
		server:  server,
		promCli: promCli,
	}
}

type binder struct {
	server  *mcp.Server
	promCli client.PrometheusClient
}

func (b *binder) Bind() error {
	if err := b.addTools(); err != nil {
		return fmt.Errorf("add tools, err: %w", err)
	}
	if err := b.addResources(); err != nil {
		return fmt.Errorf("add resources, err: %w", err)
	}
	if err := b.addPrompts(); err != nil {
		return fmt.Errorf("add prompts, err: %w", err)
	}
	return nil
}
