package main

import "fmt"

func main() {
	sayHelloTo("Rahmat", "Hidayat")
}

func sayHelloTo(firstName, lastName string) {
	fmt.Printf("Hello %s %s\n", firstName, lastName)
}
