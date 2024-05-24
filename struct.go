package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var bayu Customer
	fmt.Println(bayu)

	bayu.Name = "Bayu Sugiantoro"
	bayu.Address = "Indonesia"
	bayu.Age = 27
	fmt.Println(bayu)
	fmt.Println(bayu.Name)
	fmt.Println(bayu.Address)
	fmt.Println(bayu.Age)

	Eko := Customer{
		Name:    "Eko Kurniawan",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(Eko)
	fmt.Println(Eko.Name)
	fmt.Println(Eko.Address)
	fmt.Println(Eko.Age)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

}