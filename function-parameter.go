package main

import "fmt"

func sayHelloTo(firstName string,middleName string, lastName string) {
	fmt.Println("Hello", firstName, middleName, lastName)
}

func main() {
	sayHelloTo("Eko", "Kurniawan", "Khannedy")
	sayHelloTo("Bayu", "Sugiantoro", "Rizki")
}