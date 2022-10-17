package main

import "fmt"

func main() {
	// SCAN line
	// CONTOH PERTAMA
	var name string
	fmt.Println("Masukkan nama")
	fmt.Scan(&name)
	fmt.Println("Hallo", name)

	// CONTOH KE-DUA
	var name, alamat string
	fmt.Println("Masukkan nama")
	fmt.Scan(&name)
	fmt.Println("Masukkan alamat")
	fmt.Scan(&alamat)
	fmt.Println("Hello", name)
	fmt.Println("Alamat anda", alamat)

	// SCAN
	var name, address string
	fmt.Print("Enter your name and address : ")
	fmt.Scan(&name, &address) // set input to variable 'name' and 'address'

	fmt.Println("Hello", name)
	fmt.Println("Your address", address)

	// SCANF
	var name string
	var HP int
	fmt.Println("Masukkan nama dan no WA :")
	fmt.Scanf("%s %d", &name, &HP)
	fmt.Println("Hello", name)
	fmt.Println("WA", HP)

	var nama string
	var HP int
	println("Masukkan nama :")
	fmt.Scanln("&nama")
	println("Masukkan no WA :")
	fmt.Scanf("%d", HP)
	fmt.Println("Hello", nama)
	fmt.Println("WA", HP)
}
