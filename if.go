package main

import "fmt"

func main() {
	x := 10

	// Contoh penggunaan if sederhana
	if x > 5 {
		fmt.Println("x lebih besar dari 5")
	}

	// Penggunaan if dengan else
	if x < 5 {
		fmt.Println("x kurang dari 5")
	} else {
		fmt.Println("x tidak kurang dari 5")
	}

	// Penggunaan if dengan else if
	if x < 5 {
		fmt.Println("x kurang dari 5")
	} else if x == 5 {
		fmt.Println("x sama dengan 5")
	} else {
		fmt.Println("x lebih besar dari 5")
	}

	// Penggunaan if dengan statement pendek
	if y := 5; y > 0 {
		fmt.Println("y adalah bilangan positif")
	}
}
