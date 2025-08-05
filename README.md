# PMCP - Prometheus Model Context Protocol Server

🚀 A Model Context Protocol (MCP) server implementation for Prometheus that enables natural language interactions with Prometheus metrics and queries.

## Features

- Instant Query: Execute Prometheus queries at specific points in time
- Range Query: Execute Prometheus queries over time ranges
- Metadata Query: Discover series, label names, and label values
- MCP Integration: Seamless integration with MCP-compatible clients
- Multiple Transport Options: Support for HTTP, SSE, and stdio communication

## Installation

```bash
go install github.com/yshngg/pmcp@latest
```

## Usage

Start the PMCP server by providing your Prometheus server address:

```bash
# Basic usage with default settings
pmcp --prom-addr="http://localhost:9090"

# Using HTTP transport
pmcp --prom-addr="http://localhost:9090" --transport=http --mcp-addr="localhost:8080"

# Using stdio transport (default)
pmcp --prom-addr="http://localhost:9090" --transport=stdio
```

### Command Line Flags

- `-help`: Display help message.
- `-mcp-addr`: The address of the MCP server to listen on. (default: localhost:8080)
- `-prom-addr`: The address of the Prometheus to connect to. (default: <http://localhost:9090/>)
- `-transport`: Transport type (stdio, sse or http). (default: stdio)
- `-version`: Display the version and exit.

## Available Tools

### Expression Query

#### 1. Prometheus Instant Query

Run a Prometheus expression and get the current value for a metric or calculation at a specific time. Use this to check the latest status or value of any metric.

#### 2. Prometheus Range Query

Run a Prometheus expression over a time range to get historical values for a metric or calculation. Use this to analyze trends or patterns over time.

### Metadata Query

#### 1. Find Series by Labels

List all time series that match specific label filters. Use this to discover which series exist for given label criteria.

#### 2. List Label Names

Get all label names used in the Prometheus database. Use this to explore available labels for filtering or grouping.

#### 3. List Label Values

Get all possible values for a specific label name. Use this to see which values a label can take for filtering or selection.

### Target Descover

#### 1. Target Discovery

Get an overview of the current state of the Prometheus target discovery.

### Management API

#### 1. Health Check

Check whether Prometheus is healthy.

#### 2. Readiness Check

Check whether Prometheus is ready to serve traffic.

#### 3. Reload

Trigger a reload of the Prometheus configuration and rule files.

#### 4. Quit

Trigger a graceful shutdown of Prometheus.

## Available Prompts

#### All Available Metrics

List all available metrics in the Prometheus instance.

## Requirements

- Go 1.23.5 or higher
- Access to a running Prometheus server

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Model Context Protocol](https://github.com/modelcontextprotocol/go-sdk)
- Powered by [Prometheus](https://prometheus.io/)
