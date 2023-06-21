package config

import (
	"fmt"
	"os"
)

// All configuration is through environment variables

const POSTGRES_URI_ENV_VAR = "POSTGRES_URI"

type Config struct {
	postgresUri string
}

func NewConfigFromEnvVars() (*Config, error) {
	postgresUri, err := getPostgresUri()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting postgres uri: %v", err)
	}

	return &Config{
		postgresUri: postgresUri,
	}, nil
}

// Get postgres uri
func getPostgresUri() (string, error) {
	postgresUri, ok := os.LookupEnv(POSTGRES_URI_ENV_VAR)
	if !ok {
		return "", fmt.Errorf("%s environment variable is not defined and is required. Please define it", POSTGRES_URI_ENV_VAR)
	}

	return postgresUri, nil
}

func (c *Config) GetPostgresUri() string {
	return c.postgresUri
}
