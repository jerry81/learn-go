package main

import "fmt"

var c, python, java bool = true, false, true // scoped to package, shows variable with initializer

func main() {
	var i int            // scoped to function
	short_variable := 42 // := replaces var and does implicit type
	i = 1
	fmt.Println(i+short_variable, c, python, java)
}
