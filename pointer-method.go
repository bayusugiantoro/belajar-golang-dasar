package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	bayu := Man{"Bayu"}
	bayu.Married()

	fmt.Println(bayu.Name)
}