package main

import "fmt"

func main() {
	// Perulangan dengan satu kondisi
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Perulangan tanpa kondisi
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Perulangan tanpa kondisi dan tanpa post statement
	j := 0
	for {
		fmt.Println(i)
		j++
		if j == 5 {
			break
		}
	}

	// Perulangan menggunakan for range pada slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Perulangan menggunakan for range pada array
	number := [...]int{1, 2, 3, 4, 5}
	for index, value := range number {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Perulangan menggunakan for range pada map
	person := map[string]string{"name": "John", "age": "30"}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
