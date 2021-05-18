package main

import "fmt"

func main() {
	var name string
	fmt.Println("What is the Name?")
	fmt.Scanln(&name)
	fmt.Println("Hello ", name)
}
