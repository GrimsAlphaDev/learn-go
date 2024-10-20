package main

import "fmt"

func main() {
	name := "Andisa"

	switch name {
	case "Andi":
		println("Hello, Andi")
	case "Budi":
		println("Hello, Budi")
	default:
		println("Hello, " + name)
	}

	// if short statment
	switch length := len(name); {
	case length > 5:
		println("Terlalu panjang")
	default:
		println("Nama sudah benar")
	}

	// switch without condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
