package main

import (
	"github.com/teris-io/shortid"
)

func main() {}

func GenerateID() string {
	return shortid.MustGenerate()
}
