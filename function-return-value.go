package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Eko")
	fmt.Println(result)

	fmt.Println(getHello("Kurniawan"))
	fmt.Println(getHello("Khannedy"))
}