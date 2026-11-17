package main

import "fmt"

type Authenticator interface {
	Login()
}

type Authorizer interface {
	CheckPermission()
}

type SecureAccess interface {
	Authenticator
	Authorizer
}

type User struct{}

func (u User) Login() {
	fmt.Println("User authenticated")
}

func (u User) CheckPermission() {
	fmt.Println("Permission granted")
}

func main() {
	var access SecureAccess = User{}

	access.Login()
	access.CheckPermission()
}