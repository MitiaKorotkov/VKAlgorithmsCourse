package Hw2

import "fmt"

/**
 * Является ли одна строка исходной для другой
 *
 * В исходную строку добавили некоторое
 * количество символов.
 *
 * Необходимо выявить, была ли строка
 * a исходной для строки b
 */

func isSubsequence(s1 string, s2 string) bool {
	first := 0

	for second := 0; second < len(s2) && first < len(s1); second++ {
		if s1[first] == s2[second] {
			first++
		}
	}

	return first == len(s1)
}

func Problem5() {
	var s1, s2 string

	fmt.Print("Enter first string: ")
	fmt.Scanln(&s1)

	fmt.Print("Enter first string: ")
	fmt.Scanln(&s2)

	if isSubsequence(s1, s2) {
		fmt.Println("First string IS a subset of second string")
	} else {
		fmt.Println("First string IS NOT a subset of second string")
	}
}
