package main

import "fmt"

func main() {
	type NoKTP string

	var ktpBayu NoKTP = "1111111"

	var contoh string = "2222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpBayu)
	fmt.Println(contohKtp)
}