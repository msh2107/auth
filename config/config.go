package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	GRPC `yaml:"grpc"`
}

type GRPC struct {
	Host     string `yaml:"host" env-default:"localhost"`
	GRPCPort string `yaml:"grpc_port" env-default:"50051"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
