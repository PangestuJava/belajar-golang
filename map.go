package main

import (
	"fmt"
)

func main() {
	// Membuat map yang memetakan nama buah ke harga per kilogramnya
	fruitPrices := map[string]int{
		"Apel":   10000,
		"Pisang": 5000,
		"Mangga": 15000,
	}

	// Menampilkan harga apel
	fmt.Println("Harga Apel per kilogram:", fruitPrices["Apel"])

	// Mengubah harga mangga
	fruitPrices["Mangga"] = 12000

	// Menampilkan semua harga buah
	fmt.Println("\nHarga Semua Buah:")
	for fruit, price := range fruitPrices {
		fmt.Printf("%s: Rp%d/kg\n", fruit, price)
	}

	book := map[string]string{
		"title":  "Menampilkan buku",
		"author": "Rahmat Hidayat",
		"wrong":  "Ups",
	}

	delete(book, "wrong")

	fmt.Println("\nDetail Buku:")
	for key, value := range book {
		fmt.Printf("%s: %s\n", key, value)
	}
}
