package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}]
	// person["name"] = "Andi"
	// person["title"] = "Programmer"

	// person := map[string]string{
	// 	"name":  "Andi",
	// 	"title": "Programmer",
	// }

	// fmt.Println(person["name"])
	// fmt.Println(person["title"])
	// fmt.Println(person)

	// fmt.Println(person["salah"])

	// demo making book map
	book := make(map[string]string)
	book["title"] = "Golang Book"
	book["author"] = "Andi"
	book["wrong"] = "Ups"

	delete(book, "wrong")

	fmt.Println(book)
}
