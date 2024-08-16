package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTPHost  string
	HTTPPort  string
	RedisHost string
	RedisPort int64
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPHost = cast.ToString(getOrReturnDefaultValue("HTTPHost", "analytics-service"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8082"))
	config.RedisHost = cast.ToString(getOrReturnDefaultValue("RedisHost", "localhost"))
	config.RedisPort = cast.ToInt64(getOrReturnDefaultValue("RedisPort", 6379))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
