package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	GRPC        GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func New() *Config {
	cfg := Config{}
	if err := cleanenv.ReadConfig("./config/local.yaml", &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}
	return &cfg
}
