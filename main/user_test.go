package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	mike := User{GenerateID(), "mike", make([]MessageInterface, 0), nil}
	beeMail := App{make([]*User, 0)}

	mike.login(&beeMail)

	assert.NotEmpty(t, mike.app, "User should have an instance of App.")
}

func TestSendMessage(t *testing.T) {
	assert := assert.New(t)
	mike := User{GenerateID(), "mike", make([]MessageInterface, 0), nil}
	charlie := User{GenerateID(), "charlie", make([]MessageInterface, 0), nil}
	beeMail := App{make([]*User, 0)}
	mike.login(&beeMail)
	charlie.login(&beeMail)

	mike.sendMessage(charlie.id, "hello charlie")
	message := charlie.inbox[0].(*Message)

	assert.NotNil(message.id, "Message should have an id.")
	assert.NotNil(message.timestamp, "Message should have an timestamp.")
	assert.Equal(&mike, message.from, "Message should be from 'mike'.")
	assert.Equal(&charlie, message.to, "Message should be to 'charlie'.")
	assert.Equal("hello charlie", message.content, "Message content should match 'hello charlie'.")
	assert.Equal(true, message.delivered, "Message should be delivered.")
	assert.Equal(false, message.read, "Message should not be read.")
}

func TestReadMessage(t *testing.T) {
	assert := assert.New(t)
	mike := User{GenerateID(), "mike", make([]MessageInterface, 0), nil}
	charlie := User{GenerateID(), "charlie", make([]MessageInterface, 0), nil}
	beeMail := App{make([]*User, 0)}
	mike.login(&beeMail)
	charlie.login(&beeMail)

	mike.sendMessage(charlie.id, "hello charlie")
	message := charlie.inbox[0].(*Message)
	charlie.readMessage(message.id)

	assert.NotNil(message.id, "Message should have an id.")
	assert.NotNil(message.timestamp, "Message should have an timestamp.")
	assert.Equal(&mike, message.from, "Message should be from 'mike'.")
	assert.Equal(&charlie, message.to, "Message should be to 'charlie'.")
	assert.Equal("hello charlie", message.content, "Message content should match 'hello charlie'.")
	assert.Equal(true, message.delivered, "Message should be delivered.")
	assert.Equal(true, message.read, "Message should be read.")
}
