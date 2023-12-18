package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Hextoi(hex string) int {
	i64, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i64)
}
