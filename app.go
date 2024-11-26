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

func (app *App) findUser(userId string) (*User) {
	for _, user := range app.users {
		if user.id == userId {
			return user
		}
	}
	return nil
}

func (app *App) deliverMessage(senderId string, recipientId string, content string) {
	sender := app.findUser(senderId)
	recipient := app.findUser(recipientId)
	message := Message{"0", time.Now(), sender, recipient, false, false}
	recipient.receiveMessage(&message)
	message.markDelivered()
}
