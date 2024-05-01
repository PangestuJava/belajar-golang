package main

import "fmt"

func main()  {
	type noKtp string

	var ktpRahmat noKtp = "3173625144229998"

	var contoh string = "2122122211220009"
	var contohKtp noKtp = noKtp(contoh)

	fmt.Println(ktpRahmat)
	fmt.Println(contohKtp)
	
}