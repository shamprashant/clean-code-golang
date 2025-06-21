/*
Create a Go program that:

Takes a list of users (with name and email).

Has a UserManager that stores users.

Has a separate Notifier that sends emails (just print messages).

➡️ Make sure you do not mix responsibilities.
*/

package main

import "fmt"

type User struct {
	name  string
	email string
}

type UserManagerService struct {
	users []User
}

type EmailNotifierService struct{}

func (ums *UserManagerService) addUser(u User) {
	ums.users = append(ums.users, u)
}

func (ens EmailNotifierService) sendMail(u User) {
	fmt.Println("Email sent to: ", u.email)
}

func main() {
	fmt.Println("Welcome to SRP!!.")
	user1 := User{name: "Prashant Sharma", email: "prashant.sharma@gmail.com"}
	user2 := User{name: "Arun Sharma", email: "arun.sharma@gmail.com"}
	ums := UserManagerService{}
	ums.addUser(user1)
	ums.addUser(user2)

	ens := EmailNotifierService{}
	ens.sendMail(user1)
	ens.sendMail(user2)

	fmt.Println(ums.users)
}
