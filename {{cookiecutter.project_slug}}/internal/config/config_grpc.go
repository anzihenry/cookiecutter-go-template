package config

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	once   sync.Once
	config *Config
)

type Config struct {
	LogLevel string `mapstructure:"log_level"`
}

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetDefault("log_level", "info")
		viper.SetConfigName("config") // Looks for config.yaml
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	})
	return config
}
