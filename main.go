package main

import (
	
)

func main() {
	beeMail := App{make([]*User, 0)}

	mike := User{"0", "mike", make([]*Message, 0), nil}

	mike.login(&beeMail)
}
