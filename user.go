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

func (user *User) sendMessage(recipientId string, content string) {
	if user.app != nil {
		var recipient = user.app.findUser(recipientId)
		if recipient != nil {
			user.app.deliverMessage(user.id, recipientId, content)
		}
	}
}

func (user *User) receiveMessage(message *Message) {
	user.inbox = append(user.inbox, message)
}

func (user *User) readMessage(messageID int) {

}
