package main

import (
	"github.com/teris-io/shortid"
)

func main() {
	beeMail := App{make([]*User, 0)}

	mike := User{generateID(), "mike", make([]*Message, 0), nil}

	mike.login(&beeMail)
}

func generateID() string {
	return shortid.MustGenerate()
}
