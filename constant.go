package main

import "fmt"

func main() {
	// const tidak dapat berubah berbeda dengan variable
	// const juga tidak masalah jika tidak dipakai berbeda dengan variable
	const (
		firstName = "Rahmat"
		lastName  = "Hidayat"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
