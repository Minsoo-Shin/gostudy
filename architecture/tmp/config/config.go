package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		Mysql `yaml:"Mysql"`
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
	Mysql struct {
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
	cfg := Config{}

	// 환경 변수에 값이 없는 경우, dev.yaml 로 설정
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	viper.SetConfigName(env)         // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {
		return cfg, fmt.Errorf("fatal error config file: %w", err)
	}
	// viper 는 읽어온 설정파일의 정보를 가지고있으니, 전역변수에 언마샬링해
	// 애플리케이션의 원하는곳에서 사용하도록 합니다.
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("fatal error config file: %w", err)
	}

	return cfg, nil
}
