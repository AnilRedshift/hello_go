package main

import "fmt"

func main() {
	a, b := 20, 30
	fmt.Println("a is ", a, " and b is ", b)

	b, c := 40, 50
	fmt.Println("b is ", b, " and c is ", c)

	var foo = "hello"
	// can't redeclare a variable like this?
	// var foo = "world"
	foo, bar := "hello", "world"
	// This is weird. So if you have an extra variable in a list, you can redeclare
	// but only in this case. why?
	fmt.Println(foo + " " + bar)
}
