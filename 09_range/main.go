package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}

	for i, id := range ids {
		fmt.Printf("%d %d\n", i, id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)
}
