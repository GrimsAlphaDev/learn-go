package main

import "fmt"

func main() {
	names := [...]string{
		"1Elphonso",
		"2Cartier",
		"3Benedict",
		"4Gregori",
		"5Pierson",
		"6Cumberbatch",
	}

	var slice []string = names[4:6]
	fmt.Println(slice)

	slice2 := names[1:5]
	fmt.Println(slice2)

	slice3 := names[:4]
	fmt.Println(slice3)

	slice4 := names[2:]
	fmt.Println(slice4)

	slice5 := names[:]
	fmt.Println(slice5)

	// check slice length
	fmt.Println(len(slice))

	// check slice capacity
	fmt.Println(cap(slice))

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[5:] // Sabtu Minggu
	fmt.Println(daySlice1)
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days) // array days juga berubah

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"

	fmt.Println(daySlice1) // slice daySlice1 tidak berubah
	fmt.Println(daySlice2)
	fmt.Println(days) // array tidak berubah

	// bikin array saat membuat slice
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Muhammad"
	newSlice[1] = "Rizal"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Fahmi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "bido"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// array
	iniarray := [...]int{1, 2, 3, 4, 5}
	// slice
	inislice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniarray)
	fmt.Println(inislice)

}
