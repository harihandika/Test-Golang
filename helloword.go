package main

import "fmt"

func main1() {
	// const firsName string = "Hari"
	// const lastName = "Handika S"
	//atau
	// const (
	// 	firsName string = "Hari"
	// 	lastName        = "Handika S"
	// )
	// fmt.Println(firsName, lastName)

	// var nilai32 int32 = 129
	// var nilai64 int64 = int64(nilai32)
	// var nilai8 int8 = int8(nilai32)

	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai8)

	// name := "Eko"
	// var e byte = name[1]
	// var konversi string = string(e)

	// fmt.Println(konversi)
	// fmt.Println(name)
	// type hari string
	// type married bool
	// type nilai int

	// var nama hari = "hari handika setiawan"
	// var kebenaran married = true
	// var angka nilai = 100
	// fmt.Println(nama)
	// fmt.Println(hari(nama))
	// fmt.Println(kebenaran)
	// fmt.Println(angka)

	// var result = 10 + 10
	// var a int = 10
	// var b int = 10
	// var c = a * b
	// var d = a % 2
	// var i = 10
	// i += 10
	// a = a + 1
	// fmt.Println(c, result, d, i, a)

	// var name1 = "Hari"
	// var name2 = "Hari1"

	// var result bool = name1 < name2
	// fmt.Println(result)
	// fmt.Println(name1 != name2)

	// var nilaiAkhir = 100
	// var absensi = 80

	// var lulusAkhir bool = nilaiAkhir >= 100
	// var lulusAbsensi bool = absensi > 80
	// fmt.Println(lulusAkhir)
	// fmt.Println(lulusAbsensi)

	// var lulus = lulusAkhir || lulusAbsensi
	// fmt.Println(lulus)

	// var names [4]string
	// names[0] = "John"
	// names[1] = "Michael"
	// names[2] = "Steven"
	// names[3] = "George"

	// fmt.Println(names[0], " x ", names[1], " x ", names[2], " x ", names[3])
	// var values = [3]int{
	// 	90,
	// 	80,
	// 	70,
	// }
	// fmt.Println(values)
	// fmt.Println(len(values), values[1])
	// values[0], values[1] = 50, 60
	// fmt.Println(values)

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	var slice2 = months[:9]
	var slice3 = months[4:]
	fmt.Println(slice1, slice2, slice3)
	fmt.Println(len(slice1), cap(slice2), cap(slice3))

	// months[5] = "Edit"
	// fmt.Println(slice1, slice2)

	// slice1[0] = "Mei-Edit"
	// fmt.Println(months)

	var slice4 = append(slice2, "tambahan")
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Edit"
	newSlice[1] = "rambahan"
	// var slice4 = append(slice3, "tambahan")
	fmt.Println(slice4)

	slice4[1] = "Coba"
	fmt.Println(slice4)
	fmt.Println(slice2)
	fmt.Println(months)
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniSlice := []int{1, 2, 3, 4, 5}
	iniArray := [...]int{1, 2, 3, 4, 5}
	fmt.Println(iniSlice, iniArray)

}
