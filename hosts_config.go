package simplehttpreverseproxy

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type ConfigHost struct {
	Path    string            `json:"path" yaml:"path"`
	Host    string            `json:"host" yaml:"host"`
	Headers map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
}

type ConfigRoot struct {
	Hosts []ConfigHost `json:"hosts" yaml:"hosts"`
}

func LoadHostsConfig(path string) *ConfigRoot {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var config ConfigRoot
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	for i := range config.Hosts {
		c := &config.Hosts[i]
		c.Path = strings.TrimPrefix(c.Path, "/")
		c.Path = strings.TrimSuffix(c.Path, "/")
	}

	return &config

}
