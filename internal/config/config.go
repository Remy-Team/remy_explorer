package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"remy_explorer/pkg/logging"
	"sync"
)

// Config is the application configuration structure that is read from the config file.
type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

// StorageConfig is the database configuration structure that is read from the config file.
type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var instance *Config
var once sync.Once

// GetConfig reads the application configuration from the default path.
func GetConfig() *Config {
	return GetConfigWithPath("config.yml")
}

// GetConfigWithPath reads the application configuration from a given path.
func GetConfigWithPath(path string) *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig(path, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
