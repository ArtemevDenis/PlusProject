package config

type Config struct {
	Env     string `yaml:"env" default:"local"`
	Gateway struct {
		Port    string `yaml:"port"`
		Host    string `yaml:"host"`
		Version string `yaml:"version"`
	} `yaml:"gateway"`
	Quiz struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
	} `yaml:"quiz"`
}
