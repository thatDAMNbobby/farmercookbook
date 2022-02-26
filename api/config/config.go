package config

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var (
	Name    = "farmer-cookbook api"
	Version = "dev"
	Hash    = "dev"
)

type ESConfig struct {
	Address          string
	Sniff            bool
	IndexRetryMinS   int
	IndexRetryMaxS   int
	SearchRetryMinMS int
	SearchRetryMaxMS int
}

type DatabaseConfig struct {
	Name   string
	Region string
}

type ServerConfig struct {
	Port int
}

type Config struct {
	Name          string
	Version       string
	Hash          string
	LogLevel      log.Level
	Elasticsearch *ESConfig
	Database      *DatabaseConfig
	Server        *ServerConfig
}

func MustGetEnvString(envName string) string {
	return GetEnvString(envName, true)
}

func GetEnvString(envName string, required bool) string {
	value, present := os.LookupEnv(envName)

	if required && (!present || strings.TrimSpace(value) == "") {
		log.WithField("key", envName).Panic("Missing required environment variable")
	}

	return value
}

func MustGetEnvInt(envName string) int {
	return GetEnvInt(envName, true)
}

func GetEnvInt(envName string, required bool) int {
	value := GetEnvString(envName, false)

	entry := log.WithField("key", envName)
	if required && value == "" {
		entry.Panic("Missing required environment variable")
	}
	if !required && value == "" {
		return 0
	}

	parsed, err := strconv.Atoi(value)

	if err != nil {
		entry.Panic(err)
	}

	return parsed
}

func MustGetEnvBool(envName string) bool {
	value := MustGetEnvString(envName)

	parsed, err := strconv.ParseBool(value)

	if err != nil {
		log.WithField("key", envName).Panic(err)
	}

	return parsed
}

func Loader() *Config {
	logLevel, err := log.ParseLevel(MustGetEnvString("LOG_LEVEL"))
	if err != nil {
		log.Panic(err)
	}

	esConfig := &ESConfig{
		Address:          MustGetEnvString("SEARCH_ES_ADDRESS"),
		Sniff:            MustGetEnvBool("SEARCH_ES_SNIFF"),
		IndexRetryMinS:   MustGetEnvInt("SEARCH_ES_INDEX_RETRY_MIN_S"),
		IndexRetryMaxS:   MustGetEnvInt("SEARCH_ES_INDEX_RETRY_MAX_S"),
		SearchRetryMinMS: MustGetEnvInt("SEARCH_ES_SEARCH_RETRY_MIN_MS"),
		SearchRetryMaxMS: MustGetEnvInt("SEARCH_ES_SEARCH_RETRY_MAX_MS"),
	}

	dbConfig := &DatabaseConfig{
		Name:   MustGetEnvString("DATABASE_NAME"),
		Region: MustGetEnvString("DATABASE_REGION"),
	}

	ServerConfig := &ServerConfig{
		Port: MustGetEnvInt("SERVER_PORT"),
	}

	return &Config{
		Name:          Name,
		Version:       Version,
		Hash:          Hash,
		LogLevel:      logLevel,
		Elasticsearch: esConfig,
		Database:      dbConfig,
		Server:        ServerConfig,
	}
}
