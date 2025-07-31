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

- `--prom-addr`: Prometheus server address (default: http://localhost:9090/)
- `--mcp-addr`: MCP server listening address (default: localhost:8080)
- `--transport`: Transport mechanism (stdio, http, sse) (default: stdio)

## Available Tools

### Expression query

#### 1. Prometheus Instant Query

- Execute queries at specific timestamps
- Get current metric values
- Perfect for real-time monitoring

#### 2. Prometheus Range Query

- Query metrics over time ranges
- Analyze historical data
- Ideal for trend analysis

### Metadata query

#### 1. Find Series by Labels

- List all time series that match specific label filters
- Discover which series exist for given label criteria

#### 2. List Label Names

- Get all label names used in the Prometheus database
- Explore available labels for filtering or grouping

#### 3. List Label Values

- Get all possible values for a specific label name
- See which values a label can take for filtering or selection

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
