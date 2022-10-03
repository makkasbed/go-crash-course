package main

import "fmt"

func main() {
	i := 1
	for i <= 20 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d \n", i)
	}
}
