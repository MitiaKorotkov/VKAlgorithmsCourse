package Hw3

import (
	"fmt"
	"sort"
)

/**
 * Накормить животных
 *
 * В небольшом зоопарке есть некоторое количество животных.
 * Каждое животное потребляет какой-то объем еды, выраженный
 * в целочисленном значении. Например, еноту нужна 1-порция
 * еды, зебре 2, пантере 3, льву 4, жирафу 8 и т.д. Каждый день,
 * смотритель зоопарка привозит еду такими же порциями. То есть
 * за раз он привозит 8, 3, 9, 1, 7. Порция на 8 может накормить
 * одно животное один раз. То есть такая порция может накормить
 * либо енота, либо льва, либо жирафа, но не может накормить,
 * например зебру и енота. Только кого-то одного. Надо написать
 * функцию, которая определит, сколько из переданных животных
 * может накормить заданное количество еды.
 */

func FeedAnumals(animals []int, food []int) int {
	sort.Ints(animals)
	sort.Ints(food)

	count := 0
	for _, f := range food {
		if count == len(animals) {
			break
		}

		if f >= animals[count] {
			count++
		}
	}

	return count
}

func Problem3() {
	var n int
	fmt.Print("Enter count of animals: ")
	_, _ = fmt.Scan(&n)

	fmt.Print("Enter the food needs of the animals: ")
	animals := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&animals[i])
	}

	var k int
	fmt.Print("Enter count of portions: ")
	_, _ = fmt.Scan(&k)

	fmt.Print("Enter the portion sizes: ")
	food := make([]int, k)
	for i := 0; i < k; i++ {
		_, _ = fmt.Scan(&food[i])
	}

	fmt.Print("Managed to feed: ", FeedAnumals(animals, food))
}
