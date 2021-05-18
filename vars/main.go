package main

import "fmt"

func main() {
	// var
	var age int32 = 25
	// short hand declaration only for functions
	name := "John"
	size := 3.14
	// const
	const isCool = true

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", size)
}
