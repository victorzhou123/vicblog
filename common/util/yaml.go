package util

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadFromYAML(path string, cfg interface{}) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, cfg)
}
