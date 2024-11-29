package main

type User struct {
	id       string
	username string
	inbox    []*Message
	app      *App
}

func (user *User) login(a *App) {
	app := a
	app.addUser(user)
}

func (user *User) sendMessage(recipientID string, content string) {
	if user.app != nil {
		var recipient = user.app.findUser(recipientID)
		if recipient != nil {
			user.app.deliverMessage(user.id, recipientID, content)
		}
	}
}

func (user *User) receiveMessage(message *Message) {
	user.inbox = append(user.inbox, message)
}

func (user *User) readMessage(messageID string) {
	if user.app != nil {
		var message *Message
		for _, msg := range user.inbox {
			if msg.id == messageID {
				message = msg
			}
		}
		if message != nil {
			message.markRead()
			message.log()
		}
	}
}
