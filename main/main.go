package main

import (
	"github.com/teris-io/shortid"
)

func GenerateID() string {
	return shortid.MustGenerate()
}
