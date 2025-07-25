package client

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

const APIVersion = "/api/v1"

type PrometheusClient interface {
	v1.API
}

func New(addr string, client *http.Client, roundTripper http.RoundTripper) (PrometheusClient, error) {
	cli, err := api.NewClient(api.Config{
		Address:      addr,
		Client:       client,
		RoundTripper: roundTripper,
	})
	if err != nil {
		return nil, fmt.Errorf("new client, err: %w", err)
	}

	return v1.NewAPI(cli), nil
}
