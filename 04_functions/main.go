package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func sumNumbers(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Adu Asare Bediako"))
	fmt.Println(sumNumbers(3, 7))
}
