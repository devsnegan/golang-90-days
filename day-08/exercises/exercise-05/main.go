package main

import "fmt"

func main() {
	statusCode := 403

	switch statusCode {
	case 200:
		fmt.Println("OK")
	case 201:
		fmt.Println("Created")
	case 400:
		fmt.Println("Bad Request")
	case 401:
		fmt.Println("Unauthorized")
	case 403:
		fmt.Println("Forbidden")
	case 404:
		fmt.Println("Not Found")
	default:
		fmt.Println("Unknown status code")
	}
}