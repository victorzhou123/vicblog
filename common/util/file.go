package util

import (
	"os"
)

func createFile(path, content string) error {
	f, err := os.Create(path) //#nosec G304
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func CreateFileNamedReady() error {
	return createFile("./ready", "ready")
}
