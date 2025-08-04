package bind

import (
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
)

type Binder interface {
	Bind() error
}

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
	if err := b.addTools(b.server); err != nil {
		return fmt.Errorf("add tools, err: %w", err)
	}
	if err := b.addResources(b.server); err != nil {
		return fmt.Errorf("add resources, err: %w", err)
	}
	if err := b.addPrompts(b.server); err != nil {
		return fmt.Errorf("add prompts, err: %w", err)
	}
	return nil
}
