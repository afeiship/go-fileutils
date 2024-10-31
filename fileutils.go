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

func RmRf(path string) error {
	return os.RemoveAll(path)
}
