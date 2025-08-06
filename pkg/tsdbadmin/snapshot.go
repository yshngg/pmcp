package tsdbadmin

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// URL query parameters:
// skip_head=<bool>: Skip data present in the head block. Optional.
// curl -XPOST http://localhost:9090/api/v1/admin/tsdb/snapshot
type SnapshotParams struct {
	SkipHead bool `json:"skip_head,omitempty" jsonschema:""`
}
type SnapshotResult = v1.SnapshotResult

func (a *tsdbAdmin) SnapshotHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[SnapshotParams]) (*mcp.CallToolResultFor[SnapshotResult], error) {
	result, err := a.API.Snapshot(ctx, params.Arguments.SkipHead)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[SnapshotResult]{
		Content:           []mcp.Content{&mcp.TextContent{Text: string(content)}},
		StructuredContent: result,
	}, nil
}
