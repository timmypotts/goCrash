package main

import "fmt"

func main() {
	// var name string = "Tim"

	name, email := "Tim", "tim@tim.com"
	size := 1.3
	var age int32 = 25
	const isCool = true

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", email)
}