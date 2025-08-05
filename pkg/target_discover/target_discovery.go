package targetdiscover

import (
	"context"
	"encoding/json"
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

func (d *targetDiscoverer) TargetDiscoverHandler(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[TargetDiscoverParams]) (*mcp.CallToolResultFor[TargetDiscoverResult], error) {
	var (
		result TargetDiscoverResult
		err    error
	)
	if result, err = d.API.Targets(ctx); err != nil {
		return nil, err
	}

	// Filter by state
	if len(params.Arguments.State) == 0 {
		params.Arguments.State = TargetStateAny
	}
	switch params.Arguments.State {
	case TargetStateActive:
		result.Dropped = nil
	case TargetStateDropped:
		result.Active = nil
	case TargetStateAny:
		break
	default:
		return nil, fmt.Errorf("invalid state: %s, must be active, dropped or any", params.Arguments.State)
	}

	// Filter by scrape pool
	scrapePool := params.Arguments.ScrapePool
	if len(scrapePool) != 0 {
		// Dropped targets don't have scrape pool field
		result.Dropped = nil

		activeTargets := []v1.ActiveTarget{}
		for _, v := range result.Active {
			if v.ScrapePool == scrapePool {
				activeTargets = append(activeTargets, v)
			}
		}
		result.Active = activeTargets
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[TargetDiscoverResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
