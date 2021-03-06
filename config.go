package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type manifest struct {
	Output string
	Inputs []string
}

type Config struct {
	Manifests []manifest
}

func Load() *Config {
	viper.SetConfigName("loom")

	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("./etc")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
	}

	return &config
}
