package main

import (
	"fmt"
	"math"
	// "golang/database"	// tanpa blank identifier
	// _ "golang/database" // dengan blank identifier ( _ )
)

func main() {
	// helper.SayHello("Hari")
	// helper.sayGoodBye("Hari")  // error acces modifier
	// fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error acces modifier

	// result := database.GetDatabase()
	// fmt.Println(result)

	// args := os.Args
	// fmt.Println("ini argumen : ", args)

	// fmt.Println(args[0])
	// fmt.Println(args[1])
	// fmt.Println(args[2])

	// hostname, err := os.Hostname() // bila tidak menggunakan err bisa gunakan tanda ( _ )
	// if err == nil {
	// 	fmt.Println(hostname)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	// username := os.Getenv("APP_USERNAME")
	// password := os.Getenv("APP_PASSWORD")

	// fmt.Println(username)
	// fmt.Println(password)

	// host := flag.String("host", "localhost", "Put your database host")
	// username := flag.String("username", "root", "Put your database username")
	// password := flag.String("password", "root", "Put your database password")
	// number := flag.Int("number", 100, "Put your database number")

	// flag.Parse()

	// fmt.Println("Host : ", *host)
	// fmt.Println("User : ", *username)
	// fmt.Println("Password : ", *password)
	// fmt.Println("Number : ", *number)

	// fmt.Println(strings.Contains("Hari Handika", "Harii"))
	// fmt.Println(strings.Contains("Hari Handika", "Hari"))
	// fmt.Println(strings.Split("Hari Handika", " "))
	// fmt.Println(strings.ToLower("Hari Handika"))
	// fmt.Println(strings.ToUpper("Hari Handika"))
	// fmt.Println(strings.ToTitle("Hari Handika"))
	// fmt.Println(strings.Trim("     Hari Handika      ", " "))
	// fmt.Println(strings.ReplaceAll("Hari Hari Hari Hari", "Hari", "Handika"))

	// boolean, err := strconv.ParseBool("true")
	// if err == nil {
	// 	fmt.Println(boolean)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	// number, err := strconv.ParseInt("1000", 10, 32)
	// if err == nil {
	// 	fmt.Println(number)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	// value := strconv.FormatInt(10000, 10)
	// fmt.Println(value)

	// valueInt, err := strconv.Atoi("100")
	// if err == nil {
	// 	fmt.Println(valueInt)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	// valueStr := strconv.Itoa(100)
	// fmt.Println(valueStr)

	fmt.Println("Ini Round ", math.Round(1.7))
	fmt.Println("Ini Round ", math.Round(1.499))
	fmt.Println("Ini Floor", math.Floor(1.8))
	fmt.Println("Ini Ceil", math.Ceil(1.2))
	fmt.Println("Ini Max", math.Max(1, 8))
	fmt.Println("Ini Min", math.Min(1, 8))

}
