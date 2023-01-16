package main

import "fmt"

func main() {
	user := []int{4, 4, 1, 3, 5, 6, 6, 6, 7}
	Birthdaycakecandlesn(user)
}

func Birthdaycakecandlesn(candles []int) int {
	temp := candles[0]
	var res int
	for _, e := range candles {
		if e > temp {
			temp = e
		}
	}
	for _, e := range candles {
		if e == temp {
			res++
		}
	}
	fmt.Println(res)
	return res
}
