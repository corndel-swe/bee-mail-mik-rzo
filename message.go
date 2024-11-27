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

func (message *Message) log() {
	fmt.Printf("%+v", message)
}

func (message *Message) markDelivered() {
	message.delivered = true
}

func (message *Message) markRead() {

}
