package bindingblocks

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/pkg/prometheus/api"
)

type Binder interface {
	Bind()
}

// NewBinder returns a Binder that binds components to the given MCP server using the provided Prometheus client.
func NewBinder(server *mcp.Server, api api.PrometheusAPI) Binder {
	return &binder{
		server: server,
		api:    api,
	}
}

type binder struct {
	server *mcp.Server
	api    api.PrometheusAPI
}

func (b *binder) Bind() {
	// add tools
	b.addTools()
	// add resources
	b.addResources()
	// add prompts
	b.addPrompts()
}
