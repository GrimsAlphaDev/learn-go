package main

import (
	"fmt"
	"learn-basic-golang/database"
	_ "learn-basic-golang/internal" // menggunakan blank identifier (_) ketika ingin mengimport initnya saja tidak menggunakan method apapun agar tidak dihapus secara otomatis
)

func main() {
	fmt.Println(database.GetDatabase())
}
