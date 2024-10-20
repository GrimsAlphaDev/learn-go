package main

import "fmt"

func sayHello(firstname string, lastname string) {
	fmt.Println("Hello,", firstname, lastname)
}

func main() {
	sayHi()
	sayHello("Rizky", "Fauzan")
	var result string = getHello("Puerto Rico")
	var firstName, lastName string = getFullName()
	fmt.Println(result)
	fmt.Println(firstName, lastName)
	// ignore second return value
	var firstName2, _ = getFullName()
	fmt.Println(firstName2)
}

func sayHi() {
	fmt.Println("Hi")
}

// function return value
func getHello(name string) string {
	return "Hello " + name
}

// function with multiple return value
func getFullName() (string, string) {
	return "Rizky", "Fauzan"
}
