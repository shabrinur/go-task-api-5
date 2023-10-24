package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("properties")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}
}

func getConfigValue(configKey string) string {
	return viper.GetViper().GetString(configKey)
}
