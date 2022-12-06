package main

import "fmt"

func main() {
	// sayHelloWithFilter1("Hari", spamFilter)
	// filter := spamFilter
	// sayHelloWithFilter1("Anjing", filter)

	// blackList := func(name string) bool {
	// 	return name == "anjing"
	// }
	// registerUser("Hari", blackList)
	// registerUser("monkey", func(name string) bool {
	// 	return name == "monkey"
	// })

	// loop := factorialLoop(5)
	// fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	// recursive := factorialRecursive(5)
	// fmt.Println(recursive)
	// fmt.Println(factorialRecursive(10))

	// name := "hari"
	// counter := 0
	// increment := func() {
	// 	name = "handika" // ini case closure, penyelesaiannya dengan nama
	// 	// variabel yang berbeda atau deklara ulang seperti
	// 	name := "setiawan"
	// 	fmt.Println("Increment")
	// 	fmt.Println(name )
	// 	counter++
	// }
	// increment()
	// increment()
	// fmt.Println(counter)
	// fmt.Println(name)

	// defer
	// runApplication(0)

	//panic
	// runApp(true)

	//recover(menangkap data panic)
	runApp(false)
}

// kode recover yang benar
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("Run Application")
}

// kode recover yang salah
func endApp1() {
	fmt.Println("End App")
}

func runApp1(error bool) {
	defer endApp1()
	if error {
		panic("Error Application")
	}
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // biasakan menggunakan defer itu diatas
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive((value - 1))
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

//atau dengan declaration
func sayHelloWithFilter1(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
