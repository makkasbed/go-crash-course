package main

import "fmt"

func main() {
	// Arrays
	var fruits [2]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"

	fmt.Println(fruits)

	newFruits := [2]string{"Apple", "Orange"}
	fmt.Println(newFruits)
	fmt.Println(len(newFruits))
}
