package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	expressionquery "github.com/yshngg/pmcp/pkg/expression_query"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
)

const Schema = "prom"

var (
	promAddr      = flag.String("prom-addr", "http://localhost:9090/", "The address of the Prometheus to connect to.")
	mcpAddr       = flag.String("mcp-addr", "localhost:8080", "The address of the MCP server to listen on.")
	transportType = flag.String("transport", "stdio", "Transport type (stdio or http).\nThe mechanisms that handle the underlying communication between clients and servers.")
)

func init() {
	flag.Parse()
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "pmcp",
		Version: "0.1.0-alpha",
	}, nil)

	if err := AddTools(server); err != nil {
		log.Fatal(err)
	}

	if *transportType == "http" {
		// http.Handle("/ping", )
		http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		// Run the server over Streamable HTTP
		streamableHTTPHandler := mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server {
			return server
		}, nil)
		http.Handle("/mcp", streamableHTTPHandler)
		if err := http.ListenAndServe(*mcpAddr, nil); err != nil {
			log.Fatal(err)
		}
	}

	// Run the server over stdin/stdout, until the client disconnects
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}

func AddTools(server *mcp.Server) error {
	promCli, err := client.New(*promAddr, http.DefaultClient, nil)
	if err != nil {
		return fmt.Errorf("new prometheus client, err: %w", err)
	}
	expressionQuerier := expressionquery.NewExpressionQuerier(promCli)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "Prometheus Expression Query - Instant",
		Description: "Execute a Prometheus query expression at a specific point in time to get the current value of a metric or calculation. Useful for checking the current state of systems and applications.",
	}, expressionQuerier.InstantQueryHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "Prometheus Expression Query - Range",
		Description: "Execute a Prometheus query expression over a time range to get historical values of a metric or calculation. Useful for analyzing trends and patterns over time.",
	}, expressionQuerier.RangeQueryHandler)

	return nil
}
