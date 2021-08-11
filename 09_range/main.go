package main

import (
	"fmt"
)

func main() {
	ids := []int{1, 2, 3, 4, 5, 6}

	// Loop thru ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharona@sharon.com"}

	for k, v := range emails {
		fmt.Println(k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}