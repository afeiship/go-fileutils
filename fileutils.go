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

func WriteFile(filename string, data []byte) {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
}

func ReadContents(filename string) string {
	res := ReadFile(filename)
	return string(res)
}

func WriteContents(filename string, data string) {
	WriteFile(filename, []byte(data))
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
