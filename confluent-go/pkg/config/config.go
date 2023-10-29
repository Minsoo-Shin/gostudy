package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Env    string `yaml:"env"`
	Listen string `yaml:"listen"`
	Kafka  struct {
		Topics           []string `yaml:"topics"`
		BootstrapServers string   `yaml:"bootstrapServers"`
		GroupID          string   `yaml:"groupID"`
		AutoOffsetReset  string   `yaml:"autoOffsetReset"`
	} `yaml:"kafka"`
}

func DefaultConfig() *Config {
	// default or for test
	return &Config{
		Env:    "dev",
		Listen: ":3000",
		Kafka: struct {
			Topics           []string `yaml:"topics"`
			BootstrapServers string   `yaml:"bootstrapServers"`
			GroupID          string   `yaml:"groupID"`
			AutoOffsetReset  string   `yaml:"autoOffsetReset"`
		}{
			Topics:           []string{"eventCreated"},
			BootstrapServers: "localhost:29092",
			GroupID:          "test-group",
			AutoOffsetReset:  "earliest",
		},
	}
}

func New() *Config {
	config := &Config{}

	filePath := "conf.yaml"
	file, err := os.Open(filePath)
	if err != nil {
		config = DefaultConfig()
	} else {
		confByte := make([]byte, 0)
		_, err = file.Read(confByte)
		if err != nil {
			log.Fatalf("read file err: %v\n", err)
		}

		if err = yaml.Unmarshal(confByte, &config); err != nil {
			fmt.Println("unmarshal err:", err)
		}
	}
	defer file.Close()

	return config
}
