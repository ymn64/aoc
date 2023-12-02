package utils

import (
	"log"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
