package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(a, b)

	//use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&b)

	*b = 10
	fmt.Println(a)
}
