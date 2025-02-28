package main

import "fmt"

func getMessage() string {
	return "Hello, World"
}

func main() {
	message := getMessage()
	fmt.Println(message)
}
