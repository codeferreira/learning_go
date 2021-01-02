package main

import (
	"fmt"
)

type newInt int

var x newInt
var y int

func main() {
	fmt.Printf("%v %T\n", x, x)

	x = 42
	
	fmt.Println(x)

	y = int(x)

	fmt.Printf("%v %T\n", y, y)
}