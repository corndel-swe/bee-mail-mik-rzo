package main

type User struct {
	id       string
	username string
	inbox    []MessageInterface
	app      *App
}

func (user *User) login(a *App) {
	user.app = a
	a.addUser(user)
}

func (user *User) sendMessage(recipientID string, content string) {
	if user.app != nil {
		var recipient = user.app.findUser(recipientID)
		if recipient != nil {
			user.app.deliverMessage(user.id, recipientID, content)
		}
	}
}

func (user *User) receiveMessage(message MessageInterface) {
	user.inbox = append(user.inbox, message)
}

func (user *User) readMessage(messageID string) {
	if user.app != nil {
		var message *Message
		for _, value := range user.inbox {
			if value.(*Message).id == messageID {
				message = value.(*Message)
			}
		}
		if message != nil {
			message.markRead()
			message.log()
		}
	}
}
