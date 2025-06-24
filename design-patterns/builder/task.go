/*

Implement a builder for:

type User struct {
	Name     string
	Email    string
	Age      int
	IsAdmin  bool
	Metadata map[string]string
}
Support:

Required: Name, Email

Optional: WithAge, AsAdmin, WithMetadata
*/

package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Age      int
	IsAdmin  bool
	Metadata map[string]string
}

type UserBuilder struct {
	user User
}

func NewUserBuilder(name string, email string) *UserBuilder {
	return &UserBuilder{
		user: User{
			Name:  name,
			Email: email,
		},
	}
}

func (ub *UserBuilder) WithAge(age int) *UserBuilder {
	ub.user.Age = age
	return ub
}

func (ub *UserBuilder) AsAdmin() *UserBuilder {
	ub.user.IsAdmin = true
	return ub
}

func (ub *UserBuilder) WithMetadata(metadata map[string]string) *UserBuilder {
	ub.user.Metadata = metadata
	return ub
}

func (ub *UserBuilder) Build() *User {
	return &ub.user
}

func main() {
	user := NewUserBuilder("prashant", "sharma").
		AsAdmin().
		WithAge(27).
		WithMetadata(
			map[string]string{
				"env":  "prod",
				"tier": "gold",
			}).
		Build()
	fmt.Printf("%+v\n", user)
}
