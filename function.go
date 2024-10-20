package main

import "fmt"

func sayHello(firstname string, lastname string) {
	fmt.Println("Hello,", firstname, lastname)
}

func main() {
	sayHi()
	sayHello("Rizky", "Fauzan")
	var result string = getHello("Puerto Rico")
	fmt.Println(result)
}

func sayHi() {
	fmt.Println("Hi")
}

// function return value
func getHello(name string) string {
	return "Hello " + name
}
