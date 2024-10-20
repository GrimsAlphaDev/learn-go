package main

import "fmt"

func main() {
	// var counter int8 = 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan Ke ", counter)
	// 	counter++
	// }

	// for with statement
	for counter := 1; counter <= 100000; counter++ {
		fmt.Println("Perulangan Ke ", counter)
	}

	names := []string{"Rizky", "Fauzan", "Rahman"}
	// manual
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for range
	// for index, name := range names {
	// 	fmt.Println("Index ", index, "=", name)
	// }

	// no index
	for _, name := range names {
		fmt.Println(name)
	}

}
