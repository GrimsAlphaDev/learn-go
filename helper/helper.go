package helper // nama package sesuai folder
import "fmt"

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("test") // dapat diakses dialam package only jika huruf didepan kecil
	fmt.Println(version)
}
