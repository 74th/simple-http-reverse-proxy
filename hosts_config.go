package simplehttpreverseproxy

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type HostsConfigHost struct {
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty" `
	Path   string `json:"path" yaml:"path"`
	Host   string `json:"host" yaml:"host"`
}

type HostsConfigRoot struct {
	Hosts []HostsConfigHost `json:"hosts" yaml:"hosts"`
}

func LoadHostsConfig(path string) *HostsConfigRoot {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var config HostsConfigRoot
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	for _, c := range config.Hosts {
		c.Path = strings.Trim(c.Path, "/")
	}

	return &config

}
