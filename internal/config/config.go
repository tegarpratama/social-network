package config

import (
	"github.com/spf13/viper"
)

const (
	CONFIG_PATH = "."
	CONFIG_NAME = ".env"
	CONFIG_TYPE = "env"
)

var config ConfigTypes

func SetupConfig() (*ConfigTypes, error) {
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return &config, err
	}

	viper.Unmarshal(&config)

	return &config, nil
}
