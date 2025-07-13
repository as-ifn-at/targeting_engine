package config

import "os"

type Config struct {
	Port         string
	DatabaseName string
	Dbpath       string
}

const (
	DBNAME = "databasename"
	PORT   = "port"
	DBPATH = "dbpath"
)

var (
	DefaultDbName = "hashmap"
	DefaultPort   = "8080"
	DefaultDBPath = "data"
)

func getEnvOrDefault(env, defaultVal string) string {
	if envVal := os.Getenv(env); envVal != "" {
		return envVal
	}
	return defaultVal
}

func Load() *Config {
	config := &Config{}
	config.DatabaseName = getEnvOrDefault(DBNAME, DefaultDbName)
	config.Port = getEnvOrDefault(PORT, DefaultPort)
	config.Dbpath = getEnvOrDefault(DBPATH, DefaultDBPath)

	return config
}
