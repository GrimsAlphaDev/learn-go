package main

import (
	"fmt"
	"learn-basic-golang/database"
	_ "learn-basic-golang/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
