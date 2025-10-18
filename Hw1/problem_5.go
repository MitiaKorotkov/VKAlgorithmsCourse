package Hw1

import "fmt"

/**
 * Сортировка массива из 0 и 1
 *
 * Дан массив, содержащий только 0 и 1.
 *
 * Напишите функцию, которая сортирует
 * его так, чтобы все нули оказались
 * в начале, а все единички — в конце.
 */

func SortBinArray(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left] == 1 && arr[right] == 0 {
			arr[left], arr[right] = arr[right], arr[left]
		}
		if arr[left] == 0 {
			left++
		}
		if arr[right] == 1 {
			right--
		}
	}
}

func Task5() {
	var size int
	fmt.Print("Enter array size: ")
	fmt.Scanln(&size)

	fmt.Print("Enter array: ")
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	SortBinArray(arr)

	fmt.Print("Sorted arrays: ")
	fmt.Println(arr)
}
