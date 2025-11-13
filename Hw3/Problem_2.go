package Hw3

import "fmt"

/**
* Очень легкая задача
*
* Как быстро можно сделать N копий
* документа, используя два ксерокса,
* каждый копирует со своей скоростью
* (x и y минут)?
 */

func good(N int64, x int64, y int64, t int64) bool {
	return (t/x)+((t-x)/y) >= N
}

func AnsBinSearch(N int64, x int64, y int64) int64 {
	l := int64(0)
	r := N * max(x, y)
	for r-l > 1 {
		m := (l + r) / 2
		if good(N, min(x, y), max(x, y), m) {
			r = m
		} else {
			l = m
		}
	}

	return r
}

func Problem2() {
	fmt.Print("Enter N, x, y: ")
	var N, x, y int64
	_, _ = fmt.Scanln(&N, &x, &y)

	fmt.Println("Minimal time is: ", AnsBinSearch(N, x, y))
}
