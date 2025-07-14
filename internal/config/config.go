package config

import (
	"cmp"
	"os"
)

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

func Load() *Config {
	config := &Config{}
	config.DatabaseName = cmp.Or(os.Getenv(DBNAME), DefaultDbName)
	config.Port = cmp.Or(os.Getenv(PORT), DefaultPort)
	config.Dbpath = cmp.Or(os.Getenv(DBPATH), DefaultDBPath)

	return config
}
