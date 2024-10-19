package main

func main() {
	// type declaration
	type NoKTP string

	var noKTPRizky NoKTP = "1234567890"

	var contoh string = "213123123"
	var contohKtp NoKTP = NoKTP(contoh)

	println(noKTPRizky)
	println(contohKtp)
}
