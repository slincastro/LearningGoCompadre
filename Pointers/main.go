package main

import (
	"fmt"
)

func main() {
	var firstName *string = new(string)

	//dereferencing ? *
	*firstName = "Jorge"

	//firstName the location in memory
	//*firstName the value

	fmt.Println(*firstName)

	otherFirstName := "Raul"

	pointer := &otherFirstName

	fmt.Println(pointer, *pointer)

	//change the value but the direction in memory is the same
	otherFirstName = "Alicia"

	fmt.Println(pointer, *pointer)

}
