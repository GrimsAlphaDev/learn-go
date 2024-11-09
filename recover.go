// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
//  Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

package main

import "fmt"

func endApp() {
	fmt.Println("End APP")
	message := recover()
	fmt.Println("Pesan Errornya adalah", message)
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
