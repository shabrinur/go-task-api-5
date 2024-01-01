package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfigFile() {
	viper.SetConfigName("config-railway")
	viper.SetConfigType("properties")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}
}

func GetConfigValue(configKey string) string {
	return viper.GetViper().GetString(configKey)
}

func GetConfigIntValue(configKey string) int {
	return viper.GetViper().GetInt(configKey)
}
