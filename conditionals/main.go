package main

import "fmt"

func main() {

	// IF ELSE
	x := 14
	y := 5

	if x <= y {
		fmt.Printf("%d is less or equal to %d", x, y)
	} else {
		fmt.Printf("%d is less or equal to %d", y, x)
	}

	color := "Pink"

	if color == "Red" {
		fmt.Println("Color is red")
	} else if color == "Blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is Pink")
	}

	// Switch
	switch color {
	case "Red":
		fmt.Println("Color is red")
	case "Blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is pink")
	}
}
