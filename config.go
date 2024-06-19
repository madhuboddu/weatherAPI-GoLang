package main

import "github.com/spf13/viper"

type Config struct {
	BaseURL string `mapstructure:"BaseURL"`
	APIKey  string `mapstructure:"APIKey"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&config)
	return
}
