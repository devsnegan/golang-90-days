package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 7 {
			break
		}

		fmt.Println(i)
	}
}