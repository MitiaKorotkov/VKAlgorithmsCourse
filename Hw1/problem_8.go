package Hw1

/**
 * Нули в конец
 *
 * В школе прошел экзамен по математике. Несколько человек
 * списали решения и были замечены. Этим школьникам поставил
 * 0 баллов.
 *
 * Задача: есть массив с оценками, среди которых есть 0.
 * Необходимо все оценки, равные нулю перенести в конец
 * массива, чтобы все такие школьники оказались в конце
 * списка.
 */

import "fmt"

func ZerosLast(arr []int) {
	insertion := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[insertion] = arr[insertion], arr[i]
			insertion++
		}
	}
}

func Task8() {
	var size int
	fmt.Print("Enter count of students: ")
	fmt.Scanln(&size)

	fmt.Print("Enter scores: ")
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	ZerosLast(arr)

	fmt.Print("Sorted results: ")
	fmt.Println(arr)
}
