package Hw1

import "fmt"

/**
 * Задача 1
 *
 * Дан отсортированный по возрастанию массив целых чисел
 * и некоторое число target.
 *
 * Необходимо найти два числа в массиве, которые в сумме дают
 * заданное значение target, и вернуть их индексы.
 */

func TwoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if (nums[left] + nums[right]) == target {
			return []int{left, right}
		} else if (nums[left] + nums[right]) < target {
			left++
		} else {
			right--
		}
	}

	// Note(Dima): Error! Pair not found
	return []int{-1, -1}
}

func Problem1() {
	var target int
	fmt.Print("Enter target: ")
	fmt.Scanln(&target)

	var size int
	fmt.Print("Enter count of numbers: ")
	fmt.Scanln(&size)

	fmt.Print("Enter numbers: ")
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Print("Answer: ")
	fmt.Println(TwoSum(nums, target))
}
