package config

import (
	"os"
	"path/filepath"
)

var (
	CAAFile        = configFile("ca.pem")
	ServerCertFile = configFile("server.pem")
	ServerKeyFile  = configFile("server-kety.pem")
)

func configFile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".proglog", filename)
}
