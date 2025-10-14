package main

import "fmt"

/**
 * Задача 2
 *
 * Дан массив целых чисел. Необходимо развернуть его.
 *
 * Сделать это надо за линейное время без дополнительных
 * аллокаций.
 */

func ReverseArray(arr []int) {
	l := 0
	r := len(arr) - 1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func Task2() {
	var size int
	fmt.Print("Enter count of numbers: ")
	fmt.Scanln(&size)

	fmt.Print("Enter numbers: ")
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Original: ")
	fmt.Println(arr)

	ReverseArray(arr)

	fmt.Print("Reversed: ")
	fmt.Println(arr)
}
