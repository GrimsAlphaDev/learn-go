/*
sebelumnya untuk membuat pointer dengan menggunakan operator &
Go juga memiliki function new yang bisa digunakan untuk membuat pointer
namum function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.City = "Malang"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}