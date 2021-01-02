package main

import (
	"fmt"
)

// Package scope variable declaration
var y = "Hello World!"

func main() {
	// Auto type variable creation
	x := 10

	// Print simple line on prompt
	fmt.Println(x)
}

// Creating a new function on package
func printLine(x string) {
	// Print receaved variable on function paramaters
	fmt.Println(x)

	// Print package scope variable
	fmt.Println(y)
}