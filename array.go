package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Rizal"
	names[2] = "Fahmi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// names[3] = "joko" // error karena panjang array hanya 3
	var values = [3]int{
		90,
		95,
		80,
		// 100 // error karena panjang array hanya 3
	}

	fmt.Println(values)

	var values2 = [...]int {
		90,
		95,
		30,
		80,
		100,
	} // [...] panjang array akan menyesuaikan jumlah elemen
}
