package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetCallStackInfo(skip int) string {
	if dir, err := os.Getwd(); err == nil {
		dir = filepath.ToSlash(dir) + "/"

		if _, file, line, ok := runtime.Caller(skip); ok {
			return fmt.Sprintf("%s:%d", strings.TrimPrefix(file, dir), line)
		}
	}

	return ""
}
