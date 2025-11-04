package setup

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

var (
	once sync.Once
	cfg  *Config
)

type Config struct {
	DevicesFile string `yaml:"devices_file"`
}

func GetConfig() *Config {
	once.Do(func() {
		var err error
		cfg, err = loadYAMLConfig("config.yaml")
		if err != nil {
			// Handle error appropriately (e.g., log or panic)
			cfg = &Config{} // fallback to empty config
		}
	})
	return cfg
}

func loadYAMLConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}
