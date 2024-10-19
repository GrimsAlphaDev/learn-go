package main

import "fmt"

func main() {
	// const firstName string = "Killa"
	// const lastName = "Anderson"

	// firstName = "yolo"
	// lastName = "yoli"
	// cant reassign const

	const (
		firstName = "Killa"
		lastName  = "Anderson"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
