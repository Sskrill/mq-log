package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	URI string
}

func NewCfg() (*Config, error) {
	if err := godotenv.Load(); err != nil {

		return nil, err
	}
	cfg := new(Config)
	if err := envconfig.Process("mq", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
