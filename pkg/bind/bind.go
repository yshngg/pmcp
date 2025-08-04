package bind

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
)

type Binder interface {
	Bind()
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

func (b *binder) Bind() {
	// add tools
	b.addTools()
	// add resources
	b.addResources()
	// add prompts
	b.addPrompts()
}
