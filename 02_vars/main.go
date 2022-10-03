package main

import "fmt"

func main() {

	var name string = "Adu Asare"
	var age int = 37
	fmt.Println(name, age)
	fmt.Printf("%T\n", name)

	name, email := "Adu Asare", "aluta182004@gmail.com"
	fmt.Println(name, email)
}
