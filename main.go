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
	metadataquery "github.com/yshngg/pmcp/pkg/metadata_query"
	"github.com/yshngg/pmcp/pkg/prometheus/client"
	"github.com/yshngg/pmcp/pkg/version"
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
	// printVersion prints the version and exit.
	printVersion = flag.Bool("version", false, "Print the version and exit.")
)

func init() {
	// Parse command-line flags.
	flag.Parse()
}

// main is the entry point for the pmcp server.
// It sets up the MCP server, registers Prometheus query tools, and starts the server
// using the specified transport (stdio, http, or sse).
func main() {
	if *printVersion {
		fmt.Println(version.Info)
		os.Exit(0)
	}

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "pmcp",
		Version: version.Info.Number,
	}, nil)

	if err := AddTools(server); err != nil {
		slog.Error("add tools", "err", err)
		os.Exit(1)
	}

	if *transportType == "http" {
		http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte("pong")); err != nil {
				slog.Error("write pong", "err", err)
				os.Exit(1)
			}
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

		http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte("pong")); err != nil {
				slog.Error("write pong", "err", err)
				os.Exit(1)
			}
		})

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

	// Expression queries
	// Query language expressions may be evaluated at a single instant or over a range of time.
	{
		expressionQuerier := expressionquery.NewExpressionQuerier(promCli)
		mcp.AddTool(server, &mcp.Tool{
			Name:        "Prometheus Instant Query",
			Description: "Run a Prometheus expression and get the current value for a metric or calculation at a specific time. Use this to check the latest status or value of any metric.",
		}, expressionQuerier.InstantQueryHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "Prometheus Range Query",
			Description: "Run a Prometheus expression over a time range to get historical values for a metric or calculation. Use this to analyze trends or patterns over time.",
		}, expressionQuerier.RangeQueryHandler)
	}

	// Querying metadata
	// Prometheus offers a set of API endpoints to query metadata about series and their labels.
	{
		metadataQuerier := metadataquery.NewMetadataQuerier(promCli)
		mcp.AddTool(server, &mcp.Tool{
			Name:        "Find Series by Labels",
			Description: "List all time series that match specific label filters. Use this to discover which series exist for given label criteria.",
		}, metadataQuerier.SeriesHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "List Label Names",
			Description: "Get all label names used in the Prometheus database. Use this to explore available labels for filtering or grouping.",
		}, metadataQuerier.LabelNamesHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "List Label Values",
			Description: "Get all possible values for a specific label name. Use this to see which values a label can take for filtering or selection.",
		}, metadataQuerier.LabelValuesHandler)
	}

	return nil
}
