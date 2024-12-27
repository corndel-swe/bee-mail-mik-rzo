package main

type MessageInterface interface {
	String() string
	log()
	markDelivered()
	markRead()
}
