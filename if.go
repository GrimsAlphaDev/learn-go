package main

import "fmt"

func main() {
	name := "Andi"

	if name == "Andi" {
		fmt.Println("Hello, Andi")
	} else if name == "Budi" {
		fmt.Println("Hello, Budi")
	} else {
		fmt.Println("Hello, " + name)
	}

	// if short statment
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
