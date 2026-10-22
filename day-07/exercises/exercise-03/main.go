package main

import "fmt"

func main() {
	isLoggedIn := true
	hasPermission := true
	isBlocked := false

	canAccess := isLoggedIn && hasPermission && !isBlocked

	fmt.Println("Can access:", canAccess)
	fmt.Println("Can login:", isLoggedIn && !isBlocked)
	fmt.Println("Needs attention:", isBlocked || !hasPermission)
}