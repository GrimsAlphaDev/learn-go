/**
Interface adalah tipe data abstract. dia tidak memiliki Implementasi langsung
sebuah interface berisikan definisi - definisi method
Biasanya interface digunakan sebagai kontrak
*/

package main

import "fmt"

type HasName interface {
	getName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.getName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Gunaji"}
	SayHello(person)
	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
