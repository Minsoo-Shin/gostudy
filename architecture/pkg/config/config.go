package config

import (
	"fmt"
	"ggurugi/pkg/cli"
	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		Mongo `yaml:"mongo"`
		Jwt   `yaml:"jwt"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"`
		Version string `env-required:"true" yaml:"version"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"`
	}

	// Mongo -.
	Mongo struct {
		Host     string `env-required:"true" yaml:"host"`
		User     string `env-required:"true" yaml:"user"`
		Password string `env-required:"true" yaml:"password"`
		DbName   string `env-required:"true" yaml:"dbName"`
		Options  struct {
			MinConnections int `yaml:"minConnections"`
			MaxConnections int `yaml:"maxConnections"`
		} `yaml:"options"`
	}

	// Jwt -.
	Jwt struct {
		Secret string `env-required:"true" yaml:"secret"`
	}
)

func New() (Config, error) {
	cfg := &Config{}

	viper.SetConfigName(cli.Config)  // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {
		return *cfg, fmt.Errorf("fatal error config file: %w", err)
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return *cfg, fmt.Errorf("fatal error config file: %w", err)
	}

	return *cfg, nil
}
