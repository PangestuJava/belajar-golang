package main

import "fmt"

func main() {
	//Break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke", i)
	}
}
