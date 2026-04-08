package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	AppName         string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func Load() Config {
	loadDotEnv(".env")

	return Config{
		AppName:         getEnv("APP_NAME", "go-learning-api"),
		Port:            getEnv("APP_PORT", "8080"),
		ReadTimeout:     getEnvAsDuration("APP_READ_TIMEOUT_SECONDS", 5),
		WriteTimeout:    getEnvAsDuration("APP_WRITE_TIMEOUT_SECONDS", 10),
		IdleTimeout:     getEnvAsDuration("APP_IDLE_TIMEOUT_SECONDS", 60),
		ShutdownTimeout: getEnvAsDuration("APP_SHUTDOWN_TIMEOUT_SECONDS", 5),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func getEnvAsDuration(key string, defaultSeconds int) time.Duration {
	value := getEnv(key, strconv.Itoa(defaultSeconds))
	seconds, err := strconv.Atoi(value)
	if err != nil || seconds <= 0 {
		seconds = defaultSeconds
	}

	return time.Duration(seconds) * time.Second
}

func loadDotEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if key == "" {
			continue
		}

		if _, exists := os.LookupEnv(key); exists {
			continue
		}

		_ = os.Setenv(key, value)
	}
}
