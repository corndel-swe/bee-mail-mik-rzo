package main

import (
	"time"
)

type App struct {
	users []*User
}

func (app *App) addUser(user *User) {
	app.users = append(app.users, user)
}

func (app *App) findUser(userID string) *User {
	for _, user := range app.users {
		if user.id == userID {
			return user
		}
	}
	return nil
}

func (app *App) deliverMessage(senderID string, recipientID string, content string) {
	sender := app.findUser(senderID)
	recipient := app.findUser(recipientID)
	message := Message{GenerateID(), time.Now(), content, sender, recipient, false, false}
	recipient.receiveMessage(&message)
	message.markDelivered()
}
