package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(Config string) {
	viper.SetConfigName("edged.conf")
	viper.AddConfigPath("etc")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read configuration file")
	}
}
