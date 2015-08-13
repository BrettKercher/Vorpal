package config

import (
	"fmt"

	"code.google.com/p/gcfg"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

func (dc *DatabaseConfig) Url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dc.Username,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.Name)
}

type Config struct {
	Database DatabaseConfig
	Server   struct {
		ColorLogs bool
		Port      int
	}
}

func LoadConfig(filename string) (*Config, error) {
	var config Config
	err := gcfg.ReadFileInto(&config, filename)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
