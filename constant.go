package main

import "fmt"

func main() {
	const (
		firstName string = "Bayu"
		lastName         = "Sugiantoro"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Joko"
}