package main

import (
	"fmt"
)

func main() {
	x := (10 == 10)
	y := (10 != 10)
	z := (10 >= 10)
	a := (10 > 10)
	b := (10 <= 10)
	c := (10 < 10)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", x, y, z, a, b, c)
}