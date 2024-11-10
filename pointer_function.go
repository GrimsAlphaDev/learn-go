/*
saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
oleh karena itum jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah
hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
namum kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
untuk melakukan ini, kita juga bisa menggunakan pointer di function
untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * do parameternya
*/

package main

import "fmt"

type Address struct {
	Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indochina"
}

func main() {
	var address Address = Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
