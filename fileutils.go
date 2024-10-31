package fileutils

import (
	"log"
	"os"
)

func ReadFile(filename string) []byte {
	res, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return res
}

func GetContents(filename string) string {
	res := ReadFile(filename)
	return string(res)
}

func RmRf(path string) error {
	return os.RemoveAll(path)
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func MkdirP(path string) error {
	return os.MkdirAll(path, 0755)
}
