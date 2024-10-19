package main

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) // overflow max int16 = 32767 if more than that will be -32768 and so on

	println(nilai32)
	println(nilai64)
	println(nilai16)

	var name = "Rizky Khapidsyah"
	var e uint8 = name[0]
	var eString string = string(e)

	println(name)
	println(e)
	println(eString)
}
