package expressionquery

import (
	"testing"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type dummyClient struct {
	v1.API
}

func TestNewExpressionQuerier(t *testing.T) {
	q := NewExpressionQuerier(&dummyClient{})
	if q == nil {
		t.Fatal("expected non-nil ExpressionQuerier")
	}
}
