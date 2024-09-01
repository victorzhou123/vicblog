package util

import (
	"os"

	"sigs.k8s.io/yaml"
)

func LoadFromYAML(path string, cfg interface{}) error {
	b, err := os.ReadFile(path) //#nosec G304
	if err != nil {
		return err
	}

	// parsing environments variables in yaml
	t := []byte(os.ExpandEnv(string(b)))

	return yaml.Unmarshal(t, cfg)
}
