package main

import "fmt"

func main() {
	var a = 3
	var b = 10
	var d = 8
	var c = a * (d % b)

	fmt.Println(c)

	var i = 10
	i++
	fmt.Println(i)
}
