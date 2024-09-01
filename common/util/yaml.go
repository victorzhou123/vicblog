package util

import (
	"fmt"
	"os"

	"sigs.k8s.io/yaml"
)

func LoadFromYAML(path string, cfg interface{}) error {
	b, err := os.ReadFile(path) //#nosec G304
	if err != nil {
		return err
	}

	fmt.Printf("string(b): %v\n", string(b))

	return yaml.Unmarshal(b, cfg)
}
