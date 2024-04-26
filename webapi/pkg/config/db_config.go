package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbType string

const (
	MySQL DbType = "mysql"
	PgSQL DbType = "postgres"
	Mongo DbType = "mongodb"
)

type DbConfig struct {
	Type       DbType
	Host       string
	Port       int
	DBUser     string
	DBPassword string
	DBName     string
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetDatabaseConfig() (*DbConfig, error) {
	dbType := DbType(getEnv("DB_TYPE", "pgsql"))
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	cfg := &DbConfig{
		Type: dbType,
		Host: dbHost,
	}
	switch dbType {
	case MySQL:
		cfg = &DbConfig{
			Type:       MySQL,
			Port:       getEnv("MYSQL_DB_PORT", 3306),
			Host:       dbHost,
			DBUser:     getEnv("MYSQL_DB_USER", "root"),
			DBPassword: getEnv("MYSQL_DB_PASSWORD", "root"),
			DBName:     getEnv("MYSQL_DB_NAME", "tasks"),
		}
	case PgSQL:
		cfg = &DbConfig{
			Type:       PgSQL,
			Port:       getEnv("PG_DB_PORT", 5432),
			Host:       dbHost,
			DBUser:     getEnv("PG_DB_USER", "postgres"),
			DBPassword: getEnv("PG_DB_PASSWORD", "postgres"),
			DBName:     getEnv("PG_DB_NAME", "tasks"),
		}
	case Mongo:
		cfg = &DbConfig{
			Type:       Mongo,
			Port:       getEnv("MONGO_DB_PORT", 27017),
			Host:       dbHost,
			DBUser:     getEnv("MONGO_DB_USER", "root"),
			DBPassword: getEnv("MONGO_DB_PASSWORD", "root"),
			DBName:     getEnv("MONGO_DB_NAME", "tasks"),
		}
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	return cfg, nil
}

func getEnv[T any](key string, fallback T) T {
	if value, ok := os.LookupEnv(key); ok {
		casted, ok := any(value).(T)
		if !ok {
			return fallback
		}
		return casted
	}
	return fallback
}
