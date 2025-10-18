package Hw1

import "fmt"

/**
 * Слияние двух отсортированных массивов.
 * Без дополнительных аллокаций
 *
 * Дано два отсортированных массива.
 * Необходимо написать функцию которая
 * объединит эти два массива в один
 * отсортированный.
 *
 * Первый массив имеет размер, равный
 * результирующиму массиву, значения
 * в конце равны нулям. Мы не должны
 * создавать третий массив.
 */

func MergeSortedArrays(arr1 []int, arr2 []int) {
	first := len(arr1) - len(arr2) - 1
	second := len(arr2) - 1

	for curr_pos := len(arr1) - 1; second >= 0; curr_pos-- {
		if first >= 0 && arr1[first] > arr2[second] {
			arr1[curr_pos] = arr1[first]
			first--
		} else {
			arr1[curr_pos] = arr2[second]
			second--
		}
	}
}

func Task4() {
	var size int
	fmt.Print("Enter count of numbers in arr1: ")
	fmt.Scanln(&size)

	fmt.Print("Enter arr1: ")
	arr1 := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Print("Enter count of numbers in arr2: ")
	fmt.Scanln(&size)

	fmt.Print("Enter arr1: ")
	arr2 := make([]int, size+len(arr1))
	for i := 0; i < size; i++ {
		fmt.Scan(&arr2[i])
	}

	fmt.Print("Original: ")
	fmt.Println(arr2)

	MergeSortedArrays(arr2, arr1)

	fmt.Print("Merged arrays: ")
	fmt.Println(arr2)
}
