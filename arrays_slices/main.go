package main

import "fmt"

func main() {
	// Arrays
	// var fruit_arr [2]string

	// // Assign
	// fruit_arr[0] = "Apple"
	// fruit_arr[1] = "Mango"

	// Declare and Assign

	//fruitArr := [2]string{"Apple", "Mango"}

	//fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Mango", "Orange", "Passion"}

	fmt.Println(fruitSlice)

	fmt.Println(fruitSlice[1:3])
}
