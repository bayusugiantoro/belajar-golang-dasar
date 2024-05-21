package main

import "fmt"

func main() {
	var a = 100
	var b = 10
	var d = 5
	var e = 20
	var c = a + b - e * d
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}