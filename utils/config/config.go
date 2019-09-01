package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	HTTP      HTTPConfig
	FPL       FPLConfig
	DebugMode bool
	UserRoles []string
	MaxUsers  int
}

// FPLConfig Fantasy Premier League Config
type FPLConfig struct {
	Me       int
	VVLeague int
}

// HTTPConfig http/s defaults
type HTTPConfig struct {
	HTTPHost     string
	HTTPPort     int
	HTTPSDomains []string
	HTTPSPort    int
}

// HTTP_HOST=localhost
// HTTPS_PORT=8111
// HTTPS_DOMAINS=54apenwith.com,www.54apenwith.com
// HTTPS_PORT=443

// New returns a new Config struct
func New() *Config {
	return &Config{
		HTTP: HTTPConfig{
			HTTPHost:     getEnv("HTTP_HOST", "localhost"),
			HTTPPort:     getEnvAsInt("HTTP_PORT", 8111),
			HTTPSDomains: getEnvAsSlice("HTTPS_DOMAINS", []string{""}, ","),
			HTTPSPort:    getEnvAsInt("HTTPS_PORT", 443),
		},
		FPL: FPLConfig{
			Me:       getEnvAsInt("FPL_ME", -1),
			VVLeague: getEnvAsInt("FPL_VV_LEAGUE", -1),
		},
		DebugMode: getEnvAsBool("DEBUG_MODE", true),
		UserRoles: getEnvAsSlice("USER_ROLES", []string{"admin"}, ","),
		MaxUsers:  getEnvAsInt("MAX_USERS", 1),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
