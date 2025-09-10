Testing library/framework in use: Go standard library "testing".
We avoided new dependencies and used table-free explicit tests for clarity.
Tests added:
- internal/manage/health_check_handler_test.go

If ManagementResult isn't defined at build time, enable the test shim with:
  go test -tags testshim ./internal/manage/...