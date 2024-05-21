package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Bayu"
	names[1] = "Sugiantoro"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [...]int{
		90,
		80,
		95,
		100,
		110,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}