package main

import "fmt"

func main() {
	accountStatus := "pending"

	switch accountStatus {
	case "active":
		fmt.Println("Allow normal access")
	case "pending":
		fmt.Println("Verify email address")
	case "suspended":
		fmt.Println("Block account access")
	case "closed":
		fmt.Println("Account is closed")
	default:
		fmt.Println("Unknown account status")
	}
}