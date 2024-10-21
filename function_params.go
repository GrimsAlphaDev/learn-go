package main

import "fmt"

func sayHelloWithFilterr(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "******"
	} else if name == "babi" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilterr("anjing", spamFilter)
}
