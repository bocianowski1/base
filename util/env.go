package util

import (
	"os"
	"strings"
)

// IsDev checks if the environment is dev / development
func IsDev() bool {
	return strings.HasPrefix(getEnv(), "dev")
}

// IsProd checks if the environment is prod / production
func IsProd() bool {
	return strings.HasPrefix(getEnv(), "prod")
}

func getEnv() string {
	env := os.Getenv("ENV")
	env = strings.ToLower(env)
	if env == "" {
		env = "dev"
	}
	return env
}
