package bind

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	expressionquery "github.com/yshngg/pmcp/pkg/expression_query"
	metadataquery "github.com/yshngg/pmcp/pkg/metadata_query"
)

// addTools registers Prometheus query tools with the given MCP server.
// It creates a Prometheus client and adds instant and range query handlers.
//
// Returns an error if the Prometheus client cannot be created.
func (b *binder) addTools(server *mcp.Server) error {
	// Expression queries
	// Query language expressions may be evaluated at a single instant or over a range of time.
	{
		expressionQuerier := expressionquery.NewExpressionQuerier(b.promCli)
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
		metadataQuerier := metadataquery.NewMetadataQuerier(b.promCli)
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
