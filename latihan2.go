package main

import "fmt"

func main() {

	// //for loop
	// counter := 0
	// counter1 := 0

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	fmt.Println("INI BEDA", counter1)
	// 	counter1 = counter1 + 2
	// 	counter++
	// }

	// // bisa menggunakan sintaks ini
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke", counter)
	// }

	// slice := []string{"Hari", "Handika", "Setiawan", "Backend", "Developer"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	// for i, value := range slice {
	// 	fmt.Println("Index", i, "=", value)
	// 	fmt.Println(value)
	// }

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
