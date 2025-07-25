package client

import _ "github.com/prometheus/client_golang/api/prometheus/v1"

const (
	InstantQueryPath = "/query"
	RangeQueryPath   = "/query_range"
)

type ExpressionQuerier interface {
	InstantQuerier
	RangeQuerier
}

type InstantQuerier interface {
	Query()
}

type RangeQuerier interface{}
