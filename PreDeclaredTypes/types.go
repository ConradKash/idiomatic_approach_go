package main

import "fmt"

func vare() {
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("%v %v %v %v\n", a, b, c, d)
}

func main() {
	vare()
}