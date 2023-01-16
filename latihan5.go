package main

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
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

	// fmt.Println("Ini Round ", math.Round(1.7))
	// fmt.Println("Ini Round ", math.Round(1.499))
	// fmt.Println("Ini Floor", math.Floor(1.8))
	// fmt.Println("Ini Ceil", math.Ceil(1.2))
	// fmt.Println("Ini Max", math.Max(1, 8))
	// fmt.Println("Ini Min", math.Min(1, 8))

	// data := list.New()
	// data.PushBack("Hari")
	// data.PushBack("Handika")
	// data.PushBack("Setiawan")
	// data.PushFront("Ini")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)
	// fmt.Println("Ini next", data.Front().Next().Next())
	// fmt.Println("Ini Back", data.Back().Prev().Prev())

	// // iterate dari depan ke belakang
	// for e := data.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	// // iterate dari belakang ke depan
	// for e := data.Back(); e != nil; e = e.Prev() {
	// 	fmt.Println(e.Value)
	// }

	// var data *ring.Ring = ring.New(5)
	// data := ring.New(5) // tanpa var dengan :=
	// for i := 0; i < data.Len(); i++ {
	// 	data.Value = "Value " + strconv.FormatInt(int64(i), 10)
	// 	data = data.Next()
	// }

	// data.Do(func(value interface{}) {
	// 	fmt.Println(value)
	// })

	users := []User{
		{"Handika", 21},
		{"Setiawan", 22},
		{"Hari", 20},
		{"Ganteng", 23},
	}

	// fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)

	// now := time.Now()
	// now.Day()

	// fmt.Println(now)
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Day())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// fmt.Println(now.Nanosecond())

	// fmt.Println(now.Local())

	// utc := time.Date(2009, time.December, 10, 22, 0, 0, 0, time.UTC)
	// fmt.Println(utc)

	// parse, _ := time.Parse(time.RFC3339, "2008-02-03T15:04:05Z07:00")
	// fmt.Println(parse)

	// layout := "2006-01-02"
	// parse1, _ := time.Parse(layout, "2020-12-23")
	// fmt.Println(parse1)

	// sample := Sample{"Hari"}
	// sampleType := reflect.TypeOf(sample) // atau bisa menggunakan var
	// // var sampleType reflect.Type = reflect.TypeOf(sample)
	// structField := sampleType.Field(0)
	// required := structField.Tag.Get("required")

	// fmt.Println(sampleType.NumField())
	// fmt.Println(structField.Name)
	// fmt.Println(required)
	// fmt.Println(sampleType.Field(0).Tag.Get("max"))
	// fmt.Println(sampleType.Field(0).Tag.Get("min"))

	// sample.Name = ""
	// fmt.Println(IsValid(sample))

	// contoh := ContohSample{"Hari", "Hari"}
	// fmt.Println(IsValid(contoh))

	var regex *regexp.Regexp = regexp.MustCompile("h([a-z])i")
	fmt.Println(regex.MatchString("hai"))
	fmt.Println(regex.MatchString("hri"))
	fmt.Println(regex.MatchString("hSi"))

	search := regex.FindAllString("hai hri hsi hfi hki nhi", 20)
	fmt.Println(search)

}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

type ContohSample struct {
	Name        string `required:"true"`
	Description string `required:"true" `
}

type Sample struct {
	Name string `required:"true" max:"10"`
}

// type Sample struct {
// 	Name string
// }

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}
