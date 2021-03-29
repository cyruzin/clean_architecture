package util

import (
	"path/filepath"
	"runtime"
)

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..")
	return basepath
}

// PathBuilder builds a path.
func PathBuilder(path string) string {
	return rootDir() + path
}
