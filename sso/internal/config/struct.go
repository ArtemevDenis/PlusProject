package config

import "time"

type Config struct {
	Env            string     `yaml:"env"`
	StoragePath    string     `yaml:"storage_path"`
	GRPC           GRPCConfig `yaml:"grpc"`
	MigrationsPath string
	TokenTTL       time.Duration `yaml:"token_ttl"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
