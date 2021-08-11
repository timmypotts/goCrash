package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2] string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr)


	// Declare and assign at the same time

	carArr := [2]string{"Jeep", "Bronco"}

	fmt.Println(carArr)

	fruitSlice :=  []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])


}