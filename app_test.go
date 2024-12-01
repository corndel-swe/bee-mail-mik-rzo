package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddUser(t *testing.T) {
	beeMail := App{make([]*User, 0)}
	mike := User{generateID(), "mike", make([]*Message, 0), nil}

	beeMail.addUser(&mike)

	assert.NotEmpty(t, beeMail.users, "Users should not be empty.")
	assert.Equal(t, "mike", beeMail.users[0].username, "User added should have the username 'mike'.")
}
