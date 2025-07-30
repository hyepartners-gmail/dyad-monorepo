package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config defines the app config schema
type Config struct {
	ProjectID   string `yaml:"project_id"`
	Port        string `yaml:"port"`
	Environment string `yaml:"environment"`
	JWTSecret   string `yaml:"jwt_secret"`
}

// Load returns the app configuration from env or config.yaml
func Load() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println("No config.yaml found, falling back to env vars")
		return loadFromEnv(), nil
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// loadFromEnv builds config from environment variables
func loadFromEnv() *Config {
	return &Config{
		ProjectID:   os.Getenv("PROJECT_ID"),
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("GO_ENV", "development"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
