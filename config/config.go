package config

import (
	"os"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func GetPort() string {
	port := GetConfig("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}
