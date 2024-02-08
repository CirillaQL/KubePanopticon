package prometheus

import (
	"context"
	"github.com/CirillaQL/kubepanopticon/utils/logger"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"go.uber.org/zap"
	"time"
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

func QueryWithTimeRange(client *api.Client, starTime, endTime time.Time, step time.Duration, promeQL string) model.Matrix {
	v1api := v1.NewAPI(*client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: starTime,
		End:   endTime,
		Step:  step,
	}

	result, warnings, err := v1api.QueryRange(ctx, promeQL, r, v1.WithTimeout(10*time.Second))
	if err != nil {
		logger.Log.Error("error querying prometheus", zap.Error(err))
		return nil
	}
	if len(warnings) > 0 {
		logger.Log.Warn("warnings querying prometheus", zap.Any("warning", warnings))
		return nil
	}

	switch t := result.Type(); t {
	case model.ValMatrix:
		prometheusQueryData := result.(model.Matrix)
		return prometheusQueryData
	default:
		logger.Log.Error("Wrong data format. Can't handle this type of prometheus response.")
		return nil
	}
}
