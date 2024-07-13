package main

import (
	"fmt"
	"image/color"
)

func main() {
	x := 15
	y := 10

	// If else
	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is less than %d", y, x)
	}

	// Else if
	color := "red"

	if color == "red" {
		fmt.Println(("color is red"))
	}else if color == "green" {
		fmt.Println(("color is green"))
	} else if color == "yellow" {
		fmt.Println(("color is yellow"))
	}

	// Switch

	switch color {
		case "red":
            fmt.Println(("color is red"))
        case "green":
            fmt.Println(("color is green"))
        case "yellow":
            fmt.Println(("color is yellow"))
        default:
            fmt.Println(("color is unknown"))
	}
}