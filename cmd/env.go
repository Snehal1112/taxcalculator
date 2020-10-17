package cmd

import (
	"os"
)

// envOrDefault returns the value of an env-variable or the default if the env-var is not set
func envOrDefault(name string, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}

	return v
}