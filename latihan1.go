package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string
	//{ (ini sama dengan = person := map[string]string{)
	// person := map[string]string{
	// 	"name":    "Hari",
	// 	"address": "Subang",
	// }

	// person["title"] = "Programmer"
	// fmt.Println(person)
	// fmt.Println(person["name"], person["address"])

	// book := make(map[string]string)
	// book["title"] = "Belajar-Golang"
	// book["author"] = "Hari"
	// book["ups"] = "Salah"
	// fmt.Println(book, len(book))

	// delete(book, "ups")

	// fmt.Println(book, len(book))

	// name := "Sania hahah"

	// if name == "Hari" {
	// 	fmt.Println("ini", name)
	// } else if name == "Sania" {
	// 	fmt.Print("ini ", name)
	// } else {
	// 	fmt.Println("ini", name)
	// }

	// if length := len(name); length > 5 {
	// 	fmt.Println("terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	// name := "Hari"

	// switch name {
	// case "Hari":
	// 	fmt.Println("Hello", name)
	// 	fmt.Println("Goodbye", name)
	// case "Sania":
	// 	fmt.Println("Hello", name)
	// 	fmt.Println("Goodbye", name)
	// default:
	// 	fmt.Println("Hello " + name + ", Boleh Kenalan?")
	// 	fmt.Println("Goodbye", name)

	// }

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// default:
	// 	fmt.Println("Nama Sudah benar")
	// }

	// name := "Hari H"
	// length := len(name)
	// switch {
	// case length > 10:
	// 	fmt.Println("nama terlalu panjang")
	// case length > 5:
	// 	fmt.Println("nama lumayan panjang")
	// default:
	// 	fmt.Println("nama sudah benar")
	// }

	// min, max := findMinAndMax(a)
	// fmt.Println("Min: ", min)
	// fmt.Println("Max: ", max)
	var h = []int{11, -4, 7, 8, -10}
	var min = Shipping(10, h)
	fmt.Println(min)
}

func Shipping(N int, h []int) int {
	var min int
	var res int
	min = h[0]
	for _, value := range h {
		if value <= min {
			min = value
			res++
		}
	}
	fmt.Println(min)
	return min
}
