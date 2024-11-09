// dever function = function yang bisa kita jadwalkan untuk diekseskusi setelah fungsi yang memanggilnya selesai dieksekusi
// defer function akan selalu dieksekusi walaupun terjadi error

package main

import "fmt"

func loggin() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication() {
	defer loggin()
	fmt.Println("Run Application")
	// jika ada error
	// loggin() // tidak akan dipanggil
}

func main() {
	runApplication()
}
