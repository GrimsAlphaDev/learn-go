/**
type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
Jika salah menggunakan type assertion, maka bisa terjadi error panic
sebaiknya kita menggunakan switch expression untuk melakukan type assertion
*/

package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt int = result.(int) // panic
	// fmt.Println(resultInt)
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type", value)
	}

}
