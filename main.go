// Package main provides a server that exposes Prometheus query capabilities via the Model Context Protocol (MCP).
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/bindingblocks"
	"github.com/yshngg/pmcp/internal/prometheus/api"
	"github.com/yshngg/pmcp/internal/version"
)

// Schema is the identifier for the Prometheus schema.
const Schema = "prom"

// main starts the MCP server with Prometheus query capabilities,
// selecting the transport mechanism (stdio, HTTP, or SSE) based on command-line flags.
// It initializes the Prometheus client, binds query handlers,
// and serves requests until termination.
// The function exits the program on critical errors or when printing version information.
func main() {
	fs := flag.NewFlagSet("pmcp", flag.ExitOnError)
	var (
		// promAddr is the address of the Prometheus server to connect to.
		promAddr = fs.String("prom-addr", "http://localhost:9090/", "The address of the Prometheus to connect to.")
		// mcpAddr is the address for the MCP server to listen on.
		mcpAddr = fs.String("mcp-addr", "localhost:8080", "The address of the MCP server to listen on.")
		// transportType specifies the transport mechanism (stdio, sse, or http).
		transportType = fs.String("transport", "stdio", "Transport type (stdio, sse or http).\nThe mechanisms that handle the underlying communication between clients and servers.")
		// printVersion prints the version and exit.
		printVersion = fs.Bool("version", false, "Print the version and exit.")
	)
	fs.Usage = usageFor(fs, "pmcp [flags]")
	// Parse command-line flags.
	if err := fs.Parse(os.Args[1:]); err != nil {
		slog.Error("parse args", "err", err)
		os.Exit(1)
	}

	if *printVersion {
		fmt.Println(version.Info)
		os.Exit(0)
	}

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "pmcp",
		Version: string(version.Info.Number),
	}, nil)

	promCli, err := api.New(*promAddr, http.DefaultClient, nil)
	if err != nil {
		slog.Error("new prometheus client", "err", err)
		os.Exit(1)
	}

	binder := bindingblocks.NewBinder(server, promCli)
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

		sseHandler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server { return server }, nil)
		http.Handle("/mcp", sseHandler)

		slog.Info("Listening on http://" + *mcpAddr)
		if err := http.ListenAndServe(*mcpAddr, nil); err != nil {
			slog.Error("listen and serve with HTTP+SSE transport", "err", err)
			os.Exit(1)
		}
	}

	// Run the server over stdin/stdout, until the client disconnects
	slog.Info("Listening on stdio")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		slog.Error("run server with stdio transport", "err", err)
		os.Exit(1)
	}
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "Prometheus Model Context Protocol Server\n\n")
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			def := f.DefValue
			if def == "" {
				def = "..."
			}
			_, err := fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, def, f.Usage)
			if err != nil {
				panic(err)
			}
		})
		if err := w.Flush(); err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "VERSION\n")
		fmt.Fprintf(os.Stderr, "  %s\n", version.Info.Number)
		fmt.Fprintf(os.Stderr, "\n")
	}
}
