package bindingblocks

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yshngg/pmcp/internal/alertmanagerdiscover"
	"github.com/yshngg/pmcp/internal/alertquery"
	"github.com/yshngg/pmcp/internal/expressionquery"
	"github.com/yshngg/pmcp/internal/manage"
	"github.com/yshngg/pmcp/internal/metadataquery"
	"github.com/yshngg/pmcp/internal/rulequery"
	"github.com/yshngg/pmcp/internal/statusexpose"
	"github.com/yshngg/pmcp/internal/targetdiscover"
	"github.com/yshngg/pmcp/internal/tsdbadmin"
)

// addTools registers Prometheus query tools with the MCP server.
// It adds tools for expression queries (instant and range) and metadata queries.
func (b *binder) addTools() {
	// Expression queries
	// Query language expressions may be evaluated at a single instant or over a range of time.
	{
		expressionQuerier := expressionquery.NewExpressionQuerier(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Instant Query",
			Description: "Evaluate an instant query at a single point in time.",
		}, expressionQuerier.InstantQueryHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Range Query",
			Description: "Evaluate an expression query over a range of time.",
		}, expressionQuerier.RangeQueryHandler)
	}

	// Querying metadata
	// Prometheus offers a set of API endpoints to query metadata about series and their labels.
	{
		metadataQuerier := metadataquery.NewMetadataQuerier(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Find Series by Labels",
			Description: "Return the list of time series that match a certain label set.",
		}, metadataQuerier.SeriesHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "List Label Names",
			Description: "Return a list of label names.",
		}, metadataQuerier.LabelNamesHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "List Label Values",
			Description: "Return a list of label values for a provided label name.",
		}, metadataQuerier.LabelValuesHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Target Metadata Query",
			Description: "Return metadata about metrics currently scraped from targets.",
		}, metadataQuerier.TargetMetadataQueryHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Metric Metadata Query",
			Description: "Return metadata about metrics currently scraped from targets. However, it does not provide any target information.",
		}, metadataQuerier.MetricsMetadataQueryHandler)
	}

	// Targets
	// An overview of the current state of the Prometheus target discovery.
	{
		targetDiscoverer := targetdiscover.NewTargetDiscoverer(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Target Discovery",
			Description: "Return an overview of the current state of the Prometheus target discovery.",
		}, targetDiscoverer.TargetDiscoverHandler)
	}

	// Rules
	// A list of alerting and recording rules that are currently loaded.
	{
		ruleQuerier := rulequery.NewRuleQuerier(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Rule Query",
			Description: "Return a list of alerting and recording rules that are currently loaded. In addition it returns the currently active alerts fired by the Prometheus instance of each alerting rule.",
		}, ruleQuerier.RuleQueryHandler)
	}

	// Alerts
	// A list of all active alerts.
	{
		alertQuerier := alertquery.NewAlertQuerier(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Alert Query",
			Description: "Return a list of all active alerts.",
		}, alertQuerier.AlertQueryHandler)
	}

	// Alertmanagers
	// An overview of the current state of the Prometheus alertmanager discovery.
	{
		alertmanagerDiscoverer := alertmanagerdiscover.NewAlertmanagerDiscoverer(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Alertmanager Discovery",
			Description: "Return an overview of the current state of the Prometheus alertmanager discovery.",
		}, alertmanagerDiscoverer.AlertmanagerDiscoverHandler)
	}

	// Status
	// Expose current Prometheus configuration.
	{
		statusExposer := statusexpose.NewStatusExposer(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Config",
			Description: "Return currently loaded configuration file.",
		}, statusExposer.ConfigExposeHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Flags",
			Description: "Return flag values that Prometheus was configured with.",
		}, statusExposer.FlagsExposeHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Runtime Information",
			Description: "Return various runtime information properties about the Prometheus server.",
		}, statusExposer.RuntimeInformationExposeHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Build Information",
			Description: "Return various build information properties about the Prometheus server.",
		}, statusExposer.BuildInformationExposeHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "TSDB Stats",
			Description: "Return various cardinality statistics about the Prometheus TSDB.",
		}, statusExposer.TSDBStatsExposeHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "WAL Replay Stats",
			Description: "Return information about the WAL replay.",
		}, statusExposer.WALReplayStatsExposeHandler)
	}

	// TSDB Admin APIs
	// Expose database functionalities for the advanced user.
	{
		tsdbAdmin := tsdbadmin.NewTSDBAdmin(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "TSDB Snapshot",
			Description: "Create a snapshot of all current data into snapshots/<datetime>-<rand> under the TSDB's data directory and returns the directory as response. It will optionally skip snapshotting data that is only present in the head block, and which has not yet been compacted to disk.",
		}, tsdbAdmin.SnapshotHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Delete Series",
			Description: "Delete data for a selection of series in a time range. The actual data still exists on disk and is cleaned up in future compactions or can be explicitly cleaned up by hitting the Clean Tombstones endpoint. Not mentioning both start and end times would clear all the data for the matched series in the database.",
		}, tsdbAdmin.DeleteSeriesHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Clean Tombstones",
			Description: "Remove the deleted data from disk and cleans up the existing tombstones. This can be used after deleting series to free up space.",
		}, tsdbAdmin.CleanTombstonesHandler)
	}

	// Management API
	// Prometheus provides a set of management APIs to facilitate automation and integration.
	{
		manager := manage.NewManager(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Health Check",
			Description: "Check Prometheus health.",
		}, manager.HealthCheckHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Readiness Check",
			Description: "Check if Prometheus is ready to serve traffic (i.e. respond to queries).",
		}, manager.ReadinessCheckHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Reload",
			Description: "Trigger a reload of the Prometheus configuration and rule files.",
		}, manager.ReloadHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Quit",
			Description: "Trigger a graceful shutdown of Prometheus.",
		}, manager.QuitHandler)
	}
}
