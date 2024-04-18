package config

import "os"

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBAddress  string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() *Config {
	return &Config{
		Port:       getEnv("PORT", "3306"),
		DBUser:     getEnv("DB_USER", "mysql"),
		DBPassword: getEnv("DB_PASSWORD", "mysql"),
		DBName:     getEnv("DB_NAME", "webapi"),
		DBAddress:  getEnv("DB_ADDRESS", "localhost"),
		JWTSecret:  getEnv("JWT_SECRET", "random_secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
