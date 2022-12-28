package env

import "os"

type Environment struct {
	Name  string
	Value string
}

type EnvironmentVariables struct {
	BINANCE_API_KEY    string
	BINANCE_API_SECRET string
	BINANCE_API_USER   string
	PORT               string // default port
	BASEURL            string // base URL
	INTERVALS          string // Seconds between requests
}

func EnvironmentProvider() []*Environment {
	return []*Environment{
		MustEnv("BINANCE_API_KEY"),
		MustEnv("BINANCE_API_SECRET"),
		MustEnv("BINANCE_API_USER"),
		MustEnv("PORT"),
		MustEnv("BASEURL"),
		GetEnv("PORT", "8080"),
		GetEnv("INTERVALS", "3"),
	}
}

func MustEnv(envName string) *Environment {
	value := os.Getenv(envName)
	if value == "" {
		panic("environment variable " + envName + " is not set")
	}
	return &Environment{Name: envName, Value: value}
}

func GetEnv(envName string, defaultValue string) *Environment {
	value := os.Getenv(envName)
	if value == "" {
		return &Environment{Name: envName, Value: defaultValue}
	}
	return &Environment{Name: envName, Value: value}
}
