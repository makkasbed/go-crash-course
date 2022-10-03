package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Adu"] = "aluta182004@gmail.com"
	emails["Hannah"] = "afua1995@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Adu"])

	delete(emails, "Adu")
	fmt.Println(emails)

	emails_2 := map[string]string{"Adu": "aluta182004@gmail.com", "Hannah": "afua1995@gmail.com"}
	fmt.Println(emails_2)
}
