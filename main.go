package main

import (
	"github.com/teris-io/shortid"
)

func main() {}

func generateID() string {
	return shortid.MustGenerate()
}
