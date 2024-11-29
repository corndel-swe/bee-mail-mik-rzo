package main

import (
	"fmt"
	"time"
)

type App struct {
	users []*User
}

func (app *App) addUser(user *User) {
	app.users = append(app.users, user)
	fmt.Printf("application state: \n%+v", *app)
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
	message := Message{generateID(), time.Now(), content, sender, recipient, false, false}
	recipient.receiveMessage(&message)
	message.markDelivered()
}
