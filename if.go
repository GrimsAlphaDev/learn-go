package main

import "fmt"

func main() {
	name := "Andiga"

	if name == "Andi" {
		fmt.Println("Hello, Andi")
	} else if name == "Budi" {
		fmt.Println("Hello, Budi")
	} else {
		fmt.Println("Hello, " + name)
	}
}
