package env

import (
	"fmt"
	"os"
)

// GetenvOrDefault retrieves the value of the environment variable named by the key.
// If the value is the empty string, returns defaultValue
func GetenvOrDefault(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

// GetenvOrDie retrieves the value of the environment variable named by the key.
// This function panic() if the value match the empty string
func GetenvOrDie(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("missing required env var %s", key))
	}
	return v
}

// GetEnv retrieves the value of the environment variable named by the key
// if defined, otherwise return an error
func GetEnv(k string) (string, error) {
	v := os.Getenv(k)
	if v == "" {
		return "", fmt.Errorf("%s environment variable not set", k)
	}
	return v, nil
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
