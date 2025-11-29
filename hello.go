// Package main is the entry point for our Go program.
// Every executable Go program must have a main package.
package main

// Import statements bring in code from other packages.
// The "fmt" package provides formatted I/O functions like Println.
import (
	"fmt"
	"os"
)

// getMessage returns a greeting string.
// Functions are declared with the "func" keyword.
// The string after the parentheses is the return type.
func getMessage(name string) string {
	// If no name is provided, use a default greeting
	if name == "" {
		return "Hello, World"
	}
	// fmt.Sprintf formats a string without printing it
	return fmt.Sprintf("Hello, %s", name)
}

// main is the entry point of the program.
// It's automatically called when you run the program.
func main() {
	// os.Args is a slice (like an array) containing command-line arguments
	// os.Args[0] is the program name, os.Args[1] is the first argument
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// Call getMessage and store the result in a variable
	message := getMessage(name)

	// Print the message to standard output
	fmt.Println(message)
}
