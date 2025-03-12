package util

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	env := viper.GetString("ENVIRONMENT")
	if env == "" || env == "development" {
		viper.AddConfigPath(path) // Path to look for the file
		viper.SetConfigFile(".env")
		if err = viper.ReadInConfig(); err != nil {
			fmt.Println("No .env file found, relying on environment variables")
		} else {
			fmt.Println("Loaded .env file for local development")
		}
	}

	viper.BindEnv("ENVIRONMENT")
	viper.BindEnv("HTTP_SERVER_ADDRESS")
	viper.BindEnv("DB_SOURCE")
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Error unmarshalling config")
	}
	return
}
