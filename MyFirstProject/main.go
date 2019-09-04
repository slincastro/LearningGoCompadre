package main

import (
	"fmt"
)

func main() {
	var message string
	message = "Hi string variable"

	secondMessage := ", and implicit inicialization variable"

	fmt.Println(message + secondMessage)

	complexNumber := complex(4, 5)

	fmt.Println(complexNumber)
}
