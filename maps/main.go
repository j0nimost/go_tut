package main

import "fmt"

func main() {

	// Define map
	//emails := make(map[string]string)

	// //Assign
	// emails["J0n"] = "j0n@gmail.com"
	// emails["Aki"] = "Aki@Naukwa.com"
	// emails["OG"] = "Respect@OGs.com"

	// Declare and add

	emails := map[string]string{"j0n": "j0n@gmail.com", "Aki": "Aki@Naukwa.com", "OG": "Respect@OGs.com"}

	fmt.Println(emails)
	//fmt.Println(emails["J0n"])
	//DELETE
	delete(emails, "Aki")
	fmt.Println(emails)
}
