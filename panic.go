// Panic Function adalah function yang bisa kita gunakan untuk menghentikan program
// Panic Function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namum defer function tetap akan diekseskusi

package main

import "fmt"

func endApp() {
	fmt.Println("End APP")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
}

func main() {
	runApp(true)
}
