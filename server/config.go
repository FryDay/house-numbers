package main

import (
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	URL string `json:"url"`
}

// NewConfig ...
func NewConfig(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	config := &Config{}
	if err := v.Unmarshal(config); err != nil {
		panic(err)
	}
	return config
}
