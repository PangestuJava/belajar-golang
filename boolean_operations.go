package main

import "fmt"

func main() {
	nilaiAlhir := 90
	nilaiAbsensi := 80

	lulusNilaiAkhir := nilaiAlhir > 80
	lulusAbsensi := nilaiAbsensi > 80

	lulus := lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}
