package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Adu"] = "aluta182004@gmail.com"
	emails["Hannah"] = "afua1995@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Adu"])
}
