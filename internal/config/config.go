package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type (
	Config struct {
		Env           string `yaml:"env" env-required:"true"`
		LogPath       string `yaml:"log_path" env-required:"true"`
		StorageConfig `yaml:"storage_config" env-required:"true"`
		HttpConfig    `yaml:"http_config" env-required:"true"`
	}
	StorageConfig struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     int    `yaml:"port" env-default:"5432"`
		Username string `yaml:"username" env-default:"postgres"`
		Password string `yaml:"password" env-default:"postgres"`
		DBName   string `yaml:"dbname" env-default:"postgres"`
		SslMode  string `yaml:"sslmode" env-default:"prefer"`
	}
	HttpConfig struct {
		Host         string        `yaml:"host" env-required:"true"`
		Port         int           `yaml:"port" env-default:"8080"`
		ReadTimeout  time.Duration `yaml:"read_timeout" env-default:"5s"`
		WriteTimeout time.Duration `yaml:"write_timeout" env-default:"60s"`
	}
)

func Init() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		return nil, errors.New("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("CONFIG_PATH does not exist")
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, fmt.Errorf("error reading config: %s", err)
	}
	return &config, nil
}
