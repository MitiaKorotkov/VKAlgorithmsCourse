package Hw3

import (
	"fmt"
)

func GreaterThanSqrt(n uint64, target uint64) bool {
	return n >= target/n+1
}

func BinSearchSqrt(n uint64) (uint64, uint64) {
	var l uint64 = 0
	r := n/2 + 1
	for r > l+1 {
		m := (l + r) / 2
		if GreaterThanSqrt(m, n) {
			r = m
		} else {
			l = m
		}
	}

	return l, r
}

func Problem1() {

	// TODO(Dima): move it to test function
	//for n := 1; n < 100000000; n++ {
	//	l, r := BinSearchSqrt(uint64(n))
	//	if float64(l) > math.Sqrt(float64(n)) || float64(r) <= math.Sqrt(float64(n)) {
	//		fmt.Println(n, " : ", l, " ", r)
	//	}
	//}

	var n int64
	fmt.Print("Enter number: ")
	_, _ = fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Sqrt doesn`t exist")
	} else {
		l, r := BinSearchSqrt(uint64(n))
		fmt.Println(l, " <= sqrt(", n, ") < ", r)
	}
}
