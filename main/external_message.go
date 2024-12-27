package main

import "fmt"

type ExternalMessage struct {
	id        string
	body      string
	from      *User
	to        *User
	delivered bool
	read      bool
}

func (message *ExternalMessage) String() string {
	return fmt.Sprintf(
		"Message ID: %s\nFrom: %s\nTo: %s\nDelivered: %t\nRead: %t\nContent: %s\n",
		message.id,
		message.from.username,
		message.to.username,
		message.delivered,
		message.read,
		message.body,
	)
}

func (message *ExternalMessage) log() {
	fmt.Printf("%v", message)
}

func (message *ExternalMessage) markDelivered() {
	message.delivered = true
}

func (message *ExternalMessage) toggleRead() {
	message.read = !message.read
}
