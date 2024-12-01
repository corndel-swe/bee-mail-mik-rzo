package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	mike := User{generateID(), "mike", make([]*Message, 0), nil}
	beeMail := App{make([]*User, 0)}

	mike.login(&beeMail)

	assert.NotEmpty(t, mike.app, "User should have an instance of App.")
}
