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
	"github.com/yshngg/pmcp/pkg/bind"
	"github.com/yshngg/pmcp/pkg/prometheus/api"
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

// main starts the MCP server with Prometheus query capabilities,
// selecting the transport mechanism (stdio, HTTP, or SSE) based on command-line flags.
// It initializes the Prometheus client, binds query handlers,
// and serves requests until termination.
// The function exits the program on critical errors or when printing version information.
func main() {
	if *printVersion {
		fmt.Println(version.Info)
		os.Exit(0)
	}

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "pmcp",
		Version: version.Info.Number,
	}, nil)

	promCli, err := api.New(*promAddr, http.DefaultClient, nil)
	if err != nil {
		slog.Error("new prometheus client", "err", err)
		os.Exit(1)
	}

	binder := bind.NewBinder(server, promCli)
	binder.Bind()

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
