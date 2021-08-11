package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "red"
	
	// else if
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}


	//  Swtich
	switch color {
	case "red":
		fmt.Println("Color is red")
	
	case "blue": 
		fmt.Println("Color is blue")

	default:
		fmt.Println("Color is NOT blue or red")
	}

}