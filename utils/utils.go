package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadLines(input string) []string {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
