package config

import "github.com/spf13/viper"

type config struct {
	Server struct {
		Port int
	}
}

// C - variable containing server configuration
var C config

// ReadConfig - reads server configuration
func ReadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}
