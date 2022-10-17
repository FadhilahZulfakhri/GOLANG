package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// Printf tidak menambah new line
	fmt.Printf("Hello, everyone! I'm %s and i'm %d years old ", "dhila", 20)
	// Cara menggunakan new line pada Printf yaitu menggunakan karakter <\n> hal ini juga berlaku pada print
	fmt.Printf("Hello, everyone! I'm %s and i'm %d years old \n", "dhila", 20)
	fmt.Printf("Nice to meet you")
	// bisa juga kita tidak menggunakan karakter \n dengan menggunakan Println, tetapi hal tersebut membuat codingan menjadi panjang. Hal itu menggunakan <,> atau <+>
	fmt.Println("Hello, everyone!" + " I'm" + " dhila" + " and" + " i'm" + " 20" + "years old")
	fmt.Println("Hello, everyone!", "I am", "dhila", "and i'm", "20", "years old")
}
