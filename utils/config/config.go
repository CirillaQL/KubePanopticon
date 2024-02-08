package config

import (
	"github.com/CirillaQL/kubepanopticon/utils/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v3"
	"os"
)

var GlobalConfig *Config

type Config struct {
	PrometheusURL string `yaml:"prometheus_url"`
}

func LoadConfig() error {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		logger.Log.Error("failed to read config file", zap.Error(err))
		return err
	}

	if err := yaml.Unmarshal(data, &GlobalConfig); err != nil {
		return errors.Wrapf(err, "parse config file %s failed", GlobalConfig)
	}

	return nil
}
