package utils

import (
	"log"
	"os"
	"strings"
)

func ReadLines(input string) []string {
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
