package env

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

// pgConfig - .
type pgConfig struct {
	dsn string
}

// NewPGConfig - .
func NewPGConfig() (*pgConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("user dsn not found")
	}

	return &pgConfig{dsn: dsn}, nil
}

// DSN - .
func (cfg pgConfig) DSN() string {
	return cfg.dsn
}
