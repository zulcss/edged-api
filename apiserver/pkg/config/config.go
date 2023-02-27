package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(ConfigFile string) {
	viper.SetConfigFile(ConfigFile)

	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
