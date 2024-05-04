package main

import "fmt"

func main() {
	//  ini array
	// monthsArray := [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "oct", "Nov", "Dec"}
	// INI SLICE
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "oct", "Nov", "Dec"}

	slices1 := months[4:6]
	fmt.Println(slices1)

	slices2 := months[:3]
	fmt.Println(slices2)

	slices3 := months[3:]
	fmt.Println(slices3)

	slices4 := months[:]
	fmt.Println(slices4)

	monthsSlice1 := months[10:]
	fmt.Println(monthsSlice1)

	monthsSlice1[0] = "Nov now"
	monthsSlice1[1] = "Dec now"
	fmt.Println(monthsSlice1)
	fmt.Println(months)

	monthsSlice2 := append(monthsSlice1, "Libur Baru")
	monthsSlice2[0] = "Ups"
	fmt.Println(monthsSlice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "rahmat"
	newSlice[1] = "rahmat"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Hidyat", "anjay")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

}
