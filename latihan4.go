package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name, Address string
	Age           int
}

func main5() {
	// var hari Customer
	// hari.Name = "Hari Handika Setiawan"
	// hari.Address = "Indonesia"
	// hari.Age = 24

	// fmt.Println(hari.Name)
	// fmt.Println(hari.Address)
	// fmt.Println(hari)

	// hari := Customer{
	// 	Name:    "Hari",
	// 	Address: "Indonesia",
	// 	Age:     24,
	// }
	// fmt.Println(hari)

	// handika := Customer{"Handika", "Indonesia", 24}
	// fmt.Println(handika)
	// hari := Customer{Name: "Hari"}
	// hari.sayHello("Handika")
	// hari.sayTest()
	// person := Person{
	// 	Name: "Hari",
	// }
	// SayHello(person)

	// var hari Person
	// hari.Name = "Hari"
	// SayHello(hari)

	// animal := Animal{Name: "Cat & Bird"}
	// SayHello(animal)

	// var data interface{} = Ups(3)
	// kosong := Ups(2)
	// fmt.Println(data)

	// var person map[string]string = nil
	// fmt.Println(person)

	// person := NewMap("Hari")
	// fmt.Println(person)

	// var person map[string]string = NewMap("Hari")

	// if person == nil {
	// 	fmt.Println("Data Kosong")
	// } else {
	// 	fmt.Println(person)
	// }

	// var contohError error = errors.New("Ups Error")
	// fmt.Println(contohError.Error())

	// hasil, err := Pembagian(100, 0)
	// if err == nil {
	// 	fmt.Println("Hasil", hasil)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	// result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) //panic
	// fmt.Println(resultInt)

	// result := random()
	// switch value := result.(type) {
	// case string:
	// 	fmt.Println("String", value)
	// case int:
	// 	fmt.Println("Int", value)
	// default:
	// 	fmt.Println("Unknown", value)
	// }

	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// // address2 := address1 	// tanpa pointer
	// address2 := &address1 // dengan pointer
	// var address3 *Address = &address1

	// address2.City = "Bandung"

	// // address2 = &Address{"Jakarta", "DKI_Jakarta", "Indonesia"}	//tanpa operator *
	// *address2 = Address{"Jakarta", "DKI_Jakarta", "Indonesia"} // dengan operator * dan hapus tanda &

	// fmt.Println(address1)
	// fmt.Println(address2)
	// fmt.Println(address3)

	// var address4 *Address = new(Address) // data kosong, tidak ada data awal
	// address4.State = "SumateraSelatan"
	// fmt.Println(address4)

	// address := Address{"Subang", "Jawa Barat", ""}
	// ChangeAddressToIndonesia(&address)

	// fmt.Println(address)

	hari := Man{"Hari"}
	hari.Married()
	fmt.Println(hari.Name)

}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	// fmt.Println("Ini Method", man.Name)
}

// func ChangeAddressToIndonesia(address Address) {	//tanpa pointer *
func ChangeAddressToIndonesia(address *Address) { //dengan pointer *
	address.Country = "Indonesia"
}

type Address struct {
	City, State, Country string
}

func random() interface{} {
	return 0.1
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}

}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayTest() {
	fmt.Println("Hello", a.Name)
}
