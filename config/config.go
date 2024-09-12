package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	CPU_THRESHOLD     float64
	MEMORY_THRESHOLD  float64
	DISK_IO_THRESHOLD float64
	LOG_FILE          string
}

func NewConfig() *Config {
	//if err := godotenv.Load(".env"); err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	config := &Config{
		CPU_THRESHOLD:     cast.ToFloat64(coalesce("CPU_THRESHOLD", "3")),
		MEMORY_THRESHOLD:  cast.ToFloat64(coalesce("MEMORY_THRESHOLD", "40")),
		DISK_IO_THRESHOLD: cast.ToFloat64(coalesce("DISK_IO_THRESHOLD", "40")),
		LOG_FILE:          cast.ToString(coalesce("LOG_FILE", "dodi.log")),
	}
	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
