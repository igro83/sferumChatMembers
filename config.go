package main

import (
	"os"
)

type SferumConfig struct {
	REMIXDSID string
	CHAT      string
	V         string
	VVK       string
}

type Config struct {
	Sferum SferumConfig
}

func NewConfig() *Config {
	return &Config{
		Sferum: SferumConfig{
			REMIXDSID: getEnv("REMIXDSID", ""),
			CHAT:      getEnv("CHAT", ""),
			V:         getEnv("V", ""),
			VVK:       getEnv("VVK", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
