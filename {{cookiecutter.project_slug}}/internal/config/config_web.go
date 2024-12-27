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
	ServerPort string `mapstructure:"server_port"`
	LogLevel   string `mapstructure:"log_level"`
}

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetDefault("server_port", "8080")
		viper.SetDefault("log_level", "info")

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs")
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		config = &Config{}
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
	return config
}
