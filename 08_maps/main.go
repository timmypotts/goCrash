package main

import (
	"fmt"
)

func main() {
	// // Define map
	// emails := make(map[string]string)

	// // Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@sharon.com"
	// emails["Mike"] = "mike@gmail.com"


	// Declare map and add key/value pairs

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharona@sharon.com"}
	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}