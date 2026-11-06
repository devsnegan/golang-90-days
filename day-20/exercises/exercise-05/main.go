
package main

import "fmt"

func greetUsers(message string, names ...string) {
	for _, name := range names {
		fmt.Println(message, name)
	}
}

func main() {
	greetUsers("Welcome", "Snegan", "Alex", "John")
}
