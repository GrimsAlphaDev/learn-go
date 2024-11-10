/*
Secara default di Go-Lang semua variable itu di passing by value bukan by reference
jika kita menegirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi valuenya
|
Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variablenya
*/
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// addres1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// addres2 := addres1 // copy value not using pointer

	// addres2.City = "Bandung"
	// fmt.Println(addres1) // tidak berubah
	// fmt.Println(addres2)  // berubah menjadi bandung

	// using pointer
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

}
