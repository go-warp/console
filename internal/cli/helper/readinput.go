package helper

import "fmt"

// ReadString reads a string from the user
func ReadString(question string) string {
	fmt.Println(question)

	var input string
	fmt.Scanln(&input)

	return input
}
