package main

import "fmt"

func main() {
	x := 3

	switch x {
	case 1:
		fmt.Println("x adalah 1")
	case 2:
		fmt.Println("x adalah 2")
	default:
		fmt.Println("x bukan 1 atau 2")
	}
}
