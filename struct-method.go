package main

import "fmt"

type Customer struct {
	Name string
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name, "Ada yang bisa di bantu?")
}

func main() {
	var bayu Customer
	
	bayu.Name = "Bayu Sugiantoro"

	bayu.sayHello("Eko")
}