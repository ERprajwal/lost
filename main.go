package main

import (
	"fmt"
)

func main() {
	// Using Println to print a line with a newline at the end
	fmt.Println("Hello, World!")

	// Using Printf to print formatted strings
	name := "Prajwal"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Using Sprintf to format a string and store it in a variable
	formattedString := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(formattedString)
}
