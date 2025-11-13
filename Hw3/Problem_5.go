package Hw3

import "fmt"

/*
 * Сумма двух элементов массива
 *
 * Дан не отсортированный массив целых чисел и некоторое число target.
 * Необходимо написать функцию, которая найдет два таких элемента в
 * массиве, сумма которых будет равна target
 *
 * Один элемент можно использовать лишь один раз. В случае если два
 * таких элемента не нашлось, возвращаем пустой массив
 */

func TwoSum(arr []int, targer int) (int, int) {
	cache := make(map[int]int)

	for i, n := range arr {
		if cache[targer-n] != 0 {
			return i, cache[targer-n]
		}
		cache[n] = i
	}

	return -1, -1
}

func Problem5() {
	var n int
	fmt.Print("Enter count of nums: ")
	_, _ = fmt.Scan(&n)

	fmt.Print("Enter numbers: ")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	var target int
	fmt.Print("Enter target: ")
	_, _ = fmt.Scan(&target)

	i, j := TwoSum(arr, target)
	fmt.Print("arr[i] + arr[j] = target, i = ", i, ", j = ", j)
}
