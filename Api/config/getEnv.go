package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ApiPort       string
	UsersPort     string
	AnalyticsPort string

	DefaultOffset string
	DefaultLimit  string

	TokenKey string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ApiPort = cast.ToString(getOrReturnDefaultValue("API_PORT", ":8080"))
	config.UsersPort = cast.ToString(getOrReturnDefaultValue("USERS_PORT", ":8081"))
	config.AnalyticsPort = cast.ToString(getOrReturnDefaultValue("ANALYTICS_PORT", "8082"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))
	config.TokenKey = cast.ToString(getOrReturnDefaultValue("TOKEN_KEY", "my_secret_key"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
