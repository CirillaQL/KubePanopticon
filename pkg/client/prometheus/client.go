package prometheus

import (
	"github.com/prometheus/client_golang/api"
)

func NewPrometheusClient(prometheusAddress string) (*api.Client, error) {
	client, err := api.NewClient(api.Config{
		Address: prometheusAddress,
	})
	if err != nil {
		return nil, err
	}
	return &client, nil
}
