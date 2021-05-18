package main

import "fmt"

func main() {
	a := 5

	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use  value from pointer
	fmt.Println(*b)

	//Modify pointer

	*b = 13
	fmt.Println(a)
}
