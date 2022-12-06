package main

import "fmt"

func main3() {

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

	// person := make(map[string]string)
	// person["name"] = "Eko"
	// person["title"] = "Programmer"

	// for key, value := range person {
	// 	fmt.Println(key, "=", value)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Perulangan ke ", i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}
	// 	fmt.Println("Perulangan ke ", i)
	// }

	// function in golang
	// for i := 0; i < 10; i++ {
	// 	sayHello()
	// 	fmt.Println(i)
	// }
	// sayHelloTo("Hari", 10)
	// atau
	// firstName := "Hari"
	// nilai := 10
	// sayHelloTo(firstName, nilai)
	// result := getHelloTo(" Hari")
	// fmt.Println(result)
	// fmt.Println(getHelloTo(""))

	// firstName, lastName, nilai, bool := getFullName()
	// fmt.Println(firstName, lastName, nilai, bool)

	// firstName, _, _, _ := getFullName()
	// fmt.Println(firstName)
	// a, b, c, _ := getCompleteName() // nama variable tidak harus sama pada kasus name return value
	// fmt.Println(a, b, c)

	// total := sumAll(10, 11, 12, 13, 14, 15)
	// fmt.Println(total)

	// total := sumAll(10, 11, 12, 13, 14, 15)
	// fmt.Println(total)

	// slice := []int{11, 12, 13, 14, 15}
	// total1 := sumAll(slice...)
	// fmt.Println(total1)

	sayGoodBye := getGoodBye
	result := sayGoodBye("Hari")
	fmt.Println(result)
	fmt.Println(getGoodBye("Hari"))

}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

// untuk tipe data sama
// func getCompleteName() (firstName, middleName, lastName string) {}
func getCompleteName() (firstName string, middleName string, lastName string, nilai int) {
	firstName = "Hari"
	middleName = "Handika"
	lastName = "Setiawan"
	nilai = 10

	return
}

func sayHelloTo(firstName string, nilai int) {
	fmt.Println("Hello", firstName, nilai)
}

func sayHello() {
	fmt.Println("Hello")
}

func getHello(name string) string {
	return "Hello" + name
}

func getHelloTo(name string) string {
	if name == "" {
		return "Hello New User"
	} else {
		return "Hello" + name
	}
}

func getFullName() (string, string, int, bool) {
	return "Hari", "Handika", 10, true
}
