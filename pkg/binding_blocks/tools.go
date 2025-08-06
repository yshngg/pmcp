package bindingblocks

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	alertquery "github.com/yshngg/pmcp/pkg/alert_query"
	alertmanagerdiscover "github.com/yshngg/pmcp/pkg/alertmanager_discover"
	expressionquery "github.com/yshngg/pmcp/pkg/expression_query"
	"github.com/yshngg/pmcp/pkg/manage"
	metadataquery "github.com/yshngg/pmcp/pkg/metadata_query"
	rulequery "github.com/yshngg/pmcp/pkg/rule_query"
	targetdiscover "github.com/yshngg/pmcp/pkg/target_discover"
)

// addTools registers Prometheus query tools with the MCP server.
// It adds tools for expression queries (instant and range) and metadata queries.
func (b *binder) addTools() {
	// Expression queries
	// Query language expressions may be evaluated at a single instant or over a range of time.
	{
		expressionQuerier := expressionquery.NewExpressionQuerier(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Prometheus Instant Query",
			Description: "Run a Prometheus expression and get the current value for a metric or calculation at a specific time. Use this to check the latest status or value of any metric.",
		}, expressionQuerier.InstantQueryHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Prometheus Range Query",
			Description: "Run a Prometheus expression over a time range to get historical values for a metric or calculation. Use this to analyze trends or patterns over time.",
		}, expressionQuerier.RangeQueryHandler)
	}

	// Querying metadata
	// Prometheus offers a set of API endpoints to query metadata about series and their labels.
	{
		metadataQuerier := metadataquery.NewMetadataQuerier(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Find Series by Labels",
			Description: "List all time series that match specific label filters. Use this to discover which series exist for given label criteria.",
		}, metadataQuerier.SeriesHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "List Label Names",
			Description: "Get all label names used in the Prometheus database. Use this to explore available labels for filtering or grouping.",
		}, metadataQuerier.LabelNamesHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "List Label Values",
			Description: "Get all possible values for a specific label name. Use this to see which values a label can take for filtering or selection.",
		}, metadataQuerier.LabelValuesHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Target Metadata Query",
			Description: "Get metadata about metrics currently scraped from targets.",
		}, metadataQuerier.TargetMetadataQueryHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Metric Metadata Query",
			Description: "Get metadata about metrics currently scraped from targets. However, it does not provide any target information.",
		}, metadataQuerier.MetricsMetadataQueryHandler)
	}

	// Targets
	// An overview of the current state of the Prometheus target discovery.
	{
		targetDiscoverer := targetdiscover.NewTargetDiscoverer(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Target Discovery",
			Description: "Get an overview of the current state of the Prometheus target discovery.",
		}, targetDiscoverer.TargetDiscoverHandler)
	}

	// Rules
	// A list of alerting and recording rules that are currently loaded.
	{
		ruleQuerier := rulequery.NewRuleQuerier(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Rule Query",
			Description: "Get a list of alerting and recording rules that are currently loaded. In addition it returns the currently active alerts fired by the Prometheus instance of each alerting rule.",
		}, ruleQuerier.RuleQueryHandler)
	}

	// Alerts
	// A list of all active alerts.
	{
		alertQuerier := alertquery.NewAlertQuerier(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Alert Query",
			Description: "Get a list of all active alerts.",
		}, alertQuerier.AlertQueryHandler)
	}

	// Alertmanagers
	// An overview of the current state of the Prometheus alertmanager discovery.
	{
		alertmanagerDiscoverer := alertmanagerdiscover.NewAlertmanagerDiscoverer(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Alertmanager Discovery",
			Description: "Get an overview of the current state of the Prometheus alertmanager discovery.",
		}, alertmanagerDiscoverer.AlertmanagerDiscoverHandler)
	}

	// Management API
	// Prometheus provides a set of management APIs to facilitate automation and integration.
	{
		manager := manage.NewManager(b.api)
		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Health Check",
			Description: "Check whether Prometheus is healthy.",
		}, manager.HealthCheckHandler)

		mcp.AddTool(b.server, &mcp.Tool{
			Name:        "Readiness Check",
			Description: "Check whether Prometheus is ready to serve traffic.",
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
