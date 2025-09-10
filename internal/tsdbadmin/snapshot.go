package tsdbadmin

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// URL query parameters:
// skip_head=<bool>: Skip data present in the head block. Optional.
// curl -XPOST http://localhost:9090/api/v1/admin/tsdb/snapshot
type SnapshotParams struct {
	SkipHead bool `json:"skip_head,omitempty" jsonschema:"<bool>: Skip data present in the head block. Optional."`
}
type SnapshotResult = v1.SnapshotResult

func (a *tsdbAdmin) SnapshotHandler(ctx context.Context, request *mcp.CallToolRequest, input *SnapshotParams) (*mcp.CallToolResult, *SnapshotResult, error) {
	result, err := a.API.Snapshot(ctx, input.SkipHead)
	if err != nil {
		return nil, nil, err
	}
	return nil, &result, nil
}
