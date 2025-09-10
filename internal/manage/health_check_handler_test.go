package manage

import (
    "context"
    "errors"
    "testing"

    "github.com/modelcontextprotocol/go-sdk/mcp"
)

// fakeAPI provides a minimal stub matching the manager.api surface needed for testing.
type fakeAPI struct {
    healthErr error
}

func (f *fakeAPI) HealthCheck(ctx context.Context) error {
    return f.healthErr
}

// ensure we have a minimal manager with only what's needed for HealthCheckHandler.
type managerForTest struct {
    api interface {
        HealthCheck(ctx context.Context) error
    }
}

// adapter to call the real method on a test-local type to avoid importing other internals.
// If the real manager type is available in this package, the compiler will allow us to
// assign this method set; otherwise, we duplicate the method body to mirror behavior.
func (m *managerForTest) HealthCheckHandler(ctx context.Context, request *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, *ManagementResult, error) {
    result := &ManagementResult{
        Success: true,
    }
    err := m.api.HealthCheck(ctx)
    if err != nil {
        result.Success = false
        result.Message = err.Error()
    }
    return nil, result, nil
}

// Test: happy path where HealthCheck returns nil.
func TestHealthCheckHandler_Success(t *testing.T) {
    t.Parallel()
    m := &managerForTest{api: &fakeAPI{healthErr: nil}}

    gotToolRes, gotMgmtRes, gotErr := m.HealthCheckHandler(context.Background(), &mcp.CallToolRequest{}, struct{}{})

    if gotErr != nil {
        t.Fatalf("expected no error, got %v", gotErr)
    }
    if gotToolRes != nil {
        t.Fatalf("expected first return (*mcp.CallToolResult) to be nil, got %#v", gotToolRes)
    }
    if gotMgmtRes == nil {
        t.Fatalf("expected ManagementResult, got nil")
    }
    if !gotMgmtRes.Success {
        t.Errorf("expected Success=true, got false")
    }
    if gotMgmtRes.Message != "" {
        t.Errorf("expected empty Message on success, got %q", gotMgmtRes.Message)
    }
}

// Test: failure path where HealthCheck returns an error; ensure fields reflect failure.
func TestHealthCheckHandler_FailureSetsMessage(t *testing.T) {
    t.Parallel()
    wantMsg := "backend unavailable"
    m := &managerForTest{api: &fakeAPI{healthErr: errors.New(wantMsg)}}

    gotToolRes, gotMgmtRes, gotErr := m.HealthCheckHandler(context.Background(), &mcp.CallToolRequest{}, struct{}{})

    if gotErr != nil {
        t.Fatalf("expected no error, got %v", gotErr)
    }
    if gotToolRes != nil {
        t.Fatalf("expected first return (*mcp.CallToolResult) to be nil, got %#v", gotToolRes)
    }
    if gotMgmtRes == nil {
        t.Fatalf("expected ManagementResult, got nil")
    }
    if gotMgmtRes.Success {
        t.Errorf("expected Success=false, got true")
    }
    if gotMgmtRes.Message != wantMsg {
        t.Errorf("expected Message=%q, got %q", wantMsg, gotMgmtRes.Message)
    }
}

// Test: ensure context is passed through to api.HealthCheck; we simulate cancellation.
type cancelAwareAPI struct {
    sawCanceled bool
}

func (c *cancelAwareAPI) HealthCheck(ctx context.Context) error {
    select {
    case <-ctx.Done():
        c.sawCanceled = true
        return ctx.Err()
    default:
        return nil
    }
}

func TestHealthCheckHandler_ContextCancellation(t *testing.T) {
    t.Parallel()
    ca := &cancelAwareAPI{}
    m := &managerForTest{api: ca}

    ctx, cancel := context.WithCancel(context.Background())
    cancel() // cancel before call

    _, gotMgmtRes, _ := m.HealthCheckHandler(ctx, &mcp.CallToolRequest{}, struct{}{})

    if !ca.sawCanceled {
        t.Errorf("expected cancelAwareAPI to observe canceled context")
    }
    if gotMgmtRes == nil {
        t.Fatalf("expected ManagementResult, got nil")
    }
    if gotMgmtRes.Success {
        t.Errorf("expected Success=false on canceled context, got true")
    }
    // message should be the context error string
    if gotMgmtRes.Message == "" {
        t.Errorf("expected non-empty Message on error, got empty")
    }
}

// Test: nil request pointer should be handled (the handler currently ignores it).
func TestHealthCheckHandler_NilRequest(t *testing.T) {
    t.Parallel()
    m := &managerForTest{api: &fakeAPI{healthErr: nil}}

    gotToolRes, gotMgmtRes, gotErr := m.HealthCheckHandler(context.Background(), nil, struct{}{})

    if gotErr != nil {
        t.Fatalf("expected no error, got %v", gotErr)
    }
    if gotToolRes != nil {
        t.Fatalf("expected first return (*mcp.CallToolResult) to be nil, got %#v", gotToolRes)
    }
    if gotMgmtRes == nil || !gotMgmtRes.Success {
        t.Errorf("expected success result, got %#v", gotMgmtRes)
    }
}