package Hw1

import "fmt"

/**
 * Дан массив целых чисел.
 *
 * Необходимо повернуть (сдвинуть) справа налево
 * часть массива, которая указана вторым параметром.
 *
 * Сделать это надо за линейное время
 * без дополнительных аллокаций
 */

func Rotate(nums []int, k int) {
	n := len(nums)
	if n > 0 {
		ReverseArray(nums)
		ReverseArray(nums[:k%n])
		ReverseArray(nums[k%n:])
	}
}

func Task3() {
	var size int
	fmt.Print("Enter count of numbers: ")
	fmt.Scanln(&size)

	fmt.Print("Enter numbers: ")
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	var k int
	fmt.Print("Enter k: ")
	fmt.Scanln(&k)

	fmt.Print("Original: ")
	fmt.Println(arr)

	Rotate(arr, k)

	fmt.Print("Rotated: ")
	fmt.Println(arr)
}
