/**
Struct adalah sebuat template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct biassanya representasi data dalam program aplikasi yang kita buat
Data di struct disimpan dalam field
sederhananya struct adalah kumpulan dari field
*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var ali Customer
	ali.Name = "Ali Akbar"
	ali.Address = "Jakarta"
	ali.Age = 20
	fmt.Println(ali)

	// struct Literals
	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Surabaya", 35}
	fmt.Println(budi)

}
