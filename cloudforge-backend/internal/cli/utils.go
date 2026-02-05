package cli

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"io/ioutil"

	"cloudforge-backend/internal/providers"
	"cloudforge-backend/internal/providers/gcp"
)

type Config struct {
	Project  string `yaml:"project"`
	Provider string `yaml:"provider"`
	Region   string `yaml:"region"`
}

func readConfig() (*Config, error) {
	data, err := ioutil.ReadFile("cloudforge.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func getProvider(name string) (providers.Provider, error) {
	switch name {
	case "gcp":
		return &gcp.GCP{}, nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", name)
	}
}
