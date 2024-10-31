package fileutils

import (
	"fmt"
	"log"
	"os"
)

func SayHi() {
	fmt.Println("Hi from fileutils")
}

func ReadFile(filename string) []byte {
	res, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return res
}
