// Package main provides a server that exposes Prometheus query capabilities via the Model Context Protocol (MCP).
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	expressionquery "github.com/yshngg/pmcp/pkg/expression_query"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
)

// Schema is the identifier for the Prometheus schema.
const Schema = "prom"

var (
	// promAddr is the address of the Prometheus server to connect to.
	promAddr = flag.String("prom-addr", "http://localhost:9090/", "The address of the Prometheus to connect to.")
	// mcpAddr is the address for the MCP server to listen on.
	mcpAddr = flag.String("mcp-addr", "localhost:8080", "The address of the MCP server to listen on.")
	// transportType specifies the transport mechanism (stdio, sse, or http).
	transportType = flag.String("transport", "stdio", "Transport type (stdio, sse or http).\nThe mechanisms that handle the underlying communication between clients and servers.")
)

func init() {
	// Parse command-line flags.
	flag.Parse()
}

// main is the entry point for the pmcp server.
// It sets up the MCP server, registers Prometheus query tools, and starts the server
// using the specified transport (stdio, http, or sse).
func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "pmcp",
		Version: "0.1.0-alpha",
	}, nil)

	if err := AddTools(server); err != nil {
		slog.Error("add tools", "err", err)
		os.Exit(1)
	}

	if *transportType == "http" {
		http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		// Run the server over Streamable HTTP
		streamableHTTPHandler := mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server {
			return server
		}, nil)
		http.Handle("/mcp", streamableHTTPHandler)

		slog.Info("Listening on http://" + *mcpAddr)
		if err := http.ListenAndServe(*mcpAddr, nil); err != nil {
			slog.Error("listen and serve with Streamable HTTP transport", "err", err)
			os.Exit(1)
		}
	}

	// Backwards Compatibility
	if *transportType == "sse" {
		slog.Warn("HTTP+SSE transport is deprecated. Please use Streamable HTTP instead.")
		sseHandler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server { return server })
		http.Handle("/mcp", sseHandler)
		slog.Info("Listening on http://" + *mcpAddr)
		if err := http.ListenAndServe(*mcpAddr, nil); err != nil {
			slog.Error("listen and serve with HTTP+SSE transport", "err", err)
			os.Exit(1)
		}
	}

	// Run the server over stdin/stdout, until the client disconnects
	slog.Info("Listening on stdio")
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		slog.Error("run server with stdio transport", "err", err)
		os.Exit(1)
	}
}

// AddTools registers Prometheus query tools with the given MCP server.
// It creates a Prometheus client and adds instant and range query handlers.
//
// Returns an error if the Prometheus client cannot be created.
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
