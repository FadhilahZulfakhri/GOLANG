package main

import "fmt"

//NOTE : PADA GOLANG INI SPASI SANGATLAH BERPENGARUH
func main() {
	// fmt.Println(`
	// hello
	// world`)

	// MENGGUNAKAN LEN
	fmt.Println(len("Hai Dhila"))

	//MENGUNAKAN GET CHARACTER (hasil yang didapat berdasarkan ASCII Code)
	// urutan ke byte
	fmt.Println("DHILA"[4])
	fmt.Println(string("DHILA"[4]))

	// string
	fmt.Println(string(65))
}
