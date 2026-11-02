
package main

import (
	"errors"
	"fmt"
)

func getUser(id int) (string, error) {
	users := map[int]string{
		101: "Alex",
		102: "John",
		103: "David",
	}

	name, exists := users[id]

	if !exists {
		return "", errors.New("user not found")
	}

	return name, nil
}

func main() {
	name, err := getUser(102)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Username:", name)
}
