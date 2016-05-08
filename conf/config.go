package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ConfigManager struct
type ConfigManager struct {
	file string
}

// NewConfig constructor
func NewConfig(file string) *ConfigManager {
	return &ConfigManager{file}
}

// Load config from file
func (c *ConfigManager) Load() (*Config, error) {
	data, err := ioutil.ReadFile(c.file)
	if err != nil {
		return nil, err
	}

	var config *Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Config struct
type Config struct {
	Debug bool `yaml:"debug"`
	DB    DB   `yaml:"db"`
	Host  Host `yaml:"host"`
}

// DB struct
type DB struct {
	UserName     string `yaml:"username"`
	UserPassword string `yaml:"userpassword"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Database     string `yaml:"database"`
}

// Host struct
type Host struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}
