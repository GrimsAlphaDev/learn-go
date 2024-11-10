package main

import (
	"fmt"
	"learn-basic-golang/helper"
)

func main() {
	gerard := helper.SayHello("gerard")
	fmt.Println(gerard)

	fmt.Println(helper.Application) // huruf depan besar dapat diakses diluar package
	fmt.Println(helper.version)     // huruf depan kecil tidak bisa diakses diluar package

}
