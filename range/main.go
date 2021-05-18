package main

import "fmt"

func main() {
	ids := []int{23, 343, 5, 54, 6, 56, 3, 2, 32, 44, 53, 22, 3, 4632, 4, 653, 4}

	// Loop through
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Add all ids
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println(" Sum: ", sum)

	// Range with maps
	emails := map[string]string{"j0n": "j0n@gmail.com", "Aki": "Aki@Naukwa.com", "OG": "Respect@OGs.com"}

	for k, v := range emails {
		fmt.Printf("Name: %s Email: %s\n", k, v)
	}
}
