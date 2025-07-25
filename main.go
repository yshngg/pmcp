package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	expressionquery "github.com/yshngg/pmcp/pkg/expression_query"
)

const (
	Schema     = "prometheus"
	APIVersion = "/api/v1"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "pmcp",
		Version: "0.1.0-alpha",
	}, nil)

	AddResource(server)

	// Run the server over stdin/stdout, until the client disconnects
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}

func AddResource(server *mcp.Server) {

	server.AddTool(&mcp.Tool{
		Name:        "Expression query/Instant query",
		Description: "Query language expressions may be evaluated at a single instant.",
		InputSchema: &jsonschema.Schema{
			Defs: map[string]*jsonschema.Schema{
				"query": {
					Type:        "string",
					Description: "Prometheus expression query string.",
				},
				"time": {
					Type:        "",
					Description: "Evaluation timestamp. Optional.",
				},
				"timeout": {Description: "Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag."},
				"limit":   {Description: "Maximum number of returned series. Doesn't affect scalars or strings but truncates the number of series for matrices and vectors. Optional. 0 means disabled."},
			},
		},
	}, nil)

	server.AddResourceTemplate(&mcp.ResourceTemplate{
		Name: "Expression queries/Instant queries",
		Description: `

URL query parameters:
  - query=<string>: Prometheus expression query string.
  - time=<rfc3339 | unix_timestamp>: Evaluation timestamp. Optional.
  - timeout=<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag.
  - limit=<number>: Maximum number of returned series. Doesn't affect scalars or strings but truncates the number of series for matrices and vectors. Optional. 0 means disabled.`,
		URITemplate: "prom://api/v1/query?query={query}&time={time}&timeout={timeout}&limit={limit}",
		MIMEType:    "application/json",
	}, expressionquery.InstantQueryHandler)

	server.AddResourceTemplate(&mcp.ResourceTemplate{
		Name: "Expression queries/Range queries",
		Description: `
Query language expressions may be evaluated at a single instant or over a range of time. The sections below describe the API endpoints for each type of expression query.
URL query parameters:
  - query=<string>: Prometheus expression query string.
  - start=<rfc3339 | unix_timestamp>: Start timestamp, inclusive.
  - end=<rfc3339 | unix_timestamp>: End timestamp, inclusive.
  - step=<duration | float>: Query resolution step width in duration format or float number of seconds.
  - timeout=<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag.
  - limit=<number>: Maximum number of returned series. Optional. 0 means disabled.`,
		URITemplate: "prom://api/v1/query_range?query={query}&start={start}&end={end}&step={step}&timeout={timeout}&limit={limit}",
		MIMEType:    "application/json",
	}, expressionquery.RangeQueryHandler)

}
