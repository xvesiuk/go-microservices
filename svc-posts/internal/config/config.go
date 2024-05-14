package config

import "time"

type Config struct {
	Port     int
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string

	// Connection pool related
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
	ConnectTimeout    time.Duration
}

func New() *Config {
	return &Config{
		Port: 8001,
		Database: DatabaseConfig{
			Host:              "localhost",
			Port:              5433,
			User:              "dev",
			Password:          "dev",
			Database:          "postgres",
			MaxConns:          4,
			MinConns:          0,
			MaxConnLifetime:   time.Hour,
			MaxConnIdleTime:   time.Minute * 30,
			HealthCheckPeriod: time.Minute,
			ConnectTimeout:    time.Second * 5,
		},
	}
}
