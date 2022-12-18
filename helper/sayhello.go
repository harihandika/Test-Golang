package helper

import "fmt"

var version = "1.0.0"      // tidak bisa diakses dari luar
var Application = "golang" // bisa diakses dari luar

func sayGoodBye(name string) string {
	return "Hello " + name
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}
