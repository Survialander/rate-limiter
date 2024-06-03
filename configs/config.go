package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RequestLimit         int    `mapstructure:"REQUEST_LIMIT"`
	TimeoutTimeInSeconds int    `mapstructure:"TIMEOUT_TIME_IN_SECONDS"`
	CustomTokens         string `mapstructure:"CUSTOM_TOKENS"`
	RedisAddr            string `mapstructure:"REDIS_ADDR"`
	RedisPassword        string `mapstructure:"REDIS_PASSWORD"`
	RedisDb              int    `mapstructure:"REDIS_DB"`
}

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading configuration file %v", err.Error())
	}
}

func GetConfig() Config {
	return Config{
		RedisAddr:            viper.GetString("REDIS_ADDR"),
		RedisPassword:        viper.GetString("REDIS_PASSWORD"),
		RedisDb:              viper.GetInt("REDIS_DB"),
		RequestLimit:         viper.GetInt("REQUEST_LIMIT"),
		TimeoutTimeInSeconds: viper.GetInt("TIMEOUT_TIME_IN_SECONDS"),
		CustomTokens:         viper.GetString("CUSTOM_TOKENS"),
	}
}
