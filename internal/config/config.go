package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

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

func Load() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, errors.New("config path is empty")
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("config file not found: " + configPath)
	}

	var cfg Config

	if err := readConfig(configPath, &cfg); err != nil {
		return nil, errors.New("error in reading config: " + err.Error())
	}
	return &cfg, nil
}

func readConfig(configPath string, cfg *Config) error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}
