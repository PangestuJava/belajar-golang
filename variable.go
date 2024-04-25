package main

import "fmt"

func main() {
	// deklarasi variable dapat dilakukan menggunakan :=
	name := "Nadia Aulia"
	fmt.Println(name)

	name = "Nadia Rahmah"
	fmt.Println(name)

	// variable multiple
	var(
		firstName = "Rahmat"
		lastName = "Hidayat"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}