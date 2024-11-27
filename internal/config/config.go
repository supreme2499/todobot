package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Evn        string     `yaml:"evn" env:"ENV" env-default:"local"`
	Token      string     `yaml:"token" env:"TOKEN" evn-required:"true" `
	StorageCfg StorageCfg `yaml:"storage"`
}

type StorageCfg struct {
	StorageURL     string `yaml:"storage_url" env-required:"true"`
	MigrationsPath string `yaml:"migrations_path" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
