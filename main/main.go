package main

import (
	"github.com/teris-io/shortid"
)

func main() {
	beeMail := App{make([]*User, 0)}

	usernames := []string{"mike", "charlie", "meron", "damian", "james", "pablo", "hala", "anna", "john"}
	var users []*User = nil
	for _, username := range usernames {
		user := User{GenerateID(), username, make([]MessageInterface, 0), nil}
		user.login(&beeMail)
		users = append(users, &user)
	}

	mike := users[0]
	charlie := users[1]

	mike.sendMessage(charlie.id, "hello charlie")

	message := charlie.inbox[0].(*Message)
	charlie.readMessage(message.id)
}

func GenerateID() string {
	return shortid.MustGenerate()
}
