package main

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	timestamp time.Time
	content   string
	from      *User
	to        *User
	delivered bool
	read      bool
}

func (message *Message) String() string {
	return fmt.Sprintf(
		"Message ID: %s\nTimestamp: %s\nFrom: %s\nTo: %s\nDelivered: %t\nRead: %t\nContent: %s\n",
		message.id,
		message.timestamp.Format(time.RFC1123),
		message.from.username,
		message.to.username,
		message.delivered,
		message.read,
		message.content,
	)
}

func (message *Message) log() {
	fmt.Printf("%s", message)
}

func (message *Message) markDelivered() {
	message.delivered = true
}

func (message *Message) markRead() {
	message.read = true
}
