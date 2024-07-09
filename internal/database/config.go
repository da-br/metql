package database

import "github.com/spf13/viper"

type Config struct {
	DatabaseFile  string
	ServerAddress string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseFile:  viper.GetString("masterbase"),
		ServerAddress: viper.GetString("address"),
	}
}
