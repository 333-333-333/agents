// internal/shared/config/config.go
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env       string
	HTTP      HTTPConfig
	GRPC      GRPCConfig
	Database  DatabaseConfig
	Messaging MessagingConfig
	Storage   StorageConfig
}

type HTTPConfig struct {
	Port int
}

type GRPCConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	MaxOpen  int
	MaxIdle  int
	MaxLife  time.Duration
}

type MessagingConfig struct {
	Provider string // "nats", "rabbitmq"
	URL      string
}

type StorageConfig struct {
	Provider       string
	Endpoint       string
	Region         string
	ForcePathStyle bool
}

func Load() (*Config, error) {
	return &Config{
		Env: getEnv("APP_ENV", "development"),
		HTTP: HTTPConfig{
			Port: getEnvInt("HTTP_PORT", 8080),
		},
		GRPC: GRPCConfig{
			Port: getEnvInt("GRPC_PORT", 9090),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: requireEnv("DB_PASSWORD"),
			DBName:   getEnv("DB_NAME", "bastet"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
			MaxOpen:  getEnvInt("DB_MAX_OPEN", 25),
			MaxIdle:  getEnvInt("DB_MAX_IDLE", 5),
			MaxLife:  getEnvDuration("DB_MAX_LIFE", 5*time.Minute),
		},
		Messaging: MessagingConfig{
			Provider: getEnv("MESSAGING_PROVIDER", "nats"),
			URL:      getEnv("MESSAGING_URL", "nats://localhost:4222"),
		},
		Storage: StorageConfig{
			Provider:       getEnv("STORAGE_PROVIDER", "minio"),
			Endpoint:       getEnv("STORAGE_ENDPOINT", "http://localhost:9000"),
			Region:         getEnv("STORAGE_REGION", "us-east-1"),
			ForcePathStyle: getEnvBool("STORAGE_FORCE_PATH_STYLE", true),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("required environment variable %s is not set", key))
	}
	return v
}

func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return i
}

func getEnvBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return b
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return fallback
	}
	return d
}
