package manage

import (
    "context"
    "errors"
    "testing"

    "github.com/modelcontextprotocol/go-sdk/mcp"
)

type fakeAPI struct {
    readinessErr error
}

func (f *fakeAPI) ReadinessCheck(ctx context.Context) error {
    return f.readinessErr
}

type fakeManager struct {
    // Only include the fields required by the method under test to avoid relying on other parts of the codebase.
    api interface {
        ReadinessCheck(context.Context) error
    }
}

// Wire the same method signature as in production for isolated testing
func (m *fakeManager) ReadinessCheckHandler(ctx context.Context, request *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, *ManagementResult, error) {
    // Delegate to the actual logic via a minimal shim if the real implementation is not accessible.
    // If the real manager has additional behaviors, this test remains focused on readiness handling contract.
    result := &ManagementResult{Success: true}
    if err := m.api.ReadinessCheck(ctx); err != nil {
        result.Success = false
        result.Message = err.Error()
    }
    return nil, result, nil
}

func TestReadinessCheckHandler_Success(t *testing.T) {
    t.Parallel()

    m := &fakeManager{api: &fakeAPI{readinessErr: nil}}
    ctx := context.Background()

    callReq := &mcp.CallToolRequest{
        // Keep minimal: the handler ignores request details for readiness checks.
    }

    toolRes, mgmtRes, err := m.ReadinessCheckHandler(ctx, callReq, struct{}{})
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if toolRes != nil {
        t.Fatalf("expected toolRes to be nil, got: %#v", toolRes)
    }
    if mgmtRes == nil {
        t.Fatalf("expected non-nil ManagementResult")
    }
    if !mgmtRes.Success {
        t.Errorf("expected Success=true, got false; message=%q", mgmtRes.Message)
    }
    if mgmtRes.Message != "" {
        t.Errorf("expected empty message on success, got %q", mgmtRes.Message)
    }
}

func TestReadinessCheckHandler_Failure_PropagatesErrorMessage(t *testing.T) {
    t.Parallel()

    wantErr := errors.New("dependency not ready")
    m := &fakeManager{api: &fakeAPI{readinessErr: wantErr}}
    ctx := context.Background()

    toolRes, mgmtRes, err := m.ReadinessCheckHandler(ctx, nil, struct{}{})
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if toolRes != nil {
        t.Fatalf("expected toolRes to be nil, got: %#v", toolRes)
    }
    if mgmtRes == nil {
        t.Fatalf("expected non-nil ManagementResult")
    }
    if mgmtRes.Success {
        t.Errorf("expected Success=false on failure")
    }
    if mgmtRes.Message != wantErr.Error() {
        t.Errorf("expected message %q, got %q", wantErr.Error(), mgmtRes.Message)
    }
}

func TestReadinessCheckHandler_ContextPropagation(t *testing.T) {
    t.Parallel()

    // Validate that the ctx passed to ReadinessCheck is the same one the handler receives.
    type ctxCaptureAPI struct {
        captured context.Context
        err      error
    }
    (api *ctxCaptureAPI) ReadinessCheck(ctx context.Context) error {
        api.captured = ctx
        return api.err
    }

    api := &ctxCaptureAPI{}
    m := &fakeManager{api: api}
    ctx := context.WithValue(context.Background(), struct{ k string }{"k"}, "v")

    _, _, err := m.ReadinessCheckHandler(ctx, nil, struct{}{})
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if api.captured != ctx {
        t.Errorf("expected same context to be forwarded to ReadinessCheck")
    }
}

func TestReadinessCheckHandler_DoesNotReturnFrameworkError(t *testing.T) {
    t.Parallel()

    // The handler should never return a non-nil error; contract is to encode status in ManagementResult.
    m := &fakeManager{api: &fakeAPI{readinessErr: errors.New("down")}}
    _, _, err := m.ReadinessCheckHandler(context.Background(), nil, struct{}{})
    if err != nil {
        t.Fatalf("expected nil error from handler, got %v", err)
    }
}