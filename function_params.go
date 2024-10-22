package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilterr(name string, filter Filter) {
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
