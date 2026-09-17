package filesystem

import (
	"path/filepath"
)

func IsAbsolute(path string) bool {
	return filepath.IsAbs(path)
}

func GetFileName(path string) string {
	return filepath.Base(path)
}

func GetPath(path string) string {
	return filepath.Dir(path)
}
