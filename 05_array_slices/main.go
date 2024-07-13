package main

import "fmt"

func main() {

	//Declaring an Array
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	//Declaring and assigning an array
	myFruitArr := []string{"Apple", "Banana", "Guava"}

	fmt.Println(len(myFruitArr))

	fmt.Println(fruitArr)
}