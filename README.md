# PMCP - Prometheus Model Context Protocol Server

**🚀 A Model Context Protocol (MCP) server implementation for Prometheus that enables natural language interactions with Prometheus metrics and queries.**

---

## Table of Contents

1. [Features](#features)
2. [Requirements](#requirements)
3. [Installation](#installation)
4. [Usage](#usage)
   * [Command Line Flags](#command-line-flags)
5. [Binding Blocks](#binding-blocks)
   * [Tools](#tools)
   * [Prompts](#prompts)
6. [Contributing](#contributing)
7. [License](#license)
8. [Acknowledgments](#acknowledgments)

---

## Features

* **Instant Query**: Execute Prometheus queries at a specific point in time.
* **Range Query**: Retrieve historical metric data over defined time ranges.
* **Metadata Query**: Discover time series, label names, and label values.
* **Transport Options**: Support for HTTP, Server-Sent Events (SSE), and stdio.
* **MCP Integration**: Seamless communication with MCP-compatible clients.

---

## Requirements

* Go **1.23.5** or higher
* A running Prometheus server (v2.x)

---

## Installation

Install the `pmcp` binary via `go install`:

```bash
go install github.com/yshngg/pmcp@latest
```

Ensure `$GOPATH/bin` is in your `$PATH`.

---

## Usage

Run the server by specifying your Prometheus address and preferred transport:

```bash
# Default (stdio transport)
pmcp --prom-addr="http://localhost:9090"

# HTTP transport
pmcp --prom-addr="http://localhost:9090" --transport=http --mcp-addr="localhost:8080"

# SSE transport
pmcp --prom-addr="http://localhost:9090" --transport=sse --mcp-addr="localhost:8080"
```

### Command Line Flags

| Flag         | Description                                       | Default                 |
| ------------ | ------------------------------------------------- | ----------------------- |
| `-help`      | Show help information.                            | N/A                     |
| `-mcp-addr`  | Address for the MCP server to listen on.          | `localhost:8080`        |
| `-prom-addr` | Prometheus server URL.                            | `http://localhost:9090` |
| `-transport` | Communication transport (`stdio`, `http`, `sse`). | `stdio`                 |
| `-version`   | Print version and exit.                           | N/A                     |

---

## Binding Blocks

### Tools

* **Instant Query**: Evaluate an instant query at a single point in time.
* **Range Query**: Evaluate an expression query over a range of time.
* **Find Series by Labels**: Return the list of time series that match a certain label set.
* **List Label Names**: Return a list of label names.
* **List Label Values**: Return a list of label values for a provided label name.
* **Target Metadata Query**: Return metadata about metrics currently scraped from targets.
* **Metric Metadata Query**: Return metadata about metrics currently scraped from targets. However, it does not provide any target information.
* **Target Discovery**: Return an overview of the current state of the Prometheus target discovery.
* **Alert Query**: Return a list of all active alerts.
* **Rule Query**: Return a list of alerting and recording rules that are currently loaded. In addition it returns the currently active alerts fired by the Prometheus instance of each alerting rule.
* **Alertmanager Discovery**: Return an overview of the current state of the Prometheus alertmanager discovery.
* **Config**: Return currently loaded configuration file.
* **Flags**: Return flag values that Prometheus was configured with.
* **Runtime Information**: Return various runtime information properties about the Prometheus server.
* **Build Information**: Return various build information properties about the Prometheus server.
* **TSDB Stats**: Return various cardinality statistics about the Prometheus TSDB.
* **WAL Replay Stats**: Return information about the WAL replay.
* **TSDB Snapshot**: Create a snapshot of all current data into snapshots/<datetime>-<rand>.
* **Delete Series**: Delete data for a selection of series in a time range.
* **Clean Tombstones**: Remove the deleted data from disk and cleans up the existing tombstones.
* **Health Check**: Check Prometheus health.
* **Readiness Check**: Check if Prometheus is ready to serve traffic (i.e. respond to queries).
* **Reload**: Trigger a reload of the Prometheus configuration and rule files.
* **Quit**: Trigger a graceful shutdown of Prometheus.

### Prompts

* **All Available Metrics**: Return a list of every metric exposed by the Prometheus instance.

---

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss improvements.

---

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

* Built with [Model Context Protocol](https://github.com/modelcontextprotocol/go-sdk)
* Powered by [Prometheus](https://prometheus.io/)
