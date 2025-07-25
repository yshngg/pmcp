package client

import (
	"net/http"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type PrometheusClient interface {
}

func New(addr string, client *http.Client, roundTripper http.RoundTripper) (PrometheusClient, error) {
	cli, err := api.NewClient(api.Config{
		Address:      addr,
		Client:       client,
		RoundTripper: roundTripper,
	})
	if err != nil {
		return nil, err
	}

	api := v1.NewAPI(cli)
	return api, nil
}
