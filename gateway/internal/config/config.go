package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
)

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadConfig() Config {
	var cfg Config
	f, err := os.Open("config/config.yaml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
	ReadEnv(&cfg)

	return cfg
}

func ReadEnv(cfg *Config) {
	var err = envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
