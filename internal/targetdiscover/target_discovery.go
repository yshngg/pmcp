package targetdiscover

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type TargetState string

const (
	TargetStateActive  = "active"
	TargetStateDropped = "dropped"
	TargetStateAny     = "any"
)

type TargetDiscoverParams struct {
	State      TargetState `json:"state,omitzero" jsonschema:"Allow to filter by active or dropped targets, (e.g., state=active, state=dropped, state=any)."`
	ScrapePool string      `json:"scrapePool,omitzero" jsonschema:"Allow to filter by scrape pool name."`
}

type TargetDiscoverResult = v1.TargetsResult

func (d *targetDiscoverer) TargetDiscoverHandler(ctx context.Context, request *mcp.CallToolRequest, input *TargetDiscoverParams) (*mcp.CallToolResult, *TargetDiscoverResult, error) {
	var (
		result = &TargetDiscoverResult{}
		err    error
	)
	if *result, err = d.API.Targets(ctx); err != nil {
		return nil, nil, err
	}

	// Filter by scrape pool, which only affects active targets.
	scrapePool := input.ScrapePool
	if len(scrapePool) != 0 {
		// Dropped targets don't have scrape pool field
		result.Dropped = nil

		n := 0
		for _, v := range result.Active {
			if v.ScrapePool == scrapePool {
				result.Active[n] = v
				n++
			}
		}
		result.Active = result.Active[:n]
	}

	// Filter by state
	state := input.State
	if len(state) == 0 {
		state = TargetStateAny
	}
	switch state {
	case TargetStateActive:
		result.Dropped = nil
	case TargetStateDropped:
		result.Active = nil
	case TargetStateAny:
		// keep both active and dropped targets
	default:
		return nil, nil, fmt.Errorf("invalid state: %s, must be active, dropped or any", input.State)
	}
	return nil, result, nil
}
