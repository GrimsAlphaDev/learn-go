package main

import "fmt"

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Rizky"
	return firstName, middleName, lastName
}
