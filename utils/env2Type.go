package utils

import (
	"os"
	"strconv"
)

func GetEnvString(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return def
}

func GetEnvInt(key string, def int) int {
	if val, ok := os.LookupEnv(key); ok {
		if val, err := strconv.Atoi(val); err == nil {
			return val
		}
	}

	return def
}

func GetEnvBool(key string, def bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		return val == "true"
	}

	return def
}
