package config

import (
	"os"
)

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}

	return value
}

// DatabasePassword returns the value from environment
func DatabasePassword() string {
	return getEnv("DatabasePassword", "")
}

// DatabaseUser returns the value from environment
func DatabaseUser() string {
	return getEnv("DatabaseUser", "")
}

// DatabaseURL returns the value from environment
func DatabaseURL() string {
	return getEnv("DatabaseURL", "localhost")
}

// TwitterConsumerKey returns the value from environment
func TwitterConsumerKey() string {
	return getEnv("TwitterConsumerKey", "")
}

// TwitterConsumerSecret returns the value from environment
func TwitterConsumerSecret() string {
	return getEnv("TwitterConsumerSecret", "")
}

// AccessKey returns the value from environment
func AccessKey() string {
	return getEnv("AccessKey", "")
}

// AccessTokenSecret returns the value from environment
func AccessTokenSecret() string {
	return getEnv("AccessTokenSecret", "")
}

// JWTAlgorithm returns the value from environment
func JWTAlgorithm() string {
	return getEnv("JWTAlgorithm", "")
}

// JWTSecret returns the value from environment
func JWTSecret() string {
	return getEnv("JWTSecret", "")
}
