package main

import (
	"fmt"
)

const (
	_ = iota + 2021
	a
	b
	c
	d
)

func main() {
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}