package main

import "fmt"

func main() {
	Shipping(4, []int{1, 4, 2, 4})
}

func Shipping(N int, h []int) []int {
	temp := h[0]
	var res []int
	for _, e := range h {
		if e <= temp {
			temp = e
		}
	}
	for i, e := range h {
		if e == temp {
			res = append(res, i)
		}
	}

	// for j := 0; j < len(res); j++ {
	for i, _ := range h {
		if h[i] > temp {
			if h[i-1] == temp && temp+(i-(i-1)) < h[i] {
				h[i] = temp + (i - (i - 1))
			} else if h[i-2] == temp && temp+(i-(i-2)) < h[i] {
				h[i] = temp + (i - (i - 2))
			} else if h[i-3] == temp && temp+(i-(i-3)) < h[i] {
				h[i] = temp + (i - (i - 3))
			} else if temp+(i-(i-N)) > h[i] {
				h[i] = h[i]
			} else {
				h[i] = temp + (i - (i - N))
			}
		}
	}
	// }
	fmt.Println("ini h", h)
	return res
}
