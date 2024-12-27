package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInternalMessage_String(t *testing.T) {
	messageID := GenerateID()
	timestamp := time.Now()
	mike := User{GenerateID(), "mike", make([]MessageInterface, 0), nil}
	charlie := User{GenerateID(), "charlie", make([]MessageInterface, 0), nil}
	message := Message{messageID, timestamp, "hello", &mike, &charlie, false, false}

	actualStr := message.String()

	expectedStr := fmt.Sprintf(
		"Message ID: %s\nTimestamp: %s\nFrom: %s\nTo: %s\nDelivered: %t\nRead: %t\nContent: %s\n",
		messageID,
		timestamp.Format(time.RFC1123),
		"mike",
		"charlie",
		false,
		false,
		"hello",
	)

	assert.Equal(t, expectedStr, actualStr)
}

func TestExternalMessage_String(t *testing.T) {
	messageID := GenerateID()
	mike := User{GenerateID(), "mike", make([]MessageInterface, 0), nil}
	charlie := User{GenerateID(), "charlie", make([]MessageInterface, 0), nil}
	externalMessage := ExternalMessage{messageID, "hello", &mike, &charlie, false, false}
	adapter := MessageAdapter{&externalMessage}

	actualStr := adapter.String()

	expectedStr := fmt.Sprintf(
		"Message ID: %s\nFrom: %s\nTo: %s\nDelivered: %t\nRead: %t\nContent: %s\n",
		messageID,
		"mike",
		"charlie",
		false,
		false,
		"hello",
	)

	assert.Equal(t, expectedStr, actualStr)
}
