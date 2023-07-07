package config

import (
	"fmt"
	"os"
	"time"
)

// All configuration is through environment variables

const POSTGRES_URI_ENV_VAR = "POSTGRES_URI"
const QUERY_TIME_THRESHOLD_ENV_VAR = "QUERY_TIME_THRESHOLD"
const DEFAULT_QUERY_TIME_THRESHOLD = "5s"

type Config struct {
	postgresUri        string
	queryTimeThreshold time.Duration
}

func NewConfigFromEnvVars() (*Config, error) {
	postgresUri, err := getPostgresUri()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting postgres uri: %v", err)
	}

	queryTimeThreshold, err := getQueryTimeThreshold()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting query time threshold: %v", err)
	}

	return &Config{
		postgresUri:        postgresUri,
		queryTimeThreshold: queryTimeThreshold,
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

// Get query time threshold. Any slow running query taking time more than this threshold will be considered
// for alerting and killing
func getQueryTimeThreshold() (time.Duration, error) {
	queryTimeThresholdString, ok := os.LookupEnv(QUERY_TIME_THRESHOLD_ENV_VAR)
	if !ok {
		queryTimeThresholdString = DEFAULT_QUERY_TIME_THRESHOLD
	}

	queryTimeThreshold, err := time.ParseDuration(queryTimeThresholdString)
	if err != nil {
		return 0, fmt.Errorf("error occurred while parsing query time threshold value %s: %v. Example of valid values are `5s`, `1m`, `1h`, `24h`", queryTimeThreshold, err)
	}

	return queryTimeThreshold, nil
}

func (c *Config) GetPostgresUri() string {
	return c.postgresUri
}

func (c *Config) GetQueryTimeThreshold() time.Duration {
	return c.queryTimeThreshold
}
